package graphql

import (
	"time"
)

// Config is used for configuring a GraphQL load test.
type Config struct {
	Endpoint  string            `yaml:"endpoint"`
	Query     string            `yaml:"query"`
	Variables map[string]interface{} `yaml:"variables,omitempty"`
	Headers   map[string]string `yaml:"headers,omitempty"`
	Count     int               `yaml:"count"`
	Timeout   *time.Duration    `yaml:"timeout,omitempty"`
}
