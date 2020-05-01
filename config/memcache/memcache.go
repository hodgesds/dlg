package memcache

import (
	"github.com/rainycape/memcache"
)

// Config is used for configuring memcache.
type Config struct {
	Addrs []string `yaml:"addrs"`
	Ops   []*Op    `yaml:"ops"`
}

// Client returns a memcache Client.
func (c *Config) Client() (*memcache.Client, error) {
	return memcache.New(c.Addrs...)
}

// Op is used for configuring Memcache operations.
type Op struct {
	Get    *Get    `yaml:"get,omitempty"`
	Delete *Delete `yaml:"delete,omitempty"`
	Set    *Set    `yaml:"set,omitempty"`
}

// Get is an memcache get KV op.
type Get struct {
	Key string `yaml:"key"`
}

// Delete is an memcache delete KV op.
type Delete struct {
	Key string `yaml:"key"`
}

// Set is an memcache set KV op.
type Set struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}
