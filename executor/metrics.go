package executor

import "github.com/prometheus/client_golang/prometheus"

// metrics contains metrics.
type metrics struct {
	StagesTotal *prometheus.CounterVec
}

func newMetrics(reg *prometheus.Registry) (*metrics, error) {
	m := &metrics{
		StagesTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "plan",
			Name:      "stages_total",
			Help:      "The total number of stages.",
		}, []string{"stage"}),
	}
	return m, reg.Register(m.StagesTotal)
}
