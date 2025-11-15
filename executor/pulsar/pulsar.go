package pulsar

import (
	"context"
	"fmt"

	"github.com/apache/pulsar-client-go/pulsar"
	pulsarconfig "github.com/hodgesds/dlg/config/pulsar"
)

type pulsarExecutor struct{}

func New() *pulsarExecutor {
	return &pulsarExecutor{}
}

func (e *pulsarExecutor) Execute(ctx context.Context, config *pulsarconfig.Config) error {
	if err := config.Validate(); err != nil {
		return err
	}

	opts := pulsar.ClientOptions{
		URL:              config.URL,
		OperationTimeout: config.Timeout,
	}
	if config.AuthToken != "" {
		opts.Authentication = pulsar.NewAuthenticationToken(config.AuthToken)
	}

	client, err := pulsar.NewClient(opts)
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}
	defer client.Close()

	switch config.Operation {
	case "produce":
		producer, err := client.CreateProducer(pulsar.ProducerOptions{Topic: config.Topic})
		if err != nil {
			return err
		}
		defer producer.Close()

		_, err = producer.Send(ctx, &pulsar.ProducerMessage{Payload: []byte(config.Message)})
		return err

	case "consume":
		subscriptionType := pulsar.Exclusive
		switch config.SubscriptionType {
		case "shared":
			subscriptionType = pulsar.Shared
		case "failover":
			subscriptionType = pulsar.Failover
		case "key_shared":
			subscriptionType = pulsar.KeyShared
		}

		consumer, err := client.Subscribe(pulsar.ConsumerOptions{
			Topic:            config.Topic,
			SubscriptionName: config.SubscriptionName,
			Type:             subscriptionType,
		})
		if err != nil {
			return err
		}
		defer consumer.Close()

		msg, err := consumer.Receive(ctx)
		if err != nil {
			return err
		}
		consumer.Ack(msg)
		return nil

	default:
		return fmt.Errorf("unknown operation: %s", config.Operation)
	}
}
