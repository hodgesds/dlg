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
	"github.com/hodgesds/dlg/config/grpc"
	"github.com/hodgesds/dlg/config/http"
	"github.com/hodgesds/dlg/config/mongodb"
	"github.com/hodgesds/dlg/config/redis"
	"github.com/hodgesds/dlg/config/sql"
	"github.com/hodgesds/dlg/config/websocket"
	"github.com/hodgesds/dlg/executor"
	grpcexec "github.com/hodgesds/dlg/executor/grpc"
	httpexec "github.com/hodgesds/dlg/executor/http"
	mongodbexec "github.com/hodgesds/dlg/executor/mongodb"
	redisexec "github.com/hodgesds/dlg/executor/redis"
	sqlexec "github.com/hodgesds/dlg/executor/sql"
	stageexec "github.com/hodgesds/dlg/executor/stage"
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

// PostgresLoadTestInput defines input parameters for PostgreSQL load testing
type PostgresLoadTestInput struct {
	DSN       string   `json:"dsn" jsonschema:"required,description=PostgreSQL DSN connection string"`
	Query     string   `json:"query" jsonschema:"description=SQL query to execute"`
	Count     int      `json:"count" jsonschema:"description=Number of queries to execute"`
	MaxConns  int      `json:"max_conns" jsonschema:"description=Maximum number of connections"`
	MaxIdle   int      `json:"max_idle" jsonschema:"description=Maximum number of idle connections"`
	Operation string   `json:"operation" jsonschema:"description=Operation type (select, insert, update, delete)"`
	Table     string   `json:"table" jsonschema:"description=Table name for operations"`
}

// PostgresLoadTestOutput defines the output of PostgreSQL load testing
type PostgresLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// WebSocketLoadTestInput defines input parameters for WebSocket load testing
type WebSocketLoadTestInput struct {
	URL       string `json:"url" jsonschema:"required,description=WebSocket URL to connect to"`
	Count     int    `json:"count" jsonschema:"description=Number of messages to send"`
	Message   string `json:"message" jsonschema:"description=Message to send"`
}

// WebSocketLoadTestOutput defines the output of WebSocket load testing
type WebSocketLoadTestOutput struct {
	Message string `json:"message" jsonschema:"description=Status message"`
	Metrics string `json:"metrics" jsonschema:"description=Prometheus metrics from the load test"`
}

// GRPCLoadTestInput defines input parameters for gRPC load testing
type GRPCLoadTestInput struct {
	Target  string `json:"target" jsonschema:"required,description=gRPC server target (host:port)"`
	Method  string `json:"method" jsonschema:"description=gRPC method to call"`
	Count   int    `json:"count" jsonschema:"description=Number of requests to make"`
	TLS     bool   `json:"tls" jsonschema:"description=Use TLS connection"`
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
			Registry: reg,
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
			HTTP:          httpexec.New(reg),
			Redis:         redisexec.New(reg),
			MongoDB:       mongodbexec.New(reg),
			SQL:           sqlexec.New(reg),
			WebSocket:     websocketexec.New(reg),
			GRPC:          grpcexec.New(reg),
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
