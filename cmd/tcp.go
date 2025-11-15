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
	"strconv"
	"strings"

	"github.com/hodgesds/dlg/config"
	tcpconfig "github.com/hodgesds/dlg/config/tcp"
	stageexec "github.com/hodgesds/dlg/executor/stage"
	tcpexec "github.com/hodgesds/dlg/executor/tcp"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
)

var (
	tcpData      string
	tcpOperation string
)

// tcpCmd represents the tcp command
var tcpCmd = &cobra.Command{
	Use:   "tcp",
	Short: "TCP load generator",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("tcp requires host:port")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		plan := defaultPlan("tcp")

		children := []*config.Stage{}
		for i, arg := range args {
			parts := strings.Split(arg, ":")
			if len(parts) != 2 {
				log.Fatalf("invalid host:port format: %s", arg)
			}
			port, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatalf("invalid port: %s", parts[1])
			}

			child := &config.Stage{
				Name: fmt.Sprintf("%s-%d", plan.Stages[0].Name, i),
				Tags: tags,
				TCP: &tcpconfig.Config{
					Host:      parts[0],
					Port:      port,
					Data:      tcpData,
					Operation: tcpOperation,
				},
			}
			children = append(children, child)
		}

		plan.Stages[0].Children = children

		reg := prometheus.NewPedanticRegistry()

		stageExec, err := stageexec.New(
			stageexec.Params{
				Registry: reg,
				TCP:      tcpexec.New(),
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		execPlan(plan, reg, stageExec)
	},
}

func init() {
	RootCmd.AddCommand(tcpCmd)

	tcpCmd.PersistentFlags().AddFlagSet(planFlags())
	tcpCmd.AddCommand(newDocCmd())
	tcpCmd.PersistentFlags().
		StringVarP(&tcpData, "data", "d", "", "Data to send")
	tcpCmd.PersistentFlags().
		StringVarP(&tcpOperation, "operation", "o", "connect", "Operation (connect, send, send_receive)")
}
