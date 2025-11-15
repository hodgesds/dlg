package executor

import (
	"context"

	"github.com/hodgesds/dlg/config"
	"github.com/hodgesds/dlg/config/arangodb"
	"github.com/hodgesds/dlg/config/cassandra"
	"github.com/hodgesds/dlg/config/clickhouse"
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

// Plan is used for executing a Plan.
type Plan interface {
	Execute(context.Context, *config.Plan) error
}

// Stage is used for executing a Stage.
type Stage interface {
	Execute(context.Context, *config.Stage) error
}

// DHCP4 is uesd for executing DHCP4.
type DHCP4 interface {
	Execute(context.Context, *dhcp4.Config) error
}

// DNS is used for executing a DNS.
type DNS interface {
	Execute(context.Context, *dns.Config) error
}

// ETCD is used for executing ETCD.
type ETCD interface {
	Execute(context.Context, *etcd.Config) error
}

// HTTP is used for executing HTTP.
type HTTP interface {
	Execute(context.Context, *http.Config) error
}

// LDAP is used for executing LDAP.
type LDAP interface {
	Execute(context.Context, *ldap.Config) error
}

// Memcache is used for executing Memcache.
type Memcache interface {
	Execute(context.Context, *memcache.Config) error
}

// Redis is used for executing Redis.
type Redis interface {
	Execute(context.Context, *redis.Config) error
}

// SQL is used for executing SQL.
type SQL interface {
	Execute(context.Context, *sql.Config) error
}

// SSH is used for executing SSH.
type SSH interface {
	Execute(context.Context, *ssh.Config) error
}

// SNMP is used for executing SNMP.
type SNMP interface {
	Execute(context.Context, *snmp.Config) error
}

// UDP is used for executing UDP.
type UDP interface {
	Execute(context.Context, *udp.Config) error
}

// Websocket is used for websocket.
type Websocket interface {
	Execute(context.Context, *websocket.Config) error
}

// GRPC is used for executing gRPC.
type GRPC interface {
	Execute(context.Context, *grpc.Config) error
}

// MongoDB is used for executing MongoDB.
type MongoDB interface {
	Execute(context.Context, *mongodb.Config) error
}

// MQTT is used for executing MQTT.
type MQTT interface {
	Execute(context.Context, *mqtt.Config) error
}

// GraphQL is used for executing GraphQL.
type GraphQL interface {
	Execute(context.Context, *graphql.Config) error
}

// Cassandra is used for executing Cassandra.
type Cassandra interface {
	Execute(context.Context, *cassandra.Config) error
}

// Elasticsearch is used for executing Elasticsearch.
type Elasticsearch interface {
	Execute(context.Context, *elasticsearch.Config) error
}

// InfluxDB is used for executing InfluxDB.
type InfluxDB interface {
	Execute(context.Context, *influxdb.Config) error
}

// Kafka is used for executing Kafka.
type Kafka interface {
	Execute(context.Context, *kafka.Config) error
}

// RabbitMQ is used for executing RabbitMQ.
type RabbitMQ interface {
	Execute(context.Context, *rabbitmq.Config) error
}

// NATS is used for executing NATS.
type NATS interface {
	Execute(context.Context, *nats.Config) error
}

// Pulsar is used for executing Apache Pulsar.
type Pulsar interface {
	Execute(context.Context, *pulsar.Config) error
}

// ScyllaDB is used for executing ScyllaDB.
type ScyllaDB interface {
	Execute(context.Context, *scylladb.Config) error
}

// CouchDB is used for executing CouchDB.
type CouchDB interface {
	Execute(context.Context, *couchdb.Config) error
}

// Neo4j is used for executing Neo4j.
type Neo4j interface {
	Execute(context.Context, *neo4j.Config) error
}

// ArangoDB is used for executing ArangoDB.
type ArangoDB interface {
	Execute(context.Context, *arangodb.Config) error
}

// FTP is used for executing FTP/SFTP.
type FTP interface {
	Execute(context.Context, *ftp.Config) error
}

// TCP is used for executing TCP.
type TCP interface {
	Execute(context.Context, *tcp.Config) error
}

// ICMP is used for executing ICMP/Ping.
type ICMP interface {
	Execute(context.Context, *icmp.Config) error
}

// NTP is used for executing NTP.
type NTP interface {
	Execute(context.Context, *ntp.Config) error
}

// TFTP is used for executing TFTP.
type TFTP interface {
	Execute(context.Context, *tftp.Config) error
}

// Telnet is used for executing Telnet.
type Telnet interface {
	Execute(context.Context, *telnet.Config) error
}

// Syslog is used for executing Syslog.
type Syslog interface {
	Execute(context.Context, *syslog.Config) error
}

// ClickHouse is used for executing ClickHouse.
type ClickHouse interface {
	Execute(context.Context, *clickhouse.Config) error
}
