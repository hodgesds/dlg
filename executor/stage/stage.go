package stage

import (
	"context"
	"errors"
	"sync"

	"github.com/hodgesds/dlg/config"
	"github.com/hodgesds/dlg/executor"
	"github.com/hodgesds/dlg/executor/dhcp4"
	"github.com/hodgesds/dlg/executor/dns"
	"github.com/hodgesds/dlg/executor/etcd"
	"github.com/hodgesds/dlg/executor/http"
	"github.com/hodgesds/dlg/executor/kafka"
	"github.com/hodgesds/dlg/executor/ldap"
	"github.com/hodgesds/dlg/executor/memcache"
	"github.com/hodgesds/dlg/executor/redis"
	"github.com/hodgesds/dlg/executor/snmp"
	"github.com/hodgesds/dlg/executor/sql"
	"github.com/hodgesds/dlg/executor/ssh"
	"github.com/hodgesds/dlg/executor/udp"
	"github.com/hodgesds/dlg/executor/websocket"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/multierr"
)

var (
	// ErrNoStageExecutor is returned when a stage has no configured
	// executor.
	ErrNoStageExecutor = errors.New("no executor for stage")
)

type stageExecutor struct {
	metrics *metrics

	dhcp4     executor.DHCP4
	dns       executor.DNS
	etcd      executor.ETCD
	http      executor.HTTP
	kafka     executor.Kafka
	ldap      executor.LDAP
	memcache  executor.Memcache
	redis     executor.Redis
	sql       executor.SQL
	snmp      executor.SNMP
	ssh       executor.SSH
	udp       executor.UDP
	websocket executor.Websocket
}

// Params is used for configuring a Stage executor.
type Params struct {
	Registry *prometheus.Registry

	DHCP4     executor.DHCP4
	DNS       executor.DNS
	ETCD      executor.ETCD
	HTTP      executor.HTTP
	Kafka     executor.Kafka
	LDAP      executor.LDAP
	Memcache  executor.Memcache
	Redis     executor.Redis
	SQL       executor.SQL
	SNMP      executor.SNMP
	SSH       executor.SSH
	UDP       executor.UDP
	Websocket executor.Websocket
}

// New returns a new Stage executor.
func New(p Params) (executor.Stage, error) {
	metrics, err := newMetrics(p.Registry)
	if err != nil {
		return nil, err
	}
	return &stageExecutor{
		metrics:   metrics,
		dhcp4:     p.DHCP4,
		dns:       p.DNS,
		etcd:      p.ETCD,
		http:      p.HTTP,
		kafka:     p.Kafka,
		ldap:      p.LDAP,
		memcache:  p.Memcache,
		redis:     p.Redis,
		sql:       p.SQL,
		snmp:      p.SNMP,
		ssh:       p.SSH,
		udp:       p.UDP,
		websocket: p.Websocket,
	}, nil
}

// Default returns a default Stage executor.
func Default(reg *prometheus.Registry) (executor.Stage, error) {
	metrics, err := newMetrics(reg)
	if err != nil {
		return nil, err
	}
	return &stageExecutor{
		metrics:   metrics,
		dhcp4:     dhcp4.New(),
		dns:       dns.New(),
		etcd:      etcd.New(),
		http:      http.New(reg),
		kafka:     kafka.New(),
		ldap:      ldap.New(),
		memcache:  memcache.New(),
		redis:     redis.New(),
		sql:       sql.New(),
		snmp:      snmp.New(),
		ssh:       ssh.New(),
		udp:       udp.New(),
		websocket: websocket.New(),
	}, nil
}

// Execute implements the Stage interface.
func (e *stageExecutor) Execute(ctx context.Context, s *config.Stage) error {
	if err := s.Validate(); err != nil {
		return err
	}

	var (
		// exCtx is the context for this execution, since a stage can
		// be repeated multiple times with a timeout a copy of the
		// original context must be used.
		exCtx  context.Context
		cancel func()
	)
	if s.Timeout != nil {
		exCtx, cancel = context.WithTimeout(ctx, *s.Timeout)
	} else {
		exCtx, cancel = context.WithCancel(ctx)
	}
	defer cancel()

	if s.DHCP4 != nil {
		if e.dhcp4 == nil {
			return ErrNoStageExecutor
		}
		e.metrics.DHCP4Total.With(prometheus.Labels{"stage": s.Name}).Add(1)
		if err := e.dhcp4.Execute(exCtx, s.DHCP4); err != nil {
			return err
		}
	}
	if s.DNS != nil {
		if e.dns == nil {
			return ErrNoStageExecutor
		}
		e.metrics.DNSTotal.With(prometheus.Labels{"stage": s.Name}).Add(1)
		if err := e.dns.Execute(exCtx, s.DNS); err != nil {
			return err
		}
	}
	if s.ETCD != nil {
		if e.etcd == nil {
			return ErrNoStageExecutor
		}
		e.metrics.ETCDTotal.With(prometheus.Labels{"stage": s.Name}).Add(1)
		if err := e.etcd.Execute(exCtx, s.ETCD); err != nil {
			return err
		}
	}
	if s.HTTP != nil {
		if e.http == nil {
			return ErrNoStageExecutor
		}
		e.metrics.HTTPTotal.With(prometheus.Labels{"stage": s.Name}).Add(1)
		if err := e.http.Execute(exCtx, s.HTTP); err != nil {
			return err
		}
	}
	if s.Kafka != nil {
		if e.kafka == nil {
			return ErrNoStageExecutor
		}
		e.metrics.KafkaTotal.With(prometheus.Labels{"stage": s.Name}).Add(1)
		if err := e.kafka.Execute(exCtx, s.Kafka); err != nil {
			return err
		}
	}
	if s.LDAP != nil {
		if e.ldap == nil {
			return ErrNoStageExecutor
		}
		e.metrics.LDAPTotal.With(prometheus.Labels{"stage": s.Name}).Add(1)
		if err := e.ldap.Execute(exCtx, s.LDAP); err != nil {
			return err
		}
	}
	if s.Memcache != nil {
		if e.memcache == nil {
			return ErrNoStageExecutor
		}
		e.metrics.MemcacheTotal.With(prometheus.Labels{"stage": s.Name}).Add(1)
		if err := e.memcache.Execute(exCtx, s.Memcache); err != nil {
			return err
		}
	}
	if s.Redis != nil {
		if e.redis == nil {
			return ErrNoStageExecutor
		}
		e.metrics.RedisTotal.With(prometheus.Labels{"stage": s.Name}).Add(1)
		if err := e.redis.Execute(exCtx, s.Redis); err != nil {
			return err
		}
	}
	if s.SNMP != nil {
		if e.snmp == nil {
			return ErrNoStageExecutor
		}
		e.metrics.SNMPTotal.With(prometheus.Labels{"stage": s.Name}).Add(1)
		if err := e.snmp.Execute(exCtx, s.SNMP); err != nil {
			return err
		}
	}
	if s.SSH != nil {
		if e.ssh == nil {
			return ErrNoStageExecutor
		}
		e.metrics.SSHTotal.With(prometheus.Labels{"stage": s.Name}).Add(1)
		if err := e.ssh.Execute(exCtx, s.SSH); err != nil {
			return err
		}
	}
	if s.SQL != nil {
		if e.sql == nil {
			return ErrNoStageExecutor
		}
		e.metrics.SQLTotal.With(prometheus.Labels{"stage": s.Name}).Add(1)
		if err := e.sql.Execute(exCtx, s.SQL); err != nil {
			return err
		}
	}
	if s.UDP != nil {
		if e.udp == nil {
			return ErrNoStageExecutor
		}
		e.metrics.UDPTotal.With(prometheus.Labels{"stage": s.Name}).Add(1)
		if err := e.udp.Execute(exCtx, s.UDP); err != nil {
			return err
		}
	}
	if s.Websocket != nil {
		if e.websocket == nil {
			return ErrNoStageExecutor
		}
		e.metrics.WebsocketTotal.With(prometheus.Labels{"stage": s.Name}).Add(1)
		if err := e.websocket.Execute(exCtx, s.Websocket); err != nil {
			return err
		}
	}

	// Execute any children.
	if len(s.Children) > 1 && s.Concurrent {
		if err := e.execParallel(exCtx, s.Children); err != nil {
			return err
		}
		if s.Repeat > 0 {
			s.Repeat--
			return e.Execute(ctx, s)
		}
	}

	for _, child := range s.Children {
		if err := e.Execute(exCtx, child); err != nil {
			e.metrics.ErrorsTotal.With(prometheus.Labels{"stage": child.Name}).Add(1)
			return err
		}
	}

	if s.Repeat > 0 {
		s.Repeat--
		return e.Execute(ctx, s)
	}

	return nil
}

func (e *stageExecutor) execParallel(ctx context.Context, stages []*config.Stage) error {
	var (
		wg  sync.WaitGroup
		mu  sync.Mutex
		err error
	)
	for _, stage := range stages {
		wg.Add(1)
		go func(stage *config.Stage) {
			err2 := e.Execute(ctx, stage)
			if err2 != nil {
				e.metrics.ErrorsTotal.With(prometheus.Labels{"stage": stage.Name}).Add(1)
				mu.Lock()
				err = multierr.Append(err, err2)
				mu.Unlock()
			}
			wg.Done()
		}(stage)
	}
	wg.Wait()
	return err
}
