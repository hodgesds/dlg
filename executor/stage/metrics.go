package stage

import (
	"github.com/prometheus/client_golang/prometheus"
)

// metrics contains metrics.
type metrics struct {
	ErrorsTotal        *prometheus.CounterVec
	ArangoDBTotal      *prometheus.CounterVec
	CassandraTotal     *prometheus.CounterVec
	CouchDBTotal       *prometheus.CounterVec
	DHCP4Total         *prometheus.CounterVec
	DNSTotal           *prometheus.CounterVec
	ElasticsearchTotal *prometheus.CounterVec
	ETCDTotal          *prometheus.CounterVec
	FTPTotal           *prometheus.CounterVec
	GraphQLTotal       *prometheus.CounterVec
	GRPCTotal          *prometheus.CounterVec
	HTTPTotal          *prometheus.CounterVec
	ICMPTotal          *prometheus.CounterVec
	InfluxDBTotal      *prometheus.CounterVec
	KafkaTotal         *prometheus.CounterVec
	LDAPTotal          *prometheus.CounterVec
	MemcacheTotal      *prometheus.CounterVec
	MongoDBTotal       *prometheus.CounterVec
	MQTTTotal          *prometheus.CounterVec
	NATSTotal          *prometheus.CounterVec
	Neo4jTotal         *prometheus.CounterVec
	NTPTotal           *prometheus.CounterVec
	PulsarTotal        *prometheus.CounterVec
	RabbitMQTotal      *prometheus.CounterVec
	RedisTotal         *prometheus.CounterVec
	ScyllaDBTotal      *prometheus.CounterVec
	SQLTotal           *prometheus.CounterVec
	SSHTotal           *prometheus.CounterVec
	SNMPTotal          *prometheus.CounterVec
	SyslogTotal        *prometheus.CounterVec
	TCPTotal           *prometheus.CounterVec
	TelnetTotal        *prometheus.CounterVec
	TFTPTotal          *prometheus.CounterVec
	UDPTotal           *prometheus.CounterVec
	WebsocketTotal     *prometheus.CounterVec
}

func newMetrics(reg *prometheus.Registry) (*metrics, error) {
	m := &metrics{
		ErrorsTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "errors_total",
			Help:      "The total number and type of errors that occurred while advertising.",
		}, []string{"stage"}),
		DHCP4Total: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "dhcp4_total",
			Help:      "The total number of DHCP4 stages.",
		}, []string{"stage"}),
		DNSTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "dns_total",
			Help:      "The total number of DNS stages.",
		}, []string{"stage"}),
		ETCDTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "etcd_total",
			Help:      "The total number of ETCD stages.",
		}, []string{"stage"}),
		HTTPTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "http_total",
			Help:      "The total number of HTTP stages.",
		}, []string{"stage"}),
		LDAPTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "ldap_total",
			Help:      "The total number of LDAP stages.",
		}, []string{"stage"}),
		MemcacheTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "memcache_total",
			Help:      "The total number of Memcache stages.",
		}, []string{"stage"}),
		RedisTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "redis_total",
			Help:      "The total number of Redis stages.",
		}, []string{"stage"}),
		SNMPTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "snmp_total",
			Help:      "The total number of SNMP stages.",
		}, []string{"stage"}),
		SSHTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "ssh_total",
			Help:      "The total number of SSH stages.",
		}, []string{"stage"}),
		SQLTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "sql_total",
			Help:      "The total number of SQL stages.",
		}, []string{"stage"}),
		UDPTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "udp_total",
			Help:      "The total number of UDP stages.",
		}, []string{"stage"}),
		WebsocketTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "websocket_total",
			Help:      "The total number of Websocket stages.",
		}, []string{"stage"}),
		GraphQLTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "graphql_total",
			Help:      "The total number of GraphQL stages.",
		}, []string{"stage"}),
		GRPCTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "grpc_total",
			Help:      "The total number of gRPC stages.",
		}, []string{"stage"}),
		MongoDBTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "mongodb_total",
			Help:      "The total number of MongoDB stages.",
		}, []string{"stage"}),
		MQTTTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "mqtt_total",
			Help:      "The total number of MQTT stages.",
		}, []string{"stage"}),
		CassandraTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "cassandra_total",
			Help:      "The total number of Cassandra stages.",
		}, []string{"stage"}),
		ElasticsearchTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "elasticsearch_total",
			Help:      "The total number of Elasticsearch stages.",
		}, []string{"stage"}),
		InfluxDBTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "influxdb_total",
			Help:      "The total number of InfluxDB stages.",
		}, []string{"stage"}),
		KafkaTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "kafka_total",
			Help:      "The total number of Kafka stages.",
		}, []string{"stage"}),
		RabbitMQTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "rabbitmq_total",
			Help:      "The total number of RabbitMQ stages.",
		}, []string{"stage"}),
		NATSTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "nats_total",
			Help:      "The total number of NATS stages.",
		}, []string{"stage"}),
		PulsarTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "pulsar_total",
			Help:      "The total number of Pulsar stages.",
		}, []string{"stage"}),
		ScyllaDBTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "scylladb_total",
			Help:      "The total number of ScyllaDB stages.",
		}, []string{"stage"}),
		CouchDBTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "couchdb_total",
			Help:      "The total number of CouchDB stages.",
		}, []string{"stage"}),
		Neo4jTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "neo4j_total",
			Help:      "The total number of Neo4j stages.",
		}, []string{"stage"}),
		ArangoDBTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "arangodb_total",
			Help:      "The total number of ArangoDB stages.",
		}, []string{"stage"}),
		FTPTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "ftp_total",
			Help:      "The total number of FTP/SFTP stages.",
		}, []string{"stage"}),
		TCPTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "tcp_total",
			Help:      "The total number of TCP stages.",
		}, []string{"stage"}),
		ICMPTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "icmp_total",
			Help:      "The total number of ICMP/Ping stages.",
		}, []string{"stage"}),
		NTPTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "ntp_total",
			Help:      "The total number of NTP stages.",
		}, []string{"stage"}),
		TFTPTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "tftp_total",
			Help:      "The total number of TFTP stages.",
		}, []string{"stage"}),
		TelnetTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "telnet_total",
			Help:      "The total number of Telnet stages.",
		}, []string{"stage"}),
		SyslogTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "syslog_total",
			Help:      "The total number of Syslog stages.",
		}, []string{"stage"}),
	}
	reg.MustRegister(
		m.ErrorsTotal,
		m.ArangoDBTotal,
		m.CassandraTotal,
		m.CouchDBTotal,
		m.DHCP4Total,
		m.DNSTotal,
		m.ElasticsearchTotal,
		m.ETCDTotal,
		m.FTPTotal,
		m.GraphQLTotal,
		m.GRPCTotal,
		m.HTTPTotal,
		m.ICMPTotal,
		m.InfluxDBTotal,
		m.KafkaTotal,
		m.LDAPTotal,
		m.MemcacheTotal,
		m.MongoDBTotal,
		m.MQTTTotal,
		m.NATSTotal,
		m.Neo4jTotal,
		m.NTPTotal,
		m.PulsarTotal,
		m.RabbitMQTotal,
		m.RedisTotal,
		m.ScyllaDBTotal,
		m.SQLTotal,
		m.SSHTotal,
		m.SNMPTotal,
		m.SyslogTotal,
		m.TCPTotal,
		m.TelnetTotal,
		m.TFTPTotal,
		m.UDPTotal,
		m.WebsocketTotal,
	)

	return m, nil
}
