// Copyright © 2020 Daniel Hodges <hodges.daniel.scott@gmail.com>
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
	"errors"
	"fmt"
	"log"

	"github.com/hodgesds/dlg/config"
	clickhouseconfig "github.com/hodgesds/dlg/config/clickhouse"
	clickhouseexec "github.com/hodgesds/dlg/executor/clickhouse"
	stageexec "github.com/hodgesds/dlg/executor/stage"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
)

var (
	clickhouseDatabase  string
	clickhouseTable     string
	clickhouseOperation string
	clickhouseQuery     string
	clickhouseCount     int
)

// clickhouseCmd represents the clickhouse command
var clickhouseCmd = &cobra.Command{
	Use:   "clickhouse",
	Short: "ClickHouse load generator",
	Long:  `Generate load for ClickHouse databases with various operations like insert, select, batch_insert, count, optimize, and create_table.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("clickhouse requires a DSN")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		plan := defaultPlan("clickhouse")

		children := []*config.Stage{}
		for i, arg := range args {
			child := &config.Stage{
				Name: fmt.Sprintf("%s-%d", plan.Stages[0].Name, i),
				Tags: tags,
				ClickHouse: &clickhouseconfig.Config{
					DSN:       arg,
					Database:  clickhouseDatabase,
					Table:     clickhouseTable,
					Operation: clickhouseconfig.Operation(clickhouseOperation),
					Query:     clickhouseQuery,
					Count:     clickhouseCount,
				},
			}
			children = append(children, child)
		}

		plan.Stages[0].Children = children

		reg := prometheus.NewPedanticRegistry()

		stageExec, err := stageexec.New(
			stageexec.Params{
				Registry:   reg,
				ClickHouse: clickhouseexec.New(),
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		execPlan(plan, reg, stageExec)
	},
}

func init() {
	RootCmd.AddCommand(clickhouseCmd)

	clickhouseCmd.PersistentFlags().AddFlagSet(planFlags())
	clickhouseCmd.PersistentFlags().
		StringVarP(&clickhouseDatabase, "database", "d", "default", "ClickHouse database")
	clickhouseCmd.PersistentFlags().
		StringVarP(&clickhouseTable, "table", "t", "", "ClickHouse table")
	clickhouseCmd.PersistentFlags().
		StringVarP(&clickhouseOperation, "operation", "o", "select", "ClickHouse operation (insert, select, batch_insert, count, optimize, create_table)")
	clickhouseCmd.PersistentFlags().
		StringVarP(&clickhouseQuery, "query", "q", "", "ClickHouse query (for select/count operations)")
	clickhouseCmd.PersistentFlags().
		IntVarP(&clickhouseCount, "count", "n", 100, "Number of operations to execute")
}
