package stage

import (
	"context"
	"errors"
	"sync"

	"github.com/hodgesds/dlg/config"
	"github.com/hodgesds/dlg/executor"
	"github.com/hodgesds/dlg/executor/cassandra"
	"github.com/hodgesds/dlg/executor/dhcp4"
	"github.com/hodgesds/dlg/executor/dns"
	"github.com/hodgesds/dlg/executor/elasticsearch"
	"github.com/hodgesds/dlg/executor/etcd"
	"github.com/hodgesds/dlg/executor/graphql"
	"github.com/hodgesds/dlg/executor/grpc"
	"github.com/hodgesds/dlg/executor/http"
	"github.com/hodgesds/dlg/executor/influxdb"
	"github.com/hodgesds/dlg/executor/ldap"
	"github.com/hodgesds/dlg/executor/memcache"
	"github.com/hodgesds/dlg/executor/mongodb"
	"github.com/hodgesds/dlg/executor/mqtt"
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

	cassandra     executor.Cassandra
	dhcp4         executor.DHCP4
	dns           executor.DNS
	elasticsearch executor.Elasticsearch
	etcd          executor.ETCD
	graphql       executor.GraphQL
	grpc          executor.GRPC
	http          executor.HTTP
	influxdb      executor.InfluxDB
	ldap          executor.LDAP
	memcache      executor.Memcache
	mongodb       executor.MongoDB
	mqtt          executor.MQTT
	redis         executor.Redis
	sql           executor.SQL
	snmp          executor.SNMP
	ssh           executor.SSH
	udp           executor.UDP
	websocket     executor.Websocket
}

// Params is used for configuring a Stage executor.
type Params struct {
	Registry *prometheus.Registry

	Cassandra     executor.Cassandra
	DHCP4         executor.DHCP4
	DNS           executor.DNS
	Elasticsearch executor.Elasticsearch
	ETCD          executor.ETCD
	GraphQL       executor.GraphQL
	GRPC          executor.GRPC
	HTTP          executor.HTTP
	InfluxDB      executor.InfluxDB
	LDAP          executor.LDAP
	Memcache      executor.Memcache
	MongoDB       executor.MongoDB
	MQTT          executor.MQTT
	Redis         executor.Redis
	SQL           executor.SQL
	SNMP          executor.SNMP
	SSH           executor.SSH
	UDP           executor.UDP
	Websocket     executor.Websocket
}

// New returns a new Stage executor.
func New(p Params) (executor.Stage, error) {
	metrics, err := newMetrics(p.Registry)
	if err != nil {
		return nil, err
	}
	return &stageExecutor{
		metrics:       metrics,
		cassandra:     p.Cassandra,
		dhcp4:         p.DHCP4,
		dns:           p.DNS,
		elasticsearch: p.Elasticsearch,
		etcd:          p.ETCD,
		graphql:       p.GraphQL,
		grpc:          p.GRPC,
		http:          p.HTTP,
		influxdb:      p.InfluxDB,
		ldap:          p.LDAP,
		memcache:      p.Memcache,
		mongodb:       p.MongoDB,
		mqtt:          p.MQTT,
		redis:         p.Redis,
		sql:           p.SQL,
		snmp:          p.SNMP,
		ssh:           p.SSH,
		udp:           p.UDP,
		websocket:     p.Websocket,
	}, nil
}

// Default returns a default Stage executor.
func Default(reg *prometheus.Registry) (executor.Stage, error) {
	metrics, err := newMetrics(reg)
	if err != nil {
		return nil, err
	}
	return &stageExecutor{
		metrics:       metrics,
		cassandra:     cassandra.New(),
		dhcp4:         dhcp4.New(),
		dns:           dns.New(),
		elasticsearch: elasticsearch.New(),
		etcd:          etcd.New(),
		graphql:       graphql.New(),
		grpc:          grpc.New(),
		http:          http.New(reg),
		influxdb:      influxdb.New(),
		ldap:          ldap.New(),
		memcache:      memcache.New(),
		mongodb:       mongodb.New(),
		mqtt:          mqtt.New(),
		redis:         redis.New(),
		sql:           sql.New(),
		snmp:          snmp.New(),
		ssh:           ssh.New(),
		udp:           udp.New(),
		websocket:     websocket.New(),
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
	if s.GraphQL != nil {
		if e.graphql == nil {
			return ErrNoStageExecutor
		}
		e.metrics.GraphQLTotal.With(prometheus.Labels{"stage": s.Name}).Add(1)
		if err := e.graphql.Execute(exCtx, s.GraphQL); err != nil {
			return err
		}
	}
	if s.GRPC != nil {
		if e.grpc == nil {
			return ErrNoStageExecutor
		}
		e.metrics.GRPCTotal.With(prometheus.Labels{"stage": s.Name}).Add(1)
		if err := e.grpc.Execute(exCtx, s.GRPC); err != nil {
			return err
		}
	}
	if s.MongoDB != nil {
		if e.mongodb == nil {
			return ErrNoStageExecutor
		}
		e.metrics.MongoDBTotal.With(prometheus.Labels{"stage": s.Name}).Add(1)
		if err := e.mongodb.Execute(exCtx, s.MongoDB); err != nil {
			return err
		}
	}
	if s.MQTT != nil {
		if e.mqtt == nil {
			return ErrNoStageExecutor
		}
		e.metrics.MQTTTotal.With(prometheus.Labels{"stage": s.Name}).Add(1)
		if err := e.mqtt.Execute(exCtx, s.MQTT); err != nil {
			return err
		}
	}
	if s.Cassandra != nil {
		if e.cassandra == nil {
			return ErrNoStageExecutor
		}
		e.metrics.CassandraTotal.With(prometheus.Labels{"stage": s.Name}).Add(1)
		if err := e.cassandra.Execute(exCtx, s.Cassandra); err != nil {
			return err
		}
	}
	if s.Elasticsearch != nil {
		if e.elasticsearch == nil {
			return ErrNoStageExecutor
		}
		e.metrics.ElasticsearchTotal.With(prometheus.Labels{"stage": s.Name}).Add(1)
		if err := e.elasticsearch.Execute(exCtx, s.Elasticsearch); err != nil {
			return err
		}
	}
	if s.InfluxDB != nil {
		if e.influxdb == nil {
			return ErrNoStageExecutor
		}
		e.metrics.InfluxDBTotal.With(prometheus.Labels{"stage": s.Name}).Add(1)
		if err := e.influxdb.Execute(exCtx, s.InfluxDB); err != nil {
			return err
		}
	}

	// Execute any children.
	if len(s.Children) > 1 && s.Concurrent > 0 {
		if err := e.execParallel(exCtx, s.Concurrent, s.Children); err != nil {
			return err
		}
		if s.Repeat > 0 {
			s.Repeat--
			return e.Execute(ctx, s)
		}
		return e.execDuration(ctx, s)
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
	return e.execDuration(ctx, s)
}

// execDuration is used to execute a stage until the context is complete, this
// is used mainly for duration based tests.
func (e *stageExecutor) execDuration(ctx context.Context, stage *config.Stage) error {
	if stage.Duration == nil {
		return nil
	}
	select {
	// Just check the context to see if it is done, the initial
	// execution should set the context for the right timeout based
	// on the duration
	case <-ctx.Done():
		return nil
	default:
		return e.Execute(ctx, stage)
	}
}

func (e *stageExecutor) execParallel(ctx context.Context, concurrent int, stages []*config.Stage) error {
	var (
		wg   sync.WaitGroup
		mu   sync.Mutex
		err  error
		stop = make(chan struct{})
		work = make(chan *config.Stage)
	)

	for i := 0; i < concurrent; i++ {
		go func(work chan *config.Stage) {
			select {
			case <-stop:
				return
			case stage := <-work:
				wg.Add(1)
				err2 := e.Execute(ctx, stage)
				if err2 != nil {
					e.metrics.ErrorsTotal.With(prometheus.Labels{"stage": stage.Name}).Add(1)
					mu.Lock()
					err = multierr.Append(err, err2)
					mu.Unlock()
				}
				wg.Done()
			}
		}(work)
	}

	for _, stage := range stages {
		work <- stage
	}
	wg.Wait()
	close(stop)
	return err
}
