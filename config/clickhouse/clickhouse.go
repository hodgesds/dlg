package clickhouse

import (
	"time"
)

// Operation represents a ClickHouse operation type
type Operation string

const (
	OpInsert      Operation = "insert"
	OpSelect      Operation = "select"
	OpBatchInsert Operation = "batch_insert"
	OpCreateTable Operation = "create_table"
	OpOptimize    Operation = "optimize"
	OpCount       Operation = "count"
)

// Config is used for configuring a ClickHouse load test.
type Config struct {
	DSN            string                 `yaml:"dsn"`
	Database       string                 `yaml:"database"`
	Table          string                 `yaml:"table"`
	Operation      Operation              `yaml:"operation"`
	Count          int                    `yaml:"count"`
	BatchSize      int                    `yaml:"batchSize,omitempty"`
	Query          string                 `yaml:"query,omitempty"`
	Columns        []string               `yaml:"columns,omitempty"`
	Values         [][]interface{}        `yaml:"values,omitempty"`
	Data           map[string]interface{} `yaml:"data,omitempty"`
	TableSchema    string                 `yaml:"tableSchema,omitempty"`
	ConnectTimeout *time.Duration         `yaml:"connectTimeout,omitempty"`
	MaxOpenConns   int                    `yaml:"maxOpenConns,omitempty"`
	MaxIdleConns   int                    `yaml:"maxIdleConns,omitempty"`
}
