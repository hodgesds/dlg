package ratelimiter

import (
	prom "github.com/prometheus/client_golang/prometheus"
)

var (
	limits = prom.NewCounter(prom.CounterOpts{
		Namespace: "ratelimiter",
		Name:      "limits",
	})
)

func init() {
	prom.MustRegister(limits)
}
