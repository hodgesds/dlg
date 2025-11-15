package tcp

import (
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"net"
	"time"

	tcpconfig "github.com/hodgesds/dlg/config/tcp"
)

type tcpExecutor struct{}

func New() *tcpExecutor {
	return &tcpExecutor{}
}

func (e *tcpExecutor) Execute(ctx context.Context, config *tcpconfig.Config) error {
	if err := config.Validate(); err != nil {
		return err
	}

	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	dialer := &net.Dialer{
		Timeout:   config.Timeout,
		KeepAlive: -1,
	}
	if config.KeepAlive {
		dialer.KeepAlive = 30 * time.Second
	}

	conn, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	if tcpConn, ok := conn.(*net.TCPConn); ok {
		if config.NoDelay {
			tcpConn.SetNoDelay(true)
		}
		if config.KeepAlive {
			tcpConn.SetKeepAlive(true)
			tcpConn.SetKeepAlivePeriod(30 * time.Second)
		}
	}

	switch config.Operation {
	case "connect":
		return nil

	case "send":
		var data []byte
		if config.BinaryData != "" {
			data, err = hex.DecodeString(config.BinaryData)
			if err != nil {
				return err
			}
		} else {
			data = []byte(config.Data)
		}
		if len(data) == 0 {
			return fmt.Errorf("no data to send")
		}
		conn.SetWriteDeadline(time.Now().Add(config.WriteTimeout))
		_, err = conn.Write(data)
		return err

	case "send_receive":
		var data []byte
		if config.BinaryData != "" {
			data, err = hex.DecodeString(config.BinaryData)
			if err != nil {
				return err
			}
		} else {
			data = []byte(config.Data)
		}
		if len(data) == 0 {
			return fmt.Errorf("no data to send")
		}

		conn.SetWriteDeadline(time.Now().Add(config.WriteTimeout))
		if _, err := conn.Write(data); err != nil {
			return err
		}

		buf := make([]byte, 4096)
		conn.SetReadDeadline(time.Now().Add(config.ReadTimeout))
		_, err := conn.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		return nil

	default:
		return fmt.Errorf("unknown operation: %s", config.Operation)
	}
}
