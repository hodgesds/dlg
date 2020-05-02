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
	"net"

	"github.com/hodgesds/dlg/config"
	dhcp4conf "github.com/hodgesds/dlg/config/dhcp4"
	"github.com/hodgesds/dlg/executor"
	"github.com/hodgesds/dlg/executor/dhcp4"
	stageexec "github.com/hodgesds/dlg/executor/stage"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
)

var (
	dhcp4Iface     string
	dhcp4HwAddrStr string
)

// dhcp4Cmd represents the dhcp4 command
var dhcp4Cmd = &cobra.Command{
	Use:   "dhcp4",
	Short: "dhcp4 load generator",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		addr, err := net.ParseMAC(dhcp4HwAddrStr)
		if err != nil {
			log.Fatal(err)
		}
		plan := &config.Plan{
			Name: name,
			Tags: tags,
		}
		stage := &config.Stage{
			Name:       fmt.Sprintf("%s-dhcp4", name),
			Tags:       tags,
			Repeat:     repeat,
			Concurrent: true,
			Children:   []*config.Stage{},
			DHCP4: &dhcp4conf.Config{
				Iface:  dhcp4Iface,
				HwAddr: addr,
			},
		}
		if dur > 0 {
			stage.Duration = &dur
		}

		plan.Stages = []*config.Stage{stage}

		reg := prometheus.NewPedanticRegistry()
		reg.MustRegister(
			prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
			prometheus.NewGoCollector(),
		)

		stageExec, err := stageexec.New(
			stageexec.Params{
				Registry: reg,
				DHCP4:    dhcp4.New(),
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		planExec, err := executor.NewPlan(
			executor.Params{Registry: reg},
			stageExec,
		)
		if err != nil {
			log.Fatal(err)
		}

		err = planExec.Execute(context.Background(), plan)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(dhcp4Cmd)

	dhcp4Cmd.PersistentFlags().StringVarP(
		&dhcp4Iface,
		"iface", "i",
		"eth0",
		"Network interface",
	)
	dhcp4Cmd.PersistentFlags().StringVarP(
		&dhcp4HwAddrStr,
		"addr", "a",
		"44:85:00:17:d6:53",
		"Hardware MAC address",
	)
	dhcp4Cmd.PersistentFlags().AddFlagSet(planFlags())
	dhcp4Cmd.AddCommand(newDocCmd())
}
