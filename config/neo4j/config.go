package neo4j

import (
	"fmt"
	"time"
)

// Config represents the configuration for Neo4j operations
type Config struct {
	// URI is the Neo4j connection URI (bolt://host:port or neo4j://host:port)
	URI string `yaml:"uri"`
	// Username for authentication
	Username string `yaml:"username"`
	// Password for authentication
	Password string `yaml:"password"`
	// Database is the database name
	Database string `yaml:"database,omitempty"`
	// Query is the Cypher query to execute
	Query string `yaml:"query"`
	// Parameters are the query parameters
	Parameters map[string]interface{} `yaml:"parameters,omitempty"`
	// Timeout for operations
	Timeout time.Duration `yaml:"timeout,omitempty"`
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.URI == "" {
		return fmt.Errorf("uri must be specified")
	}
	if c.Username == "" {
		return fmt.Errorf("username must be specified")
	}
	if c.Password == "" {
		return fmt.Errorf("password must be specified")
	}
	if c.Query == "" {
		return fmt.Errorf("query must be specified")
	}
	if c.Database == "" {
		c.Database = "neo4j"
	}
	if c.Timeout == 0 {
		c.Timeout = 10 * time.Second
	}
	return nil
}
