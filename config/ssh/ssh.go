package ssh

import (
	"time"

	"golang.org/x/crypto/ssh"
)

// Config is a ssh config.
type Config struct {
	Addr    string
	User    string
	Timeout time.Duration
}

// SSHClient is used to create a SSH client.
func (c *Config) SSHClient() (*ssh.Client, error) {
	config := &ssh.ClientConfig{}
	return ssh.Dial("tcp", c.Addr, config)
}
