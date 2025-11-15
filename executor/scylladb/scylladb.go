package scylladb

import (
	"context"
	"fmt"

	"github.com/gocql/gocql"
	scyllaconfig "github.com/hodgesds/dlg/config/scylladb"
)

type scyllaExecutor struct{}

func New() *scyllaExecutor {
	return &scyllaExecutor{}
}

func (e *scyllaExecutor) Execute(ctx context.Context, config *scyllaconfig.Config) error {
	if err := config.Validate(); err != nil {
		return err
	}

	cluster := gocql.NewCluster(config.Hosts...)
	cluster.Keyspace = config.Keyspace
	cluster.Port = config.Port
	cluster.Timeout = config.Timeout

	if config.Username != "" && config.Password != "" {
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: config.Username,
			Password: config.Password,
		}
	}

	switch config.Consistency {
	case "one":
		cluster.Consistency = gocql.One
	case "quorum":
		cluster.Consistency = gocql.Quorum
	case "all":
		cluster.Consistency = gocql.All
	case "local_quorum":
		cluster.Consistency = gocql.LocalQuorum
	default:
		cluster.Consistency = gocql.Quorum
	}

	session, err := cluster.CreateSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}
	defer session.Close()

	return session.Query(config.Query, config.Values...).WithContext(ctx).Exec()
}
