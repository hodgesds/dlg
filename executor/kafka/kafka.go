package kafka

import (
	"context"

	kafkaconfig "github.com/hodgesds/dlg/config/kafka"
	"github.com/hodgesds/dlg/executor"
)

type kafkaExecutor struct{}

// New returns a new Kafka executor.
func New() executor.Kafka {
	return &kafkaExecutor{}
}

// Execute implements the Kafka executor interface.
func (e *kafkaExecutor) Execute(ctx context.Context, config *kafkaconfig.Config) error {
	return nil
}
