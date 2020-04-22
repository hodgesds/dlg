package redis

import (
	"context"

	redisconf "github.com/hodgesds/dlg/config/redis"
	"github.com/hodgesds/dlg/executor"
)

type redisExecutor struct{}

// New returns a new Redis executor.
func New() executor.Redis {
	return &redisExecutor{}
}

// Execute implements the Redis interface.
func (e *redisExecutor) Execute(ctx context.Context, conf *redisconf.Config) error {
	return conf.Execute(ctx)
}
