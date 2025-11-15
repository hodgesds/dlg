package scylladb

import (
	"fmt"
	"time"
)

// Config represents the configuration for ScyllaDB operations
type Config struct {
	// Hosts is the list of ScyllaDB host addresses
	Hosts []string `yaml:"hosts"`
	// Keyspace is the keyspace to use
	Keyspace string `yaml:"keyspace"`
	// Query is the CQL query to execute
	Query string `yaml:"query"`
	// Values are the query parameters
	Values []interface{} `yaml:"values,omitempty"`
	// Consistency is the consistency level (one, quorum, all, etc.)
	Consistency string `yaml:"consistency,omitempty"`
	// Username for authentication
	Username string `yaml:"username,omitempty"`
	// Password for authentication
	Password string `yaml:"password,omitempty"`
	// Timeout for operations
	Timeout time.Duration `yaml:"timeout,omitempty"`
	// Port is the ScyllaDB port
	Port int `yaml:"port,omitempty"`
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if len(c.Hosts) == 0 {
		return fmt.Errorf("hosts must be specified")
	}
	if c.Keyspace == "" {
		return fmt.Errorf("keyspace must be specified")
	}
	if c.Query == "" {
		return fmt.Errorf("query must be specified")
	}
	if c.Consistency == "" {
		c.Consistency = "quorum"
	}
	if c.Timeout == 0 {
		c.Timeout = 10 * time.Second
	}
	if c.Port == 0 {
		c.Port = 9042
	}
	return nil
}
