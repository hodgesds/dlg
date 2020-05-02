package stage

import (
	"github.com/prometheus/client_golang/prometheus"
)

// metrics contains metrics.
type metrics struct {
	ErrorsTotal *prometheus.CounterVec
}

func newMetrics(reg *prometheus.Registry) (*metrics, error) {
	m := &metrics{
		ErrorsTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "stage",
			Name:      "errors_total",

			Help: "The total number and type of errors that occurred while advertising.",
		}, []string{"interface", "error"}),
	}
	return m, reg.Register(m.ErrorsTotal)
}
