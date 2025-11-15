package tftp

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pin/tftp/v3"
	tftpconfig "github.com/hodgesds/dlg/config/tftp"
)

type tftpExecutor struct{}

// New returns a TFTP executor.
func New() *tftpExecutor {
	return &tftpExecutor{}
}

// Execute executes a TFTP operation.
func (e *tftpExecutor) Execute(ctx context.Context, config *tftpconfig.Config) error {
	if err := config.Validate(); err != nil {
		return err
	}

	client, err := tftp.NewClient(fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		return fmt.Errorf("failed to create TFTP client: %w", err)
	}

	client.SetTimeout(config.Timeout)
	client.SetRetries(config.Retries)
	client.SetBlockSize(config.BlockSize)

	switch config.Operation {
	case "read":
		reader, err := client.Receive(config.RemotePath, config.Mode)
		if err != nil {
			return fmt.Errorf("failed to receive file: %w", err)
		}
		defer reader.Close()

		if config.LocalPath != "" {
			f, err := os.Create(config.LocalPath)
			if err != nil {
				return err
			}
			defer f.Close()
			_, err = io.Copy(f, reader)
			return err
		}

		_, err = io.Copy(io.Discard, reader)
		return err

	case "write":
		writer, err := client.Send(config.RemotePath, config.Mode)
		if err != nil {
			return fmt.Errorf("failed to send file: %w", err)
		}
		defer writer.Close()

		if config.Data != "" {
			_, err = io.Copy(writer, strings.NewReader(config.Data))
			return err
		} else if config.LocalPath != "" {
			f, err := os.Open(config.LocalPath)
			if err != nil {
				return err
			}
			defer f.Close()
			_, err = io.Copy(writer, f)
			return err
		}
		return fmt.Errorf("data or local_path required for write operation")

	default:
		return fmt.Errorf("unknown operation: %s", config.Operation)
	}
}
