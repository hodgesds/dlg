package executor

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
)

func TestNewMetrics(t *testing.T) {
	reg := prometheus.NewPedanticRegistry()
	m, err := newMetrics(reg)
	require.NoError(t, err)
	require.NotNil(t, m)
}
