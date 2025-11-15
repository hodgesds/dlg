package mongodb

import (
	"time"
)

// Operation represents a MongoDB operation type
type Operation string

const (
	OpInsert  Operation = "insert"
	OpFind    Operation = "find"
	OpUpdate  Operation = "update"
	OpDelete  Operation = "delete"
	OpCount   Operation = "count"
	OpAggregate Operation = "aggregate"
)

// Config is used for configuring a MongoDB load test.
type Config struct {
	URI            string            `yaml:"uri"`
	Database       string            `yaml:"database"`
	Collection     string            `yaml:"collection"`
	Operation      Operation         `yaml:"operation"`
	Count          int               `yaml:"count"`
	Document       map[string]interface{} `yaml:"document,omitempty"`
	Filter         map[string]interface{} `yaml:"filter,omitempty"`
	Update         map[string]interface{} `yaml:"update,omitempty"`
	ConnectTimeout *time.Duration    `yaml:"connectTimeout,omitempty"`
	MaxPoolSize    *uint64           `yaml:"maxPoolSize,omitempty"`
	MinPoolSize    *uint64           `yaml:"minPoolSize,omitempty"`
}
