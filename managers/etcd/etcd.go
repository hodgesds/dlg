package etcd

import (
	"github.com/hodgesds/dlg"
	etcdconfig "github.com/hodgesds/dlg/config/etcd"
	"go.etcd.io/etcd/clientv3"
)

type etcdManager struct {
	c *clientv3.Client
}

// NewManager returns a new manager.
func NewManager(config *etcdconfig.Config) (dlg.Manager, error) {
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   config.Endpoints,
		DialTimeout: config.DialTimeout,
	})
	if err != nil {
		return nil, err
	}

	return &etcdManager{c: c}, nil
}
