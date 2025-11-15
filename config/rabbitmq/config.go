package rabbitmq

import (
	"fmt"
	"time"
)

// Config represents the configuration for RabbitMQ operations
type Config struct {
	// URL is the RabbitMQ connection URL (amqp://user:pass@host:port/)
	URL string `yaml:"url"`
	// Queue is the queue name
	Queue string `yaml:"queue"`
	// Exchange is the exchange name
	Exchange string `yaml:"exchange,omitempty"`
	// RoutingKey is the routing key
	RoutingKey string `yaml:"routing_key,omitempty"`
	// Operation is the type of operation (publish, consume)
	Operation string `yaml:"operation"`
	// Message is the message to publish
	Message string `yaml:"message,omitempty"`
	// Durable makes the queue durable
	Durable bool `yaml:"durable,omitempty"`
	// AutoDelete removes the queue when unused
	AutoDelete bool `yaml:"auto_delete,omitempty"`
	// Exclusive makes the queue exclusive
	Exclusive bool `yaml:"exclusive,omitempty"`
	// NoWait doesn't wait for server confirmation
	NoWait bool `yaml:"no_wait,omitempty"`
	// Timeout for operations
	Timeout time.Duration `yaml:"timeout,omitempty"`
	// AutoAck enables automatic acknowledgment
	AutoAck bool `yaml:"auto_ack,omitempty"`
	// Prefetch is the prefetch count for consumers
	Prefetch int `yaml:"prefetch,omitempty"`
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.URL == "" {
		return fmt.Errorf("url must be specified")
	}
	if c.Queue == "" && c.Exchange == "" {
		return fmt.Errorf("either queue or exchange must be specified")
	}
	if c.Operation == "" {
		c.Operation = "publish"
	}
	if c.Timeout == 0 {
		c.Timeout = 10 * time.Second
	}
	if c.Prefetch == 0 {
		c.Prefetch = 1
	}
	return nil
}
