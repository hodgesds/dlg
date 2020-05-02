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
	dnsconfig "github.com/hodgesds/dlg/config/dns"
	dnsexec "github.com/hodgesds/dlg/executor/dns"
	stageexec "github.com/hodgesds/dlg/executor/stage"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
)

// dnsCmd represents the dns command
var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "dns load generator",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		plan := defaultPlan("dns")

		children := []*config.Stage{}
		for i, arg := range args {
			child := &config.Stage{
				Name: fmt.Sprintf("%s-%d", plan.Stages[0].Name, i),
				Tags: tags,
				DNS: &dnsconfig.Config{
					ResourceRecords: []string{arg},
				},
			}
			children = append(children, child)
		}

		plan.Stages[0].Children = children

		reg := prometheus.NewPedanticRegistry()

		stageExec, err := stageexec.New(
			stageexec.Params{
				Registry: reg,
				DNS:      dnsexec.New(),
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		execPlan(plan, reg, stageExec)
	},
}

func init() {
	RootCmd.AddCommand(dnsCmd)

	dnsCmd.PersistentFlags().AddFlagSet(planFlags())
	dnsCmd.AddCommand(newDocCmd())
}
