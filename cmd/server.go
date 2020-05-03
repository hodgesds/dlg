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
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hodgesds/dlg"
	etcdconf "github.com/hodgesds/dlg/config/etcd"
	"github.com/hodgesds/dlg/executor"
	"github.com/hodgesds/dlg/executor/stage"
	"github.com/hodgesds/dlg/manager/etcd"
	xhttp "github.com/hodgesds/dlg/util/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
)

var (
	serverHTTPBind string
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		reg := prometheus.NewPedanticRegistry()
		stageExec, err := stage.Default(reg)
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

		managerConfig := &etcdconf.Config{
			Endpoints: etcdEndpoints,
		}
		m, err := etcd.NewManager(managerConfig, planExec)
		if err != nil {
			log.Fatal(err)
		}

		p := defaultPlan("server")
		m.Add(context.Background(), p)

		r := gin.Default()
		dlg.NewManagerRouter(r, m)

		r.Use(gin.WrapH(xhttp.StageMiddleware(nil)))
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		r.Run(serverHTTPBind)
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
	serverCmd.PersistentFlags().
		StringVarP(&serverHTTPBind, "bind", "b", ":8333", "HTTP address")
	serverCmd.PersistentFlags().StringSliceVarP(
		&etcdEndpoints,
		"endpoints", "e",
		[]string{},
		"ETCD endpoints",
	)
}
