package ssh

import (
	"context"

	sshconf "github.com/hodgesds/dlg/config/ssh"
	"github.com/hodgesds/dlg/executor"
)

type sshExecutor struct{}

// New returns a new SSH.
func New() executor.SSH {
	return &sshExecutor{}
}

// Execute implements the SSH interface.
func (e *sshExecutor) Execute(ctx context.Context, config *sshconf.Config) error {
	return nil
}
