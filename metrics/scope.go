package metrics

import (
	"net/http"

	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Scope is a metrics scope.
type Scope interface {
	Register(prom.Collector) error
}

type scope struct {
	registry *prom.Registry
}

// NewScope returns a new scope.
func NewScope() Scope {
	return &scope{
		registry: prom.NewRegistry(),
	}
}

// Register implements the Scope interface.
func (s *scope) Register(c prom.Collector) error {
	return s.registry.Register(c)
}

// Handler returns a http handler.
func (s *scope) Handler() http.Handler {
	return promhttp.InstrumentMetricHandler(s.registry, nil)
}
