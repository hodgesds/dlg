package pulsar

import (
	"fmt"
	"time"
)

// Config represents the configuration for Apache Pulsar operations
type Config struct {
	// URL is the Pulsar broker URL (pulsar://host:port)
	URL string `yaml:"url"`
	// Topic is the Pulsar topic
	Topic string `yaml:"topic"`
	// Operation is the type of operation (produce, consume)
	Operation string `yaml:"operation"`
	// Message is the message to produce
	Message string `yaml:"message,omitempty"`
	// SubscriptionName is the subscription name for consumers
	SubscriptionName string `yaml:"subscription_name,omitempty"`
	// SubscriptionType is the subscription type (exclusive, shared, failover, key_shared)
	SubscriptionType string `yaml:"subscription_type,omitempty"`
	// Timeout for operations
	Timeout time.Duration `yaml:"timeout,omitempty"`
	// TLSEnabled enables TLS
	TLSEnabled bool `yaml:"tls_enabled,omitempty"`
	// AuthToken for authentication
	AuthToken string `yaml:"auth_token,omitempty"`
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.URL == "" {
		return fmt.Errorf("url must be specified")
	}
	if c.Topic == "" {
		return fmt.Errorf("topic must be specified")
	}
	if c.Operation == "" {
		c.Operation = "produce"
	}
	if c.Operation == "consume" && c.SubscriptionName == "" {
		return fmt.Errorf("subscription_name must be specified for consume operation")
	}
	if c.SubscriptionType == "" {
		c.SubscriptionType = "exclusive"
	}
	if c.Timeout == 0 {
		c.Timeout = 10 * time.Second
	}
	return nil
}
