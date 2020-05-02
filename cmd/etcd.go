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
	etcdconf "github.com/hodgesds/dlg/config/etcd"
	"github.com/hodgesds/dlg/executor"
	etcdexec "github.com/hodgesds/dlg/executor/etcd"
	stageexec "github.com/hodgesds/dlg/executor/stage"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
)

var (
	etcdEndpoints = []string{}
	etcdKeys      = []string{}
)

// etcdCmd represents the etcd command
var etcdCmd = &cobra.Command{
	Use:   "etcd",
	Short: "etcd load generator",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		plan := &config.Plan{
			Name: name,
			Tags: tags,
		}
		stage := &config.Stage{
			Name:     fmt.Sprintf("%s-etcd", name),
			Tags:     tags,
			Repeat:   repeat,
			Children: []*config.Stage{},
		}
		if dur > 0 {
			stage.Duration = &dur
		}
		for _, key := range etcdKeys {
			kv, err := etcdconf.ParseKV(key)
			if err != nil {
				log.Fatal(err)
			}
			stage.Children = append(
				stage.Children,
				&config.Stage{
					ETCD: &etcdconf.Config{
						Endpoints: etcdEndpoints,
						KV:        []*etcdconf.KV{kv},
					},
				})
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
				ETCD:     etcdexec.New(),
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

		if err := util.RegistryGather(reg, os.Stdout); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(etcdCmd)
	etcdCmd.PersistentFlags().StringSliceVarP(
		&etcdEndpoints,
		"endpoints", "e",
		[]string{},
		"ETCD endpoints",
	)
	etcdCmd.PersistentFlags().StringSliceVarP(
		&etcdKeys,
		"keys", "k",
		[]string{},
		"ETCD keys",
	)

	etcdCmd.PersistentFlags().AddFlagSet(planFlags())
	etcdCmd.AddCommand(newDocCmd())
}
