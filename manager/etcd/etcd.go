package etcd

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/hodgesds/dlg"
	"github.com/hodgesds/dlg/config"
	etcdconfig "github.com/hodgesds/dlg/config/etcd"
	"github.com/hodgesds/dlg/executor"
	"go.etcd.io/etcd/clientv3"
	"gopkg.in/yaml.v2"
)

type manager struct {
	mu       sync.RWMutex
	planExec executor.Plan
	c        *clientv3.Client
	stop     chan struct{}
	stopped  bool
}

// NewManager returns a new manager.
func NewManager(
	config *etcdconfig.Config,
	planExec executor.Plan,
) (dlg.Manager, error) {
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   config.Endpoints,
		DialTimeout: config.DialTimeout,
	})
	if err != nil {
		return nil, err
	}

	m := &manager{
		c:        c,
		planExec: planExec,
		stop:     make(chan struct{}),
	}
	return m, m.start()
}

func (m *manager) start() error {
	go m.watch()
	return nil
}

func (m *manager) watch() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	watch := m.c.Watch(ctx, "plan")
	for {
		select {
		case res := <-watch:
			if res.Canceled {
				log.Println(res.Err().Error())
			}
			if !res.Created {
				m.handleEvents(res.Events...)
			}
		case <-m.stop:
			return
		}
	}
}

func (m *manager) handleEvents(events ...*clientv3.Event) error {
	return nil
}

// Get implements the Manager interface.
func (m *manager) Get(ctx context.Context, name string) (*config.Plan, error) {
	res, err := m.c.Get(ctx, fmt.Sprintf("plan-%s", name))
	if err != nil {
		return nil, err
	}
	for _, kv := range res.Kvs {
		var p config.Plan
		if err := yaml.Unmarshal(kv.Value, &p); err != nil {
			return nil, err
		}
		return &p, nil
	}
	return nil, fmt.Errorf("no plan with name: %q", name)
}

// Add implements the Manager interface.
func (m *manager) Add(ctx context.Context, plan *config.Plan) error {
	b, err := yaml.Marshal(plan)
	if err != nil {
		return err
	}
	_, err = m.c.Put(ctx, fmt.Sprintf("plan-%s", plan.Name), string(b))
	return err
}

// Delete implments the Manager interface.
func (m *manager) Delete(ctx context.Context, name string) error {
	_, err := m.c.Delete(ctx, fmt.Sprintf("plan-%s", name), clientv3.WithPrevKV())
	return err
}

// Plans implements the Manager interface.
func (m *manager) Plans(ctx context.Context) ([]*config.Plan, error) {
	res, err := m.c.Get(ctx, "plan-")
	if err != nil {
		return nil, err
	}
	plans := make([]*config.Plan, 0, len(res.Kvs))
	for _, kv := range res.Kvs {
		var p config.Plan
		if err := yaml.Unmarshal(kv.Value, &p); err != nil {
			return nil, err
		}
		plans = append(plans, &p)
	}
	return plans, nil
}

// Execute implements the Executor interface.
func (m *manager) Execute(ctx context.Context, plan *config.Plan) error {
	return m.planExec.Execute(ctx, plan)
}
