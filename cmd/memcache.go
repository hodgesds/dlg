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
	memcacheconfig "github.com/hodgesds/dlg/config/memcache"
	memcacheexec "github.com/hodgesds/dlg/executor/memcache"
	stageexec "github.com/hodgesds/dlg/executor/stage"
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

// memcacheGetCmd represents the memcache get command.
var memcacheGetCmd = &cobra.Command{
	Use:   "get",
	Short: "memcache get generator",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		plan := &config.Plan{
			Name: name,
			Tags: tags,
		}
		stage := &config.Stage{
			Name:       fmt.Sprintf("%s-memcache", name),
			Tags:       tags,
			Repeat:     repeat,
			Concurrent: true,
			Children:   []*config.Stage{},
		}
		if dur > 0 {
			stage.Duration = &dur
		}

		ops := []*memcacheconfig.Op{}
		for _, arg := range args {
			ops = append(ops, &memcacheconfig.Op{
				Get: &memcacheconfig.Get{
					Key: arg,
				},
			},
			)
		}
		stage.Memcache = &memcacheconfig.Config{
			Addrs: memcacheEndpoint,
			Ops:   ops,
		}

		plan.Stages = []*config.Stage{stage}

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
}
