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
	pulsarconfig "github.com/hodgesds/dlg/config/pulsar"
	pulsarexec "github.com/hodgesds/dlg/executor/pulsar"
	stageexec "github.com/hodgesds/dlg/executor/stage"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
)

var (
	pulsarTopic   string
	pulsarMessage string
)

// pulsarCmd represents the pulsar command
var pulsarCmd = &cobra.Command{
	Use:   "pulsar",
	Short: "Apache Pulsar load generator",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("pulsar requires a service URL")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		plan := defaultPlan("pulsar")

		children := []*config.Stage{}
		for i, arg := range args {
			child := &config.Stage{
				Name: fmt.Sprintf("%s-%d", plan.Stages[0].Name, i),
				Tags: tags,
				Pulsar: &pulsarconfig.Config{
					ServiceURL: arg,
					Topic:      pulsarTopic,
					Message:    pulsarMessage,
				},
			}
			children = append(children, child)
		}

		plan.Stages[0].Children = children

		reg := prometheus.NewPedanticRegistry()

		stageExec, err := stageexec.New(
			stageexec.Params{
				Registry: reg,
				Pulsar:   pulsarexec.New(),
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		execPlan(plan, reg, stageExec)
	},
}

func init() {
	RootCmd.AddCommand(pulsarCmd)

	pulsarCmd.PersistentFlags().AddFlagSet(planFlags())
	pulsarCmd.AddCommand(newDocCmd())
	pulsarCmd.PersistentFlags().
		StringVarP(&pulsarTopic, "topic", "t", "test", "Pulsar topic")
	pulsarCmd.PersistentFlags().
		StringVarP(&pulsarMessage, "message", "m", "", "Message to produce")
}
