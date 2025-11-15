package tftp

import (
	"fmt"
	"time"
)

// Config represents the configuration for TFTP operations
type Config struct {
	// Host is the TFTP server host
	Host string `yaml:"host"`
	// Port is the TFTP server port
	Port int `yaml:"port,omitempty"`
	// Operation is the type of operation (read, write)
	Operation string `yaml:"operation"`
	// RemotePath is the remote file path
	RemotePath string `yaml:"remote_path"`
	// LocalPath is the local file path (for read/write)
	LocalPath string `yaml:"local_path,omitempty"`
	// Data is the data to write (alternative to LocalPath)
	Data string `yaml:"data,omitempty"`
	// Mode is the transfer mode (octet, netascii)
	Mode string `yaml:"mode,omitempty"`
	// BlockSize is the TFTP block size
	BlockSize int `yaml:"block_size,omitempty"`
	// Timeout for operations
	Timeout time.Duration `yaml:"timeout,omitempty"`
	// Retries is the number of retries
	Retries int `yaml:"retries,omitempty"`
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("host must be specified")
	}
	if c.Port == 0 {
		c.Port = 69
	}
	if c.Operation == "" {
		c.Operation = "read"
	}
	if c.RemotePath == "" {
		return fmt.Errorf("remote_path must be specified")
	}
	if c.Mode == "" {
		c.Mode = "octet"
	}
	if c.BlockSize == 0 {
		c.BlockSize = 512
	}
	if c.Timeout == 0 {
		c.Timeout = 5 * time.Second
	}
	if c.Retries == 0 {
		c.Retries = 3
	}
	return nil
}
