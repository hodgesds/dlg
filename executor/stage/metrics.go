package stage

import (
	"github.com/prometheus/client_golang/prometheus"
)

// metrics contains metrics.
type metrics struct {
	ErrorsTotal    *prometheus.CounterVec
	DHCP4Total     *prometheus.CounterVec
	DNSTotal       *prometheus.CounterVec
	ETCDTotal      *prometheus.CounterVec
	HTTPTotal      *prometheus.CounterVec
	KafkaTotal     *prometheus.CounterVec
	LDAPTotal      *prometheus.CounterVec
	MemcacheTotal  *prometheus.CounterVec
	RedisTotal     *prometheus.CounterVec
	SQLTotal       *prometheus.CounterVec
	SSHTotal       *prometheus.CounterVec
	SNMPTotal      *prometheus.CounterVec
	UDPTotal       *prometheus.CounterVec
	WebsocketTotal *prometheus.CounterVec
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
		KafkaTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "kafka_total",
			Help:      "The total number of Kafka stages.",
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
	}
	return m, reg.Register(m.ErrorsTotal)
}
