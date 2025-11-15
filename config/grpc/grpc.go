package grpc

import (
	"time"
)

// Config is used for configuring a gRPC load test.
type Config struct {
	Target          string            `yaml:"target"`
	Method          string            `yaml:"method"`
	Count           int               `yaml:"count"`
	Data            string            `yaml:"data,omitempty"`
	Metadata        map[string]string `yaml:"metadata,omitempty"`
	Insecure        bool              `yaml:"insecure,omitempty"`
	Timeout         *time.Duration    `yaml:"timeout,omitempty"`
	MaxConcurrent   *int              `yaml:"maxConcurrent,omitempty"`
	EnableKeepalive bool              `yaml:"enableKeepalive,omitempty"`
}
