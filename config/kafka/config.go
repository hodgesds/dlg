package kafka

import (
	"fmt"
	"time"
)

// Config represents the configuration for Kafka operations
type Config struct {
	// Brokers is the list of Kafka broker addresses
	Brokers []string `yaml:"brokers"`
	// Topic is the Kafka topic to produce/consume from
	Topic string `yaml:"topic"`
	// Operation is the type of Kafka operation (produce, consume)
	Operation string `yaml:"operation"`
	// Message is the message to produce
	Message string `yaml:"message,omitempty"`
	// Key is the message key
	Key string `yaml:"key,omitempty"`
	// Partition is the partition to produce/consume from (-1 for auto)
	Partition int32 `yaml:"partition,omitempty"`
	// GroupID is the consumer group ID
	GroupID string `yaml:"group_id,omitempty"`
	// Timeout for operations
	Timeout time.Duration `yaml:"timeout,omitempty"`
	// SASLEnabled enables SASL authentication
	SASLEnabled bool `yaml:"sasl_enabled,omitempty"`
	// SASLUsername for authentication
	SASLUsername string `yaml:"sasl_username,omitempty"`
	// SASLPassword for authentication
	SASLPassword string `yaml:"sasl_password,omitempty"`
	// TLSEnabled enables TLS
	TLSEnabled bool `yaml:"tls_enabled,omitempty"`
	// Async enables async production
	Async bool `yaml:"async,omitempty"`
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if len(c.Brokers) == 0 {
		return fmt.Errorf("brokers must be specified")
	}
	if c.Topic == "" {
		return fmt.Errorf("topic must be specified")
	}
	if c.Operation == "" {
		c.Operation = "produce"
	}
	if c.Timeout == 0 {
		c.Timeout = 10 * time.Second
	}
	return nil
}
