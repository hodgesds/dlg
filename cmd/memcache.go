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

	memcacheconfig "github.com/hodgesds/dlg/config/memcache"
	memcacheexec "github.com/hodgesds/dlg/executor/memcache"
	stageexec "github.com/hodgesds/dlg/executor/stage"
	"github.com/hodgesds/dlg/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
)

var (
	memcacheEndpoint = []string{}
)

// memcacheCmd represents the memcache command.
var memcacheCmd = &cobra.Command{
	Use:   "memcache",
	Short: "memcache load generator",
	Long:  ``,
}

// memcacheSetCmd represents the memcache set command.
var memcacheSetCmd = &cobra.Command{
	Use:   "set",
	Short: "memcache set generator",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		plan := defaultPlan("memcache")

		ops := []*memcacheconfig.Op{}
		for _, arg := range args {
			k, v, err := util.ParseMemcacheKV(arg)
			if err != nil {
				log.Fatal(err)
			}
			ops = append(ops, &memcacheconfig.Op{
				Set: &memcacheconfig.Set{
					Key:   k,
					Value: v,
				},
			},
			)
		}
		plan.Stages[0].Memcache = &memcacheconfig.Config{
			Addrs: memcacheEndpoint,
			Ops:   ops,
		}

		reg := prometheus.NewPedanticRegistry()

		stageExec, err := stageexec.New(
			stageexec.Params{
				Registry: reg,
				Memcache: memcacheexec.New(),
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		execPlan(plan, reg, stageExec)
	},
}

// memcacheGetCmd represents the memcache get command.
var memcacheGetCmd = &cobra.Command{
	Use:   "get",
	Short: "memcache get generator",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		plan := defaultPlan("memcache")

		ops := []*memcacheconfig.Op{}
		for _, arg := range args {
			ops = append(ops, &memcacheconfig.Op{
				Get: &memcacheconfig.Get{
					Key: arg,
				},
			},
			)
		}
		plan.Stages[0].Memcache = &memcacheconfig.Config{
			Addrs: memcacheEndpoint,
			Ops:   ops,
		}

		reg := prometheus.NewPedanticRegistry()

		stageExec, err := stageexec.New(
			stageexec.Params{
				Registry: reg,
				Memcache: memcacheexec.New(),
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		execPlan(plan, reg, stageExec)
	},
}

// memcacheDeleteCmd represents the memcache get command.
var memcacheDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "memcache delete generator",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		plan := defaultPlan("memcache")

		ops := []*memcacheconfig.Op{}
		for _, arg := range args {
			ops = append(ops, &memcacheconfig.Op{
				Delete: &memcacheconfig.Delete{
					Key: arg,
				},
			},
			)
		}
		plan.Stages[0].Memcache = &memcacheconfig.Config{
			Addrs: memcacheEndpoint,
			Ops:   ops,
		}

		reg := prometheus.NewPedanticRegistry()

		stageExec, err := stageexec.New(
			stageexec.Params{
				Registry: reg,
				Memcache: memcacheexec.New(),
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		execPlan(plan, reg, stageExec)
	},
}

func init() {
	RootCmd.AddCommand(memcacheCmd)

	memcacheCmd.PersistentFlags().AddFlagSet(planFlags())
	memcacheCmd.PersistentFlags().StringSliceVarP(
		&memcacheEndpoint,
		"endpoint", "e",
		[]string{},
		"Memcache endpoint(s)",
	)
	memcacheCmd.MarkFlagRequired("endpoint")

	memcacheCmd.AddCommand(newDocCmd())
	memcacheCmd.AddCommand(memcacheGetCmd)
	memcacheCmd.AddCommand(memcacheSetCmd)
	memcacheCmd.AddCommand(memcacheDeleteCmd)
}
