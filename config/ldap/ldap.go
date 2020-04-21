package ldap

import (
	"crypto/tls"

	v3 "github.com/go-ldap/ldap/v3"
)

// Config is a config struct for LDAP.
type Config struct {
	Addr   string          `yaml:"addr"`
	TLS    bool            `yaml:"tls"`
	User   string          `yaml:"user"`
	Pass   string          `yaml:"pass"`
	Search []*SearchConfig `yaml:"search"`
}

// https://godoc.org/github.com/go-ldap/ldap#SearchRequest

// SearchConfig is configuration for a search request.
type SearchConfig struct {
	BaseDN       string   `yaml:"baseDN"`
	Scope        int      `yaml:"scope"`
	DerefAliases int      `yaml:"derefAliases"`
	SizeLimit    int      `yaml:"sizeLimit"`
	TimeLimit    int      `yaml:"timeLimit"`
	Filter       string   `yaml:"filter"`
	Attributes   []string `yaml:"attributes"`
}

// Conn returns a configured client connection.
func (c *Config) Conn() (*v3.Conn, error) {
	opts := []v3.DialOpt{}
	if c.TLS {
		opts = append(opts, v3.DialWithTLSConfig(
			&tls.Config{InsecureSkipVerify: true},
		))
	}
	return v3.DialURL(c.Addr, opts...)
}
