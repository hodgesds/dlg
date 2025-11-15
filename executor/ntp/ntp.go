package ntp

import (
	"context"
	"fmt"

	"github.com/beevik/ntp"
	ntpconfig "github.com/hodgesds/dlg/config/ntp"
)

type ntpExecutor struct{}

// New returns an NTP executor.
func New() *ntpExecutor {
	return &ntpExecutor{}
}

// Execute executes an NTP operation.
func (e *ntpExecutor) Execute(ctx context.Context, config *ntpconfig.Config) error {
	if err := config.Validate(); err != nil {
		return err
	}

	options := ntp.QueryOptions{
		Timeout: config.Timeout,
		Version: config.Version,
		Port:    config.Port,
	}

	response, err := ntp.QueryWithOptions(config.Host, options)
	if err != nil {
		return fmt.Errorf("NTP query failed: %w", err)
	}

	if response == nil {
		return fmt.Errorf("NTP query returned nil response")
	}

	if err := response.Validate(); err != nil {
		return fmt.Errorf("NTP response validation failed: %w", err)
	}

	return nil
}
