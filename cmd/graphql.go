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
	graphqlconfig "github.com/hodgesds/dlg/config/graphql"
	graphqlexec "github.com/hodgesds/dlg/executor/graphql"
	stageexec "github.com/hodgesds/dlg/executor/stage"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
)

var (
	graphqlQuery string
)

// graphqlCmd represents the graphql command
var graphqlCmd = &cobra.Command{
	Use:   "graphql",
	Short: "GraphQL load generator",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("graphql requires an endpoint URL")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		plan := defaultPlan("graphql")

		children := []*config.Stage{}
		for i, arg := range args {
			child := &config.Stage{
				Name: fmt.Sprintf("%s-%d", plan.Stages[0].Name, i),
				Tags: tags,
				GraphQL: &graphqlconfig.Config{
					Endpoint: arg,
					Query:    graphqlQuery,
				},
			}
			children = append(children, child)
		}

		plan.Stages[0].Children = children

		reg := prometheus.NewPedanticRegistry()

		stageExec, err := stageexec.New(
			stageexec.Params{
				Registry: reg,
				GraphQL:  graphqlexec.New(),
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		execPlan(plan, reg, stageExec)
	},
}

func init() {
	RootCmd.AddCommand(graphqlCmd)

	graphqlCmd.PersistentFlags().AddFlagSet(planFlags())
	graphqlCmd.AddCommand(newDocCmd())
	graphqlCmd.PersistentFlags().
		StringVarP(&graphqlQuery, "query", "q", "", "GraphQL query")
}
