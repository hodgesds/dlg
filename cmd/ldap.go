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

// ldapCmd represents the ldap command
var ldapCmd = &cobra.Command{
	Use:   "ldap",
	Short: "ldap load generator",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("ldap called")
	},
}

func init() {
	RootCmd.AddCommand(ldapCmd)

	ldapCmd.PersistentFlags().AddFlagSet(planFlags())
	ldapCmd.AddCommand(newDocCmd())
}
