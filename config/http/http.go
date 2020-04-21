package http

import (
	"bytes"
	"context"
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
	BodyBase64 string       `yaml:"bodyBase64,omitempty"`
}

// Request returns a new http.Request from a payload config.
func (p *Payload) Request(ctx context.Context) (*ghttp.Request, error) {
	if p.BodyFile != nil {
		b, err := ioutil.ReadFile(*p.BodyFile)
		if err != nil {
			return nil, err
		}
		p.Body = b
	}
	r, err := ghttp.NewRequest(p.Method, p.URL, bytes.NewReader(p.Body))
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
