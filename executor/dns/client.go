package dns

import (
	"context"

	dnsconfig "github.com/hodgesds/dlg/config/dns"
	"github.com/hodgesds/dlg/executor"
)

type dnsExecutor struct{}

// New returns a new DNS executor.
func New() executor.DNS {
	return &dnsExecutor{}
}

// Execute implements the DNS executor interface.
func (e *dnsExecutor) Execute(ctx context.Context, config *dnsconfig.Config) error {
	return nil
}
