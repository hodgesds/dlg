package executor

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMetrics(t *testing.T) {
	reg := prometheus.NewPedanticRegistry()
	m, err := newMetrics(reg)
	require.NoError(t, err)
	require.NotNil(t, m)
}

// TestMetricsStagesTotal tests the StagesTotal counter.
func TestMetricsStagesTotal(t *testing.T) {
	reg := prometheus.NewPedanticRegistry()
	m, err := newMetrics(reg)
	require.NoError(t, err)

	// Increment counter for a stage
	m.StagesTotal.WithLabelValues("test-stage").Inc()

	// Verify the metric was registered
	metrics, err := reg.Gather()
	require.NoError(t, err)
	require.NotEmpty(t, metrics)

	// Find the stages_total metric
	var found bool
	for _, metric := range metrics {
		if metric.GetName() == "executor_plan_stages_total" {
			found = true
			assert.Equal(t, 1, len(metric.GetMetric()))
			assert.Equal(t, float64(1), metric.GetMetric()[0].GetCounter().GetValue())
		}
	}
	assert.True(t, found, "stages_total metric not found")
}

// TestMetricsStageDuration tests the StageDuration histogram.
func TestMetricsStageDuration(t *testing.T) {
	reg := prometheus.NewPedanticRegistry()
	m, err := newMetrics(reg)
	require.NoError(t, err)

	// Record a duration observation
	m.StageDuration.WithLabelValues("test-stage").Observe(1.5)

	// Verify the metric was registered
	metrics, err := reg.Gather()
	require.NoError(t, err)
	require.NotEmpty(t, metrics)

	// Find the stage_duration metric
	var found bool
	for _, metric := range metrics {
		if metric.GetName() == "executor_plan_stage_duration" {
			found = true
			assert.Equal(t, 1, len(metric.GetMetric()))
			assert.Equal(t, uint64(1), metric.GetMetric()[0].GetHistogram().GetSampleCount())
		}
	}
	assert.True(t, found, "stage_duration metric not found")
}

// TestMetricsMultipleStages tests metrics with multiple stages.
func TestMetricsMultipleStages(t *testing.T) {
	reg := prometheus.NewPedanticRegistry()
	m, err := newMetrics(reg)
	require.NoError(t, err)

	// Increment counters for multiple stages
	stages := []string{"http", "redis", "mongodb"}
	for _, stage := range stages {
		m.StagesTotal.WithLabelValues(stage).Inc()
		m.StageDuration.WithLabelValues(stage).Observe(2.0)
	}

	// Verify metrics
	metrics, err := reg.Gather()
	require.NoError(t, err)
	require.NotEmpty(t, metrics)

	// Verify we have metrics for all stages
	for _, metric := range metrics {
		if metric.GetName() == "executor_plan_stages_total" {
			assert.Equal(t, len(stages), len(metric.GetMetric()))
		}
	}
}

// TestMetricsDuplicateRegistration tests that registering metrics twice fails.
func TestMetricsDuplicateRegistration(t *testing.T) {
	reg := prometheus.NewPedanticRegistry()

	_, err := newMetrics(reg)
	require.NoError(t, err)

	// Attempt to register again should fail
	_, err = newMetrics(reg)
	assert.Error(t, err)
}

// TestMetricsNamespace tests that metrics have the correct namespace.
func TestMetricsNamespace(t *testing.T) {
	reg := prometheus.NewPedanticRegistry()
	m, err := newMetrics(reg)
	require.NoError(t, err)

	m.StagesTotal.WithLabelValues("test").Inc()

	metrics, err := reg.Gather()
	require.NoError(t, err)

	for _, metric := range metrics {
		// All metrics should have the executor_plan namespace
		assert.Contains(t, metric.GetName(), "executor_plan")
	}
}
