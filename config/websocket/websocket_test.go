package websocket

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestConfigConn tests creating a WebSocket connection.
func TestConfigConn(t *testing.T) {
	// Create a test WebSocket server
	upgrader := websocket.Upgrader{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
	}))
	defer server.Close()

	// Convert http:// to ws://
	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")

	config := &Config{
		URL: wsURL,
	}

	conn, resp, err := config.Conn(context.Background())
	require.NoError(t, err)
	require.NotNil(t, conn)
	require.NotNil(t, resp)
	defer conn.Close()

	assert.Equal(t, http.StatusSwitchingProtocols, resp.StatusCode)
}

// TestConfigConnWithHeaders tests creating a WebSocket connection with custom headers.
func TestConfigConnWithHeaders(t *testing.T) {
	upgrader := websocket.Upgrader{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify custom header
		assert.Equal(t, "test-value", r.Header.Get("X-Custom-Header"))
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()
	}))
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")

	header := http.Header{}
	header.Set("X-Custom-Header", "test-value")

	config := &Config{
		URL:    wsURL,
		Header: header,
	}

	conn, resp, err := config.Conn(context.Background())
	require.NoError(t, err)
	require.NotNil(t, conn)
	require.NotNil(t, resp)
	defer conn.Close()
}

// TestConfigConnInvalidURL tests connection with invalid URL.
func TestConfigConnInvalidURL(t *testing.T) {
	config := &Config{
		URL: "invalid-url",
	}

	conn, _, err := config.Conn(context.Background())
	assert.Error(t, err)
	assert.Nil(t, conn)
}

// TestConfigConnContextCancellation tests context cancellation.
func TestConfigConnContextCancellation(t *testing.T) {
	config := &Config{
		URL: "ws://nonexistent.example.com",
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	conn, _, err := config.Conn(ctx)
	assert.Error(t, err)
	assert.Nil(t, conn)
}

// TestOpReadWrite tests Op structure.
func TestOpReadWrite(t *testing.T) {
	op := &Op{
		Read:  true,
		Write: "test message",
	}

	assert.True(t, op.Read)
	assert.Equal(t, "test message", op.Write)
}

// TestConfigWithOps tests Config with operations.
func TestConfigWithOps(t *testing.T) {
	ops := []*Op{
		{Read: false, Write: "hello"},
		{Read: true},
		{Read: false, Write: "world"},
	}

	config := &Config{
		URL: "ws://localhost:8080",
		Ops: ops,
	}

	assert.Len(t, config.Ops, 3)
	assert.Equal(t, "hello", config.Ops[0].Write)
	assert.True(t, config.Ops[1].Read)
	assert.Equal(t, "world", config.Ops[2].Write)
}
