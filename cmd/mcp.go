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

package cmd

import (
	"context"
	"log"

	"github.com/hodgesds/dlg/mcpserver"
	"github.com/spf13/cobra"
)

// mcpCmd represents the mcp command
var mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Start MCP server for AI agent integration",
	Long: `Start the Model Context Protocol (MCP) server to enable AI agents to use DLG for load testing.

The MCP server exposes load generation capabilities as tools that can be called by AI agents:
- http_load_test: Generate HTTP load
- redis_load_test: Generate Redis load
- mongodb_load_test: Generate MongoDB load
- postgres_load_test: Generate PostgreSQL load
- websocket_load_test: Generate WebSocket load
- grpc_load_test: Generate gRPC load
- run_load_plan: Execute a YAML-based load test plan

The server communicates over stdio, which is the standard transport for MCP servers.`,
	Run: func(cmd *cobra.Command, args []string) {
		server := mcpserver.New()
		if err := server.Run(context.Background()); err != nil {
			log.Fatalf("MCP server error: %v", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(mcpCmd)
}
