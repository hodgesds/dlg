package mqtt

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"

	mqttconfig "github.com/hodgesds/dlg/config/mqtt"
	"github.com/hodgesds/dlg/executor"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type mqttExecutor struct{}

// New returns a new MQTT executor.
func New() executor.MQTT {
	return &mqttExecutor{}
}

// Execute implements the MQTT executor interface.
func (e *mqttExecutor) Execute(ctx context.Context, config *mqttconfig.Config) error {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(config.Broker)

	if config.ClientID != "" {
		opts.SetClientID(config.ClientID)
	} else {
		opts.SetClientID(fmt.Sprintf("dlg-%d", time.Now().UnixNano()))
	}

	if config.Username != "" {
		opts.SetUsername(config.Username)
	}

	if config.Password != "" {
		opts.SetPassword(config.Password)
	}

	opts.SetCleanSession(config.CleanSession)

	if config.ConnectTimeout != nil {
		opts.SetConnectTimeout(*config.ConnectTimeout)
	} else {
		opts.SetConnectTimeout(30 * time.Second)
	}

	if config.KeepAlive != nil {
		opts.SetKeepAlive(time.Duration(*config.KeepAlive) * time.Second)
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	defer client.Disconnect(250)

	// Prepare payload - priority: Payload > PayloadHex > PayloadBase64 > PayloadString > default
	payload := config.Payload
	if len(payload) == 0 && config.PayloadHex != "" {
		var err error
		payload, err = hex.DecodeString(config.PayloadHex)
		if err != nil {
			return fmt.Errorf("failed to decode hex payload: %w", err)
		}
	}
	if len(payload) == 0 && config.PayloadBase64 != "" {
		var err error
		payload, err = base64.StdEncoding.DecodeString(config.PayloadBase64)
		if err != nil {
			return fmt.Errorf("failed to decode base64 payload: %w", err)
		}
	}
	if len(payload) == 0 && config.PayloadString != "" {
		payload = []byte(config.PayloadString)
	}
	if len(payload) == 0 {
		payload = []byte("test payload")
	}

	// Execute the configured number of publishes
	for i := 0; i < config.Count; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			token := client.Publish(
				config.Topic,
				byte(config.QoS),
				config.Retained,
				payload,
			)
			if token.Wait() && token.Error() != nil {
				return token.Error()
			}
		}
	}

	return nil
}
