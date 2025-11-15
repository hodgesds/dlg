package clickhouse

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/ClickHouse/clickhouse-go/v2"
	clickhouseconfig "github.com/hodgesds/dlg/config/clickhouse"
	"github.com/hodgesds/dlg/executor"
)

type clickhouseExecutor struct{}

// New returns a new ClickHouse executor.
func New() executor.ClickHouse {
	return &clickhouseExecutor{}
}

// Execute implements the ClickHouse executor interface.
func (e *clickhouseExecutor) Execute(ctx context.Context, config *clickhouseconfig.Config) error {
	db, err := sql.Open("clickhouse", config.DSN)
	if err != nil {
		return fmt.Errorf("failed to open connection: %w", err)
	}
	defer db.Close()

	if config.MaxOpenConns > 0 {
		db.SetMaxOpenConns(config.MaxOpenConns)
	}

	if config.MaxIdleConns > 0 {
		db.SetMaxIdleConns(config.MaxIdleConns)
	}

	if config.ConnectTimeout != nil {
		db.SetConnMaxLifetime(*config.ConnectTimeout)
	}

	// Ping to verify connection
	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	// Execute the configured number of operations
	for i := 0; i < config.Count; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := e.executeOperation(ctx, db, config); err != nil {
				return err
			}
		}
	}

	return nil
}

func (e *clickhouseExecutor) executeOperation(ctx context.Context, db *sql.DB, config *clickhouseconfig.Config) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	switch config.Operation {
	case clickhouseconfig.OpInsert:
		return e.executeInsert(ctxWithTimeout, db, config)
	case clickhouseconfig.OpBatchInsert:
		return e.executeBatchInsert(ctxWithTimeout, db, config)
	case clickhouseconfig.OpSelect:
		return e.executeSelect(ctxWithTimeout, db, config)
	case clickhouseconfig.OpCount:
		return e.executeCount(ctxWithTimeout, db, config)
	case clickhouseconfig.OpCreateTable:
		return e.executeCreateTable(ctxWithTimeout, db, config)
	case clickhouseconfig.OpOptimize:
		return e.executeOptimize(ctxWithTimeout, db, config)
	default:
		return fmt.Errorf("unsupported operation: %s", config.Operation)
	}
}

func (e *clickhouseExecutor) executeInsert(ctx context.Context, db *sql.DB, config *clickhouseconfig.Config) error {
	if config.Data == nil || len(config.Data) == 0 {
		return fmt.Errorf("data is required for insert operation")
	}

	columns := make([]string, 0, len(config.Data))
	placeholders := make([]string, 0, len(config.Data))
	values := make([]interface{}, 0, len(config.Data))

	i := 1
	for col, val := range config.Data {
		columns = append(columns, col)
		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
		values = append(values, val)
		i++
	}

	query := fmt.Sprintf(
		"INSERT INTO %s.%s (%s) VALUES (%s)",
		config.Database,
		config.Table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	_, err := db.ExecContext(ctx, query, values...)
	return err
}

func (e *clickhouseExecutor) executeBatchInsert(ctx context.Context, db *sql.DB, config *clickhouseconfig.Config) error {
	if config.Values == nil || len(config.Values) == 0 {
		return fmt.Errorf("values are required for batch insert operation")
	}

	if config.Columns == nil || len(config.Columns) == 0 {
		return fmt.Errorf("columns are required for batch insert operation")
	}

	batchSize := config.BatchSize
	if batchSize <= 0 {
		batchSize = len(config.Values)
	}

	for i := 0; i < len(config.Values); i += batchSize {
		end := i + batchSize
		if end > len(config.Values) {
			end = len(config.Values)
		}

		batch := config.Values[i:end]
		if err := e.insertBatch(ctx, db, config, batch); err != nil {
			return err
		}
	}

	return nil
}

func (e *clickhouseExecutor) insertBatch(ctx context.Context, db *sql.DB, config *clickhouseconfig.Config, batch [][]interface{}) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	placeholders := make([]string, len(config.Columns))
	for i := range config.Columns {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	query := fmt.Sprintf(
		"INSERT INTO %s.%s (%s) VALUES (%s)",
		config.Database,
		config.Table,
		strings.Join(config.Columns, ", "),
		strings.Join(placeholders, ", "),
	)

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, row := range batch {
		if _, err := stmt.ExecContext(ctx, row...); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (e *clickhouseExecutor) executeSelect(ctx context.Context, db *sql.DB, config *clickhouseconfig.Config) error {
	query := config.Query
	if query == "" {
		query = fmt.Sprintf("SELECT * FROM %s.%s LIMIT 100", config.Database, config.Table)
	}

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Consume all rows
	for rows.Next() {
		// Just iterate, don't need to scan
	}

	return rows.Err()
}

func (e *clickhouseExecutor) executeCount(ctx context.Context, db *sql.DB, config *clickhouseconfig.Config) error {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s.%s", config.Database, config.Table)
	if config.Query != "" {
		query = config.Query
	}

	var count int64
	err := db.QueryRowContext(ctx, query).Scan(&count)
	return err
}

func (e *clickhouseExecutor) executeCreateTable(ctx context.Context, db *sql.DB, config *clickhouseconfig.Config) error {
	if config.TableSchema == "" {
		return fmt.Errorf("tableSchema is required for create_table operation")
	}

	_, err := db.ExecContext(ctx, config.TableSchema)
	return err
}

func (e *clickhouseExecutor) executeOptimize(ctx context.Context, db *sql.DB, config *clickhouseconfig.Config) error {
	query := fmt.Sprintf("OPTIMIZE TABLE %s.%s", config.Database, config.Table)
	_, err := db.ExecContext(ctx, query)
	return err
}
