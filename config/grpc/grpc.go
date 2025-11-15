package grpc

import (
	"time"
)

// Config is used for configuring a gRPC load test.
type Config struct {
	Target          string            `yaml:"target"`
	Method          string            `yaml:"method"`
	Count           int               `yaml:"count"`
	Data            string            `yaml:"data,omitempty"`        // String data
	DataHex         string            `yaml:"dataHex,omitempty"`     // Hex-encoded binary data
	DataBase64      string            `yaml:"dataBase64,omitempty"`  // Base64-encoded binary data
	Metadata        map[string]string `yaml:"metadata,omitempty"`
	Insecure        bool              `yaml:"insecure,omitempty"`
	Timeout         *time.Duration    `yaml:"timeout,omitempty"`
	MaxConcurrent   *int              `yaml:"maxConcurrent,omitempty"`
	EnableKeepalive bool              `yaml:"enableKeepalive,omitempty"`
}
