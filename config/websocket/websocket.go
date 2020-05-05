package websocket

import (
	"context"
	"net/http"

	"github.com/gorilla/websocket"
)

// Config is a websocket config.
type Config struct {
	URL    string      `yaml:"url"`
	Header http.Header `yaml:"header"`
	Ops    []*Op       `yaml:"ops"`
}

// Conn returns a websocket connection.
func (c *Config) Conn(ctx context.Context) (*websocket.Conn, *http.Response, error) {
	d := &websocket.Dialer{}
	return d.DialContext(ctx, c.URL, c.Header)
}

// Op is a operation.
type Op struct {
	Read  bool   `yaml:"read"`
	Write string `yaml:"write,omitempty"`
}
