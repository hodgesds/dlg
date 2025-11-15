package mqtt

import (
	"time"
)

// QoS represents MQTT Quality of Service levels
type QoS byte

const (
	QoS0 QoS = 0 // At most once delivery
	QoS1 QoS = 1 // At least once delivery
	QoS2 QoS = 2 // Exactly once delivery
)

// Config is used for configuring an MQTT load test.
type Config struct {
	Broker         string         `yaml:"broker"`
	ClientID       string         `yaml:"clientId,omitempty"`
	Username       string         `yaml:"username,omitempty"`
	Password       string         `yaml:"password,omitempty"`
	Topic          string         `yaml:"topic"`
	QoS            QoS            `yaml:"qos,omitempty"`
	Count          int            `yaml:"count"`
	Payload        []byte         `yaml:"payload,omitempty"`
	PayloadString  string         `yaml:"payloadString,omitempty"`
	Retained       bool           `yaml:"retained,omitempty"`
	CleanSession   bool           `yaml:"cleanSession,omitempty"`
	ConnectTimeout *time.Duration `yaml:"connectTimeout,omitempty"`
	KeepAlive      *int64         `yaml:"keepAlive,omitempty"`
}
