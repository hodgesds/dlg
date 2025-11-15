# DLG User Guide

## Table of Contents

- [Introduction](#introduction)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Protocol Guides](#protocol-guides)
- [Configuration Files](#configuration-files)
- [MCP Server](#mcp-server)
- [Metrics and Monitoring](#metrics-and-monitoring)
- [Advanced Usage](#advanced-usage)
- [Troubleshooting](#troubleshooting)

## Introduction

DLG (Da Load Generator) is a universal load testing tool that supports 35+ protocols including databases, messaging systems, and network protocols. It's designed to be:

- **Easy to use**: Simple CLI interface or YAML configuration
- **Powerful**: Support for complex multi-stage load tests
- **Observable**: Built-in Prometheus metrics
- **AI-friendly**: MCP server for AI agent integration

## Installation

### From Source

```bash
git clone https://github.com/hodgesds/dlg.git
cd dlg
make build
```

The binary will be available at `./dlg`.

### Prerequisites

- Go 1.24.0 or later
- Access to target systems/services you want to test

## Quick Start

### Simple HTTP Load Test

```bash
./dlg http --url https://api.example.com/users --method GET --count 1000
```

This will send 1000 GET requests to the specified URL.

### MongoDB Load Test

```bash
./dlg mongodb \
  --uri mongodb://localhost:27017 \
  --database testdb \
  --collection users \
  --operation insert \
  --count 1000
```

### Using a Configuration File

Create a file `plan.yaml`:

```yaml
plans:
  - name: "api-load-test"
    description: "Test API endpoints"
    stages:
      - http:
          url: "https://api.example.com/users"
          method: "POST"
          count: 100
          header:
            Content-Type: "application/json"
          body: '{"name": "test", "email": "test@example.com"}'
```

Run it:

```bash
./dlg plan -f plan.yaml
```

## Protocol Guides

### HTTP/HTTPS

#### Basic GET Request

```bash
./dlg http \
  --url https://api.example.com/data \
  --method GET \
  --count 500
```

#### POST Request with Body

```bash
./dlg http \
  --url https://api.example.com/users \
  --method POST \
  --count 100 \
  --body '{"name": "John Doe", "age": 30}'
```

#### With Custom Headers

YAML configuration:

```yaml
stages:
  - http:
      url: "https://api.example.com/protected"
      method: "GET"
      count: 100
      header:
        Authorization: "Bearer your-token-here"
        X-API-Key: "your-api-key"
```

#### Connection Pooling

```yaml
stages:
  - http:
      url: "https://api.example.com/data"
      method: "GET"
      count: 1000
      maxIdleConns: 50
      maxConns: 100
```

### MongoDB

#### Insert Operations

```bash
./dlg mongodb \
  --uri mongodb://localhost:27017 \
  --database mydb \
  --collection users \
  --operation insert \
  --count 1000
```

With YAML:

```yaml
stages:
  - mongodb:
      uri: "mongodb://localhost:27017"
      database: "mydb"
      collection: "users"
      operation: "insert"
      count: 1000
      document:
        name: "Test User"
        email: "test@example.com"
        timestamp: "2024-01-01T00:00:00Z"
```

#### Find Operations

```yaml
stages:
  - mongodb:
      uri: "mongodb://localhost:27017"
      database: "mydb"
      collection: "users"
      operation: "find"
      count: 500
      filter:
        status: "active"
```

#### Update Operations

```yaml
stages:
  - mongodb:
      uri: "mongodb://localhost:27017"
      database: "mydb"
      collection: "users"
      operation: "update"
      count: 100
      filter:
        _id: 123
      update:
        $set:
          status: "updated"
```

### gRPC

#### Basic gRPC Call

```yaml
stages:
  - grpc:
      target: "localhost:50051"
      method: "helloworld.Greeter/SayHello"
      count: 100
      data: '{"name": "World"}'
      insecure: true
```

#### With Metadata

```yaml
stages:
  - grpc:
      target: "api.example.com:443"
      method: "api.v1.UserService/GetUser"
      count: 100
      data: '{"user_id": 123}'
      metadata:
        authorization: "Bearer token"
        x-request-id: "test-123"
```

#### With Timeout and Concurrency

```yaml
stages:
  - grpc:
      target: "localhost:50051"
      method: "api.Service/SlowOperation"
      count: 50
      timeout: 30s
      maxConcurrent: 10
      enableKeepalive: true
```

### WebSocket

```yaml
stages:
  - websocket:
      url: "ws://localhost:8080/ws"
      header:
        Origin: "http://localhost"
      ops:
        - write: "Hello"
        - read: true
        - write: "World"
        - read: true
```

### Redis

```bash
./dlg redis \
  --addr localhost:6379 \
  --operation set \
  --count 1000
```

### PostgreSQL

```yaml
stages:
  - postgresql:
      dsn: "host=localhost port=5432 user=postgres password=secret dbname=testdb sslmode=disable"
      query: "SELECT * FROM users WHERE status = $1"
      args: ["active"]
      count: 100
```

### ClickHouse

```yaml
stages:
  - clickhouse:
      dsn: "tcp://localhost:9000?username=default&password=&database=default"
      query: "SELECT count(*) FROM events WHERE date >= '2024-01-01'"
      count: 50
```

### Kafka

```yaml
stages:
  - kafka:
      brokers: ["localhost:9092"]
      topic: "test-topic"
      operation: "produce"
      count: 1000
      message:
        key: "test-key"
        value: "test message"
```

### TCP

```yaml
stages:
  - tcp:
      address: "localhost:8080"
      count: 100
      payload: "Hello TCP"
      payloadHex: "48656c6c6f"  # Alternative: hex-encoded
```

### DNS

```bash
./dlg dns \
  --server 8.8.8.8:53 \
  --domain example.com \
  --record-type A \
  --count 100
```

## Configuration Files

### Structure

A complete configuration file contains one or more plans:

```yaml
plans:
  - name: "plan-1"
    description: "First test plan"
    stages:
      - http:
          # HTTP configuration
      - mongodb:
          # MongoDB configuration

  - name: "plan-2"
    description: "Second test plan"
    stages:
      - grpc:
          # gRPC configuration
```

### Multi-Stage Plans

Plans can contain multiple stages that execute sequentially:

```yaml
plans:
  - name: "complex-test"
    description: "Multi-stage load test"
    stages:
      # Stage 1: Populate database
      - mongodb:
          uri: "mongodb://localhost:27017"
          database: "testdb"
          collection: "users"
          operation: "insert"
          count: 1000
          document:
            name: "Test User"
            status: "active"

      # Stage 2: Query database
      - mongodb:
          uri: "mongodb://localhost:27017"
          database: "testdb"
          collection: "users"
          operation: "find"
          count: 5000
          filter:
            status: "active"

      # Stage 3: Test API
      - http:
          url: "https://api.example.com/users"
          method: "GET"
          count: 2000
```

### Environment Variables

Use environment variables in configurations:

```yaml
stages:
  - mongodb:
      uri: "${MONGODB_URI}"
      database: "${DB_NAME}"
      collection: "users"
      operation: "insert"
      count: 1000
```

### Data Encoding Options

DLG supports multiple data encoding formats:

```yaml
stages:
  # Plain text
  - tcp:
      address: "localhost:8080"
      payload: "Hello"

  # Hex-encoded
  - tcp:
      address: "localhost:8080"
      payloadHex: "48656c6c6f"

  # Base64-encoded
  - http:
      url: "https://api.example.com"
      method: "POST"
      bodyBase64: "SGVsbG8gV29ybGQ="
```

## MCP Server

The MCP (Model Context Protocol) server allows AI agents to use DLG for automated load testing.

### Starting the MCP Server

```bash
./dlg mcp
```

The server communicates over stdio using JSON-RPC 2.0.

### Available Tools

The MCP server exposes 35+ tools, including:

- `run_http_load` - Execute HTTP load tests
- `run_mongodb_load` - Execute MongoDB load tests
- `run_grpc_load` - Execute gRPC load tests
- `run_websocket_load` - Execute WebSocket load tests
- `create_plan` - Create a multi-stage plan
- `execute_plan` - Execute a saved plan
- `list_plans` - List all available plans
- `delete_plan` - Delete a plan

### Integration with Claude Desktop

Add to your Claude Desktop configuration:

```json
{
  "mcpServers": {
    "dlg": {
      "command": "/path/to/dlg",
      "args": ["mcp"]
    }
  }
}
```

### Example AI Interaction

```
User: Test my HTTP API at https://api.example.com/users with 1000 requests

AI: I'll use DLG to run a load test on your API endpoint.
    [Calls run_http_load tool with appropriate parameters]

    Results: Successfully completed 1000 requests to https://api.example.com/users
```

## Metrics and Monitoring

### Prometheus Metrics

DLG automatically collects Prometheus metrics for all operations.

#### Exposing Metrics

Metrics are exposed on the `/metrics` endpoint when running the MCP server or using plan execution.

#### Key Metrics

**Plan Execution:**
- `executor_plan_stages_total` - Total number of stages executed
- `executor_plan_stage_duration` - Duration of stage execution

**HTTP Executor:**
- `client_in_flight_requests` - Currently active requests
- `client_api_requests_total` - Total requests by status code and method
- `dns_duration_ms` - DNS lookup latency
- `tls_duration_ms` - TLS handshake latency
- `request_duration_ms` - Request latency

### Prometheus Configuration

Add to your `prometheus.yml`:

```yaml
scrape_configs:
  - job_name: 'dlg'
    static_configs:
      - targets: ['localhost:9090']
```

### Grafana Dashboards

Create dashboards to visualize:

- Request rate
- Error rate
- Latency percentiles (p50, p95, p99)
- Success vs. failure ratio
- Concurrent connections

## Advanced Usage

### Rate Limiting

Control the rate of operations:

```yaml
stages:
  - http:
      url: "https://api.example.com"
      count: 10000
      rateLimit: 100  # 100 requests per second
```

### Connection Pooling

Configure connection pool settings:

```yaml
stages:
  - http:
      url: "https://api.example.com"
      count: 5000
      maxIdleConns: 100
      maxConns: 200
```

### Timeout Configuration

Set timeouts for operations:

```yaml
stages:
  - grpc:
      target: "localhost:50051"
      method: "service.Method"
      count: 100
      timeout: 30s
```

### Concurrent Execution

Control concurrency:

```yaml
stages:
  - grpc:
      target: "localhost:50051"
      method: "service.Method"
      count: 1000
      maxConcurrent: 50
```

### Authentication

#### HTTP Bearer Token

```yaml
stages:
  - http:
      url: "https://api.example.com"
      header:
        Authorization: "Bearer YOUR_TOKEN"
```

#### MongoDB Authentication

```yaml
stages:
  - mongodb:
      uri: "mongodb://username:password@localhost:27017"
      database: "mydb"
```

#### gRPC Metadata

```yaml
stages:
  - grpc:
      target: "api.example.com:443"
      metadata:
        authorization: "Bearer YOUR_TOKEN"
```

## Troubleshooting

### Common Issues

#### Connection Refused

```
Error: connection refused
```

**Solution:** Verify the target service is running and accessible.

#### Timeout Errors

```
Error: context deadline exceeded
```

**Solution:** Increase timeout or reduce load:

```yaml
stages:
  - http:
      timeout: 60s
      maxConcurrent: 10
```

#### Out of Memory

**Solution:** Reduce concurrency and request count:

```yaml
stages:
  - http:
      count: 1000  # Reduced from 100000
      maxConcurrent: 10  # Limit concurrent requests
```

#### TLS Certificate Errors

```
Error: x509: certificate signed by unknown authority
```

**Solution:** Use insecure mode (testing only):

```yaml
stages:
  - grpc:
      insecure: true
```

### Debugging

Enable verbose logging:

```bash
./dlg http --url https://api.example.com --count 10 --verbose
```

Check configuration:

```bash
./dlg plan -f plan.yaml --dry-run
```

### Performance Tips

1. **Connection Pooling:** Always configure appropriate pool sizes
2. **Rate Limiting:** Use rate limits to avoid overwhelming targets
3. **Metrics:** Monitor metrics to identify bottlenecks
4. **Concurrency:** Tune concurrency based on system resources
5. **Timeouts:** Set realistic timeouts for operations

## Best Practices

### Load Test Design

1. **Start Small:** Begin with low load and gradually increase
2. **Realistic Data:** Use production-like data
3. **Monitor Both Sides:** Watch both DLG and target system metrics
4. **Baseline First:** Establish baseline performance before optimization
5. **Automate:** Use configuration files for repeatable tests

### Safety

1. **Test Environment:** Always test in non-production first
2. **Rate Limits:** Don't overwhelm production systems
3. **Cleanup:** Clean up test data after completion
4. **Monitoring:** Watch for resource exhaustion
5. **Gradual Ramp:** Increase load gradually

### CI/CD Integration

Run load tests in CI/CD:

```bash
#!/bin/bash
./dlg plan -f tests/load-test.yaml
if [ $? -eq 0 ]; then
  echo "Load test passed"
  exit 0
else
  echo "Load test failed"
  exit 1
fi
```

## Examples

### Complete E2E Test

```yaml
plans:
  - name: "e2e-test"
    description: "End-to-end application test"
    stages:
      # 1. Setup: Create test data
      - mongodb:
          uri: "mongodb://localhost:27017"
          database: "testdb"
          collection: "users"
          operation: "insert"
          count: 100
          document:
            name: "Test User"
            email: "test@example.com"
            status: "active"

      # 2. Load test API
      - http:
          url: "https://api.example.com/users"
          method: "GET"
          count: 1000
          header:
            Accept: "application/json"

      # 3. Test WebSocket
      - websocket:
          url: "ws://localhost:8080/ws"
          ops:
            - write: '{"type": "subscribe", "channel": "users"}'
            - read: true

      # 4. Cleanup: Remove test data
      - mongodb:
          uri: "mongodb://localhost:27017"
          database: "testdb"
          collection: "users"
          operation: "delete"
          count: 1
          filter:
            name: "Test User"
```

## Resources

- [Architecture Documentation](ARCHITECTURE.md)
- [API Documentation](API.md)
- [Contributing Guidelines](CONTRIBUTING.md)
- [GitHub Repository](https://github.com/hodgesds/dlg)

## Getting Help

- Report issues: https://github.com/hodgesds/dlg/issues
- Discussions: https://github.com/hodgesds/dlg/discussions
- Examples: See `/config/example.yaml`
