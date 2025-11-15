package tcp

import (
	"fmt"
	"time"
)

// Config represents the configuration for TCP operations
type Config struct {
	// Host is the TCP server host
	Host string `yaml:"host"`
	// Port is the TCP server port
	Port int `yaml:"port"`
	// Operation is the type of operation (connect, send, send_receive)
	Operation string `yaml:"operation,omitempty"`
	// Data is the data to send
	Data string `yaml:"data,omitempty"`
	// BinaryData is hex-encoded binary data to send
	BinaryData string `yaml:"binary_data,omitempty"`
	// KeepAlive enables TCP keep-alive
	KeepAlive bool `yaml:"keep_alive,omitempty"`
	// NoDelay disables Nagle's algorithm
	NoDelay bool `yaml:"no_delay,omitempty"`
	// Timeout for operations
	Timeout time.Duration `yaml:"timeout,omitempty"`
	// ReadTimeout for read operations
	ReadTimeout time.Duration `yaml:"read_timeout,omitempty"`
	// WriteTimeout for write operations
	WriteTimeout time.Duration `yaml:"write_timeout,omitempty"`
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("host must be specified")
	}
	if c.Port == 0 {
		return fmt.Errorf("port must be specified")
	}
	if c.Operation == "" {
		c.Operation = "connect"
	}
	if c.Timeout == 0 {
		c.Timeout = 10 * time.Second
	}
	if c.ReadTimeout == 0 {
		c.ReadTimeout = c.Timeout
	}
	if c.WriteTimeout == 0 {
		c.WriteTimeout = c.Timeout
	}
	return nil
}
