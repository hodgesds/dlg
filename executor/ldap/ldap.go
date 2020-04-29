package ldap

import (
	"context"

	ldapconf "github.com/hodgesds/dlg/config/ldap"
	"github.com/hodgesds/dlg/executor"
)

type ldapExecutor struct{}

// New returns a new executor.
func New() executor.LDAP {
	return &ldapExecutor{}
}

// Execute implements the LDAP interface.
func (e *ldapExecutor) Execute(ctx context.Context, config *ldapconf.Config) error {
	return nil
}
