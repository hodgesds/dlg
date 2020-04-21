package sql

import (
	"context"
	"database/sql"
	"sync"

	sqlconf "github.com/hodgesds/dlg/config/sql"
	"github.com/hodgesds/dlg/executor"
	"go.uber.org/multierr"
)

type sqlExecutor struct {
}

// New returns a SQL executor.
func New() executor.SQL {
	return &sqlExecutor{}
}

// Execute implements the SQL executor interface.
func (e *sqlExecutor) Execute(ctx context.Context, c *sqlconf.Config) error {
	if err := c.Validate(); err != nil {
		return err
	}
	db, err := c.DB()
	if err != nil {
		return nil
	}
	if err := setupDB(db, c); err != nil {
		return err
	}
	if c.Concurrent {
		return e.execParallel(ctx, db, c)
	}
	for _, payload := range c.Payloads {
		_, err := db.ExecContext(ctx, payload.Exec)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *sqlExecutor) execParallel(
	ctx context.Context,
	db *sql.DB,
	c *sqlconf.Config,
) error {
	var (
		wg  sync.WaitGroup
		mu  sync.Mutex
		err error
	)
	for _, payload := range c.Payloads {
		wg.Add(1)
		go func(payload *sqlconf.Payload) {
			_, err2 := db.ExecContext(ctx, payload.Exec)
			if err2 != nil {
				mu.Lock()
				err = multierr.Append(err, err2)
				mu.Unlock()
			}
			wg.Done()
		}(payload)
	}
	wg.Wait()

	return err
}

func setupDB(db *sql.DB, c *sqlconf.Config) error {
	if err := db.Ping(); err != nil {
		return err
	}
	if c.MaxConns > 0 {
		db.SetMaxOpenConns(c.MaxConns)
	}
	if c.MaxIdleConns > 0 {
		db.SetMaxIdleConns(c.MaxIdleConns)
	}
	return nil
}
