package nats

import (
	"fmt"
	"time"
)

// Config represents the configuration for NATS operations
type Config struct {
	// URL is the NATS server URL (nats://host:port)
	URL string `yaml:"url"`
	// Subject is the NATS subject
	Subject string `yaml:"subject"`
	// Operation is the type of operation (publish, subscribe, request)
	Operation string `yaml:"operation"`
	// Message is the message to publish/request
	Message string `yaml:"message,omitempty"`
	// Queue is the queue group name for subscribers
	Queue string `yaml:"queue,omitempty"`
	// Timeout for operations
	Timeout time.Duration `yaml:"timeout,omitempty"`
	// Username for authentication
	Username string `yaml:"username,omitempty"`
	// Password for authentication
	Password string `yaml:"password,omitempty"`
	// Token for authentication
	Token string `yaml:"token,omitempty"`
	// TLSEnabled enables TLS
	TLSEnabled bool `yaml:"tls_enabled,omitempty"`
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.URL == "" {
		return fmt.Errorf("url must be specified")
	}
	if c.Subject == "" {
		return fmt.Errorf("subject must be specified")
	}
	if c.Operation == "" {
		c.Operation = "publish"
	}
	if c.Timeout == 0 {
		c.Timeout = 10 * time.Second
	}
	return nil
}
