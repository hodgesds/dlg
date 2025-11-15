package config

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	ghttp "net/http"
	"sync"
	"time"

	"github.com/hodgesds/dlg/config/arangodb"
	"github.com/hodgesds/dlg/config/cassandra"
	"github.com/hodgesds/dlg/config/couchdb"
	"github.com/hodgesds/dlg/config/dhcp4"
	"github.com/hodgesds/dlg/config/dns"
	"github.com/hodgesds/dlg/config/elasticsearch"
	"github.com/hodgesds/dlg/config/etcd"
	"github.com/hodgesds/dlg/config/ftp"
	"github.com/hodgesds/dlg/config/graphql"
	"github.com/hodgesds/dlg/config/grpc"
	"github.com/hodgesds/dlg/config/http"
	"github.com/hodgesds/dlg/config/icmp"
	"github.com/hodgesds/dlg/config/influxdb"
	"github.com/hodgesds/dlg/config/kafka"
	"github.com/hodgesds/dlg/config/ldap"
	"github.com/hodgesds/dlg/config/memcache"
	"github.com/hodgesds/dlg/config/mongodb"
	"github.com/hodgesds/dlg/config/mqtt"
	"github.com/hodgesds/dlg/config/nats"
	"github.com/hodgesds/dlg/config/neo4j"
	"github.com/hodgesds/dlg/config/ntp"
	"github.com/hodgesds/dlg/config/pulsar"
	"github.com/hodgesds/dlg/config/rabbitmq"
	"github.com/hodgesds/dlg/config/redis"
	"github.com/hodgesds/dlg/config/scylladb"
	"github.com/hodgesds/dlg/config/snmp"
	"github.com/hodgesds/dlg/config/sql"
	"github.com/hodgesds/dlg/config/ssh"
	"github.com/hodgesds/dlg/config/syslog"
	"github.com/hodgesds/dlg/config/tcp"
	"github.com/hodgesds/dlg/config/telnet"
	"github.com/hodgesds/dlg/config/tftp"
	"github.com/hodgesds/dlg/config/udp"
	"github.com/hodgesds/dlg/config/websocket"
)

// ExecutionState is a execution state
type ExecutionState int

const (
	// Waiting is the waiting state.
	Waiting ExecutionState = iota
	// Running is when something is running.
	Running
	// Paused is when something is paused.
	Paused
	// Complete is when something is complete.
	Complete
)

// Config is used for running a load test.
type Config struct {
	Plan Plan `yaml:"plan"`
}

// Limiter is used for configuring ratelimiters.
type Limiter struct {
	Bytes      *int           `yaml:"bytes,omitempty"`
	Ops        *int           `yaml:"ops,omitempty"`
	Downsample bool           `yaml:"downsample"`
	SlowStart  *time.Duration `yaml:"slowStart,omitempty"`
}

// Distributed is configuration for distributed generators.
type Distributed struct {
	Manager string `yaml:"manager"`
}

// Plan is a load testing plan.
type Plan struct {
	mu    sync.RWMutex   `yaml:"-"`
	state ExecutionState `yaml:"-"`

	Name      string   `yaml:"name"`
	Executors int      `yaml:"executors"`
	Stages    []*Stage `yaml:"stages"`
	Tags      []string `yaml:"tags,omitempty"`

	Repeat   int            `yaml:"repeat,omitempty"`
	Duration *time.Duration `yaml:"duration,omitempty"`
	Start    *time.Time     `yaml:"start,omitempty"`
}

// WaitStart is used to wait until the start of the plan if configured.
func (p *Plan) WaitStart(ctx context.Context) error {
	if p.Start == nil {
		return nil
	}
	dur := p.Start.Sub(time.Now())
	select {
	case <-time.After(dur):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Validate is used to validate a Plan.
func (p *Plan) Validate() error {
	if len(p.Stages) == 0 {
		return errors.New("plan has no stages")
	}
	names := map[string]struct{}{}
	for _, stage := range p.Stages {
		if stage.validateName(names) {
			return fmt.Errorf("stage with duplicate name %q", stage.Name)
		}
		if err := stage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Stage is a part of a plan.
type Stage struct {
	// Internal fields for handling state.
	mu sync.RWMutex

	State ExecutionState `yaml:"state"`

	Name       string   `yaml:"name"`
	Tags       []string `yaml:"tags,omitempty"`
	Children   []*Stage `yaml:"children,omitempty"`
	Concurrent int      `yaml:"concurrent"` // if children should execute concurrent
	Repeat     int      `yaml:"repeat"`     // number of times to repeat the stage

	Duration *time.Duration `yaml:"duration,omitempty"`
	Timeout  *time.Duration `yaml:"timeout,omitempty"`

	// Stage types
	ArangoDB      *arangodb.Config      `yaml:"arangodb,omitempty"`
	Cassandra     *cassandra.Config     `yaml:"cassandra,omitempty"`
	CouchDB       *couchdb.Config       `yaml:"couchdb,omitempty"`
	DHCP4         *dhcp4.Config         `yaml:"dhcp4,omitempty"`
	DNS           *dns.Config           `yaml:"dns,omitempty"`
	Elasticsearch *elasticsearch.Config `yaml:"elasticsearch,omitempty"`
	ETCD          *etcd.Config          `yaml:"etcd,omitempty"`
	FTP           *ftp.Config           `yaml:"ftp,omitempty"`
	GraphQL       *graphql.Config       `yaml:"graphql,omitempty"`
	GRPC          *grpc.Config          `yaml:"grpc,omitempty"`
	HTTP          *http.Config          `yaml:"http,omitempty"`
	ICMP          *icmp.Config          `yaml:"icmp,omitempty"`
	InfluxDB      *influxdb.Config      `yaml:"influxdb,omitempty"`
	Kafka         *kafka.Config         `yaml:"kafka,omitempty"`
	LDAP          *ldap.Config          `yaml:"ldap,omitempty"`
	Memcache      *memcache.Config      `yaml:"memcache,omitempty"`
	MongoDB       *mongodb.Config       `yaml:"mongodb,omitempty"`
	MQTT          *mqtt.Config          `yaml:"mqtt,omitempty"`
	NATS          *nats.Config          `yaml:"nats,omitempty"`
	Neo4j         *neo4j.Config         `yaml:"neo4j,omitempty"`
	NTP           *ntp.Config           `yaml:"ntp,omitempty"`
	Pulsar        *pulsar.Config        `yaml:"pulsar,omitempty"`
	RabbitMQ      *rabbitmq.Config      `yaml:"rabbitmq,omitempty"`
	Redis         *redis.Config         `yaml:"redis,omitempty"`
	ScyllaDB      *scylladb.Config      `yaml:"scylladb,omitempty"`
	SNMP          *snmp.Config          `yaml:"snmp,omitempty"`
	SQL           *sql.Config           `yaml:"sql,omitempty"`
	SSH           *ssh.Config           `yaml:"ssh,omitempty"`
	Syslog        *syslog.Config        `yaml:"syslog,omitempty"`
	TCP           *tcp.Config           `yaml:"tcp,omitempty"`
	Telnet        *telnet.Config        `yaml:"telnet,omitempty"`
	TFTP          *tftp.Config          `yaml:"tftp,omitempty"`
	UDP           *udp.Config           `yaml:"udp,omitempty"`
	Websocket     *websocket.Config     `yaml:"websocket,omitempty"`
}

func (s *Stage) validateName(names map[string]struct{}) bool {
	if _, ok := names[s.Name]; ok {
		return true
	}
	names[s.Name] = struct{}{}
	for _, child := range s.Children {
		if child.validateName(names) {
			return true
		}
	}
	return false
}

// Validate is used to validate a stage.
func (s *Stage) Validate() error {
	if s.Repeat < 0 {
		return errors.New("invalid number of repeats")
	}
	stageTypes := 0
	if s.ArangoDB != nil {
		stageTypes++
	}
	if s.Cassandra != nil {
		stageTypes++
	}
	if s.CouchDB != nil {
		stageTypes++
	}
	if s.DHCP4 != nil {
		stageTypes++
	}
	if s.DNS != nil {
		stageTypes++
	}
	if s.Elasticsearch != nil {
		stageTypes++
	}
	if s.ETCD != nil {
		stageTypes++
	}
	if s.FTP != nil {
		stageTypes++
	}
	if s.GraphQL != nil {
		stageTypes++
	}
	if s.GRPC != nil {
		stageTypes++
	}
	if s.HTTP != nil {
		stageTypes++
	}
	if s.ICMP != nil {
		stageTypes++
	}
	if s.InfluxDB != nil {
		stageTypes++
	}
	if s.Kafka != nil {
		stageTypes++
	}
	if s.LDAP != nil {
		stageTypes++
	}
	if s.Memcache != nil {
		stageTypes++
	}
	if s.MongoDB != nil {
		stageTypes++
	}
	if s.MQTT != nil {
		stageTypes++
	}
	if s.NATS != nil {
		stageTypes++
	}
	if s.Neo4j != nil {
		stageTypes++
	}
	if s.NTP != nil {
		stageTypes++
	}
	if s.Pulsar != nil {
		stageTypes++
	}
	if s.RabbitMQ != nil {
		stageTypes++
	}
	if s.Redis != nil {
		stageTypes++
	}
	if s.ScyllaDB != nil {
		stageTypes++
	}
	if s.SNMP != nil {
		stageTypes++
	}
	if s.SQL != nil {
		stageTypes++
	}
	if s.SSH != nil {
		stageTypes++
	}
	if s.Syslog != nil {
		stageTypes++
	}
	if s.TCP != nil {
		stageTypes++
	}
	if s.Telnet != nil {
		stageTypes++
	}
	if s.TFTP != nil {
		stageTypes++
	}
	if s.UDP != nil {
		stageTypes++
	}
	if s.Websocket != nil {
		stageTypes++
	}
	if stageTypes == 0 && len(s.Children) == 0 {
		return errors.New("expected exactly one stage config value or at least one child")
	}
	return nil
}

// StageFrom is used to convert a HTTP request to a Stage config.
func StageFrom(r *ghttp.Request) (*Stage, error) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return &Stage{
		Name: r.URL.EscapedPath(),
		HTTP: &http.Config{
			Payload: http.Payload{
				URL:        r.URL.String(),
				Method:     r.Method,
				Header:     r.Header,
				BodyBase64: base64.StdEncoding.EncodeToString(body),
			},
		},
	}, nil
}
