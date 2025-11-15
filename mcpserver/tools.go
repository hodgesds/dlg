// Copyright © 2025 Daniel Hodges <hodges.daniel.scott@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mcpserver

import (
	"bytes"
	"context"
	"fmt"
	"os"

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
	"github.com/hodgesds/dlg/executor"
	arangodbexec "github.com/hodgesds/dlg/executor/arangodb"
	cassandraexec "github.com/hodgesds/dlg/executor/cassandra"
	clickhouseexec "github.com/hodgesds/dlg/executor/clickhouse"
	couchdbexec "github.com/hodgesds/dlg/executor/couchdb"
	dhcp4exec "github.com/hodgesds/dlg/executor/dhcp4"
	dnsexec "github.com/hodgesds/dlg/executor/dns"
	elasticsearchexec "github.com/hodgesds/dlg/executor/elasticsearch"
	etcdexec "github.com/hodgesds/dlg/executor/etcd"
	ftpexec "github.com/hodgesds/dlg/executor/ftp"
	graphqlexec "github.com/hodgesds/dlg/executor/graphql"
	grpcexec "github.com/hodgesds/dlg/executor/grpc"
	httpexec "github.com/hodgesds/dlg/executor/http"
	icmpexec "github.com/hodgesds/dlg/executor/icmp"
	influxdbexec "github.com/hodgesds/dlg/executor/influxdb"
	kafkaexec "github.com/hodgesds/dlg/executor/kafka"
	ldapexec "github.com/hodgesds/dlg/executor/ldap"
	memcacheexec "github.com/hodgesds/dlg/executor/memcache"
	mongodbexec "github.com/hodgesds/dlg/executor/mongodb"
	mqttexec "github.com/hodgesds/dlg/executor/mqtt"
	natsexec "github.com/hodgesds/dlg/executor/nats"
	neo4jexec "github.com/hodgesds/dlg/executor/neo4j"
	ntpexec "github.com/hodgesds/dlg/executor/ntp"
	pulsarexec "github.com/hodgesds/dlg/executor/pulsar"
	rabbitmqexec "github.com/hodgesds/dlg/executor/rabbitmq"
	redisexec "github.com/hodgesds/dlg/executor/redis"
	scylladbexec "github.com/hodgesds/dlg/executor/scylladb"
	snmpexec "github.com/hodgesds/dlg/executor/snmp"
	sqlexec "github.com/hodgesds/dlg/executor/sql"
	sshexec "github.com/hodgesds/dlg/executor/ssh"
	stageexec "github.com/hodgesds/dlg/executor/stage"
	syslogexec "github.com/hodgesds/dlg/executor/syslog"
	tcpexec "github.com/hodgesds/dlg/executor/tcp"
	telnetexec "github.com/hodgesds/dlg/executor/telnet"
	tftpexec "github.com/hodgesds/dlg/executor/tftp"
	udpexec "github.com/hodgesds/dlg/executor/udp"
	websocketexec "github.com/hodgesds/dlg/executor/websocket"
	"github.com/hodgesds/dlg/util"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/prometheus/client_golang/prometheus"
	"gopkg.in/yaml.v2"
)

// HTTPLoadTestInput defines input parameters for HTTP load testing
type HTTPLoadTestInput struct {
	URL          string            `json:"url" jsonschema:"required,description=Target URL to test"`
	Method       string            `json:"method" jsonschema:"description=HTTP method (GET, POST, PUT, DELETE, etc.)"`
	Count        int               `json:"count" jsonschema:"description=Number of requests to make"`
	Concurrent   int               `json:"concurrent" jsonschema:"description=Number of concurrent requests"`
	MaxConns     int               `json:"max_conns" jsonschema:"description=Maximum number of connections"`
	MaxIdleConns int               `json:"max_idle_conns" jsonschema:"description=Maximum number of idle connections"`
	Headers      map[string]string `json:"headers" jsonschema:"description=HTTP headers to include"`
	Body         string            `json:"body" jsonschema:"description=Request body"`
	Timeout      int               `json:"timeout_seconds" jsonschema:"description=Timeout in seconds"`
}

// HTTPLoadTestOutput defines the output of HTTP load testing
type HTTPLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// RedisLoadTestInput defines input parameters for Redis load testing
type RedisLoadTestInput struct {
	Addr     string   `json:"addr" jsonschema:"required,description=Redis server address (host:port)"`
	DB       int      `json:"db" jsonschema:"description=Redis database number"`
	Password string   `json:"password" jsonschema:"description=Redis password"`
	Count    int      `json:"count" jsonschema:"description=Number of operations to perform"`
	Keys     []string `json:"keys" jsonschema:"description=Keys to operate on"`
}

// RedisLoadTestOutput defines the output of Redis load testing
type RedisLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// MongoDBLoadTestInput defines input parameters for MongoDB load testing
type MongoDBLoadTestInput struct {
	URI        string `json:"uri" jsonschema:"required,description=MongoDB connection URI"`
	Database   string `json:"database" jsonschema:"required,description=Database name"`
	Collection string `json:"collection" jsonschema:"required,description=Collection name"`
	Operation  string `json:"operation" jsonschema:"description=Operation type (find, insert, update, delete, aggregate)"`
	Count      int    `json:"count" jsonschema:"description=Number of operations to perform"`
}

// MongoDBLoadTestOutput defines the output of MongoDB load testing
type MongoDBLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// ClickHouseLoadTestInput defines input parameters for ClickHouse load testing
type ClickHouseLoadTestInput struct {
	DSN       string `json:"dsn" jsonschema:"required,description=ClickHouse DSN connection string"`
	Database  string `json:"database" jsonschema:"required,description=Database name"`
	Table     string `json:"table" jsonschema:"description=Table name"`
	Operation string `json:"operation" jsonschema:"description=Operation type (insert, select, batch_insert, count, optimize, create_table)"`
	Query     string `json:"query" jsonschema:"description=Custom SQL query to execute"`
	Count     int    `json:"count" jsonschema:"description=Number of operations to perform"`
}

// ClickHouseLoadTestOutput defines the output of ClickHouse load testing
type ClickHouseLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// PostgresLoadTestInput defines input parameters for PostgreSQL load testing
type PostgresLoadTestInput struct {
	DSN       string `json:"dsn" jsonschema:"required,description=PostgreSQL DSN connection string"`
	Query     string `json:"query" jsonschema:"description=SQL query to execute"`
	Count     int    `json:"count" jsonschema:"description=Number of queries to execute"`
	MaxConns  int    `json:"max_conns" jsonschema:"description=Maximum number of connections"`
	MaxIdle   int    `json:"max_idle" jsonschema:"description=Maximum number of idle connections"`
	Operation string `json:"operation" jsonschema:"description=Operation type (select, insert, update, delete)"`
	Table     string `json:"table" jsonschema:"description=Table name for operations"`
}

// PostgresLoadTestOutput defines the output of PostgreSQL load testing
type PostgresLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// WebSocketLoadTestInput defines input parameters for WebSocket load testing
type WebSocketLoadTestInput struct {
	URL     string `json:"url" jsonschema:"required,description=WebSocket URL to connect to"`
	Count   int    `json:"count" jsonschema:"description=Number of messages to send"`
	Message string `json:"message" jsonschema:"description=Message to send"`
}

// WebSocketLoadTestOutput defines the output of WebSocket load testing
type WebSocketLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// GRPCLoadTestInput defines input parameters for gRPC load testing
type GRPCLoadTestInput struct {
	Target string `json:"target" jsonschema:"required,description=gRPC server target (host:port)"`
	Method string `json:"method" jsonschema:"description=gRPC method to call"`
	Count  int    `json:"count" jsonschema:"description=Number of requests to make"`
	TLS    bool   `json:"tls" jsonschema:"description=Use TLS connection"`
}

// GRPCLoadTestOutput defines the output of gRPC load testing
type GRPCLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// RunLoadPlanInput defines input parameters for running a YAML load plan
type RunLoadPlanInput struct {
	YAMLConfig string `json:"yaml_config" jsonschema:"required,description=YAML configuration for the load test plan"`
}

// RunLoadPlanOutput defines the output of running a load plan
type RunLoadPlanOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// ArangoDBLoadTestInput defines input parameters for ArangoDB load testing
type ArangoDBLoadTestInput struct {
	Endpoints []string `json:"endpoints" jsonschema:"required,description=ArangoDB endpoints"`
	Database  string   `json:"database" jsonschema:"required,description=Database name"`
	Count     int      `json:"count" jsonschema:"description=Number of operations to perform"`
}

// ArangoDBLoadTestOutput defines the output of ArangoDB load testing
type ArangoDBLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// CassandraLoadTestInput defines input parameters for Cassandra load testing
type CassandraLoadTestInput struct {
	Hosts    []string `json:"hosts" jsonschema:"required,description=Cassandra host addresses"`
	Keyspace string   `json:"keyspace" jsonschema:"required,description=Keyspace name"`
	Count    int      `json:"count" jsonschema:"description=Number of operations to perform"`
}

// CassandraLoadTestOutput defines the output of Cassandra load testing
type CassandraLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// CouchDBLoadTestInput defines input parameters for CouchDB load testing
type CouchDBLoadTestInput struct {
	URL      string `json:"url" jsonschema:"required,description=CouchDB URL"`
	Database string `json:"database" jsonschema:"required,description=Database name"`
	Count    int    `json:"count" jsonschema:"description=Number of operations to perform"`
}

// CouchDBLoadTestOutput defines the output of CouchDB load testing
type CouchDBLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// DHCP4LoadTestInput defines input parameters for DHCP4 load testing
type DHCP4LoadTestInput struct {
	Server string `json:"server" jsonschema:"required,description=DHCP server address"`
	Count  int    `json:"count" jsonschema:"description=Number of DHCP requests"`
}

// DHCP4LoadTestOutput defines the output of DHCP4 load testing
type DHCP4LoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// DNSLoadTestInput defines input parameters for DNS load testing
type DNSLoadTestInput struct {
	Server string `json:"server" jsonschema:"required,description=DNS server address"`
	Domain string `json:"domain" jsonschema:"required,description=Domain to query"`
	Count  int    `json:"count" jsonschema:"description=Number of DNS queries"`
}

// DNSLoadTestOutput defines the output of DNS load testing
type DNSLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// ElasticsearchLoadTestInput defines input parameters for Elasticsearch load testing
type ElasticsearchLoadTestInput struct {
	Addresses []string `json:"addresses" jsonschema:"required,description=Elasticsearch node addresses"`
	Index     string   `json:"index" jsonschema:"required,description=Index name"`
	Count     int      `json:"count" jsonschema:"description=Number of operations to perform"`
}

// ElasticsearchLoadTestOutput defines the output of Elasticsearch load testing
type ElasticsearchLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// ETCDLoadTestInput defines input parameters for ETCD load testing
type ETCDLoadTestInput struct {
	Endpoints []string `json:"endpoints" jsonschema:"required,description=ETCD endpoints"`
	Count     int      `json:"count" jsonschema:"description=Number of operations to perform"`
}

// ETCDLoadTestOutput defines the output of ETCD load testing
type ETCDLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// FTPLoadTestInput defines input parameters for FTP load testing
type FTPLoadTestInput struct {
	Host  string `json:"host" jsonschema:"required,description=FTP server host"`
	Port  int    `json:"port" jsonschema:"description=FTP server port"`
	Count int    `json:"count" jsonschema:"description=Number of operations to perform"`
}

// FTPLoadTestOutput defines the output of FTP load testing
type FTPLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// GraphQLLoadTestInput defines input parameters for GraphQL load testing
type GraphQLLoadTestInput struct {
	URL   string `json:"url" jsonschema:"required,description=GraphQL endpoint URL"`
	Query string `json:"query" jsonschema:"required,description=GraphQL query"`
	Count int    `json:"count" jsonschema:"description=Number of queries to execute"`
}

// GraphQLLoadTestOutput defines the output of GraphQL load testing
type GraphQLLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// ICMPLoadTestInput defines input parameters for ICMP/Ping load testing
type ICMPLoadTestInput struct {
	Host  string `json:"host" jsonschema:"required,description=Host to ping"`
	Count int    `json:"count" jsonschema:"description=Number of ping requests"`
}

// ICMPLoadTestOutput defines the output of ICMP load testing
type ICMPLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// InfluxDBLoadTestInput defines input parameters for InfluxDB load testing
type InfluxDBLoadTestInput struct {
	URL    string `json:"url" jsonschema:"required,description=InfluxDB URL"`
	Bucket string `json:"bucket" jsonschema:"required,description=Bucket name"`
	Count  int    `json:"count" jsonschema:"description=Number of operations to perform"`
}

// InfluxDBLoadTestOutput defines the output of InfluxDB load testing
type InfluxDBLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// KafkaLoadTestInput defines input parameters for Kafka load testing
type KafkaLoadTestInput struct {
	Brokers []string `json:"brokers" jsonschema:"required,description=Kafka broker addresses"`
	Topic   string   `json:"topic" jsonschema:"required,description=Topic name"`
	Count   int      `json:"count" jsonschema:"description=Number of messages to produce/consume"`
}

// KafkaLoadTestOutput defines the output of Kafka load testing
type KafkaLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// LDAPLoadTestInput defines input parameters for LDAP load testing
type LDAPLoadTestInput struct {
	Server string `json:"server" jsonschema:"required,description=LDAP server address"`
	BaseDN string `json:"base_dn" jsonschema:"required,description=Base DN for searches"`
	Count  int    `json:"count" jsonschema:"description=Number of LDAP operations"`
}

// LDAPLoadTestOutput defines the output of LDAP load testing
type LDAPLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// MemcacheLoadTestInput defines input parameters for Memcache load testing
type MemcacheLoadTestInput struct {
	Servers []string `json:"servers" jsonschema:"required,description=Memcache server addresses"`
	Count   int      `json:"count" jsonschema:"description=Number of operations to perform"`
}

// MemcacheLoadTestOutput defines the output of Memcache load testing
type MemcacheLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// MQTTLoadTestInput defines input parameters for MQTT load testing
type MQTTLoadTestInput struct {
	Broker string `json:"broker" jsonschema:"required,description=MQTT broker address"`
	Topic  string `json:"topic" jsonschema:"required,description=Topic name"`
	Count  int    `json:"count" jsonschema:"description=Number of messages to publish"`
}

// MQTTLoadTestOutput defines the output of MQTT load testing
type MQTTLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// NATSLoadTestInput defines input parameters for NATS load testing
type NATSLoadTestInput struct {
	URL     string `json:"url" jsonschema:"required,description=NATS server URL"`
	Subject string `json:"subject" jsonschema:"required,description=Subject name"`
	Count   int    `json:"count" jsonschema:"description=Number of messages to publish"`
}

// NATSLoadTestOutput defines the output of NATS load testing
type NATSLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// Neo4jLoadTestInput defines input parameters for Neo4j load testing
type Neo4jLoadTestInput struct {
	URI   string `json:"uri" jsonschema:"required,description=Neo4j connection URI"`
	Query string `json:"query" jsonschema:"description=Cypher query to execute"`
	Count int    `json:"count" jsonschema:"description=Number of queries to execute"`
}

// Neo4jLoadTestOutput defines the output of Neo4j load testing
type Neo4jLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// NTPLoadTestInput defines input parameters for NTP load testing
type NTPLoadTestInput struct {
	Server string `json:"server" jsonschema:"required,description=NTP server address"`
	Count  int    `json:"count" jsonschema:"description=Number of NTP requests"`
}

// NTPLoadTestOutput defines the output of NTP load testing
type NTPLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// PulsarLoadTestInput defines input parameters for Pulsar load testing
type PulsarLoadTestInput struct {
	URL   string `json:"url" jsonschema:"required,description=Pulsar service URL"`
	Topic string `json:"topic" jsonschema:"required,description=Topic name"`
	Count int    `json:"count" jsonschema:"description=Number of messages to produce"`
}

// PulsarLoadTestOutput defines the output of Pulsar load testing
type PulsarLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// RabbitMQLoadTestInput defines input parameters for RabbitMQ load testing
type RabbitMQLoadTestInput struct {
	URL   string `json:"url" jsonschema:"required,description=RabbitMQ connection URL"`
	Queue string `json:"queue" jsonschema:"required,description=Queue name"`
	Count int    `json:"count" jsonschema:"description=Number of messages to publish"`
}

// RabbitMQLoadTestOutput defines the output of RabbitMQ load testing
type RabbitMQLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// ScyllaDBLoadTestInput defines input parameters for ScyllaDB load testing
type ScyllaDBLoadTestInput struct {
	Hosts    []string `json:"hosts" jsonschema:"required,description=ScyllaDB host addresses"`
	Keyspace string   `json:"keyspace" jsonschema:"required,description=Keyspace name"`
	Count    int      `json:"count" jsonschema:"description=Number of operations to perform"`
}

// ScyllaDBLoadTestOutput defines the output of ScyllaDB load testing
type ScyllaDBLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// SNMPLoadTestInput defines input parameters for SNMP load testing
type SNMPLoadTestInput struct {
	Target string `json:"target" jsonschema:"required,description=SNMP target address"`
	OID    string `json:"oid" jsonschema:"required,description=SNMP OID to query"`
	Count  int    `json:"count" jsonschema:"description=Number of SNMP requests"`
}

// SNMPLoadTestOutput defines the output of SNMP load testing
type SNMPLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// SSHLoadTestInput defines input parameters for SSH load testing
type SSHLoadTestInput struct {
	Host    string `json:"host" jsonschema:"required,description=SSH server host"`
	Port    int    `json:"port" jsonschema:"description=SSH server port"`
	Command string `json:"command" jsonschema:"description=Command to execute"`
	Count   int    `json:"count" jsonschema:"description=Number of SSH connections/commands"`
}

// SSHLoadTestOutput defines the output of SSH load testing
type SSHLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// SyslogLoadTestInput defines input parameters for Syslog load testing
type SyslogLoadTestInput struct {
	Server string `json:"server" jsonschema:"required,description=Syslog server address"`
	Count  int    `json:"count" jsonschema:"description=Number of syslog messages to send"`
}

// SyslogLoadTestOutput defines the output of Syslog load testing
type SyslogLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// TCPLoadTestInput defines input parameters for TCP load testing
type TCPLoadTestInput struct {
	Host  string `json:"host" jsonschema:"required,description=TCP server host"`
	Port  int    `json:"port" jsonschema:"required,description=TCP server port"`
	Count int    `json:"count" jsonschema:"description=Number of TCP connections"`
}

// TCPLoadTestOutput defines the output of TCP load testing
type TCPLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// TelnetLoadTestInput defines input parameters for Telnet load testing
type TelnetLoadTestInput struct {
	Host  string `json:"host" jsonschema:"required,description=Telnet server host"`
	Port  int    `json:"port" jsonschema:"description=Telnet server port"`
	Count int    `json:"count" jsonschema:"description=Number of Telnet connections"`
}

// TelnetLoadTestOutput defines the output of Telnet load testing
type TelnetLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// TFTPLoadTestInput defines input parameters for TFTP load testing
type TFTPLoadTestInput struct {
	Host  string `json:"host" jsonschema:"required,description=TFTP server host"`
	Count int    `json:"count" jsonschema:"description=Number of TFTP operations"`
}

// TFTPLoadTestOutput defines the output of TFTP load testing
type TFTPLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// UDPLoadTestInput defines input parameters for UDP load testing
type UDPLoadTestInput struct {
	Host  string `json:"host" jsonschema:"required,description=UDP server host"`
	Port  int    `json:"port" jsonschema:"required,description=UDP server port"`
	Count int    `json:"count" jsonschema:"description=Number of UDP packets to send"`
}

// UDPLoadTestOutput defines the output of UDP load testing
type UDPLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// handleHTTPLoadTest executes an HTTP load test
func handleHTTPLoadTest(ctx context.Context, req *mcp.CallToolRequest, input HTTPLoadTestInput) (*mcp.CallToolResult, HTTPLoadTestOutput, error) {
	// Set defaults
	if input.Method == "" {
		input.Method = "GET"
	}
	if input.Count == 0 {
		input.Count = 100
	}

	// Build HTTP config
	httpConf := &http.Config{
		Count: input.Count,
		Payload: http.Payload{
			URL:    input.URL,
			Method: input.Method,
			Body:   util.StrPtr(input.Body),
		},
	}

	if input.MaxConns > 0 {
		httpConf.MaxConns = util.IntPtr(input.MaxConns)
	}
	if input.MaxIdleConns > 0 {
		httpConf.MaxIdleConns = util.IntPtr(input.MaxIdleConns)
	}
	if len(input.Headers) > 0 {
		httpConf.Payload.Header = make(map[string][]string)
		for k, v := range input.Headers {
			httpConf.Payload.Header[k] = []string{v}
		}
	}

	// Create stage
	stage := &config.Stage{
		Name: "http-load-test",
		HTTP: httpConf,
	}
	if input.Concurrent > 0 {
		stage.Concurrent = input.Concurrent
	}

	// Create plan
	plan := &config.Plan{
		Name:   "mcp-http-test",
		Stages: []*config.Stage{stage},
	}

	// Execute load test
	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			HTTP:     httpexec.New(reg),
		})
	})
	if err != nil {
		return nil, HTTPLoadTestOutput{}, fmt.Errorf("failed to execute HTTP load test: %w", err)
	}

	output := HTTPLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed HTTP load test: %d requests to %s", input.Count, input.URL),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleRedisLoadTest executes a Redis load test
func handleRedisLoadTest(ctx context.Context, req *mcp.CallToolRequest, input RedisLoadTestInput) (*mcp.CallToolResult, RedisLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	// Build Redis config
	redisConf := &redis.Config{
		Addr:  input.Addr,
		DB:    input.DB,
		Count: input.Count,
	}
	if input.Password != "" {
		redisConf.Password = util.StrPtr(input.Password)
	}

	// Create basic GET commands for the provided keys
	if len(input.Keys) > 0 {
		redisConf.Commands = make([]redis.Command, len(input.Keys))
		for i, key := range input.Keys {
			redisConf.Commands[i] = redis.Command{
				Get: &redis.Get{Key: key},
			}
		}
	}

	// Create stage
	stage := &config.Stage{
		Name:  "redis-load-test",
		Redis: redisConf,
	}

	// Create plan
	plan := &config.Plan{
		Name:   "mcp-redis-test",
		Stages: []*config.Stage{stage},
	}

	// Execute load test
	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			Redis:    redisexec.New(reg),
		})
	})
	if err != nil {
		return nil, RedisLoadTestOutput{}, fmt.Errorf("failed to execute Redis load test: %w", err)
	}

	output := RedisLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed Redis load test: %d operations to %s", input.Count, input.Addr),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleMongoDBLoadTest executes a MongoDB load test
func handleMongoDBLoadTest(ctx context.Context, req *mcp.CallToolRequest, input MongoDBLoadTestInput) (*mcp.CallToolResult, MongoDBLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}
	if input.Operation == "" {
		input.Operation = "find"
	}

	// Build MongoDB config
	mongoConf := &mongodb.Config{
		URI:        input.URI,
		Database:   input.Database,
		Collection: input.Collection,
		Operation:  input.Operation,
		Count:      input.Count,
	}

	// Create stage
	stage := &config.Stage{
		Name:    "mongodb-load-test",
		MongoDB: mongoConf,
	}

	// Create plan
	plan := &config.Plan{
		Name:   "mcp-mongodb-test",
		Stages: []*config.Stage{stage},
	}

	// Execute load test
	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			MongoDB:  mongodbexec.New(reg),
		})
	})
	if err != nil {
		return nil, MongoDBLoadTestOutput{}, fmt.Errorf("failed to execute MongoDB load test: %w", err)
	}

	output := MongoDBLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed MongoDB load test: %d %s operations on %s.%s",
			input.Count, input.Operation, input.Database, input.Collection),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleClickHouseLoadTest executes a ClickHouse load test
func handleClickHouseLoadTest(ctx context.Context, req *mcp.CallToolRequest, input ClickHouseLoadTestInput) (*mcp.CallToolResult, ClickHouseLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}
	if input.Operation == "" {
		input.Operation = "select"
	}

	// Build ClickHouse config
	clickhouseConf := &clickhouse.Config{
		DSN:       input.DSN,
		Database:  input.Database,
		Table:     input.Table,
		Operation: clickhouse.Operation(input.Operation),
		Query:     input.Query,
		Count:     input.Count,
	}

	// Create stage
	stage := &config.Stage{
		Name:       "clickhouse-load-test",
		ClickHouse: clickhouseConf,
	}

	// Create plan
	plan := &config.Plan{
		Name:   "mcp-clickhouse-test",
		Stages: []*config.Stage{stage},
	}

	// Execute load test
	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry:   reg,
			ClickHouse: clickhouseexec.New(),
		})
	})
	if err != nil {
		return nil, ClickHouseLoadTestOutput{}, fmt.Errorf("failed to execute ClickHouse load test: %w", err)
	}

	output := ClickHouseLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed ClickHouse load test: %d %s operations on %s.%s",
			input.Count, input.Operation, input.Database, input.Table),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handlePostgresLoadTest executes a PostgreSQL load test
func handlePostgresLoadTest(ctx context.Context, req *mcp.CallToolRequest, input PostgresLoadTestInput) (*mcp.CallToolResult, PostgresLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	// Build SQL config for PostgreSQL
	sqlConf := &sql.Config{
		PostgresDSN: util.StrPtr(input.DSN),
		Count:       input.Count,
	}

	if input.Query != "" {
		sqlConf.Query = util.StrPtr(input.Query)
	}
	if input.MaxConns > 0 {
		sqlConf.MaxConns = util.IntPtr(input.MaxConns)
	}
	if input.MaxIdle > 0 {
		sqlConf.MaxIdle = util.IntPtr(input.MaxIdle)
	}

	// Create stage
	stage := &config.Stage{
		Name: "postgres-load-test",
		SQL:  sqlConf,
	}

	// Create plan
	plan := &config.Plan{
		Name:   "mcp-postgres-test",
		Stages: []*config.Stage{stage},
	}

	// Execute load test
	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			SQL:      sqlexec.New(reg),
		})
	})
	if err != nil {
		return nil, PostgresLoadTestOutput{}, fmt.Errorf("failed to execute PostgreSQL load test: %w", err)
	}

	output := PostgresLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed PostgreSQL load test: %d queries", input.Count),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleWebSocketLoadTest executes a WebSocket load test
func handleWebSocketLoadTest(ctx context.Context, req *mcp.CallToolRequest, input WebSocketLoadTestInput) (*mcp.CallToolResult, WebSocketLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	// Build WebSocket config
	wsConf := &websocket.Config{
		URL:   input.URL,
		Count: input.Count,
	}
	if input.Message != "" {
		wsConf.Message = util.StrPtr(input.Message)
	}

	// Create stage
	stage := &config.Stage{
		Name:      "websocket-load-test",
		WebSocket: wsConf,
	}

	// Create plan
	plan := &config.Plan{
		Name:   "mcp-websocket-test",
		Stages: []*config.Stage{stage},
	}

	// Execute load test
	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry:  reg,
			WebSocket: websocketexec.New(reg),
		})
	})
	if err != nil {
		return nil, WebSocketLoadTestOutput{}, fmt.Errorf("failed to execute WebSocket load test: %w", err)
	}

	output := WebSocketLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed WebSocket load test: %d messages to %s", input.Count, input.URL),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleGRPCLoadTest executes a gRPC load test
func handleGRPCLoadTest(ctx context.Context, req *mcp.CallToolRequest, input GRPCLoadTestInput) (*mcp.CallToolResult, GRPCLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	// Build gRPC config
	grpcConf := &grpc.Config{
		Target: input.Target,
		Count:  input.Count,
		TLS:    input.TLS,
	}
	if input.Method != "" {
		grpcConf.Method = util.StrPtr(input.Method)
	}

	// Create stage
	stage := &config.Stage{
		Name: "grpc-load-test",
		GRPC: grpcConf,
	}

	// Create plan
	plan := &config.Plan{
		Name:   "mcp-grpc-test",
		Stages: []*config.Stage{stage},
	}

	// Execute load test
	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			GRPC:     grpcexec.New(reg),
		})
	})
	if err != nil {
		return nil, GRPCLoadTestOutput{}, fmt.Errorf("failed to execute gRPC load test: %w", err)
	}

	output := GRPCLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed gRPC load test: %d requests to %s", input.Count, input.Target),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleRunLoadPlan executes a load test plan from YAML configuration
func handleRunLoadPlan(ctx context.Context, req *mcp.CallToolRequest, input RunLoadPlanInput) (*mcp.CallToolResult, RunLoadPlanOutput, error) {
	// Parse YAML config
	var plan config.Plan
	if err := yaml.Unmarshal([]byte(input.YAMLConfig), &plan); err != nil {
		return nil, RunLoadPlanOutput{}, fmt.Errorf("failed to parse YAML config: %w", err)
	}

	// Validate plan
	if err := plan.Validate(); err != nil {
		return nil, RunLoadPlanOutput{}, fmt.Errorf("invalid plan configuration: %w", err)
	}

	// Execute load test with all executors
	metrics, err := executePlan(ctx, &plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry:      reg,
			ArangoDB:      arangodbexec.New(),
			Cassandra:     cassandraexec.New(),
			CouchDB:       couchdbexec.New(),
			DHCP4:         dhcp4exec.New(),
			DNS:           dnsexec.New(),
			Elasticsearch: elasticsearchexec.New(),
			ETCD:          etcdexec.New(),
			FTP:           ftpexec.New(),
			GraphQL:       graphqlexec.New(),
			GRPC:          grpcexec.New(reg),
			HTTP:          httpexec.New(reg),
			ICMP:          icmpexec.New(),
			InfluxDB:      influxdbexec.New(),
			Kafka:         kafkaexec.New(),
			LDAP:          ldapexec.New(),
			Memcache:      memcacheexec.New(),
			MongoDB:       mongodbexec.New(reg),
			MQTT:          mqttexec.New(),
			NATS:          natsexec.New(),
			Neo4j:         neo4jexec.New(),
			NTP:           ntpexec.New(),
			Pulsar:        pulsarexec.New(),
			RabbitMQ:      rabbitmqexec.New(),
			Redis:         redisexec.New(reg),
			ScyllaDB:      scylladbexec.New(),
			SQL:           sqlexec.New(reg),
			SNMP:          snmpexec.New(),
			SSH:           sshexec.New(),
			Syslog:        syslogexec.New(),
			TCP:           tcpexec.New(),
			Telnet:        telnetexec.New(),
			TFTP:          tftpexec.New(),
			UDP:           udpexec.New(),
			Websocket:     websocketexec.New(reg),
		})
	})
	if err != nil {
		return nil, RunLoadPlanOutput{}, fmt.Errorf("failed to execute load plan: %w", err)
	}

	output := RunLoadPlanOutput{
		Message: fmt.Sprintf("Successfully executed load plan: %s", plan.Name),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleArangoDBLoadTest executes an ArangoDB load test
func handleArangoDBLoadTest(ctx context.Context, req *mcp.CallToolRequest, input ArangoDBLoadTestInput) (*mcp.CallToolResult, ArangoDBLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	arangoConf := &arangodb.Config{
		Endpoints: input.Endpoints,
		Database:  input.Database,
		Count:     input.Count,
	}

	stage := &config.Stage{
		Name:     "arangodb-load-test",
		ArangoDB: arangoConf,
	}

	plan := &config.Plan{
		Name:   "mcp-arangodb-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			ArangoDB: arangodbexec.New(),
		})
	})
	if err != nil {
		return nil, ArangoDBLoadTestOutput{}, fmt.Errorf("failed to execute ArangoDB load test: %w", err)
	}

	output := ArangoDBLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed ArangoDB load test: %d operations on %s", input.Count, input.Database),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleCassandraLoadTest executes a Cassandra load test
func handleCassandraLoadTest(ctx context.Context, req *mcp.CallToolRequest, input CassandraLoadTestInput) (*mcp.CallToolResult, CassandraLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	cassandraConf := &cassandra.Config{
		Hosts:    input.Hosts,
		Keyspace: input.Keyspace,
		Count:    input.Count,
	}

	stage := &config.Stage{
		Name:      "cassandra-load-test",
		Cassandra: cassandraConf,
	}

	plan := &config.Plan{
		Name:   "mcp-cassandra-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry:  reg,
			Cassandra: cassandraexec.New(),
		})
	})
	if err != nil {
		return nil, CassandraLoadTestOutput{}, fmt.Errorf("failed to execute Cassandra load test: %w", err)
	}

	output := CassandraLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed Cassandra load test: %d operations on keyspace %s", input.Count, input.Keyspace),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleCouchDBLoadTest executes a CouchDB load test
func handleCouchDBLoadTest(ctx context.Context, req *mcp.CallToolRequest, input CouchDBLoadTestInput) (*mcp.CallToolResult, CouchDBLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	couchdbConf := &couchdb.Config{
		URL:      input.URL,
		Database: input.Database,
		Count:    input.Count,
	}

	stage := &config.Stage{
		Name:    "couchdb-load-test",
		CouchDB: couchdbConf,
	}

	plan := &config.Plan{
		Name:   "mcp-couchdb-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			CouchDB:  couchdbexec.New(),
		})
	})
	if err != nil {
		return nil, CouchDBLoadTestOutput{}, fmt.Errorf("failed to execute CouchDB load test: %w", err)
	}

	output := CouchDBLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed CouchDB load test: %d operations on %s", input.Count, input.Database),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleDHCP4LoadTest executes a DHCP4 load test
func handleDHCP4LoadTest(ctx context.Context, req *mcp.CallToolRequest, input DHCP4LoadTestInput) (*mcp.CallToolResult, DHCP4LoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	dhcp4Conf := &dhcp4.Config{
		Server: input.Server,
		Count:  input.Count,
	}

	stage := &config.Stage{
		Name:  "dhcp4-load-test",
		DHCP4: dhcp4Conf,
	}

	plan := &config.Plan{
		Name:   "mcp-dhcp4-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			DHCP4:    dhcp4exec.New(),
		})
	})
	if err != nil {
		return nil, DHCP4LoadTestOutput{}, fmt.Errorf("failed to execute DHCP4 load test: %w", err)
	}

	output := DHCP4LoadTestOutput{
		Message: fmt.Sprintf("Successfully executed DHCP4 load test: %d requests to %s", input.Count, input.Server),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleDNSLoadTest executes a DNS load test
func handleDNSLoadTest(ctx context.Context, req *mcp.CallToolRequest, input DNSLoadTestInput) (*mcp.CallToolResult, DNSLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	dnsConf := &dns.Config{
		Server: input.Server,
		Domain: input.Domain,
		Count:  input.Count,
	}

	stage := &config.Stage{
		Name: "dns-load-test",
		DNS:  dnsConf,
	}

	plan := &config.Plan{
		Name:   "mcp-dns-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			DNS:      dnsexec.New(),
		})
	})
	if err != nil {
		return nil, DNSLoadTestOutput{}, fmt.Errorf("failed to execute DNS load test: %w", err)
	}

	output := DNSLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed DNS load test: %d queries for %s", input.Count, input.Domain),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleElasticsearchLoadTest executes an Elasticsearch load test
func handleElasticsearchLoadTest(ctx context.Context, req *mcp.CallToolRequest, input ElasticsearchLoadTestInput) (*mcp.CallToolResult, ElasticsearchLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	elasticsearchConf := &elasticsearch.Config{
		Addresses: input.Addresses,
		Index:     input.Index,
		Count:     input.Count,
	}

	stage := &config.Stage{
		Name:          "elasticsearch-load-test",
		Elasticsearch: elasticsearchConf,
	}

	plan := &config.Plan{
		Name:   "mcp-elasticsearch-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry:      reg,
			Elasticsearch: elasticsearchexec.New(),
		})
	})
	if err != nil {
		return nil, ElasticsearchLoadTestOutput{}, fmt.Errorf("failed to execute Elasticsearch load test: %w", err)
	}

	output := ElasticsearchLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed Elasticsearch load test: %d operations on index %s", input.Count, input.Index),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleETCDLoadTest executes an ETCD load test
func handleETCDLoadTest(ctx context.Context, req *mcp.CallToolRequest, input ETCDLoadTestInput) (*mcp.CallToolResult, ETCDLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	etcdConf := &etcd.Config{
		Endpoints: input.Endpoints,
		Count:     input.Count,
	}

	stage := &config.Stage{
		Name: "etcd-load-test",
		ETCD: etcdConf,
	}

	plan := &config.Plan{
		Name:   "mcp-etcd-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			ETCD:     etcdexec.New(),
		})
	})
	if err != nil {
		return nil, ETCDLoadTestOutput{}, fmt.Errorf("failed to execute ETCD load test: %w", err)
	}

	output := ETCDLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed ETCD load test: %d operations", input.Count),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleFTPLoadTest executes an FTP load test
func handleFTPLoadTest(ctx context.Context, req *mcp.CallToolRequest, input FTPLoadTestInput) (*mcp.CallToolResult, FTPLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}
	if input.Port == 0 {
		input.Port = 21
	}

	ftpConf := &ftp.Config{
		Host:  input.Host,
		Port:  input.Port,
		Count: input.Count,
	}

	stage := &config.Stage{
		Name: "ftp-load-test",
		FTP:  ftpConf,
	}

	plan := &config.Plan{
		Name:   "mcp-ftp-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			FTP:      ftpexec.New(),
		})
	})
	if err != nil {
		return nil, FTPLoadTestOutput{}, fmt.Errorf("failed to execute FTP load test: %w", err)
	}

	output := FTPLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed FTP load test: %d operations", input.Count),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleGraphQLLoadTest executes a GraphQL load test
func handleGraphQLLoadTest(ctx context.Context, req *mcp.CallToolRequest, input GraphQLLoadTestInput) (*mcp.CallToolResult, GraphQLLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	graphqlConf := &graphql.Config{
		URL:   input.URL,
		Query: input.Query,
		Count: input.Count,
	}

	stage := &config.Stage{
		Name:    "graphql-load-test",
		GraphQL: graphqlConf,
	}

	plan := &config.Plan{
		Name:   "mcp-graphql-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			GraphQL:  graphqlexec.New(),
		})
	})
	if err != nil {
		return nil, GraphQLLoadTestOutput{}, fmt.Errorf("failed to execute GraphQL load test: %w", err)
	}

	output := GraphQLLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed GraphQL load test: %d queries to %s", input.Count, input.URL),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleICMPLoadTest executes an ICMP/Ping load test
func handleICMPLoadTest(ctx context.Context, req *mcp.CallToolRequest, input ICMPLoadTestInput) (*mcp.CallToolResult, ICMPLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	icmpConf := &icmp.Config{
		Host:  input.Host,
		Count: input.Count,
	}

	stage := &config.Stage{
		Name: "icmp-load-test",
		ICMP: icmpConf,
	}

	plan := &config.Plan{
		Name:   "mcp-icmp-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			ICMP:     icmpexec.New(),
		})
	})
	if err != nil {
		return nil, ICMPLoadTestOutput{}, fmt.Errorf("failed to execute ICMP load test: %w", err)
	}

	output := ICMPLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed ICMP load test: %d pings to %s", input.Count, input.Host),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleInfluxDBLoadTest executes an InfluxDB load test
func handleInfluxDBLoadTest(ctx context.Context, req *mcp.CallToolRequest, input InfluxDBLoadTestInput) (*mcp.CallToolResult, InfluxDBLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	influxdbConf := &influxdb.Config{
		URL:    input.URL,
		Bucket: input.Bucket,
		Count:  input.Count,
	}

	stage := &config.Stage{
		Name:     "influxdb-load-test",
		InfluxDB: influxdbConf,
	}

	plan := &config.Plan{
		Name:   "mcp-influxdb-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			InfluxDB: influxdbexec.New(),
		})
	})
	if err != nil {
		return nil, InfluxDBLoadTestOutput{}, fmt.Errorf("failed to execute InfluxDB load test: %w", err)
	}

	output := InfluxDBLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed InfluxDB load test: %d operations on bucket %s", input.Count, input.Bucket),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleKafkaLoadTest executes a Kafka load test
func handleKafkaLoadTest(ctx context.Context, req *mcp.CallToolRequest, input KafkaLoadTestInput) (*mcp.CallToolResult, KafkaLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	kafkaConf := &kafka.Config{
		Brokers: input.Brokers,
		Topic:   input.Topic,
		Count:   input.Count,
	}

	stage := &config.Stage{
		Name:  "kafka-load-test",
		Kafka: kafkaConf,
	}

	plan := &config.Plan{
		Name:   "mcp-kafka-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			Kafka:    kafkaexec.New(),
		})
	})
	if err != nil {
		return nil, KafkaLoadTestOutput{}, fmt.Errorf("failed to execute Kafka load test: %w", err)
	}

	output := KafkaLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed Kafka load test: %d messages to topic %s", input.Count, input.Topic),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleLDAPLoadTest executes an LDAP load test
func handleLDAPLoadTest(ctx context.Context, req *mcp.CallToolRequest, input LDAPLoadTestInput) (*mcp.CallToolResult, LDAPLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	ldapConf := &ldap.Config{
		Server: input.Server,
		BaseDN: input.BaseDN,
		Count:  input.Count,
	}

	stage := &config.Stage{
		Name: "ldap-load-test",
		LDAP: ldapConf,
	}

	plan := &config.Plan{
		Name:   "mcp-ldap-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			LDAP:     ldapexec.New(),
		})
	})
	if err != nil {
		return nil, LDAPLoadTestOutput{}, fmt.Errorf("failed to execute LDAP load test: %w", err)
	}

	output := LDAPLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed LDAP load test: %d operations", input.Count),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleMemcacheLoadTest executes a Memcache load test
func handleMemcacheLoadTest(ctx context.Context, req *mcp.CallToolRequest, input MemcacheLoadTestInput) (*mcp.CallToolResult, MemcacheLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	memcacheConf := &memcache.Config{
		Servers: input.Servers,
		Count:   input.Count,
	}

	stage := &config.Stage{
		Name:     "memcache-load-test",
		Memcache: memcacheConf,
	}

	plan := &config.Plan{
		Name:   "mcp-memcache-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			Memcache: memcacheexec.New(),
		})
	})
	if err != nil {
		return nil, MemcacheLoadTestOutput{}, fmt.Errorf("failed to execute Memcache load test: %w", err)
	}

	output := MemcacheLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed Memcache load test: %d operations", input.Count),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleMQTTLoadTest executes an MQTT load test
func handleMQTTLoadTest(ctx context.Context, req *mcp.CallToolRequest, input MQTTLoadTestInput) (*mcp.CallToolResult, MQTTLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	mqttConf := &mqtt.Config{
		Broker: input.Broker,
		Topic:  input.Topic,
		Count:  input.Count,
	}

	stage := &config.Stage{
		Name: "mqtt-load-test",
		MQTT: mqttConf,
	}

	plan := &config.Plan{
		Name:   "mcp-mqtt-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			MQTT:     mqttexec.New(),
		})
	})
	if err != nil {
		return nil, MQTTLoadTestOutput{}, fmt.Errorf("failed to execute MQTT load test: %w", err)
	}

	output := MQTTLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed MQTT load test: %d messages to topic %s", input.Count, input.Topic),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleNATSLoadTest executes a NATS load test
func handleNATSLoadTest(ctx context.Context, req *mcp.CallToolRequest, input NATSLoadTestInput) (*mcp.CallToolResult, NATSLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	natsConf := &nats.Config{
		URL:     input.URL,
		Subject: input.Subject,
		Count:   input.Count,
	}

	stage := &config.Stage{
		Name: "nats-load-test",
		NATS: natsConf,
	}

	plan := &config.Plan{
		Name:   "mcp-nats-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			NATS:     natsexec.New(),
		})
	})
	if err != nil {
		return nil, NATSLoadTestOutput{}, fmt.Errorf("failed to execute NATS load test: %w", err)
	}

	output := NATSLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed NATS load test: %d messages to subject %s", input.Count, input.Subject),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleNeo4jLoadTest executes a Neo4j load test
func handleNeo4jLoadTest(ctx context.Context, req *mcp.CallToolRequest, input Neo4jLoadTestInput) (*mcp.CallToolResult, Neo4jLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	neo4jConf := &neo4j.Config{
		URI:   input.URI,
		Query: input.Query,
		Count: input.Count,
	}

	stage := &config.Stage{
		Name:  "neo4j-load-test",
		Neo4j: neo4jConf,
	}

	plan := &config.Plan{
		Name:   "mcp-neo4j-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			Neo4j:    neo4jexec.New(),
		})
	})
	if err != nil {
		return nil, Neo4jLoadTestOutput{}, fmt.Errorf("failed to execute Neo4j load test: %w", err)
	}

	output := Neo4jLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed Neo4j load test: %d queries", input.Count),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleNTPLoadTest executes an NTP load test
func handleNTPLoadTest(ctx context.Context, req *mcp.CallToolRequest, input NTPLoadTestInput) (*mcp.CallToolResult, NTPLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	ntpConf := &ntp.Config{
		Server: input.Server,
		Count:  input.Count,
	}

	stage := &config.Stage{
		Name: "ntp-load-test",
		NTP:  ntpConf,
	}

	plan := &config.Plan{
		Name:   "mcp-ntp-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			NTP:      ntpexec.New(),
		})
	})
	if err != nil {
		return nil, NTPLoadTestOutput{}, fmt.Errorf("failed to execute NTP load test: %w", err)
	}

	output := NTPLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed NTP load test: %d requests to %s", input.Count, input.Server),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handlePulsarLoadTest executes a Pulsar load test
func handlePulsarLoadTest(ctx context.Context, req *mcp.CallToolRequest, input PulsarLoadTestInput) (*mcp.CallToolResult, PulsarLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	pulsarConf := &pulsar.Config{
		URL:   input.URL,
		Topic: input.Topic,
		Count: input.Count,
	}

	stage := &config.Stage{
		Name:   "pulsar-load-test",
		Pulsar: pulsarConf,
	}

	plan := &config.Plan{
		Name:   "mcp-pulsar-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			Pulsar:   pulsarexec.New(),
		})
	})
	if err != nil {
		return nil, PulsarLoadTestOutput{}, fmt.Errorf("failed to execute Pulsar load test: %w", err)
	}

	output := PulsarLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed Pulsar load test: %d messages to topic %s", input.Count, input.Topic),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleRabbitMQLoadTest executes a RabbitMQ load test
func handleRabbitMQLoadTest(ctx context.Context, req *mcp.CallToolRequest, input RabbitMQLoadTestInput) (*mcp.CallToolResult, RabbitMQLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	rabbitmqConf := &rabbitmq.Config{
		URL:   input.URL,
		Queue: input.Queue,
		Count: input.Count,
	}

	stage := &config.Stage{
		Name:     "rabbitmq-load-test",
		RabbitMQ: rabbitmqConf,
	}

	plan := &config.Plan{
		Name:   "mcp-rabbitmq-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			RabbitMQ: rabbitmqexec.New(),
		})
	})
	if err != nil {
		return nil, RabbitMQLoadTestOutput{}, fmt.Errorf("failed to execute RabbitMQ load test: %w", err)
	}

	output := RabbitMQLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed RabbitMQ load test: %d messages to queue %s", input.Count, input.Queue),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleScyllaDBLoadTest executes a ScyllaDB load test
func handleScyllaDBLoadTest(ctx context.Context, req *mcp.CallToolRequest, input ScyllaDBLoadTestInput) (*mcp.CallToolResult, ScyllaDBLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	scylladbConf := &scylladb.Config{
		Hosts:    input.Hosts,
		Keyspace: input.Keyspace,
		Count:    input.Count,
	}

	stage := &config.Stage{
		Name:     "scylladb-load-test",
		ScyllaDB: scylladbConf,
	}

	plan := &config.Plan{
		Name:   "mcp-scylladb-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			ScyllaDB: scylladbexec.New(),
		})
	})
	if err != nil {
		return nil, ScyllaDBLoadTestOutput{}, fmt.Errorf("failed to execute ScyllaDB load test: %w", err)
	}

	output := ScyllaDBLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed ScyllaDB load test: %d operations on keyspace %s", input.Count, input.Keyspace),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleSNMPLoadTest executes an SNMP load test
func handleSNMPLoadTest(ctx context.Context, req *mcp.CallToolRequest, input SNMPLoadTestInput) (*mcp.CallToolResult, SNMPLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	snmpConf := &snmp.Config{
		Target: input.Target,
		OID:    input.OID,
		Count:  input.Count,
	}

	stage := &config.Stage{
		Name: "snmp-load-test",
		SNMP: snmpConf,
	}

	plan := &config.Plan{
		Name:   "mcp-snmp-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			SNMP:     snmpexec.New(),
		})
	})
	if err != nil {
		return nil, SNMPLoadTestOutput{}, fmt.Errorf("failed to execute SNMP load test: %w", err)
	}

	output := SNMPLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed SNMP load test: %d requests to %s", input.Count, input.Target),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleSSHLoadTest executes an SSH load test
func handleSSHLoadTest(ctx context.Context, req *mcp.CallToolRequest, input SSHLoadTestInput) (*mcp.CallToolResult, SSHLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}
	if input.Port == 0 {
		input.Port = 22
	}

	sshConf := &ssh.Config{
		Host:    input.Host,
		Port:    input.Port,
		Command: input.Command,
		Count:   input.Count,
	}

	stage := &config.Stage{
		Name: "ssh-load-test",
		SSH:  sshConf,
	}

	plan := &config.Plan{
		Name:   "mcp-ssh-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			SSH:      sshexec.New(),
		})
	})
	if err != nil {
		return nil, SSHLoadTestOutput{}, fmt.Errorf("failed to execute SSH load test: %w", err)
	}

	output := SSHLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed SSH load test: %d connections to %s", input.Count, input.Host),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleSyslogLoadTest executes a Syslog load test
func handleSyslogLoadTest(ctx context.Context, req *mcp.CallToolRequest, input SyslogLoadTestInput) (*mcp.CallToolResult, SyslogLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	syslogConf := &syslog.Config{
		Server: input.Server,
		Count:  input.Count,
	}

	stage := &config.Stage{
		Name:   "syslog-load-test",
		Syslog: syslogConf,
	}

	plan := &config.Plan{
		Name:   "mcp-syslog-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			Syslog:   syslogexec.New(),
		})
	})
	if err != nil {
		return nil, SyslogLoadTestOutput{}, fmt.Errorf("failed to execute Syslog load test: %w", err)
	}

	output := SyslogLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed Syslog load test: %d messages to %s", input.Count, input.Server),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleTCPLoadTest executes a TCP load test
func handleTCPLoadTest(ctx context.Context, req *mcp.CallToolRequest, input TCPLoadTestInput) (*mcp.CallToolResult, TCPLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	tcpConf := &tcp.Config{
		Host:  input.Host,
		Port:  input.Port,
		Count: input.Count,
	}

	stage := &config.Stage{
		Name: "tcp-load-test",
		TCP:  tcpConf,
	}

	plan := &config.Plan{
		Name:   "mcp-tcp-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			TCP:      tcpexec.New(),
		})
	})
	if err != nil {
		return nil, TCPLoadTestOutput{}, fmt.Errorf("failed to execute TCP load test: %w", err)
	}

	output := TCPLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed TCP load test: %d connections to %s:%d", input.Count, input.Host, input.Port),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleTelnetLoadTest executes a Telnet load test
func handleTelnetLoadTest(ctx context.Context, req *mcp.CallToolRequest, input TelnetLoadTestInput) (*mcp.CallToolResult, TelnetLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}
	if input.Port == 0 {
		input.Port = 23
	}

	telnetConf := &telnet.Config{
		Host:  input.Host,
		Port:  input.Port,
		Count: input.Count,
	}

	stage := &config.Stage{
		Name:   "telnet-load-test",
		Telnet: telnetConf,
	}

	plan := &config.Plan{
		Name:   "mcp-telnet-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			Telnet:   telnetexec.New(),
		})
	})
	if err != nil {
		return nil, TelnetLoadTestOutput{}, fmt.Errorf("failed to execute Telnet load test: %w", err)
	}

	output := TelnetLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed Telnet load test: %d connections to %s:%d", input.Count, input.Host, input.Port),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleTFTPLoadTest executes a TFTP load test
func handleTFTPLoadTest(ctx context.Context, req *mcp.CallToolRequest, input TFTPLoadTestInput) (*mcp.CallToolResult, TFTPLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	tftpConf := &tftp.Config{
		Host:  input.Host,
		Count: input.Count,
	}

	stage := &config.Stage{
		Name: "tftp-load-test",
		TFTP: tftpConf,
	}

	plan := &config.Plan{
		Name:   "mcp-tftp-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			TFTP:     tftpexec.New(),
		})
	})
	if err != nil {
		return nil, TFTPLoadTestOutput{}, fmt.Errorf("failed to execute TFTP load test: %w", err)
	}

	output := TFTPLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed TFTP load test: %d operations", input.Count),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// handleUDPLoadTest executes a UDP load test
func handleUDPLoadTest(ctx context.Context, req *mcp.CallToolRequest, input UDPLoadTestInput) (*mcp.CallToolResult, UDPLoadTestOutput, error) {
	if input.Count == 0 {
		input.Count = 100
	}

	udpConf := &udp.Config{
		Host:  input.Host,
		Port:  input.Port,
		Count: input.Count,
	}

	stage := &config.Stage{
		Name: "udp-load-test",
		UDP:  udpConf,
	}

	plan := &config.Plan{
		Name:   "mcp-udp-test",
		Stages: []*config.Stage{stage},
	}

	metrics, err := executePlan(ctx, plan, func(reg *prometheus.Registry) (executor.Stage, error) {
		return stageexec.New(stageexec.Params{
			Registry: reg,
			UDP:      udpexec.New(),
		})
	})
	if err != nil {
		return nil, UDPLoadTestOutput{}, fmt.Errorf("failed to execute UDP load test: %w", err)
	}

	output := UDPLoadTestOutput{
		Message: fmt.Sprintf("Successfully executed UDP load test: %d packets to %s:%d", input.Count, input.Host, input.Port),
		Metrics: metrics,
	}

	return &mcp.CallToolResult{
		Content: []interface{}{
			&mcp.TextContent{
				Type: "text",
				Text: output.Message,
			},
		},
	}, output, nil
}

// executePlan is a helper function to execute a load test plan and capture metrics
func executePlan(ctx context.Context, plan *config.Plan, stageFactory func(*prometheus.Registry) (executor.Stage, error)) (string, error) {
	reg := prometheus.NewPedanticRegistry()

	// Create stage executor
	stageExec, err := stageFactory(reg)
	if err != nil {
		return "", fmt.Errorf("failed to create stage executor: %w", err)
	}

	// Create plan executor
	planExec, err := executor.NewPlan(
		executor.Params{Registry: reg},
		stageExec,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create plan executor: %w", err)
	}

	// Execute the plan
	if err := planExec.Execute(ctx, plan); err != nil {
		return "", fmt.Errorf("failed to execute plan: %w", err)
	}

	// Gather metrics
	var buf bytes.Buffer
	if err := util.RegistryGather(reg, &buf); err != nil {
		return "", fmt.Errorf("failed to gather metrics: %w", err)
	}

	return buf.String(), nil
}
