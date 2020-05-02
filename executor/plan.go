package executor

import (
	"context"

	"github.com/hodgesds/dlg/config"
	"github.com/prometheus/client_golang/prometheus"
)

// Params are used for configuring a Plan.
type Params struct {
	Registry *prometheus.Registry
}

type planExecutor struct {
	stage   Stage
	metrics *metrics
}

// NewPlan returns a new Plan executor.
func NewPlan(p Params, s Stage) (Plan, error) {
	m, err := newMetrics(p.Registry)
	if err != nil {
		return nil, err
	}
	return &planExecutor{
		stage:   s,
		metrics: m,
	}, nil
}

// Executor implements the Plan interface.
func (e *planExecutor) Execute(ctx context.Context, p *config.Plan) error {
	if err := p.Validate(); err != nil {
		return err
	}
	if err := p.WaitStart(ctx); err != nil {
		return err
	}
	var cancel func()
	if p.Duration != nil {
		ctx, cancel = context.WithTimeout(ctx, *p.Duration)
	} else {
		ctx, cancel = context.WithCancel(ctx)
	}
	defer cancel()

	for _, stage := range p.Stages {
		e.metrics.StagesTotal.With(prometheus.Labels{"stage": stage.Name}).Add(1)
		if err := e.stage.Execute(ctx, stage); err != nil {
			return err
		}
	}
	return nil
}
