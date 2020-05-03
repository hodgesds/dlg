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
	"log"

	snmpconf "github.com/hodgesds/dlg/config/snmp"
	snmpexec "github.com/hodgesds/dlg/executor/snmp"
	stageexec "github.com/hodgesds/dlg/executor/stage"
	"github.com/hodgesds/dlg/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
)

var (
	snmpEndpoint  string
	snmpTransport string
	snmpVersion   string
	snmpCommunity string
	snmpOids      = []string{}
)

// snmpCmd represents the snmp command
var snmpCmd = &cobra.Command{
	Use:   "snmp",
	Short: "SNMP load generator",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		plan := defaultPlan("snmp")
		host, port, err := util.ParseSNMPEndpoint(snmpEndpoint)
		if err != nil {
			log.Fatal(err)
		}
		plan.Stages[0].SNMP = &snmpconf.Config{
			Addr:      host,
			Port:      port,
			Transport: snmpTransport,
			Version:   snmpVersion,
			Community: snmpCommunity,
			Oids:      snmpOids,
		}

		reg := prometheus.NewPedanticRegistry()

		stageExec, err := stageexec.New(
			stageexec.Params{
				Registry: reg,
				SNMP:     snmpexec.New(),
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		execPlan(plan, reg, stageExec)
	},
}

func init() {
	RootCmd.AddCommand(snmpCmd)
	snmpCmd.PersistentFlags().AddFlagSet(planFlags())
	snmpCmd.PersistentFlags().StringVarP(
		&snmpEndpoint,
		"endpoint", "e",
		"",
		"SNMP endpoints",
	)
	snmpCmd.PersistentFlags().StringSliceVarP(
		&snmpOids,
		"oids", "o",
		[]string{},
		"SNMP oids",
	)
	snmpCmd.PersistentFlags().StringVar(
		&snmpCommunity,
		"community",
		"",
		"SNMP community",
	)
	snmpCmd.PersistentFlags().StringVarP(
		&snmpTransport,
		"transport", "t",
		"tcp",
		"SNMP transport",
	)
	snmpCmd.PersistentFlags().StringVarP(
		&snmpVersion,
		"version", "v",
		"2",
		"SNMP version",
	)
	snmpCmd.AddCommand(newDocCmd())
}
