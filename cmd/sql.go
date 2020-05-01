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
	"context"
	"fmt"
	"log"

	"github.com/hodgesds/dlg/config"
	sqlconf "github.com/hodgesds/dlg/config/sql"
	"github.com/hodgesds/dlg/executor"
	sqlexec "github.com/hodgesds/dlg/executor/sql"
	stageexec "github.com/hodgesds/dlg/executor/stage"
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
		plan := &config.Plan{
			Name: name,
			Tags: tags,
		}
		stage := &config.Stage{
			Name:       fmt.Sprintf("%s-http", name),
			Tags:       tags,
			Repeat:     repeat,
			Concurrent: true,
			Children:   []*config.Stage{},
		}
		if dur > 0 {
			stage.Duration = &dur
		}

		for i, arg := range args {
			child := &config.Stage{
				Name: fmt.Sprintf("%s-%d", stage.Name, i),
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
			stage.Children = append(stage.Children, child)
		}

		plan.Stages = []*config.Stage{stage}

		planExec := executor.NewPlan(stageexec.New(
			stageexec.Params{
				SQL: sqlexec.New(),
			},
		))

		err := planExec.Execute(context.Background(), plan)
		if err != nil {
			log.Fatal(err)
		}
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
