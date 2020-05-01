package snmp

import (
	"context"

	config "github.com/hodgesds/dlg/config/snmp"
)

type snmpExecutor struct{}

func (e *snmpExecutor) Execute(ctx context.Context, config *config.Config) error {
	return nil
}
