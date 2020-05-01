package ssh

import (
	"context"
	"io/ioutil"
	"time"

	"golang.org/x/crypto/ssh"
)

// Config is a ssh config.
type Config struct {
	Addr              string   `yaml:"addr"`
	User              string   `yaml:"user"`
	Password          string   `yaml:"password,omitempty"`
	KeyFile           string   `yaml:"keyFile,omitempty"`
	ClientVersion     string   `yaml:"clientVersion,omitempty"`
	HostKeyAlgorithms []string `yaml:"hostKeyAlgorithms,omitempty"`
	Cmd               *string  `yaml:"cmd,omitempty"`
}

// SSHClient is used to create a SSH client.
func (c *Config) SSHClient(ctx context.Context) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User:              c.User,
		ClientVersion:     c.ClientVersion,
		HostKeyAlgorithms: c.HostKeyAlgorithms,
		HostKeyCallback:   ssh.InsecureIgnoreHostKey(),
	}
	if c.Password != "" {
		config.Auth = []ssh.AuthMethod{ssh.Password(c.Password)}
	}
	if c.KeyFile != "" {
		key, err := ioutil.ReadFile(c.KeyFile)
		if err != nil {
			return nil, err
		}
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			return nil, err
		}
		config.Auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
	}
	if t, ok := ctx.Deadline(); ok {
		config.Timeout = time.Until(t)
	}
	return ssh.Dial("tcp", c.Addr, config)
}
