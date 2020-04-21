package sql

import (
	gsql "database/sql"
	"errors"

	// db imports
	_ "github.com/ClickHouse/clickhouse-go"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// Config is used for configuring SQL generators.
type Config struct {
	PostgresDSN   string     `yaml:"postgresDsn,omitempty"`
	MysqlDSN      string     `yaml:"mysqlDSN,omitempty"`
	ClickHouseDSN string     `yaml:"clickhouseDSN,omitempty"`
	MaxConns      int        `yaml:"maxConns"`
	MaxIdleConns  int        `yaml:"maxIdleConns"`
	Concurrent    bool       `yaml:"concurrent"`
	Payloads      []*Payload `yaml:"payloads"`
}

// DB returns a sql DB from a SQL config.
func (c *Config) DB() (*gsql.DB, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}
	if c.MysqlDSN != "" {
		return gsql.Open("mysql", c.MysqlDSN)
	}
	if c.PostgresDSN != "" {
		return gsql.Open("postgres", c.PostgresDSN)
	}
	return gsql.Open("clickhouse", c.ClickHouseDSN)
}

// Payload is used for configuring a SQL payload.
type Payload struct {
	Exec string `yaml:"exec"`
}

// Validate is used to validate a SQL configuration.
func (c *Config) Validate() error {
	configuredDSNs := 0
	if c.PostgresDSN != "" {
		configuredDSNs++
	}
	if c.MysqlDSN != "" {
		configuredDSNs++
	}
	if c.ClickHouseDSN != "" {
		configuredDSNs++
	}
	if configuredDSNs != 1 {
		return errors.New("expected exactly one configured DSN")
	}
	return nil
}
