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

	"github.com/hodgesds/dlg/config"
	"github.com/hodgesds/dlg/executor"
	"github.com/hodgesds/dlg/util"
	"github.com/prometheus/client_golang/prometheus"
)

func defaultPlan(planType string) *config.Plan {
	stage := &config.Stage{
		Name:       fmt.Sprintf("%s-%s", name, planType),
		Tags:       tags,
		Repeat:     repeat,
		Concurrent: concurrent,
		Children:   []*config.Stage{},
	}
	if dur > 0 {
		stage.Duration = &dur
	}

	return &config.Plan{
		Name:   name,
		Tags:   tags,
		Stages: []*config.Stage{stage},
	}
}

func execPlan(
	plan *config.Plan,
	reg *prometheus.Registry,
	stageExec executor.Stage,
) {
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
}
