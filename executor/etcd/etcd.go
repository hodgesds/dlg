package etcd

import (
	"context"

	etcdconf "github.com/hodgesds/dlg/config/etcd"
	"github.com/hodgesds/dlg/executor"
	"go.etcd.io/etcd/clientv3"
)

type etcdExecutor struct{}

// New returns a ETCD executor.
func New() executor.ETCD {
	return &etcdExecutor{}
}

// Execute implements the ETCD executor interface.
func (e *etcdExecutor) Execute(ctx context.Context, c *etcdconf.Config) error {
	client, err := clientv3.New(c.ClientConfig())
	if err != nil {
		return err
	}
	for _, kv := range c.KV {
		if err := e.execKv(ctx, client, kv); err != nil {
			return err
		}
	}
	return nil
}

func (e *etcdExecutor) execKv(
	ctx context.Context,
	client *clientv3.Client,
	kv *etcdconf.KV,
) error {
	if kv.Compact != nil {
		if _, err := client.Compact(ctx, kv.Compact.Rev); err != nil {
			return err
		}
	}
	if kv.Delete != nil {
		opts := kv.Delete.Opts.Opts()
		if _, err := client.Delete(ctx, kv.Delete.Key, opts...); err != nil {
			return err
		}
	}
	if kv.Get != nil {
		opts := kv.Get.Opts.Opts()
		_, err := client.Get(ctx, kv.Get.Key, opts...)
		if err != nil {
			return err
		}
	}
	if kv.Put != nil {
		opts := kv.Put.Opts.Opts()
		_, err := client.Put(ctx, kv.Put.Key, kv.Put.Value, opts...)
		if err != nil {
			return err
		}
	}
	return nil
}
