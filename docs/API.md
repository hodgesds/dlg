# DLG API Documentation

## Table of Contents

- [Core Interfaces](#core-interfaces)
- [Manager API](#manager-api)
- [Executor API](#executor-api)
- [Configuration API](#configuration-api)
- [Rate Limiter API](#rate-limiter-api)
- [MCP Server API](#mcp-server-api)

## Core Interfaces

### Executor Interface

All protocol executors implement a common interface pattern:

```go
package executor

import "context"

type Executor interface {
    Execute(ctx context.Context, config interface{}) error
}
```

### Protocol-Specific Executors

Each protocol defines its own executor interface:

```go
// HTTP Executor
type HTTP interface {
    Execute(ctx context.Context, conf *httpconf.Config) error
}

// MongoDB Executor
type MongoDB interface {
    Execute(ctx context.Context, conf *mongodbconf.Config) error
}

// gRPC Executor
type GRPC interface {
    Execute(ctx context.Context, conf *grpcconf.Config) error
}
```

## Manager API

The Manager orchestrates multi-stage load test plans.

### Interface

```go
package dlg

import (
    "context"
    "github.com/hodgesds/dlg/config"
)

type Manager interface {
    // Get retrieves a plan by name
    Get(ctx context.Context, name string) (*config.Plan, error)

    // Add adds or updates a plan
    Add(ctx context.Context, plan *config.Plan) error

    // Delete removes a plan
    Delete(ctx context.Context, name string) error

    // Plans returns all registered plans
    Plans(ctx context.Context) ([]*config.Plan, error)

    // Execute runs a plan
    Execute(ctx context.Context, plan *config.Plan) error
}
```

### Creating a Manager

```go
import (
    "github.com/hodgesds/dlg"
    "github.com/hodgesds/dlg/executor"
)

// Create a plan executor
planExec, err := executor.NewPlan(executor.Params{}, nil)
if err != nil {
    // handle error
}

// Create a manager
mgr := dlg.NewManager(planExec)
```

### Using the Manager

#### Adding a Plan

```go
import "github.com/hodgesds/dlg/config"

plan := &config.Plan{
    Name:        "my-test",
    Description: "Test plan description",
    Stages: []config.Stage{
        {
            HTTP: &httpconf.Config{
                Count: 1000,
                Payload: httpconf.Payload{
                    URL:    "https://api.example.com",
                    Method: "GET",
                },
            },
        },
    },
}

err := mgr.Add(ctx, plan)
```

#### Executing a Plan

```go
plan, err := mgr.Get(ctx, "my-test")
if err != nil {
    // handle error
}

err = mgr.Execute(ctx, plan)
```

#### Listing Plans

```go
plans, err := mgr.Plans(ctx)
for _, plan := range plans {
    fmt.Printf("Plan: %s - %s\n", plan.Name, plan.Description)
}
```

#### Deleting a Plan

```go
err := mgr.Delete(ctx, "my-test")
```

## Executor API

### HTTP Executor

```go
package http

import (
    "context"
    "github.com/prometheus/client_golang/prometheus"
)

// Create a new HTTP executor
func New(reg *prometheus.Registry) executor.HTTP

// Execute HTTP load test
func (e *httpExecutor) Execute(ctx context.Context, conf *httpconf.Config) error
```

#### HTTP Configuration

```go
package http

type Config struct {
    Count        int       // Number of requests
    MaxIdleConns *int      // Maximum idle connections
    MaxConns     *int      // Maximum connections per host
    Payload      Payload   // Request configuration
}

type Payload struct {
    URL        string           // Target URL
    Header     http.Header      // Request headers
    Method     string           // HTTP method (GET, POST, etc.)
    Body       []byte           // Request body
    BodyFile   *string          // Path to body file
    BodyHex    string           // Hex-encoded body
    BodyBase64 string           // Base64-encoded body
}
```

#### Example Usage

```go
import (
    "context"
    "github.com/hodgesds/dlg/config/http"
    httpexec "github.com/hodgesds/dlg/executor/http"
    "github.com/prometheus/client_golang/prometheus"
)

// Create executor
reg := prometheus.NewRegistry()
executor := httpexec.New(reg)

// Create configuration
count := 1000
config := &http.Config{
    Count: count,
    Payload: http.Payload{
        URL:    "https://api.example.com/users",
        Method: "GET",
    },
}

// Execute
err := executor.Execute(context.Background(), config)
```

### MongoDB Executor

```go
package mongodb

// Create a new MongoDB executor
func New(reg *prometheus.Registry) executor.MongoDB

// Execute MongoDB load test
func (e *mongodbExecutor) Execute(ctx context.Context, conf *mongodbconf.Config) error
```

#### MongoDB Configuration

```go
package mongodb

import "time"

type Operation string

const (
    OpInsert    Operation = "insert"
    OpFind      Operation = "find"
    OpUpdate    Operation = "update"
    OpDelete    Operation = "delete"
    OpCount     Operation = "count"
    OpAggregate Operation = "aggregate"
)

type Config struct {
    URI            string                 // MongoDB connection URI
    Database       string                 // Database name
    Collection     string                 // Collection name
    Operation      Operation              // Operation type
    Count          int                    // Number of operations
    Document       map[string]interface{} // Document for insert
    Filter         map[string]interface{} // Filter for queries
    Update         map[string]interface{} // Update document
    ConnectTimeout *time.Duration         // Connection timeout
    MaxPoolSize    *uint64                // Maximum pool size
    MinPoolSize    *uint64                // Minimum pool size
}
```

### gRPC Executor

```go
package grpc

// Create a new gRPC executor
func New(reg *prometheus.Registry) executor.GRPC

// Execute gRPC load test
func (e *grpcExecutor) Execute(ctx context.Context, conf *grpcconf.Config) error
```

#### gRPC Configuration

```go
package grpc

import "time"

type Config struct {
    Target          string            // Server address
    Method          string            // gRPC method to call
    Count           int               // Number of calls
    Data            string            // Request data (JSON)
    DataHex         string            // Hex-encoded request data
    DataBase64      string            // Base64-encoded request data
    Metadata        map[string]string // gRPC metadata
    Insecure        bool              // Skip TLS verification
    Timeout         *time.Duration    // Request timeout
    MaxConcurrent   *int              // Max concurrent requests
    EnableKeepalive bool              // Enable keepalive
}
```

### WebSocket Executor

```go
package websocket

// Create a new WebSocket executor
func New(reg *prometheus.Registry) executor.WebSocket

// Execute WebSocket load test
func (e *websocketExecutor) Execute(ctx context.Context, conf *websocketconf.Config) error
```

#### WebSocket Configuration

```go
package websocket

import "net/http"

type Config struct {
    URL    string      // WebSocket URL
    Header http.Header // Connection headers
    Ops    []*Op       // Operations to perform
}

type Op struct {
    Read  bool   // Read from connection
    Write string // Write to connection
}
```

## Configuration API

### Plan Structure

```go
package config

type Plan struct {
    Name        string  // Plan name
    Description string  // Plan description
    Stages      []Stage // Execution stages
}

type Stage struct {
    HTTP       *httpconf.Config
    MongoDB    *mongodbconf.Config
    GRPC       *grpcconf.Config
    WebSocket  *websocketconf.Config
    Redis      *redisconf.Config
    PostgreSQL *postgresqlconf.Config
    // ... other protocols
}
```

### Loading from YAML

```go
import (
    "github.com/hodgesds/dlg/config"
    "gopkg.in/yaml.v3"
)

// Read YAML file
data, err := ioutil.ReadFile("plan.yaml")
if err != nil {
    // handle error
}

// Parse YAML
var cfg config.Config
err = yaml.Unmarshal(data, &cfg)
if err != nil {
    // handle error
}

// Access plans
for _, plan := range cfg.Plans {
    fmt.Printf("Plan: %s\n", plan.Name)
}
```

### Generating YAML

```go
import (
    "github.com/hodgesds/dlg/config"
    "gopkg.in/yaml.v3"
)

plan := &config.Plan{
    Name: "test-plan",
    Stages: []config.Stage{
        {
            HTTP: &httpconf.Config{
                Count: 1000,
                Payload: httpconf.Payload{
                    URL:    "https://api.example.com",
                    Method: "GET",
                },
            },
        },
    },
}

cfg := &config.Config{
    Plans: []*config.Plan{plan},
}

// Marshal to YAML
data, err := yaml.Marshal(cfg)
if err != nil {
    // handle error
}

// Write to file
err = ioutil.WriteFile("plan.yaml", data, 0644)
```

## Rate Limiter API

### Interface

```go
package ratelimiter

import "context"

type RateLimiter interface {
    // Wait blocks until a single operation can proceed
    Wait(ctx context.Context) error

    // WaitBytes blocks until the specified bytes can be sent
    WaitBytes(ctx context.Context, bytes int) error

    // Reset resets the rate limiter
    Reset()
}
```

### Creating a Rate Limiter

```go
import "github.com/hodgesds/dlg/ratelimiter"

limiter := ratelimiter.NewLimiter()
```

### Using a Rate Limiter

```go
// Wait for permission to proceed
err := limiter.Wait(ctx)
if err != nil {
    // handle error
}

// Perform operation
// ...

// Wait for permission to send bytes
err = limiter.WaitBytes(ctx, 1024)
if err != nil {
    // handle error
}
```

## MCP Server API

The MCP server exposes DLG functionality as tools via the Model Context Protocol.

### Tool Categories

#### Load Testing Tools

- `run_http_load` - Execute HTTP load test
- `run_mongodb_load` - Execute MongoDB load test
- `run_redis_load` - Execute Redis load test
- `run_grpc_load` - Execute gRPC load test
- `run_websocket_load` - Execute WebSocket load test
- ... (35+ total tools)

#### Plan Management Tools

- `create_plan` - Create a new plan
- `execute_plan` - Execute a plan
- `list_plans` - List all plans
- `get_plan` - Get plan details
- `delete_plan` - Delete a plan

### Tool Input Schema

Each tool accepts a JSON input matching its configuration structure.

#### Example: run_http_load

```json
{
  "url": "https://api.example.com/users",
  "method": "GET",
  "count": 1000,
  "headers": {
    "Authorization": "Bearer token"
  },
  "body": "{\"name\": \"test\"}"
}
```

#### Example: create_plan

```json
{
  "name": "my-plan",
  "description": "Test plan",
  "stages": [
    {
      "http": {
        "url": "https://api.example.com",
        "method": "GET",
        "count": 100
      }
    }
  ]
}
```

### Starting the MCP Server

```go
import "github.com/hodgesds/dlg/mcpserver"

// Start server (communicates over stdio)
err := mcpserver.Run()
```

Or via CLI:

```bash
./dlg mcp
```

### MCP Protocol Messages

#### Tool Discovery

Request:
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "tools/list"
}
```

Response:
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "tools": [
      {
        "name": "run_http_load",
        "description": "Execute an HTTP load test",
        "inputSchema": { ... }
      }
    ]
  }
}
```

#### Tool Execution

Request:
```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "method": "tools/call",
  "params": {
    "name": "run_http_load",
    "arguments": {
      "url": "https://api.example.com",
      "method": "GET",
      "count": 100
    }
  }
}
```

Response:
```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "result": {
    "content": [
      {
        "type": "text",
        "text": "Successfully executed 100 HTTP requests"
      }
    ]
  }
}
```

## Metrics API

### Accessing Metrics

```go
import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "net/http"
)

// Create registry
reg := prometheus.NewRegistry()

// Create executor with metrics
executor := http.New(reg)

// Expose metrics endpoint
http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
http.ListenAndServe(":9090", nil)
```

### Plan Metrics

```go
package executor

type metrics struct {
    StagesTotal   *prometheus.CounterVec   // Stage execution count
    StageDuration *prometheus.HistogramVec // Stage duration
}
```

Access:
- `executor_plan_stages_total{stage="http"}` - Counter
- `executor_plan_stage_duration{stage="http"}` - Histogram

### HTTP Metrics

- `client_in_flight_requests` - Gauge
- `client_api_requests_total{code,method}` - Counter
- `dns_duration_ms{event}` - Histogram
- `tls_duration_ms{event}` - Histogram
- `request_duration_ms{method}` - Histogram

## Error Handling

### Common Errors

```go
// Plan not found
_, err := mgr.Get(ctx, "nonexistent")
// Error: no such plan: "nonexistent"

// Invalid configuration
err := executor.Execute(ctx, invalidConfig)
// Error: invalid URL

// Context cancelled
ctx, cancel := context.WithCancel(context.Background())
cancel()
err := executor.Execute(ctx, config)
// Error: context canceled
```

### Error Types

DLG uses standard Go error handling. Errors are returned from functions and should be checked:

```go
if err != nil {
    log.Fatalf("Error: %v", err)
}
```

For multiple errors (e.g., from concurrent operations), DLG uses `go.uber.org/multierr`:

```go
import "go.uber.org/multierr"

var errors error
errors = multierr.Append(errors, err1)
errors = multierr.Append(errors, err2)
return errors
```

## Best Practices

### Context Usage

Always use context for cancellation and timeouts:

```go
import "time"

// With timeout
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

err := executor.Execute(ctx, config)
```

### Resource Cleanup

Use defer for cleanup:

```go
conn, err := dial()
if err != nil {
    return err
}
defer conn.Close()
```

### Metrics Registration

Create one registry per application:

```go
// Application-wide registry
var Registry = prometheus.NewRegistry()

// Use same registry for all executors
httpExec := http.New(Registry)
grpcExec := grpc.New(Registry)
```

### Concurrency

Handle concurrency with proper synchronization:

```go
var (
    wg  sync.WaitGroup
    mu  sync.Mutex
    err error
)

for i := 0; i < count; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()

        if err2 := doWork(); err2 != nil {
            mu.Lock()
            err = multierr.Append(err, err2)
            mu.Unlock()
        }
    }()
}

wg.Wait()
return err
```

## Examples

### Complete Application

```go
package main

import (
    "context"
    "log"

    "github.com/hodgesds/dlg"
    "github.com/hodgesds/dlg/config"
    "github.com/hodgesds/dlg/config/http"
    "github.com/hodgesds/dlg/executor"
    "github.com/prometheus/client_golang/prometheus"
)

func main() {
    // Create plan executor
    planExec, err := executor.NewPlan(executor.Params{}, prometheus.NewRegistry())
    if err != nil {
        log.Fatal(err)
    }

    // Create manager
    mgr := dlg.NewManager(planExec)

    // Create plan
    plan := &config.Plan{
        Name: "api-test",
        Stages: []config.Stage{
            {
                HTTP: &http.Config{
                    Count: 1000,
                    Payload: http.Payload{
                        URL:    "https://api.example.com/users",
                        Method: "GET",
                    },
                },
            },
        },
    }

    // Add plan
    if err := mgr.Add(context.Background(), plan); err != nil {
        log.Fatal(err)
    }

    // Execute plan
    if err := mgr.Execute(context.Background(), plan); err != nil {
        log.Fatal(err)
    }

    log.Println("Load test completed successfully")
}
```

## References

- [User Guide](USER_GUIDE.md)
- [Architecture Documentation](ARCHITECTURE.md)
- [Contributing Guidelines](CONTRIBUTING.md)
- [Go Documentation](https://pkg.go.dev/github.com/hodgesds/dlg)
