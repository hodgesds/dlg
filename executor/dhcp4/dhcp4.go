package dhcp4

import (
	"context"

	dhcp4config "github.com/hodgesds/dlg/config/dhcp4"
	"github.com/hodgesds/dlg/executor"
)

type dhcp4Executor struct{}

// New returns a new Executor.
func New() executor.DHCP4 {
	return &dhcp4Executor{}
}

// Execute implements the DHCP4 interface.
func (e *dhcp4Executor) Execute(ctx context.Context, config *dhcp4config.Config) error {
	return nil
}
