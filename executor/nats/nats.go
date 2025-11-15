package nats

import (
	"context"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	natsconfig "github.com/hodgesds/dlg/config/nats"
)

type natsExecutor struct{}

// New returns a new NATS executor.
func New() *natsExecutor {
	return &natsExecutor{}
}

// Execute implements the NATS executor interface.
func (e *natsExecutor) Execute(ctx context.Context, config *natsconfig.Config) error {
	if err := config.Validate(); err != nil {
		return err
	}

	var opts []nats.Option
	if config.Username != "" && config.Password != "" {
		opts = append(opts, nats.UserInfo(config.Username, config.Password))
	}
	if config.Token != "" {
		opts = append(opts, nats.Token(config.Token))
	}
	if config.TLSEnabled {
		opts = append(opts, nats.Secure())
	}
	opts = append(opts, nats.Timeout(config.Timeout))

	conn, err := nats.Connect(config.URL, opts...)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	defer conn.Close()

	switch config.Operation {
	case "publish":
		return conn.Publish(config.Subject, []byte(config.Message))

	case "subscribe":
		ch := make(chan *nats.Msg, 1)
		var sub *nats.Subscription
		if config.Queue != "" {
			sub, err = conn.QueueSubscribeSyncWithContext(ctx, config.Subject, config.Queue, ch)
		} else {
			sub, err = conn.SubscribeSyncWithContext(ctx, config.Subject, ch)
		}
		if err != nil {
			return err
		}
		defer sub.Unsubscribe()

		select {
		case <-ch:
			return nil
		case <-time.After(config.Timeout):
			return fmt.Errorf("subscribe timeout")
		case <-ctx.Done():
			return ctx.Err()
		}

	case "request":
		_, err := conn.RequestWithContext(ctx, config.Subject, []byte(config.Message))
		return err

	default:
		return fmt.Errorf("unknown operation: %s", config.Operation)
	}
}
