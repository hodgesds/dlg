package syslog

import (
	"fmt"
	"time"
)

// Config represents the configuration for Syslog operations
type Config struct {
	// Host is the Syslog server host
	Host string `yaml:"host"`
	// Port is the Syslog server port
	Port int `yaml:"port,omitempty"`
	// Protocol is the transport protocol (tcp, udp)
	Protocol string `yaml:"protocol,omitempty"`
	// Message is the log message to send
	Message string `yaml:"message"`
	// Severity is the syslog severity level (0-7)
	Severity int `yaml:"severity,omitempty"`
	// Facility is the syslog facility (0-23)
	Facility int `yaml:"facility,omitempty"`
	// Tag is the application tag
	Tag string `yaml:"tag,omitempty"`
	// Hostname is the hostname to report
	Hostname string `yaml:"hostname,omitempty"`
	// Format is the syslog format (rfc3164, rfc5424)
	Format string `yaml:"format,omitempty"`
	// Timeout for operations
	Timeout time.Duration `yaml:"timeout,omitempty"`
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("host must be specified")
	}
	if c.Message == "" {
		return fmt.Errorf("message must be specified")
	}
	if c.Port == 0 {
		c.Port = 514
	}
	if c.Protocol == "" {
		c.Protocol = "udp"
	}
	if c.Protocol != "tcp" && c.Protocol != "udp" {
		return fmt.Errorf("protocol must be tcp or udp")
	}
	if c.Severity < 0 || c.Severity > 7 {
		c.Severity = 6 // Info
	}
	if c.Facility < 0 || c.Facility > 23 {
		c.Facility = 1 // User-level
	}
	if c.Tag == "" {
		c.Tag = "dlg"
	}
	if c.Format == "" {
		c.Format = "rfc3164"
	}
	if c.Timeout == 0 {
		c.Timeout = 5 * time.Second
	}
	return nil
}
