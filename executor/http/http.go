package http

import (
	"context"
	"net/http"
	"sync"

	httpconf "github.com/hodgesds/dlg/config/http"
	"github.com/hodgesds/dlg/executor"
	"github.com/prometheus/client_golang/prometheus"
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/multierr"
)

// httpExecutor is a HTTP executor.
type httpExecutor struct {
	mu      sync.Mutex
	client  *http.Client
	trans   *http.Transport
	reqPool sync.Pool
	reg     *prometheus.Registry
}

// New returns a new HTTP executor.
func New(reg *prometheus.Registry) executor.HTTP {
	t := &http.Transport{}
	c := &http.Client{Transport: t}
	c = makeInstrumentedClient(reg, c)
	return &httpExecutor{
		client: c,
		trans:  t,
	}
}

// Execute implements the HTTP interface.
func (e *httpExecutor) Execute(ctx context.Context, conf *httpconf.Config) error {
	if conf.MaxIdleConns != nil {
		e.mu.Lock()
		e.trans.MaxIdleConns = *conf.MaxIdleConns
		e.mu.Unlock()
	}
	if conf.MaxConns != nil {
		e.mu.Lock()
		e.trans.MaxConnsPerHost = *conf.MaxConns
		e.mu.Unlock()
	}
	var (
		wg  sync.WaitGroup
		mu  sync.Mutex
		err error
	)
	for count := conf.Count; count > 0; count-- {
		wg.Add(1)
		go func(ctx context.Context) {
			ctx2, cancel := context.WithCancel(ctx)
			defer cancel()
			req, err2 := conf.Payload.Request(ctx2)
			if err2 != nil {
				mu.Lock()
				err = multierr.Append(err, err2)
				mu.Unlock()
				return
			}

			_, err2 = e.client.Do(req)
			if err2 != nil {
				mu.Lock()
				err = multierr.Append(err, err2)
				mu.Unlock()
			}
			wg.Done()
		}(ctx)
	}
	wg.Wait()
	return err
}

func makeInstrumentedClient(reg *prom.Registry, client *http.Client) *http.Client {
	inFlightGauge := prom.NewGauge(prom.GaugeOpts{
		Name: "client_in_flight_requests",
		Help: "A gauge of in-flight requests for the wrapped client.",
	})

	counter := prom.NewCounterVec(
		prom.CounterOpts{
			Name: "client_api_requests_total",
			Help: "A counter for requests from the wrapped client.",
		},
		[]string{"code", "method"},
	)

	dnsLatencyVec := prom.NewHistogramVec(
		prom.HistogramOpts{
			Name:    "dns_duration_ms",
			Help:    "Trace dns latency histogram.",
			Buckets: []float64{.005, .01, .025, .05, 1},
		},
		[]string{"event"},
	)

	tlsLatencyVec := prom.NewHistogramVec(
		prom.HistogramOpts{
			Name:    "tls_duration_ms",
			Help:    "Trace tls latency histogram.",
			Buckets: []float64{.05, .1, .25, .5, 1},
		},
		[]string{"event"},
	)

	histVec := prom.NewHistogramVec(
		prom.HistogramOpts{
			Name:    "request_duration_ms",
			Help:    "A histogram of request latencies.",
			Buckets: prom.DefBuckets,
		},
		[]string{"method"},
	)

	reg.MustRegister(counter, tlsLatencyVec, dnsLatencyVec, histVec, inFlightGauge)

	trace := &promhttp.InstrumentTrace{
		DNSStart: func(t float64) {
			dnsLatencyVec.WithLabelValues("dns_start").Observe(t)
		},
		DNSDone: func(t float64) {
			dnsLatencyVec.WithLabelValues("dns_done").Observe(t)
		},
		TLSHandshakeStart: func(t float64) {
			tlsLatencyVec.WithLabelValues("tls_handshake_start").Observe(t)
		},
		TLSHandshakeDone: func(t float64) {
			tlsLatencyVec.WithLabelValues("tls_handshake_done").Observe(t)
		},
	}

	client.Transport = promhttp.InstrumentRoundTripperInFlight(inFlightGauge,
		promhttp.InstrumentRoundTripperCounter(counter,
			promhttp.InstrumentRoundTripperTrace(trace,
				promhttp.InstrumentRoundTripperDuration(histVec, http.DefaultTransport),
			),
		),
	)
	return client
}
