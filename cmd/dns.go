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
	dnsconfig "github.com/hodgesds/dlg/config/dns"
	"github.com/hodgesds/dlg/executor"
	dnsexec "github.com/hodgesds/dlg/executor/dns"
	stageexec "github.com/hodgesds/dlg/executor/stage"
	"github.com/spf13/cobra"
)

// dnsCmd represents the dns command
var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "dns load generator",
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
				DNS: &dnsconfig.Config{
					ResourceRecords: []string{arg},
				},
			}
			stage.Children = append(stage.Children, child)
		}

		plan.Stages = []*config.Stage{stage}

		planExec := executor.NewPlan(stageexec.New(
			stageexec.Params{
				DNS: dnsexec.New(),
			},
		))

		err := planExec.Execute(context.Background(), plan)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(dnsCmd)

	dnsCmd.PersistentFlags().AddFlagSet(planFlags())
}
