package memcache

import (
	"fmt"
	"strings"

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

// ParseOp is used to parse an Op.
func ParseOp(s string) (*Op, error) {
	split := strings.Split(s, ":")
	if len(split) == 1 {
		return nil, fmt.Errorf("expected command (get,set,delete), got: %q", s)
	}
	switch split[0] {
	case "get":
		return &Op{
			Get: &Get{
				Key: split[1],
			},
		}, nil
	case "set":
		if len(split) != 3 {
			return nil, fmt.Errorf("invalid set: %q", s)
		}
		return &Op{
			Set: &Set{
				Key:   split[1],
				Value: split[2],
			},
		}, nil
	case "delete":
		return &Op{
			Delete: &Delete{
				Key: split[1],
			},
		}, nil
	default:
		return nil, fmt.Errorf("unknown command: %q", s)
	}
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
