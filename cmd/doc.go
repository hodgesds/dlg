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
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var (
	docFormat string
	docOutput string
)

// docCmd represents the docs command
var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "",
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
		switch f := strings.ToLower(docFormat); f {
		case "md", "markdown":
			if err := doc.GenMarkdown(cmd, w); err != nil {
				log.Fatal(err)
			}
		case "man":
			header := &doc.GenManHeader{
				Title:   cmd.Use,
				Section: "8",
			}
			if err := doc.GenMan(cmd, header, w); err != nil {
				log.Fatal(err)
			}
		default:
			log.Fatalf("Unknown format: %q", f)
		}
	},
}

func init() {
	RootCmd.AddCommand(docCmd)

	docCmd.PersistentFlags().StringVarP(
		&docFormat,
		"format", "f",
		"md",
		"Documentation format",
	)

	docCmd.PersistentFlags().StringVarP(
		&docOutput,
		"output", "o",
		"-",
		"Output",
	)
}
