package memcache

import (
	"context"

	memcacheconf "github.com/hodgesds/dlg/config/memcache"
	"github.com/hodgesds/dlg/executor"
)

type memcacheExecutor struct{}

// New returns a new Memcache Executor.
func New() executor.Memcache {
	return &memcacheExecutor{}
}

// Execute implements the Memcache interface.
func (e *memcacheExecutor) Execute(ctx context.Context, config *memcacheconf.Config) error {
	return nil
}
