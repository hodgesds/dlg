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
	"time"

	"github.com/spf13/pflag"
)

var (
	count  int
	debug  bool
	dur    time.Duration
	name   string
	repeat int
	tags   = []string{}

	planFlags = pflag.NewFlagSet("", pflag.ExitOnError)
)

func init() {
	planFlags.BoolVar(&debug, "debug", false, "debug")
	planFlags.StringVar(&name, "name", "cli", "execution name")
	planFlags.IntVarP(&repeat, "repeat", "r", 0, "number of times to repeat")
	planFlags.IntVarP(&count, "count", "c", 100, "number of times to execute")
	planFlags.DurationVarP(&dur, "duration", "d", 0, "execution duration")
	planFlags.StringSliceVar(&tags, "tags", nil, "metrics tags")
}
