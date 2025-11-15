package rabbitmq

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	rabbitmqconfig "github.com/hodgesds/dlg/config/rabbitmq"
)

type rabbitmqExecutor struct{}

// New returns a new RabbitMQ executor.
func New() *rabbitmqExecutor {
	return &rabbitmqExecutor{}
}

// Execute implements the RabbitMQ executor interface.
func (e *rabbitmqExecutor) Execute(ctx context.Context, config *rabbitmqconfig.Config) error {
	if err := config.Validate(); err != nil {
		return err
	}

	conn, err := amqp.Dial(config.URL)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel: %w", err)
	}
	defer ch.Close()

	if config.Queue != "" {
		_, err = ch.QueueDeclare(config.Queue, config.Durable, config.AutoDelete, config.Exclusive, config.NoWait, nil)
		if err != nil {
			return fmt.Errorf("failed to declare queue: %w", err)
		}
	}

	switch config.Operation {
	case "publish":
		routingKey := config.RoutingKey
		if routingKey == "" {
			routingKey = config.Queue
		}
		return ch.PublishWithContext(ctx, config.Exchange, routingKey, false, false,
			amqp.Publishing{ContentType: "text/plain", Body: []byte(config.Message)})

	case "consume":
		msgs, err := ch.Consume(config.Queue, "", config.AutoAck, config.Exclusive, false, config.NoWait, nil)
		if err != nil {
			return fmt.Errorf("failed to consume: %w", err)
		}

		select {
		case msg := <-msgs:
			if !config.AutoAck {
				msg.Ack(false)
			}
			return nil
		case <-time.After(config.Timeout):
			return fmt.Errorf("consume timeout")
		case <-ctx.Done():
			return ctx.Err()
		}

	default:
		return fmt.Errorf("unknown operation: %s", config.Operation)
	}
}
