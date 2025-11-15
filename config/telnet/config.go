package telnet

import (
	"fmt"
	"time"
)

// Config represents the configuration for Telnet operations
type Config struct {
	// Host is the Telnet server host
	Host string `yaml:"host"`
	// Port is the Telnet server port
	Port int `yaml:"port,omitempty"`
	// Username for authentication
	Username string `yaml:"username,omitempty"`
	// Password for authentication
	Password string `yaml:"password,omitempty"`
	// Commands are the commands to execute
	Commands []string `yaml:"commands,omitempty"`
	// ExpectPrompt is the prompt to expect after login
	ExpectPrompt string `yaml:"expect_prompt,omitempty"`
	// Timeout for operations
	Timeout time.Duration `yaml:"timeout,omitempty"`
	// ConnectOnly only connects without executing commands
	ConnectOnly bool `yaml:"connect_only,omitempty"`
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("host must be specified")
	}
	if c.Port == 0 {
		c.Port = 23
	}
	if c.Timeout == 0 {
		c.Timeout = 10 * time.Second
	}
	if c.ExpectPrompt == "" {
		c.ExpectPrompt = "$"
	}
	return nil
}
