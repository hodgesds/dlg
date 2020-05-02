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
	"os"
	"os/user"

	"github.com/hodgesds/dlg/config"
	sshconf "github.com/hodgesds/dlg/config/ssh"
	"github.com/hodgesds/dlg/executor"
	"github.com/hodgesds/dlg/executor/ssh"
	stageexec "github.com/hodgesds/dlg/executor/stage"
	"github.com/hodgesds/dlg/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
)

var (
	sshAddr    string
	sshUser    string
	sshKeyFile string
	sshExec    string
)

// sshCmd represents the ssh command
var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "ssh load generator",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		plan := &config.Plan{
			Name: name,
			Tags: tags,
		}
		stage := &config.Stage{
			Name:       fmt.Sprintf("%s-ssh", name),
			Tags:       tags,
			Repeat:     repeat,
			Concurrent: true,
			Children:   []*config.Stage{},
			SSH: &sshconf.Config{
				Addr:    sshAddr,
				User:    sshUser,
				KeyFile: sshKeyFile,
			},
		}
		if sshExec != "" {
			stage.SSH.Cmd = &sshExec
		}
		if dur > 0 {
			stage.Duration = &dur
		}

		plan.Stages = []*config.Stage{stage}

		reg := prometheus.NewPedanticRegistry()

		stageExec, err := stageexec.New(
			stageexec.Params{
				Registry: reg,
				SSH:      ssh.New(),
			},
		)

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
	RootCmd.AddCommand(sshCmd)

	sshCmd.PersistentFlags().StringVarP(
		&sshAddr,
		"addr", "a",
		"127.0.0.1:22",
		"SSH remote address",
	)
	sshCmd.PersistentFlags().StringVarP(
		&sshKeyFile,
		"key", "k",
		"",
		"SSH key file",
	)
	sshCmd.PersistentFlags().StringVarP(
		&sshExec,
		"exec", "e",
		"",
		"SSH command",
	)

	u, err := user.Current()
	if err == nil {
		sshUser = u.Username
	}
	sshCmd.PersistentFlags().StringVarP(
		&sshUser,
		"user", "u",
		sshUser,
		"SSH user",
	)

	sshCmd.PersistentFlags().AddFlagSet(planFlags())
	sshCmd.AddCommand(newDocCmd())
}
