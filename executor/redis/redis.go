package redis

import (
	"context"

	v7 "github.com/go-redis/redis/v7"
	redisconf "github.com/hodgesds/dlg/config/redis"
)

type redisExecutor struct{}

// New returns a new Redis executor.
func New() executor.Redis {
	return &redisExecutor{}
}

// Execute implements the Redis interface.
func (e *httpExecutor) Execute(ctx context.Context, conf *redisconf.Config) error {
	return conf.Execute(ctx)
}
