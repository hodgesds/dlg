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
}
