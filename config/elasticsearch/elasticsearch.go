package elasticsearch

import (
	"time"
)

// Operation represents an Elasticsearch operation type
type Operation string

const (
	OpIndex  Operation = "index"
	OpGet    Operation = "get"
	OpSearch Operation = "search"
	OpUpdate Operation = "update"
	OpDelete Operation = "delete"
	OpBulk   Operation = "bulk"
)

// Config is used for configuring an Elasticsearch load test.
type Config struct {
	Addresses      []string               `yaml:"addresses"`
	Username       string                 `yaml:"username,omitempty"`
	Password       string                 `yaml:"password,omitempty"`
	Index          string                 `yaml:"index"`
	Operation      Operation              `yaml:"operation"`
	Count          int                    `yaml:"count"`
	DocumentID     string                 `yaml:"documentId,omitempty"`
	Document       map[string]interface{} `yaml:"document,omitempty"`
	Query          map[string]interface{} `yaml:"query,omitempty"`
	CloudID        string                 `yaml:"cloudId,omitempty"`
	APIKey         string                 `yaml:"apiKey,omitempty"`
	Timeout        *time.Duration         `yaml:"timeout,omitempty"`
	MaxRetries     *int                   `yaml:"maxRetries,omitempty"`
	EnableMetrics  bool                   `yaml:"enableMetrics,omitempty"`
}
