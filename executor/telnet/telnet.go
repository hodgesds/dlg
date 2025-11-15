package telnet

import (
	"context"
	"fmt"
	"time"

	"github.com/ziutek/telnet"
	telnetconfig "github.com/hodgesds/dlg/config/telnet"
)

type telnetExecutor struct{}

// New returns a Telnet executor.
func New() *telnetExecutor {
	return &telnetExecutor{}
}

// Execute executes a Telnet operation.
func (e *telnetExecutor) Execute(ctx context.Context, config *telnetconfig.Config) error {
	if err := config.Validate(); err != nil {
		return err
	}

	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	conn, err := telnet.DialTimeout("tcp", addr, config.Timeout)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(config.Timeout))
	conn.SetWriteDeadline(time.Now().Add(config.Timeout))

	if config.Username != "" {
		conn.SkipUntil("login:")
		conn.Write([]byte(config.Username + "\n"))

		if config.Password != "" {
			conn.SkipUntil("Password:")
			conn.Write([]byte(config.Password + "\n"))
			conn.SkipUntil(config.ExpectPrompt)
		}
	}

	if config.ConnectOnly {
		return nil
	}

	for _, cmd := range config.Commands {
		conn.SetWriteDeadline(time.Now().Add(config.Timeout))
		_, err := conn.Write([]byte(cmd + "\n"))
		if err != nil {
			return fmt.Errorf("failed to send command '%s': %w", cmd, err)
		}

		conn.SetReadDeadline(time.Now().Add(config.Timeout))
		conn.SkipUntil(config.ExpectPrompt)
	}

	return nil
}
