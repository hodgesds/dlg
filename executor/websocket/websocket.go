package websocket

import (
	"context"

	websocketconf "github.com/hodgesds/dlg/config/websocket"
	"github.com/hodgesds/dlg/executor"
)

type websocketExecutor struct{}

// New returns a new Websocket interface.
func New() executor.Websocket {
	return &websocketExecutor{}
}

// Execute implements the Websocket interface.
func (e *websocketExecutor) Execute(ctx context.Context, config *websocketconf.Config) error {
	return nil
}
