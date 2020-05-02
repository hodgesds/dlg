package executor

import "github.com/prometheus/client_golang/prometheus"

// metrics contains metrics.
type metrics struct {
	StagesTotal   *prometheus.CounterVec
	StageDuration *prometheus.HistogramVec
}

func newMetrics(reg *prometheus.Registry) (*metrics, error) {
	m := &metrics{
		StagesTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "executor",
			Subsystem: "plan",
			Name:      "stages_total",
			Help:      "The total number of stages.",
		}, []string{"stage"}),
		StageDuration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "executor",
			Subsystem: "plan",
			Name:      "stage_duration",
			Help:      "The duration of stages.",
		}, []string{"stage"}),
	}

	reg.MustRegister(
		m.StagesTotal,
		m.StageDuration,
	)
	return m, nil
}
