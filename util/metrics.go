package util

import (
	"io"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"
)

// RegistryGather is used to gather metrics for a registry.
func RegistryGather(reg *prometheus.Registry, w io.Writer) error {
	metrics, err := reg.Gather()
	if err != nil {
		return err
	}
	for _, m := range metrics {
		_, err := expfmt.MetricFamilyToText(w, m)
		if err != nil {
			return err
		}
	}
	return nil
}
