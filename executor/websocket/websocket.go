package websocket

import (
	"context"
	"io/ioutil"

	"github.com/gorilla/websocket"
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
	conn, _, err := config.Conn(ctx)
	if err != nil {
		return err
	}
	for _, op := range config.Ops {
		if op.Read {
			_, r, err := conn.NextReader()
			if err != nil {
				return err
			}
			ioutil.ReadAll(r)
		}
		if op.Write != "" {
			w, err := conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return err
			}
			if _, err := w.Write([]byte(op.Write)); err != nil {
				return err
			}
		}
	}
	return nil
}
