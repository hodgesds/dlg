package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	httpconf "github.com/hodgesds/dlg/config/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNew tests creating a new HTTP executor.
func TestNew(t *testing.T) {
	reg := prometheus.NewRegistry()
	exec := New(reg)
	require.NotNil(t, exec)
}

// TestExecuteSimpleGET tests executing a simple GET request.
func TestExecuteSimpleGET(t *testing.T) {
	// Create a test server
	var requestCount int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&requestCount, 1)
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	// Create executor
	reg := prometheus.NewRegistry()
	exec := New(reg)

	// Execute request
	conf := &httpconf.Config{
		Count: 1,
		Payload: httpconf.Payload{
			URL:    server.URL,
			Method: "GET",
		},
	}

	err := exec.Execute(context.Background(), conf)
	assert.NoError(t, err)
	assert.Equal(t, int32(1), atomic.LoadInt32(&requestCount))
}

// TestExecuteMultipleRequests tests executing multiple concurrent requests.
func TestExecuteMultipleRequests(t *testing.T) {
	var requestCount int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&requestCount, 1)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	reg := prometheus.NewRegistry()
	exec := New(reg)

	conf := &httpconf.Config{
		Count: 10,
		Payload: httpconf.Payload{
			URL:    server.URL,
			Method: "GET",
		},
	}

	err := exec.Execute(context.Background(), conf)
	assert.NoError(t, err)
	assert.Equal(t, int32(10), atomic.LoadInt32(&requestCount))
}

// TestExecutePOST tests executing a POST request with body.
func TestExecutePOST(t *testing.T) {
	expectedBody := []byte("test data")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	reg := prometheus.NewRegistry()
	exec := New(reg)

	conf := &httpconf.Config{
		Count: 1,
		Payload: httpconf.Payload{
			URL:    server.URL,
			Method: "POST",
			Body:   expectedBody,
		},
	}

	err := exec.Execute(context.Background(), conf)
	assert.NoError(t, err)
}

// TestExecuteWithHeaders tests executing a request with custom headers.
func TestExecuteWithHeaders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "test-value", r.Header.Get("X-Custom-Header"))
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	reg := prometheus.NewRegistry()
	exec := New(reg)

	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("X-Custom-Header", "test-value")

	conf := &httpconf.Config{
		Count: 1,
		Payload: httpconf.Payload{
			URL:    server.URL,
			Method: "GET",
			Header: headers,
		},
	}

	err := exec.Execute(context.Background(), conf)
	assert.NoError(t, err)
}

// TestExecuteWithMaxIdleConns tests setting max idle connections.
func TestExecuteWithMaxIdleConns(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	reg := prometheus.NewRegistry()
	exec := New(reg)

	maxIdleConns := 5
	conf := &httpconf.Config{
		Count:        1,
		MaxIdleConns: &maxIdleConns,
		Payload: httpconf.Payload{
			URL:    server.URL,
			Method: "GET",
		},
	}

	err := exec.Execute(context.Background(), conf)
	assert.NoError(t, err)
}

// TestExecuteWithMaxConns tests setting max connections per host.
func TestExecuteWithMaxConns(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	reg := prometheus.NewRegistry()
	exec := New(reg)

	maxConns := 10
	conf := &httpconf.Config{
		Count:    1,
		MaxConns: &maxConns,
		Payload: httpconf.Payload{
			URL:    server.URL,
			Method: "GET",
		},
	}

	err := exec.Execute(context.Background(), conf)
	assert.NoError(t, err)
}

// TestExecuteContextCancellation tests that context cancellation is handled.
func TestExecuteContextCancellation(t *testing.T) {
	// Create a server that never responds
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Never respond
		select {}
	}))
	defer server.Close()

	reg := prometheus.NewRegistry()
	exec := New(reg)

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	conf := &httpconf.Config{
		Count: 1,
		Payload: httpconf.Payload{
			URL:    server.URL,
			Method: "GET",
		},
	}

	// Should handle cancellation gracefully
	_ = exec.Execute(ctx, conf)
}

// TestExecuteMetrics tests that Prometheus metrics are collected.
func TestExecuteMetrics(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	reg := prometheus.NewRegistry()
	exec := New(reg)

	conf := &httpconf.Config{
		Count: 1,
		Payload: httpconf.Payload{
			URL:    server.URL,
			Method: "GET",
		},
	}

	err := exec.Execute(context.Background(), conf)
	require.NoError(t, err)

	// Verify metrics were registered
	metrics, err := reg.Gather()
	require.NoError(t, err)
	assert.NotEmpty(t, metrics)
}

// TestExecuteServerError tests handling of server errors.
func TestExecuteServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	reg := prometheus.NewRegistry()
	exec := New(reg)

	conf := &httpconf.Config{
		Count: 1,
		Payload: httpconf.Payload{
			URL:    server.URL,
			Method: "GET",
		},
	}

	// Should not error on server errors (only network errors)
	err := exec.Execute(context.Background(), conf)
	assert.NoError(t, err)
}

// TestExecuteInvalidURL tests handling of invalid URLs.
func TestExecuteInvalidURL(t *testing.T) {
	reg := prometheus.NewRegistry()
	exec := New(reg)

	conf := &httpconf.Config{
		Count: 1,
		Payload: httpconf.Payload{
			URL:    "://invalid-url",
			Method: "GET",
		},
	}

	err := exec.Execute(context.Background(), conf)
	assert.Error(t, err)
}
