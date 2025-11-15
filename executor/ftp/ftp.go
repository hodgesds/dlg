package ftp

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	goftp "github.com/jlaffaye/go-ftp"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	ftpconfig "github.com/hodgesds/dlg/config/ftp"
)

type ftpExecutor struct{}

func New() *ftpExecutor {
	return &ftpExecutor{}
}

func (e *ftpExecutor) Execute(ctx context.Context, config *ftpconfig.Config) error {
	if err := config.Validate(); err != nil {
		return err
	}

	if config.Protocol == "ftp" {
		return e.executeFTP(ctx, config)
	}
	return e.executeSFTP(ctx, config)
}

func (e *ftpExecutor) executeFTP(ctx context.Context, config *ftpconfig.Config) error {
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	conn, err := goftp.Dial(addr, goftp.DialWithTimeout(config.Timeout))
	if err != nil {
		return err
	}
	defer conn.Quit()

	if err := conn.Login(config.Username, config.Password); err != nil {
		return err
	}

	switch config.Operation {
	case "list":
		_, err := conn.List(config.RemotePath)
		return err
	case "upload":
		var reader io.Reader
		if config.Data != "" {
			reader = strings.NewReader(config.Data)
		} else if config.LocalPath != "" {
			f, err := os.Open(config.LocalPath)
			if err != nil {
				return err
			}
			defer f.Close()
			reader = f
		} else {
			return fmt.Errorf("data or local_path required")
		}
		return conn.Stor(config.RemotePath, reader)
	case "download":
		resp, err := conn.Retr(config.RemotePath)
		if err != nil {
			return err
		}
		defer resp.Close()
		_, err = io.Copy(io.Discard, resp)
		return err
	case "delete":
		return conn.Delete(config.RemotePath)
	default:
		return fmt.Errorf("unknown operation: %s", config.Operation)
	}
}

func (e *ftpExecutor) executeSFTP(ctx context.Context, config *ftpconfig.Config) error {
	sshConfig := &ssh.ClientConfig{
		User:            config.Username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         config.Timeout,
	}

	if config.Password != "" {
		sshConfig.Auth = append(sshConfig.Auth, ssh.Password(config.Password))
	}

	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	sshClient, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return err
	}
	defer sshClient.Close()

	client, err := sftp.NewClient(sshClient)
	if err != nil {
		return err
	}
	defer client.Close()

	switch config.Operation {
	case "list":
		_, err := client.ReadDir(config.RemotePath)
		return err
	case "upload":
		dstFile, err := client.Create(config.RemotePath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		if config.Data != "" {
			_, err = io.Copy(dstFile, strings.NewReader(config.Data))
			return err
		} else if config.LocalPath != "" {
			srcFile, err := os.Open(config.LocalPath)
			if err != nil {
				return err
			}
			defer srcFile.Close()
			_, err = io.Copy(dstFile, srcFile)
			return err
		}
		return fmt.Errorf("data or local_path required")
	case "download":
		srcFile, err := client.Open(config.RemotePath)
		if err != nil {
			return err
		}
		defer srcFile.Close()
		_, err = io.Copy(io.Discard, srcFile)
		return err
	case "delete":
		return client.Remove(config.RemotePath)
	default:
		return fmt.Errorf("unknown operation: %s", config.Operation)
	}
}
