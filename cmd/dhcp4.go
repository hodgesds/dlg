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

	"github.com/spf13/cobra"
)

// dhcp4Cmd represents the dhcp4 command
var dhcp4Cmd = &cobra.Command{
	Use:   "dhcp4",
	Short: "dhcp4 load generator",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dhcp4 called")
	},
}

func init() {
	RootCmd.AddCommand(dhcp4Cmd)

	dhcp4Cmd.PersistentFlags().AddFlagSet(planFlags())
	dhcp4Cmd.AddCommand(newDocCmd())
}
