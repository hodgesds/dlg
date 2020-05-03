package udp

import (
	"encoding/base64"
	"io/ioutil"
	"net"
)

// Config is used for configuring a UDP load test.
type Config struct {
	Endpoint      string  `yaml:"endpoint"`
	Payload       []byte  `yaml:"payload,omitempty"`
	PayloadFile   *string `yaml:"payloadFile,omitempty"`
	PayloadBase64 *string `yaml:"payloadBase64,omitempty"`
}

// Conn is used to return a UDP conn.
func (c *Config) Conn() (net.Conn, error) {
	return net.Dial("udp", c.Endpoint)
}

// GetPayload is used to either return a payload from the configured file or
// the configured field.
func (c *Config) GetPayload() ([]byte, error) {
	if c.PayloadFile != nil {
		b, err := ioutil.ReadFile(*c.PayloadFile)
		if err != nil {
			return nil, err
		}
		c.Payload = b
	}
	if c.PayloadBase64 != nil {
		b, err := base64.StdEncoding.DecodeString(*c.PayloadBase64)
		if err != nil {
			return nil, err
		}
		c.Payload = b
	}
	return c.Payload, nil
}
