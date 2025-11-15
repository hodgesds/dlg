# DLG Architecture Documentation

## Overview

DLG (Da Load Generator) is a universal, multi-protocol load testing tool designed to generate load against various protocols, databases, and services. The architecture follows a modular, plugin-based design that makes it easy to add support for new protocols.

## Core Components

### 1. Command Layer (`/cmd`)

The command layer provides CLI interfaces for each supported protocol using the Cobra framework. Each protocol has its own subcommand that:

- Parses command-line arguments
- Constructs configuration objects
- Invokes the appropriate executor
- Handles output and error reporting

**Location:** `/cmd/{protocol}/`

**Example:**
```bash
dlg http --url https://api.example.com --count 1000
dlg mongodb --uri mongodb://localhost:27017 --operation insert
```

### 2. Configuration Layer (`/config`)

The configuration layer defines data structures for each protocol's load test parameters. Configurations can be:

- Loaded from YAML files
- Constructed programmatically via the CLI
- Used by executors to run load tests

**Location:** `/config/{protocol}/`

**Key Features:**
- YAML serialization support
- Validation logic
- Type-safe configuration structures
- Support for multiple data encoding formats (hex, base64, etc.)

**Example Configuration:**
```yaml
plans:
  - name: "http-load-test"
    description: "Test HTTP endpoint"
    stages:
      - http:
          url: "https://api.example.com/users"
          method: "POST"
          count: 1000
          header:
            Content-Type: "application/json"
          body: '{"name": "test"}'
```

### 3. Executor Layer (`/executor`)

The executor layer implements the actual load generation logic for each protocol. Executors:

- Receive configuration objects
- Establish connections
- Generate concurrent load
- Collect metrics
- Handle errors and retries

**Location:** `/executor/{protocol}/`

**Interface:**
```go
type Executor interface {
    Execute(ctx context.Context, config *Config) error
}
```

**Key Responsibilities:**
- Connection pooling and management
- Concurrent request execution
- Rate limiting integration
- Prometheus metrics instrumentation
- Context-aware cancellation

### 4. Manager (`/manager.go`)

The manager orchestrates execution of complex load test plans containing multiple stages and protocols. It provides:

- Plan storage and retrieval
- Concurrent execution coordination
- Thread-safe plan management
- Integration with the plan executor

**Interface:**
```go
type Manager interface {
    Get(ctx context.Context, name string) (*config.Plan, error)
    Add(ctx context.Context, plan *config.Plan) error
    Delete(ctx context.Context, name string) error
    Plans(ctx context.Context) ([]*config.Plan, error)
    Execute(ctx context.Context, plan *config.Plan) error
}
```

### 5. Rate Limiter (`/ratelimiter`)

The rate limiter provides token bucket-based rate limiting for both operations per second (ops/sec) and bytes per second (bytes/sec).

**Key Features:**
- Configurable rate limits
- Context-aware waiting
- Prometheus metrics integration
- Safe for concurrent use

### 6. Metrics (`/metrics`)

DLG uses Prometheus for comprehensive metrics collection:

**Executor Metrics:**
- `executor_plan_stages_total` - Counter for stage executions
- `executor_plan_stage_duration` - Histogram of stage durations

**Protocol-Specific Metrics (example: HTTP):**
- `client_in_flight_requests` - Gauge of active requests
- `client_api_requests_total` - Counter of total requests by code and method
- `dns_duration_ms` - Histogram of DNS lookup times
- `tls_duration_ms` - Histogram of TLS handshake times
- `request_duration_ms` - Histogram of request latencies

### 7. MCP Server (`/mcpserver`)

The Model Context Protocol (MCP) server enables AI agents to use DLG for automated load testing workflows.

**Location:** `/mcpserver/`

**Key Features:**
- Exposes DLG functionality as MCP tools (35+ tools)
- JSON-RPC 2.0 protocol over stdio
- Tool discovery and documentation
- Support for all protocols
- Plan management and execution

**Example Tools:**
- `run_http_load` - Execute HTTP load tests
- `run_mongodb_load` - Execute MongoDB load tests
- `create_plan` - Create multi-stage load test plans
- `list_plans` - List available plans
- `execute_plan` - Execute a saved plan

## Architecture Patterns

### 1. Plugin Architecture

Each protocol implementation follows a consistent pattern:

```
/config/{protocol}/     - Configuration structures
/executor/{protocol}/   - Execution logic
/cmd/{protocol}/        - CLI interface
```

This makes it easy to add new protocols by following the established patterns.

### 2. Dependency Injection

Components use dependency injection for testability and flexibility:

```go
executor := http.New(prometheusRegistry)
manager := dlg.NewManager(planExecutor)
```

### 3. Context Propagation

All operations support context for:
- Cancellation
- Timeouts
- Deadlines
- Request-scoped values

### 4. Concurrent Execution

Load generation uses goroutines with:
- WaitGroups for synchronization
- Mutexes for shared state
- Channels for communication
- Context for cancellation

### 5. Metrics-First Design

All executors are instrumented with Prometheus metrics from the start, enabling:
- Performance monitoring
- Debugging
- Capacity planning
- SLA validation

## Data Flow

```
1. CLI Input / YAML Config / MCP Tool Call
   ↓
2. Configuration Parsing & Validation
   ↓
3. Executor Selection & Initialization
   ↓
4. Concurrent Load Generation
   ↓
5. Metrics Collection
   ↓
6. Results Aggregation & Output
```

## Supported Protocols

### Databases (13)
- MongoDB
- Redis
- PostgreSQL
- MySQL
- ClickHouse
- Cassandra
- ScyllaDB
- Elasticsearch
- InfluxDB
- CouchDB
- ArangoDB
- Neo4j
- Memcache

### Messaging (5)
- Kafka
- RabbitMQ
- MQTT
- NATS
- Apache Pulsar

### Network Protocols (17)
- HTTP/HTTPS
- gRPC
- GraphQL
- WebSocket
- TCP
- UDP
- DNS
- DHCP4
- ICMP
- SSH
- FTP
- TFTP
- Telnet
- NTP
- SNMP
- Syslog
- LDAP
- ETCD

## Extension Points

### Adding a New Protocol

1. Create configuration structure in `/config/{protocol}/`
2. Implement executor in `/executor/{protocol}/`
3. Add CLI command in `/cmd/{protocol}/`
4. Register MCP tools in `/mcpserver/tools.go`
5. Add tests for all components
6. Update documentation

### Adding New Metrics

1. Define metrics in executor's `metrics.go`
2. Register with Prometheus registry
3. Instrument relevant code paths
4. Add metric documentation

### Custom Rate Limiting

Implement the `RateLimiter` interface:

```go
type RateLimiter interface {
    Wait(ctx context.Context) error
    WaitBytes(ctx context.Context, bytes int) error
    Reset()
}
```

## Performance Considerations

### Connection Pooling

Executors use connection pooling to:
- Reuse connections
- Minimize connection overhead
- Configure max idle/active connections

### Goroutine Management

- Bounded concurrency using semaphores or worker pools
- WaitGroups for proper cleanup
- Context for cancellation

### Memory Management

- Object pooling for frequently allocated objects
- Streaming for large payloads
- Proper resource cleanup with defer

### Metrics Overhead

- Pre-allocated metric vectors
- Efficient label usage
- Batch metric updates where possible

## Security Considerations

### Credentials

- Support for environment variables
- File-based credential loading
- Never log sensitive data

### TLS/SSL

- Configurable TLS settings
- Certificate validation options
- Support for custom CA certificates

### Network Isolation

- Configurable timeouts
- Connection limits
- Rate limiting

## Testing Strategy

### Unit Tests

- Test each component in isolation
- Mock external dependencies
- Focus on edge cases and error handling

### Integration Tests

- Test protocol interactions
- Use test servers (httptest, etc.)
- Validate metrics collection

### End-to-End Tests

- Test complete workflows
- Validate CLI commands
- Test plan execution

## Future Enhancements

### Planned Features

1. Distributed load generation
2. Real-time result streaming
3. Advanced result analysis
4. Custom scripting support
5. Web UI for visualization
6. Enhanced error reporting
7. Automatic baseline comparison
8. CI/CD integration plugins

### Protocol Additions

- More database systems
- Additional messaging protocols
- Cloud provider APIs
- Custom protocol support

## References

- [User Guide](USER_GUIDE.md)
- [API Documentation](API.md)
- [Contributing Guidelines](CONTRIBUTING.md)
- [Prometheus Documentation](https://prometheus.io/docs/)
- [Model Context Protocol](https://modelcontextprotocol.io/)
