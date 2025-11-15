package cassandra

import (
	"time"
)

// Consistency represents Cassandra consistency levels
type Consistency string

const (
	ConsistencyAny         Consistency = "ANY"
	ConsistencyOne         Consistency = "ONE"
	ConsistencyTwo         Consistency = "TWO"
	ConsistencyThree       Consistency = "THREE"
	ConsistencyQuorum      Consistency = "QUORUM"
	ConsistencyAll         Consistency = "ALL"
	ConsistencyLocalQuorum Consistency = "LOCAL_QUORUM"
	ConsistencyEachQuorum  Consistency = "EACH_QUORUM"
	ConsistencyLocalOne    Consistency = "LOCAL_ONE"
)

// Query represents a CQL query to execute
type Query struct {
	CQL    string                 `yaml:"cql"`
	Values []interface{}          `yaml:"values,omitempty"`
	Scan   bool                   `yaml:"scan,omitempty"` // If true, scan results
}

// Config is used for configuring a Cassandra load test.
type Config struct {
	Hosts          []string       `yaml:"hosts"`
	Keyspace       string         `yaml:"keyspace"`
	Consistency    Consistency    `yaml:"consistency,omitempty"`
	Username       string         `yaml:"username,omitempty"`
	Password       string         `yaml:"password,omitempty"`
	Count          int            `yaml:"count"`
	Queries        []Query        `yaml:"queries"`
	ConnectTimeout *time.Duration `yaml:"connectTimeout,omitempty"`
	Timeout        *time.Duration `yaml:"timeout,omitempty"`
	NumConns       *int           `yaml:"numConns,omitempty"`
}
