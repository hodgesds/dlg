package memcache

import (
	"context"

	memcacheconf "github.com/hodgesds/dlg/config/memcache"
	"github.com/hodgesds/dlg/executor"
	"github.com/rainycape/memcache"
)

type memcacheExecutor struct{}

// New returns a new Memcache Executor.
func New() executor.Memcache {
	return &memcacheExecutor{}
}

// Execute implements the Memcache interface.
func (e *memcacheExecutor) Execute(ctx context.Context, config *memcacheconf.Config) error {
	client, err := config.Client()
	if err != nil {
		return err
	}
	for _, op := range config.Ops {
		if op.Get != nil {
			_, err := client.Get(op.Get.Key)
			if err != nil {
				return err
			}
		}
		if op.Delete != nil {
			if err := client.Delete(op.Delete.Key); err != nil {
				return err
			}
		}
		if op.Set != nil {
			err := client.Set(&memcache.Item{
				Key:   op.Set.Key,
				Value: []byte(op.Set.Value),
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}
