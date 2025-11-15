package icmp

import (
	"fmt"
	"time"
)

// Config represents the configuration for ICMP ping operations
type Config struct {
	// Host is the target host to ping
	Host string `yaml:"host"`
	// Count is the number of ping packets to send
	Count int `yaml:"count,omitempty"`
	// Size is the size of the ping packet payload
	Size int `yaml:"size,omitempty"`
	// Timeout for ping operation
	Timeout time.Duration `yaml:"timeout,omitempty"`
	// Interval between ping packets
	Interval time.Duration `yaml:"interval,omitempty"`
	// Privileged mode (requires root/admin)
	Privileged bool `yaml:"privileged,omitempty"`
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("host must be specified")
	}
	if c.Count == 0 {
		c.Count = 1
	}
	if c.Size == 0 {
		c.Size = 56
	}
	if c.Timeout == 0 {
		c.Timeout = 10 * time.Second
	}
	if c.Interval == 0 {
		c.Interval = 1 * time.Second
	}
	return nil
}
