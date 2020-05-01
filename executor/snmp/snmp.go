package snmp

import (
	"context"
	"fmt"

	config "github.com/hodgesds/dlg/config/snmp"
)

type snmpExecutor struct {
	debug bool
}

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
