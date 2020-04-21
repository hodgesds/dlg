package executor

import (
	"context"

	"github.com/hodgesds/dlg/config"
	"github.com/hodgesds/dlg/config/etcd"
	"github.com/hodgesds/dlg/config/http"
	"github.com/hodgesds/dlg/config/redis"
	"github.com/hodgesds/dlg/config/sql"
	"github.com/hodgesds/dlg/config/udp"
)

// Plan is used for executing a Plan.
type Plan interface {
	Execute(context.Context, *config.Plan) error
}

// Stage is used for executing a Stage.
type Stage interface {
	Execute(context.Context, *config.Stage) error
}

// HTTP is used for executing HTTP.
type HTTP interface {
	Execute(context.Context, *http.Config) error
}

// Redis is used for executing Redis.
type Redis interface {
	Execute(context.Context, *redis.Config) error
}

// ETCD is used for executing ETCD.
type ETCD interface {
	Execute(context.Context, *etcd.Config) error
}

// SQL is used for executing SQL.
type SQL interface {
	Execute(context.Context, *sql.Config) error
}

// UDP is used for executing UDP.
type UDP interface {
	Execute(context.Context, *udp.Config) error
}
