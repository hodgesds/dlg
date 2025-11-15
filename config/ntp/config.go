package ntp

import (
	"fmt"
	"time"
)

// Config represents the configuration for NTP operations
type Config struct {
	// Host is the NTP server host
	Host string `yaml:"host"`
	// Port is the NTP server port
	Port int `yaml:"port,omitempty"`
	// Version is the NTP protocol version (3 or 4)
	Version int `yaml:"version,omitempty"`
	// Timeout for NTP query
	Timeout time.Duration `yaml:"timeout,omitempty"`
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("host must be specified")
	}
	if c.Port == 0 {
		c.Port = 123
	}
	if c.Version == 0 {
		c.Version = 4
	}
	if c.Version < 3 || c.Version > 4 {
		return fmt.Errorf("version must be 3 or 4")
	}
	if c.Timeout == 0 {
		c.Timeout = 5 * time.Second
	}
	return nil
}
