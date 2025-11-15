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
	syslogconfig "github.com/hodgesds/dlg/config/syslog"
	stageexec "github.com/hodgesds/dlg/executor/stage"
	syslogexec "github.com/hodgesds/dlg/executor/syslog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
)

var (
	syslogMessage  string
	syslogProtocol string
	syslogSeverity int
)

// syslogCmd represents the syslog command
var syslogCmd = &cobra.Command{
	Use:   "syslog",
	Short: "Syslog load generator",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("syslog requires a server address")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		plan := defaultPlan("syslog")

		children := []*config.Stage{}
		for i, arg := range args {
			child := &config.Stage{
				Name: fmt.Sprintf("%s-%d", plan.Stages[0].Name, i),
				Tags: tags,
				Syslog: &syslogconfig.Config{
					Server:   arg,
					Message:  syslogMessage,
					Protocol: syslogProtocol,
					Severity: syslogSeverity,
				},
			}
			children = append(children, child)
		}

		plan.Stages[0].Children = children

		reg := prometheus.NewPedanticRegistry()

		stageExec, err := stageexec.New(
			stageexec.Params{
				Registry: reg,
				Syslog:   syslogexec.New(),
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		execPlan(plan, reg, stageExec)
	},
}

func init() {
	RootCmd.AddCommand(syslogCmd)

	syslogCmd.PersistentFlags().AddFlagSet(planFlags())
	syslogCmd.AddCommand(newDocCmd())
	syslogCmd.PersistentFlags().
		StringVarP(&syslogMessage, "message", "m", "test message", "Syslog message")
	syslogCmd.PersistentFlags().
		StringVarP(&syslogProtocol, "protocol", "p", "udp", "Protocol (udp or tcp)")
	syslogCmd.PersistentFlags().
		IntVarP(&syslogSeverity, "severity", "s", 6, "Severity level (0-7)")
}
