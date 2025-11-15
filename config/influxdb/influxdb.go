package influxdb

import (
	"time"
)

// Operation represents an InfluxDB operation type
type Operation string

const (
	OpWrite Operation = "write"
	OpQuery Operation = "query"
)

// Point represents an InfluxDB data point
type Point struct {
	Measurement string                 `yaml:"measurement"`
	Tags        map[string]string      `yaml:"tags,omitempty"`
	Fields      map[string]interface{} `yaml:"fields"`
	Timestamp   *time.Time             `yaml:"timestamp,omitempty"`
}

// Config is used for configuring an InfluxDB load test.
type Config struct {
	URL          string         `yaml:"url"`
	Token        string         `yaml:"token,omitempty"`
	Username     string         `yaml:"username,omitempty"`
	Password     string         `yaml:"password,omitempty"`
	Organization string         `yaml:"organization"`
	Bucket       string         `yaml:"bucket"`
	Database     string         `yaml:"database,omitempty"` // For InfluxDB 1.x
	Operation    Operation      `yaml:"operation"`
	Count        int            `yaml:"count"`
	Points       []Point        `yaml:"points,omitempty"`
	Query        string         `yaml:"query,omitempty"`
	Precision    string         `yaml:"precision,omitempty"` // ns, us, ms, s
	Timeout      *time.Duration `yaml:"timeout,omitempty"`
	BatchSize    *int           `yaml:"batchSize,omitempty"`
}
