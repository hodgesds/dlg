package couchdb

import (
	"fmt"
	"time"
)

// Config represents the configuration for CouchDB operations
type Config struct {
	// URL is the CouchDB server URL (http://host:port)
	URL string `yaml:"url"`
	// Database is the database name
	Database string `yaml:"database"`
	// Operation is the type of operation (get, put, delete, query)
	Operation string `yaml:"operation"`
	// DocumentID is the document ID
	DocumentID string `yaml:"document_id,omitempty"`
	// Document is the document data (JSON string)
	Document string `yaml:"document,omitempty"`
	// Query is the Mango query (JSON string)
	Query string `yaml:"query,omitempty"`
	// Username for authentication
	Username string `yaml:"username,omitempty"`
	// Password for authentication
	Password string `yaml:"password,omitempty"`
	// Timeout for operations
	Timeout time.Duration `yaml:"timeout,omitempty"`
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.URL == "" {
		return fmt.Errorf("url must be specified")
	}
	if c.Database == "" {
		return fmt.Errorf("database must be specified")
	}
	if c.Operation == "" {
		c.Operation = "get"
	}
	if c.Timeout == 0 {
		c.Timeout = 10 * time.Second
	}
	return nil
}
