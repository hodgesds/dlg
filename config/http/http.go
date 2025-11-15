package http

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	ghttp "net/http"
)

// Payload is a payload for a HTTP load test.
type Payload struct {
	URL        string       `yaml:"url"`
	Header     ghttp.Header `yaml:"header"`
	Method     string       `yaml:"method"`
	Body       []byte       `yaml:"body,omitempty"`
	BodyFile   *string      `yaml:"bodyFile,omitempty"`
	BodyHex    string       `yaml:"bodyHex,omitempty"`     // Hex-encoded body
	BodyBase64 string       `yaml:"bodyBase64,omitempty"`  // Base64-encoded body
}

// Request returns a new http.Request from a payload config.
func (p *Payload) Request(ctx context.Context) (*ghttp.Request, error) {
	// Priority: Body > BodyFile > BodyHex > BodyBase64
	body := p.Body

	if len(body) == 0 && p.BodyFile != nil {
		b, err := ioutil.ReadFile(*p.BodyFile)
		if err != nil {
			return nil, err
		}
		body = b
	}

	if len(body) == 0 && p.BodyHex != "" {
		b, err := hex.DecodeString(p.BodyHex)
		if err != nil {
			return nil, err
		}
		body = b
	}

	if len(body) == 0 && p.BodyBase64 != "" {
		b, err := base64.StdEncoding.DecodeString(p.BodyBase64)
		if err != nil {
			return nil, err
		}
		body = b
	}

	r, err := ghttp.NewRequest(p.Method, p.URL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	r.Header = p.Header
	return r.WithContext(ctx), nil
}

// Config is used for configuring a HTTP load test.
type Config struct {
	MaxConns     *int    `yaml:"maxConns,omitempty"`
	MaxIdleConns *int    `yaml:"maxIdleConns,omitempty"`
	Count        int     `yaml:"count"`
	Payload      Payload `yaml:"payload"`
}
