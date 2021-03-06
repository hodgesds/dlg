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
	"fmt"
	"log"

	"github.com/hodgesds/dlg/config"
	sqlconf "github.com/hodgesds/dlg/config/sql"
	sqlexec "github.com/hodgesds/dlg/executor/sql"
	stageexec "github.com/hodgesds/dlg/executor/stage"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
)

var (
	sqlPgDSN string
	sqlMyDSN string
	sqlChDSN string
)

// sqlCmd represents the sql command
var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "SQL load generator",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		plan := defaultPlan("sql")
		for i, arg := range args {
			child := &config.Stage{
				Name: fmt.Sprintf("%s-%d", plan.Stages[0].Name, i),
				Tags: tags,
				SQL: &sqlconf.Config{
					Payloads: []*sqlconf.Payload{
						{
							Exec: arg,
						},
					},
				},
			}
			if sqlChDSN != "" {
				child.SQL.ClickHouseDSN = sqlChDSN
			}
			if sqlMyDSN != "" {
				child.SQL.MysqlDSN = sqlMyDSN
			}
			if sqlPgDSN != "" {
				child.SQL.PostgresDSN = sqlPgDSN
			}
			plan.Stages[0].Children = append(plan.Stages[0].Children, child)
		}

		reg := prometheus.NewPedanticRegistry()

		stageExec, err := stageexec.New(
			stageexec.Params{
				Registry: reg,
				SQL:      sqlexec.New(),
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		execPlan(plan, reg, stageExec)
	},
}

func init() {
	RootCmd.AddCommand(sqlCmd)
	sqlCmd.PersistentFlags().StringVarP(
		&sqlChDSN,
		"ch-dsn", "k",
		"",
		"ClickHouse DSN",
	)
	sqlCmd.PersistentFlags().StringVarP(
		&sqlMyDSN,
		"my-dsn", "m",
		"",
		"MySQL DSN",
	)
	sqlCmd.PersistentFlags().StringVarP(
		&sqlPgDSN,
		"pg-dsn", "p",
		"",
		"Postgres DSN",
	)
	sqlCmd.PersistentFlags().AddFlagSet(planFlags())
	sqlCmd.AddCommand(newDocCmd())
}
