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
	client, err := config.SSHClient(ctx)
	if err != nil {
		return err
	}
	if config.Cmd != nil {
		s, err := client.NewSession()
		if err != nil {
			return err
		}
		if err := s.Run(*config.Cmd); err != nil {
			return err
		}
	}
	return nil
}
