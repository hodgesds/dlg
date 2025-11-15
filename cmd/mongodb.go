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
	mongodbconfig "github.com/hodgesds/dlg/config/mongodb"
	mongodbexec "github.com/hodgesds/dlg/executor/mongodb"
	stageexec "github.com/hodgesds/dlg/executor/stage"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
)

var (
	mongoDatabase   string
	mongoCollection string
	mongoOperation  string
)

// mongodbCmd represents the mongodb command
var mongodbCmd = &cobra.Command{
	Use:   "mongodb",
	Short: "MongoDB load generator",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("mongodb requires a URI")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		plan := defaultPlan("mongodb")

		children := []*config.Stage{}
		for i, arg := range args {
			child := &config.Stage{
				Name: fmt.Sprintf("%s-%d", plan.Stages[0].Name, i),
				Tags: tags,
				MongoDB: &mongodbconfig.Config{
					URI:        arg,
					Database:   mongoDatabase,
					Collection: mongoCollection,
					Operation:  mongodbconfig.Operation(mongoOperation),
				},
			}
			children = append(children, child)
		}

		plan.Stages[0].Children = children

		reg := prometheus.NewPedanticRegistry()

		stageExec, err := stageexec.New(
			stageexec.Params{
				Registry: reg,
				MongoDB:  mongodbexec.New(),
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		execPlan(plan, reg, stageExec)
	},
}

func init() {
	RootCmd.AddCommand(mongodbCmd)

	mongodbCmd.PersistentFlags().AddFlagSet(planFlags())
	mongodbCmd.AddCommand(newDocCmd())
	mongodbCmd.PersistentFlags().
		StringVarP(&mongoDatabase, "database", "d", "test", "MongoDB database")
	mongodbCmd.PersistentFlags().
		StringVarP(&mongoCollection, "collection", "c", "test", "MongoDB collection")
	mongodbCmd.PersistentFlags().
		StringVarP(&mongoOperation, "operation", "o", "find", "MongoDB operation (insert, find, update, delete, count, aggregate)")
}
