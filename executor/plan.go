package executor

import (
	"context"

	"github.com/hodgesds/dlg/config"
)

type planExecutor struct {
	stage Stage
}

// NewPlan returns a new Plan executor.
func NewPlan(s Stage) Plan {
	return &planExecutor{stage: s}
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
		if err := e.stage.Execute(ctx, stage); err != nil {
			return err
		}
	}
	return nil
}
