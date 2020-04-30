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

	redisconf "github.com/hodgesds/dlg/config/redis"
	"github.com/spf13/cobra"
)

var (
	redisConf    = redisconf.Config{Commands: []*redisconf.Command{}}
	redisGetConf = &redisconf.Command{
		Get: &redisconf.Get{},
	}
)

// redisCmd represents the redis command
var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "redis load generator",
	Long:  ``,
}

// redisGetCmd represents the redis command
var redisGetCmd = &cobra.Command{
	Use:   "get",
	Short: "redis get",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		redisConf.Commands = []*redisconf.Command{redisGetConf}
		if err := redisConf.Execute(context.Background()); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(redisCmd)
	redisCmd.AddCommand(redisGetCmd)
	redisCmd.AddCommand(newDocCmd())
	redisGetCmd.PersistentFlags().AddFlagSet(planFlags())
	redisCmd.PersistentFlags().StringVar(
		&redisConf.Addr, "addr",
		"",
		"",
	)
	redisCmd.PersistentFlags().IntVar(
		&redisConf.DB, "db",
		0,
		"database",
	)
	redisCmd.PersistentFlags().StringVar(
		&redisConf.Password, "password",
		"",
		"",
	)
	redisGetCmd.PersistentFlags().StringVar(
		&redisGetConf.Get.Key, "key",
		"",
		"",
	)

	/* TODO: dynamically generate all these subcommands, redis is bloat.
	cmds, err := commandsFromStruct(&redisconf.Command{})
	if err != nil {
		panic(err.Error())
	}
	for _, cmd := range cmds {
		redisCmd.AddCommand(cmd)
	}
	*/
}
