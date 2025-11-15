package ftp

import (
	"fmt"
	"time"
)

// Config represents the configuration for FTP/SFTP operations
type Config struct {
	// Host is the FTP/SFTP server host
	Host string `yaml:"host"`
	// Port is the server port
	Port int `yaml:"port,omitempty"`
	// Username for authentication
	Username string `yaml:"username"`
	// Password for authentication
	Password string `yaml:"password,omitempty"`
	// PrivateKey path for SFTP authentication
	PrivateKey string `yaml:"private_key,omitempty"`
	// Operation is the type of operation (list, upload, download, delete)
	Operation string `yaml:"operation"`
	// RemotePath is the remote file/directory path
	RemotePath string `yaml:"remote_path"`
	// LocalPath is the local file path (for upload/download)
	LocalPath string `yaml:"local_path,omitempty"`
	// Data is the data to upload (alternative to LocalPath)
	Data string `yaml:"data,omitempty"`
	// Protocol is either "ftp" or "sftp"
	Protocol string `yaml:"protocol,omitempty"`
	// Timeout for operations
	Timeout time.Duration `yaml:"timeout,omitempty"`
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("host must be specified")
	}
	if c.Username == "" {
		return fmt.Errorf("username must be specified")
	}
	if c.Operation == "" {
		c.Operation = "list"
	}
	if c.Protocol == "" {
		c.Protocol = "sftp"
	}
	if c.Protocol != "ftp" && c.Protocol != "sftp" {
		return fmt.Errorf("protocol must be either ftp or sftp")
	}
	if c.Port == 0 {
		if c.Protocol == "ftp" {
			c.Port = 21
		} else {
			c.Port = 22
		}
	}
	if c.Timeout == 0 {
		c.Timeout = 30 * time.Second
	}
	return nil
}
