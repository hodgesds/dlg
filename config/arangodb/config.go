package arangodb

import (
	"fmt"
	"time"
)

// Config represents the configuration for ArangoDB operations
type Config struct {
	// Endpoints is the list of ArangoDB server endpoints
	Endpoints []string `yaml:"endpoints"`
	// Database is the database name
	Database string `yaml:"database"`
	// Collection is the collection name
	Collection string `yaml:"collection,omitempty"`
	// Operation is the type of operation (query, insert, update, delete)
	Operation string `yaml:"operation"`
	// Query is the AQL query to execute
	Query string `yaml:"query,omitempty"`
	// BindVars are the query bind variables
	BindVars map[string]interface{} `yaml:"bind_vars,omitempty"`
	// Document is the document data for insert/update
	Document map[string]interface{} `yaml:"document,omitempty"`
	// Key is the document key
	Key string `yaml:"key,omitempty"`
	// Username for authentication
	Username string `yaml:"username,omitempty"`
	// Password for authentication
	Password string `yaml:"password,omitempty"`
	// Timeout for operations
	Timeout time.Duration `yaml:"timeout,omitempty"`
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if len(c.Endpoints) == 0 {
		return fmt.Errorf("endpoints must be specified")
	}
	if c.Database == "" {
		return fmt.Errorf("database must be specified")
	}
	if c.Operation == "" {
		c.Operation = "query"
	}
	if c.Timeout == 0 {
		c.Timeout = 10 * time.Second
	}
	return nil
}
