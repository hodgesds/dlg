package snmp

import (
	"context"
	"fmt"

	config "github.com/hodgesds/dlg/config/snmp"
	"github.com/hodgesds/dlg/executor"
)

type snmpExecutor struct {
	debug bool
}

// New returns a new SNMP generator.
func New() executor.SNMP {
	return &snmpExecutor{debug: true}
}

// Execute implements the SNMP executor interface.
func (e *snmpExecutor) Execute(ctx context.Context, config *config.Config) error {
	snmp := config.SNMP()
	if err := snmp.Connect(); err != nil {
		return err
	}
	oids, err := snmp.Get(config.Oids)
	if err != nil {
	}
	if e.debug {
		fmt.Printf("%+v\n", oids)
	}

	return snmp.Conn.Close()
}
