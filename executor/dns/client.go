package dns

import (
	"context"

	dnsconfig "github.com/hodgesds/dlg/config/dns"
	"github.com/hodgesds/dlg/executor"
	"github.com/miekg/dns"
)

type dnsExecutor struct{}

// New returns a new DNS executor.
func New() executor.DNS {
	return &dnsExecutor{}
}

// Execute implements the DNS executor interface.
func (e *dnsExecutor) Execute(ctx context.Context, config *dnsconfig.Config) error {
	c := new(dns.Client)
	m1 := new(dns.Msg)
	_, _, err := c.Exchange(m1, "127.0.0.1:53")
	return err
}
