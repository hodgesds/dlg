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
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Server represents the MCP server for DLG load generator
type Server struct {
	server *mcp.Server
}

// New creates a new MCP server instance
func New() *Server {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "dlg-load-generator",
		Version: "1.0.0",
	}, nil)

	s := &Server{
		server: server,
	}

	// Register all tools
	s.registerTools()

	return s
}

// Run starts the MCP server using stdio transport
func (s *Server) Run(ctx context.Context) error {
	log.Println("Starting DLG MCP server...")
	if err := s.server.Run(ctx, &mcp.StdioTransport{}); err != nil {
		return err
	}
	return nil
}

// registerTools registers all available load generation tools
func (s *Server) registerTools() {
	// HTTP load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "http_load_test",
		Description: "Generate HTTP load against a target URL with configurable parameters",
	}, handleHTTPLoadTest)

	// Redis load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "redis_load_test",
		Description: "Generate Redis load against a target server with configurable commands",
	}, handleRedisLoadTest)

	// Generic YAML-based load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "run_load_plan",
		Description: "Execute a load test plan from a YAML configuration",
	}, handleRunLoadPlan)

	// MongoDB load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "mongodb_load_test",
		Description: "Generate MongoDB load against a target database with configurable operations",
	}, handleMongoDBLoadTest)

	// ClickHouse load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "clickhouse_load_test",
		Description: "Generate ClickHouse load against a target database with configurable operations",
	}, handleClickHouseLoadTest)

	// PostgreSQL load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "postgres_load_test",
		Description: "Generate PostgreSQL load against a target database with configurable queries",
	}, handlePostgresLoadTest)

	// WebSocket load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "websocket_load_test",
		Description: "Generate WebSocket load against a target endpoint",
	}, handleWebSocketLoadTest)

	// gRPC load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "grpc_load_test",
		Description: "Generate gRPC load against a target service",
	}, handleGRPCLoadTest)

	// ArangoDB load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "arangodb_load_test",
		Description: "Generate ArangoDB load against a target database",
	}, handleArangoDBLoadTest)

	// Cassandra load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "cassandra_load_test",
		Description: "Generate Cassandra load against a target cluster",
	}, handleCassandraLoadTest)

	// CouchDB load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "couchdb_load_test",
		Description: "Generate CouchDB load against a target database",
	}, handleCouchDBLoadTest)

	// DHCP4 load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "dhcp4_load_test",
		Description: "Generate DHCP4 load against a DHCP server",
	}, handleDHCP4LoadTest)

	// DNS load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "dns_load_test",
		Description: "Generate DNS load against a DNS server",
	}, handleDNSLoadTest)

	// Elasticsearch load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "elasticsearch_load_test",
		Description: "Generate Elasticsearch load against a target cluster",
	}, handleElasticsearchLoadTest)

	// ETCD load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "etcd_load_test",
		Description: "Generate ETCD load against a target cluster",
	}, handleETCDLoadTest)

	// FTP load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "ftp_load_test",
		Description: "Generate FTP load against an FTP server",
	}, handleFTPLoadTest)

	// GraphQL load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "graphql_load_test",
		Description: "Generate GraphQL load against a target endpoint",
	}, handleGraphQLLoadTest)

	// ICMP load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "icmp_load_test",
		Description: "Generate ICMP/Ping load against a target host",
	}, handleICMPLoadTest)

	// InfluxDB load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "influxdb_load_test",
		Description: "Generate InfluxDB load against a target database",
	}, handleInfluxDBLoadTest)

	// Kafka load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "kafka_load_test",
		Description: "Generate Kafka load against a target cluster",
	}, handleKafkaLoadTest)

	// LDAP load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "ldap_load_test",
		Description: "Generate LDAP load against a directory server",
	}, handleLDAPLoadTest)

	// Memcache load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "memcache_load_test",
		Description: "Generate Memcache load against a target server",
	}, handleMemcacheLoadTest)

	// MQTT load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "mqtt_load_test",
		Description: "Generate MQTT load against a broker",
	}, handleMQTTLoadTest)

	// NATS load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "nats_load_test",
		Description: "Generate NATS load against a NATS server",
	}, handleNATSLoadTest)

	// Neo4j load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "neo4j_load_test",
		Description: "Generate Neo4j load against a graph database",
	}, handleNeo4jLoadTest)

	// NTP load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "ntp_load_test",
		Description: "Generate NTP load against a time server",
	}, handleNTPLoadTest)

	// Pulsar load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "pulsar_load_test",
		Description: "Generate Apache Pulsar load against a target cluster",
	}, handlePulsarLoadTest)

	// RabbitMQ load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "rabbitmq_load_test",
		Description: "Generate RabbitMQ load against a message broker",
	}, handleRabbitMQLoadTest)

	// ScyllaDB load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "scylladb_load_test",
		Description: "Generate ScyllaDB load against a target cluster",
	}, handleScyllaDBLoadTest)

	// SNMP load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "snmp_load_test",
		Description: "Generate SNMP load against a target device",
	}, handleSNMPLoadTest)

	// SSH load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "ssh_load_test",
		Description: "Generate SSH load against a target server",
	}, handleSSHLoadTest)

	// Syslog load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "syslog_load_test",
		Description: "Generate Syslog load against a syslog server",
	}, handleSyslogLoadTest)

	// TCP load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "tcp_load_test",
		Description: "Generate TCP load against a target server",
	}, handleTCPLoadTest)

	// Telnet load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "telnet_load_test",
		Description: "Generate Telnet load against a target server",
	}, handleTelnetLoadTest)

	// TFTP load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "tftp_load_test",
		Description: "Generate TFTP load against a TFTP server",
	}, handleTFTPLoadTest)

	// UDP load test tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "udp_load_test",
		Description: "Generate UDP load against a target server",
	}, handleUDPLoadTest)
}
