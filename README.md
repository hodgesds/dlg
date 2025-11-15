# DLG - Da Load Generator
`dlg` is a universal load generator that can be used to generate load for various protocols and services.

## Supported Protocols

DLG supports the following protocols:

- **HTTP/HTTPS** - HTTP/HTTPS load testing with customizable headers, methods, and payloads
- **gRPC** - gRPC service load testing with metadata and connection options
- **WebSocket** - WebSocket connection testing with read/write operations
- **GraphQL** - GraphQL query execution with variables and custom headers
- **MongoDB** - MongoDB database operations (insert, find, update, delete, count, aggregate)
- **Cassandra** - Cassandra CQL queries with configurable consistency levels
- **Elasticsearch** - Elasticsearch index, search, and document operations
- **InfluxDB** - InfluxDB time-series data writes and queries
- **Redis** - Redis commands and operations
- **SQL** - MySQL, PostgreSQL, and ClickHouse database testing
- **MQTT** - MQTT publish/subscribe messaging with QoS levels
- **ETCD** - ETCD key-value operations
- **DNS** - DNS query load testing
- **LDAP** - LDAP authentication and query testing
- **Memcache** - Memcache get/set/delete operations
- **SNMP** - SNMP protocol testing
- **SSH** - SSH command execution
- **UDP** - UDP protocol testing
- **DHCP4** - DHCP4 protocol testing

## Features

- **Multi-protocol support** - Test multiple protocols in a single plan
- **Staged execution** - Organize tests into stages with configurable repeats and durations
- **Concurrent execution** - Run stages concurrently for maximum load
- **Prometheus metrics** - Built-in Prometheus metrics for monitoring
- **Rate limiting** - Control load with configurable rate limiters
- **Flexible configuration** - YAML-based configuration with extensive options
- **MCP Server** - Model Context Protocol server for AI agent integration

## MCP Server for AI Agents

DLG includes a built-in Model Context Protocol (MCP) server that allows AI agents to use DLG for load testing. This enables AI-powered workflows to automatically generate and execute load tests.

### Starting the MCP Server

```bash
dlg mcp
```

The MCP server runs over stdio and exposes the following tools:

- **http_load_test** - Generate HTTP load against a target URL
- **redis_load_test** - Generate Redis load against a server
- **mongodb_load_test** - Generate MongoDB load against a database
- **postgres_load_test** - Generate PostgreSQL load against a database
- **websocket_load_test** - Generate WebSocket load against an endpoint
- **grpc_load_test** - Generate gRPC load against a service
- **run_load_plan** - Execute a YAML-based load test plan

### Example AI Agent Usage

AI agents can use these tools to perform load testing as part of automated workflows. For example:
- Performance testing during CI/CD pipelines
- Automated capacity planning
- Dynamic stress testing based on monitoring data
- Interactive load testing through natural language

### MCP Integration

To integrate DLG with an AI agent:

1. Start the MCP server: `dlg mcp`
2. Configure your AI agent to connect to the MCP server via stdio
3. The agent can now call load testing tools with natural language

Each tool returns Prometheus metrics from the load test execution.
