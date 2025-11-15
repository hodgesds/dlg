package icmp

import (
	"context"
	"fmt"

	"github.com/go-ping/ping"
	icmpconfig "github.com/hodgesds/dlg/config/icmp"
)

type icmpExecutor struct{}

// New returns an ICMP/Ping executor.
func New() *icmpExecutor {
	return &icmpExecutor{}
}

// Execute executes an ICMP/Ping operation.
func (e *icmpExecutor) Execute(ctx context.Context, config *icmpconfig.Config) error {
	if err := config.Validate(); err != nil {
		return err
	}

	pinger, err := ping.NewPinger(config.Host)
	if err != nil {
		return fmt.Errorf("failed to create pinger: %w", err)
	}

	pinger.Count = config.Count
	pinger.Size = config.Size
	pinger.Timeout = config.Timeout
	pinger.Interval = config.Interval
	pinger.SetPrivileged(config.Privileged)

	done := make(chan error, 1)
	go func() {
		done <- pinger.Run()
	}()

	select {
	case err := <-done:
		if err != nil {
			return fmt.Errorf("ping failed: %w", err)
		}
		stats := pinger.Statistics()
		if stats.PacketsRecv == 0 {
			return fmt.Errorf("no packets received from %s", config.Host)
		}
		return nil
	case <-ctx.Done():
		pinger.Stop()
		return ctx.Err()
	}
}
