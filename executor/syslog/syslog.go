package syslog

import (
	"context"
	"fmt"
	"log/syslog"
	"net"
	"os"
	"time"

	syslogconfig "github.com/hodgesds/dlg/config/syslog"
)

type syslogExecutor struct{}

// New returns a Syslog executor.
func New() *syslogExecutor {
	return &syslogExecutor{}
}

// Execute executes a Syslog operation.
func (e *syslogExecutor) Execute(ctx context.Context, config *syslogconfig.Config) error {
	if err := config.Validate(); err != nil {
		return err
	}

	if config.Format == "rfc5424" || config.Format == "rfc3164" {
		return e.executeCustom(ctx, config)
	}

	network := config.Protocol
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	priority := syslog.Priority(config.Facility*8 + config.Severity)

	writer, err := syslog.Dial(network, addr, priority, config.Tag)
	if err != nil {
		return fmt.Errorf("failed to connect to syslog server: %w", err)
	}
	defer writer.Close()

	switch config.Severity {
	case 0:
		return writer.Emerg(config.Message)
	case 1:
		return writer.Alert(config.Message)
	case 2:
		return writer.Crit(config.Message)
	case 3:
		return writer.Err(config.Message)
	case 4:
		return writer.Warning(config.Message)
	case 5:
		return writer.Notice(config.Message)
	case 6:
		return writer.Info(config.Message)
	case 7:
		return writer.Debug(config.Message)
	default:
		return writer.Info(config.Message)
	}
}

func (e *syslogExecutor) executeCustom(ctx context.Context, config *syslogconfig.Config) error {
	network := config.Protocol
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)

	dialer := &net.Dialer{Timeout: config.Timeout}
	conn, err := dialer.DialContext(ctx, network, addr)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	defer conn.Close()

	hostname := config.Hostname
	if hostname == "" {
		hostname, _ = os.Hostname()
	}

	var message string
	if config.Format == "rfc5424" {
		priority := config.Facility*8 + config.Severity
		timestamp := time.Now().Format(time.RFC3339)
		message = fmt.Sprintf("<%d>1 %s %s %s - - - %s\n",
			priority, timestamp, hostname, config.Tag, config.Message)
	} else {
		priority := config.Facility*8 + config.Severity
		timestamp := time.Now().Format("Jan 02 15:04:05")
		message = fmt.Sprintf("<%d>%s %s %s: %s\n",
			priority, timestamp, hostname, config.Tag, config.Message)
	}

	conn.SetWriteDeadline(time.Now().Add(config.Timeout))
	_, err = conn.Write([]byte(message))
	return err
}
