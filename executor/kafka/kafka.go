package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/IBM/sarama"
	kafkaconfig "github.com/hodgesds/dlg/config/kafka"
)

type kafkaExecutor struct{}

// New returns a new Kafka executor.
func New() *kafkaExecutor {
	return &kafkaExecutor{}
}

// Execute implements the Kafka executor interface.
func (e *kafkaExecutor) Execute(ctx context.Context, config *kafkaconfig.Config) error {
	if err := config.Validate(); err != nil {
		return err
	}

	saramaCfg := sarama.NewConfig()
	saramaCfg.Producer.Return.Successes = true
	saramaCfg.Producer.Timeout = config.Timeout
	saramaCfg.Net.DialTimeout = config.Timeout
	saramaCfg.Net.ReadTimeout = config.Timeout
	saramaCfg.Net.WriteTimeout = config.Timeout

	if config.SASLEnabled {
		saramaCfg.Net.SASL.Enable = true
		saramaCfg.Net.SASL.User = config.SASLUsername
		saramaCfg.Net.SASL.Password = config.SASLPassword
	}

	if config.TLSEnabled {
		saramaCfg.Net.TLS.Enable = true
	}

	switch config.Operation {
	case "produce":
		producer, err := sarama.NewSyncProducer(config.Brokers, saramaCfg)
		if err != nil {
			return fmt.Errorf("failed to create producer: %w", err)
		}
		defer producer.Close()

		msg := &sarama.ProducerMessage{
			Topic: config.Topic,
			Value: sarama.StringEncoder(config.Message),
		}
		if config.Key != "" {
			msg.Key = sarama.StringEncoder(config.Key)
		}
		if config.Partition >= 0 {
			msg.Partition = config.Partition
		}

		_, _, err = producer.SendMessage(msg)
		return err

	case "consume":
		consumer, err := sarama.NewConsumer(config.Brokers, saramaCfg)
		if err != nil {
			return fmt.Errorf("failed to create consumer: %w", err)
		}
		defer consumer.Close()

		partition := config.Partition
		if partition < 0 {
			partition = 0
		}

		partitionConsumer, err := consumer.ConsumePartition(config.Topic, partition, sarama.OffsetNewest)
		if err != nil {
			return fmt.Errorf("failed to create partition consumer: %w", err)
		}
		defer partitionConsumer.Close()

		select {
		case <-partitionConsumer.Messages():
			return nil
		case err := <-partitionConsumer.Errors():
			return err
		case <-time.After(config.Timeout):
			return fmt.Errorf("consume timeout")
		case <-ctx.Done():
			return ctx.Err()
		}

	default:
		return fmt.Errorf("unknown operation: %s", config.Operation)
	}
}
