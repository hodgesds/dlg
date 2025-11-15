package cassandra

import (
	"context"
	"time"

	cassandraconfig "github.com/hodgesds/dlg/config/cassandra"
	"github.com/hodgesds/dlg/executor"
	"github.com/gocql/gocql"
)

type cassandraExecutor struct{}

// New returns a new Cassandra executor.
func New() executor.Cassandra {
	return &cassandraExecutor{}
}

// Execute implements the Cassandra executor interface.
func (e *cassandraExecutor) Execute(ctx context.Context, config *cassandraconfig.Config) error {
	cluster := gocql.NewCluster(config.Hosts...)
	cluster.Keyspace = config.Keyspace

	if config.Username != "" && config.Password != "" {
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: config.Username,
			Password: config.Password,
		}
	}

	if config.ConnectTimeout != nil {
		cluster.ConnectTimeout = *config.ConnectTimeout
	}

	if config.Timeout != nil {
		cluster.Timeout = *config.Timeout
	}

	if config.NumConns != nil {
		cluster.NumConns = *config.NumConns
	}

	// Set consistency level
	switch config.Consistency {
	case cassandraconfig.ConsistencyAny:
		cluster.Consistency = gocql.Any
	case cassandraconfig.ConsistencyOne:
		cluster.Consistency = gocql.One
	case cassandraconfig.ConsistencyTwo:
		cluster.Consistency = gocql.Two
	case cassandraconfig.ConsistencyThree:
		cluster.Consistency = gocql.Three
	case cassandraconfig.ConsistencyQuorum:
		cluster.Consistency = gocql.Quorum
	case cassandraconfig.ConsistencyAll:
		cluster.Consistency = gocql.All
	case cassandraconfig.ConsistencyLocalQuorum:
		cluster.Consistency = gocql.LocalQuorum
	case cassandraconfig.ConsistencyEachQuorum:
		cluster.Consistency = gocql.EachQuorum
	case cassandraconfig.ConsistencyLocalOne:
		cluster.Consistency = gocql.LocalOne
	default:
		cluster.Consistency = gocql.Quorum
	}

	session, err := cluster.CreateSession()
	if err != nil {
		return err
	}
	defer session.Close()

	// Execute the configured number of iterations
	for i := 0; i < config.Count; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// Execute all queries in sequence
			for _, q := range config.Queries {
				query := session.Query(q.CQL, q.Values...)

				if q.Scan {
					// For SELECT queries, scan results
					iter := query.Iter()
					var row map[string]interface{}
					for iter.MapScan(row) {
						// Consume results
						row = make(map[string]interface{})
					}
					if err := iter.Close(); err != nil {
						return err
					}
				} else {
					// For INSERT/UPDATE/DELETE queries
					if err := query.WithContext(ctx).Exec(); err != nil {
						return err
					}
				}

				// Small delay between queries
				time.Sleep(10 * time.Millisecond)
			}
		}
	}

	return nil
}
