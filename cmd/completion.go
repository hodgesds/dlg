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
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	completionShell  string
	completionOutput string
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate shell autocompletion",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var w io.Writer
		if docOutput == "-" {
			w = os.Stdout
		} else {
			f, err := os.OpenFile(docOutput, os.O_RDWR|os.O_CREATE, 0755)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			w = f
		}

		switch completionShell {
		case "bash":
			if err := cmd.Parent().GenBashCompletion(w); err != nil {
				log.Fatal(err)
			}
		case "fish":
			if err := cmd.Parent().GenFishCompletion(w, true); err != nil {
				log.Fatal(err)
			}
		case "powershell":
			if err := cmd.Parent().GenPowerShellCompletion(w); err != nil {
				log.Fatal(err)
			}
		case "zsh":
			if err := cmd.Parent().GenZshCompletion(w); err != nil {
				log.Fatal(err)
			}
		default:
			log.Fatalf("Unknown shell: %q", completionShell)
		}
	},
}

func init() {
	RootCmd.AddCommand(completionCmd)

	completionCmd.PersistentFlags().StringVarP(
		&completionShell,
		"shell", "s",
		"bash",
		"Shell to generation completion for",
	)

	completionCmd.PersistentFlags().StringVarP(
		&completionOutput,
		"output", "o",
		"-",
		"Output",
	)
}
