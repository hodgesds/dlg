package udp

import (
	"context"

	udpconf "github.com/hodgesds/dlg/config/udp"
	"github.com/hodgesds/dlg/executor"
)

type udpExecutor struct {
}

// New returns a UDP executor.
func New() executor.UDP {
	return &udpExecutor{}
}

// Execute implements the UDP executor interface.
func (e *udpExecutor) Execute(ctx context.Context, c *udpconf.Config) error {
	conn, err := c.Conn()
	if err != nil {
		return err
	}
	payload, err := c.GetPayload()
	if err != nil {
		return err
	}
	_, err = conn.Write(payload)
	return err
}
