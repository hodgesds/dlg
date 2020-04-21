package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// Config is a websocket config.
type Config struct {
	URL               string      `yaml:"url"`
	Header            http.Header `yaml:"header"`
	EnableCompression bool
}

// Conn returns a websocket connection.
func (c *Config) Conn() (*websocket.Conn, *http.Response, error) {
	return nil, nil, nil
}
