package influxdb

import (
	"context"
	"time"

	influxdbconfig "github.com/hodgesds/dlg/config/influxdb"
	"github.com/hodgesds/dlg/executor"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type influxdbExecutor struct{}

// New returns a new InfluxDB executor.
func New() executor.InfluxDB {
	return &influxdbExecutor{}
}

// Execute implements the InfluxDB executor interface.
func (e *influxdbExecutor) Execute(ctx context.Context, config *influxdbconfig.Config) error {
	var client influxdb2.Client

	if config.Token != "" {
		client = influxdb2.NewClient(config.URL, config.Token)
	} else {
		client = influxdb2.NewClientWithOptions(
			config.URL,
			"",
			influxdb2.DefaultOptions().
				SetUseGZip(true),
		)
	}
	defer client.Close()

	// Execute the configured number of operations
	for i := 0; i < config.Count; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := e.executeOperation(ctx, client, config); err != nil {
				return err
			}
		}
	}

	return nil
}

func (e *influxdbExecutor) executeOperation(ctx context.Context, client influxdb2.Client, config *influxdbconfig.Config) error {
	switch config.Operation {
	case influxdbconfig.OpWrite:
		writeAPI := client.WriteAPIBlocking(config.Organization, config.Bucket)

		for _, pointConfig := range config.Points {
			p := influxdb2.NewPoint(
				pointConfig.Measurement,
				pointConfig.Tags,
				pointConfig.Fields,
				time.Now(),
			)

			if pointConfig.Timestamp != nil {
				p = influxdb2.NewPoint(
					pointConfig.Measurement,
					pointConfig.Tags,
					pointConfig.Fields,
					*pointConfig.Timestamp,
				)
			}

			if err := writeAPI.WritePoint(ctx, p); err != nil {
				return err
			}
		}

	case influxdbconfig.OpQuery:
		if config.Query != "" {
			queryAPI := client.QueryAPI(config.Organization)

			result, err := queryAPI.Query(ctx, config.Query)
			if err != nil {
				return err
			}

			// Consume results
			for result.Next() {
				// Process record
				_ = result.Record()
			}

			if result.Err() != nil {
				return result.Err()
			}
		}
	}

	return nil
}
