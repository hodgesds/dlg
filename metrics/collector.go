package metrics

import (
	"github.com/hodgesds/dlg/config"
	prom "github.com/prometheus/client_golang/prometheus"
)

// NewPlanCollector returns a prometheus collector for a execution plan.
func NewPlanCollector(p *config.Plan) prom.Collector {
	return nil
}

// NewStageCollector returns a prometheus collector for a execution stage.
func NewStageCollector(stage *config.Stage) prom.Collector {
	return nil
}
