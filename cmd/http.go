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
	"errors"
	"fmt"
	"log"

	"github.com/hodgesds/dlg/config"
	"github.com/hodgesds/dlg/config/http"
	"github.com/hodgesds/dlg/executor"
	httpexec "github.com/hodgesds/dlg/executor/http"
	stageexec "github.com/hodgesds/dlg/executor/stage"
	"github.com/hodgesds/dlg/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
)

var (
	httpConf = http.Config{
		Payload: http.Payload{
			BodyFile: util.StrPtr(""),
		},
		MaxConns:     util.IntPtr(0),
		MaxIdleConns: util.IntPtr(0),
	}
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "HTTP load generator", Long: ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("http requires a URL")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
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
				HTTP: &http.Config{
					Payload: http.Payload{
						URL:    arg,
						Method: httpConf.Payload.Method,
					},
					Count:        count,
					MaxConns:     httpConf.MaxConns,
					MaxIdleConns: httpConf.MaxIdleConns,
				},
			}
			if httpConf.Payload.BodyFile != nil && *httpConf.Payload.BodyFile != "" {
				child.HTTP.Payload.BodyFile = httpConf.Payload.BodyFile
			}
			stage.Children = append(stage.Children, child)
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
				HTTP:     httpexec.New(),
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
		return nil
	},
}

func init() {
	RootCmd.AddCommand(httpCmd)

	httpCmd.PersistentFlags().AddFlagSet(planFlags())
	httpCmd.AddCommand(newDocCmd())
	httpCmd.PersistentFlags().
		StringVarP(&httpConf.Payload.Method, "method", "m", "GET", "HTTP method")
	httpCmd.PersistentFlags().
		StringVarP(httpConf.Payload.BodyFile, "body-file", "b", "", "HTTP Body")
	//httpCmd.PersistentFlags().
	//	StringVar(httpConf.Payload.BodyBase64, "body-b64", nil, "HTTP Body (base64)")
	httpCmd.PersistentFlags().
		IntVar(httpConf.MaxConns, "max-conn", 0, "Max connections")
	httpCmd.PersistentFlags().
		IntVarP(httpConf.MaxIdleConns, "max-idle", "i", 0, "Max idle connections")
}
