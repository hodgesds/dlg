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
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/hodgesds/dlg/config"
	kafkaconfig "github.com/hodgesds/dlg/config/kafka"
	kafkaexec "github.com/hodgesds/dlg/executor/kafka"
	stageexec "github.com/hodgesds/dlg/executor/stage"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
)

var (
	kafkaTopic     string
	kafkaMessage   string
	kafkaOperation string
)

// kafkaCmd represents the kafka command
var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Kafka load generator",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("kafka requires broker addresses")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		plan := defaultPlan("kafka")

		brokers := strings.Split(args[0], ",")
		child := &config.Stage{
			Name: plan.Stages[0].Name,
			Tags: tags,
			Kafka: &kafkaconfig.Config{
				Brokers:   brokers,
				Topic:     kafkaTopic,
				Message:   kafkaMessage,
				Operation: kafkaOperation,
			},
		}

		plan.Stages[0].Children = []*config.Stage{child}

		reg := prometheus.NewPedanticRegistry()

		stageExec, err := stageexec.New(
			stageexec.Params{
				Registry: reg,
				Kafka:    kafkaexec.New(),
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		execPlan(plan, reg, stageExec)
	},
}

func init() {
	RootCmd.AddCommand(kafkaCmd)

	kafkaCmd.PersistentFlags().AddFlagSet(planFlags())
	kafkaCmd.AddCommand(newDocCmd())
	kafkaCmd.PersistentFlags().
		StringVarP(&kafkaTopic, "topic", "t", "test", "Kafka topic")
	kafkaCmd.PersistentFlags().
		StringVarP(&kafkaMessage, "message", "m", "", "Message to produce")
	kafkaCmd.PersistentFlags().
		StringVarP(&kafkaOperation, "operation", "o", "produce", "Operation (produce or consume)")
}
