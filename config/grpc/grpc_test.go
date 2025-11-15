package grpc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestConfigBasic tests basic gRPC configuration.
func TestConfigBasic(t *testing.T) {
	config := &Config{
		Target: "localhost:50051",
		Method: "helloworld.Greeter/SayHello",
		Count:  100,
		Data:   `{"name": "test"}`,
	}

	assert.Equal(t, "localhost:50051", config.Target)
	assert.Equal(t, "helloworld.Greeter/SayHello", config.Method)
	assert.Equal(t, 100, config.Count)
	assert.Equal(t, `{"name": "test"}`, config.Data)
}

// TestConfigWithMetadata tests configuration with metadata.
func TestConfigWithMetadata(t *testing.T) {
	metadata := map[string]string{
		"authorization": "Bearer token123",
		"x-request-id":  "test-id",
	}

	config := &Config{
		Target:   "localhost:50051",
		Method:   "test.Service/Method",
		Count:    1,
		Metadata: metadata,
	}

	assert.Equal(t, "Bearer token123", config.Metadata["authorization"])
	assert.Equal(t, "test-id", config.Metadata["x-request-id"])
}

// TestConfigInsecure tests insecure connection configuration.
func TestConfigInsecure(t *testing.T) {
	config := &Config{
		Target:   "localhost:50051",
		Method:   "test.Service/Method",
		Count:    1,
		Insecure: true,
	}

	assert.True(t, config.Insecure)
}

// TestConfigWithTimeout tests configuration with timeout.
func TestConfigWithTimeout(t *testing.T) {
	timeout := 30 * time.Second
	config := &Config{
		Target:  "localhost:50051",
		Method:  "test.Service/Method",
		Count:   1,
		Timeout: &timeout,
	}

	assert.NotNil(t, config.Timeout)
	assert.Equal(t, 30*time.Second, *config.Timeout)
}

// TestConfigWithMaxConcurrent tests configuration with max concurrent requests.
func TestConfigWithMaxConcurrent(t *testing.T) {
	maxConcurrent := 50
	config := &Config{
		Target:        "localhost:50051",
		Method:        "test.Service/Method",
		Count:         1,
		MaxConcurrent: &maxConcurrent,
	}

	assert.NotNil(t, config.MaxConcurrent)
	assert.Equal(t, 50, *config.MaxConcurrent)
}

// TestConfigWithKeepalive tests configuration with keepalive enabled.
func TestConfigWithKeepalive(t *testing.T) {
	config := &Config{
		Target:          "localhost:50051",
		Method:          "test.Service/Method",
		Count:           1,
		EnableKeepalive: true,
	}

	assert.True(t, config.EnableKeepalive)
}

// TestConfigDataFormats tests different data format configurations.
func TestConfigDataFormats(t *testing.T) {
	tests := []struct {
		name   string
		config *Config
	}{
		{
			name: "String data",
			config: &Config{
				Target: "localhost:50051",
				Method: "test.Service/Method",
				Count:  1,
				Data:   `{"message": "hello"}`,
			},
		},
		{
			name: "Hex data",
			config: &Config{
				Target:  "localhost:50051",
				Method:  "test.Service/Method",
				Count:   1,
				DataHex: "48656c6c6f",
			},
		},
		{
			name: "Base64 data",
			config: &Config{
				Target:     "localhost:50051",
				Method:     "test.Service/Method",
				Count:      1,
				DataBase64: "SGVsbG8gV29ybGQ=",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, "localhost:50051", tt.config.Target)
			assert.Equal(t, "test.Service/Method", tt.config.Method)
			assert.Equal(t, 1, tt.config.Count)
		})
	}
}

// TestConfigEmpty tests empty configuration.
func TestConfigEmpty(t *testing.T) {
	config := &Config{}

	assert.Empty(t, config.Target)
	assert.Empty(t, config.Method)
	assert.Zero(t, config.Count)
	assert.Empty(t, config.Data)
	assert.Empty(t, config.DataHex)
	assert.Empty(t, config.DataBase64)
	assert.Nil(t, config.Metadata)
	assert.False(t, config.Insecure)
	assert.Nil(t, config.Timeout)
	assert.Nil(t, config.MaxConcurrent)
	assert.False(t, config.EnableKeepalive)
}

// TestConfigComplex tests a complex configuration with multiple options.
func TestConfigComplex(t *testing.T) {
	timeout := 60 * time.Second
	maxConcurrent := 100
	metadata := map[string]string{
		"auth": "token",
	}

	config := &Config{
		Target:          "grpc.example.com:443",
		Method:          "api.v1.Service/DoWork",
		Count:           1000,
		Data:            `{"id": 123, "action": "process"}`,
		Metadata:        metadata,
		Insecure:        false,
		Timeout:         &timeout,
		MaxConcurrent:   &maxConcurrent,
		EnableKeepalive: true,
	}

	assert.Equal(t, "grpc.example.com:443", config.Target)
	assert.Equal(t, "api.v1.Service/DoWork", config.Method)
	assert.Equal(t, 1000, config.Count)
	assert.NotEmpty(t, config.Data)
	assert.NotNil(t, config.Metadata)
	assert.False(t, config.Insecure)
	assert.Equal(t, 60*time.Second, *config.Timeout)
	assert.Equal(t, 100, *config.MaxConcurrent)
	assert.True(t, config.EnableKeepalive)
}
