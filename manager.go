package dlg

import (
	"context"
	"fmt"
	"sync"

	"github.com/hodgesds/dlg/config"
	"github.com/hodgesds/dlg/executor"
)

// Manager is used for scheduling execution plans.
type Manager interface {
	executor.Plan

	// Get is used to return a plan by name.
	Get(context.Context, string) (*config.Plan, error)

	// Add is used to add a new plan.
	Add(context.Context, *config.Plan) error

	// Delete is used to remove a plan.
	Delete(context.Context, string) error

	// Plans returns a list of all known plans.
	Plans(context.Context) ([]*config.Plan, error)
}

type manager struct {
	planMu   sync.RWMutex
	plans    map[string]*config.Plan
	planExec executor.Plan
}

// NewManager returns a new Manager.
func NewManager(planExec executor.Plan) Manager {
	return &manager{
		plans:    map[string]*config.Plan{},
		planExec: planExec,
	}
}

// Get implements the Manager interface.
func (m *manager) Get(ctx context.Context, name string) (*config.Plan, error) {
	m.planMu.RLock()
	defer m.planMu.RUnlock()
	p, ok := m.plans[name]
	if !ok {
		return nil, fmt.Errorf("no such plan: %q", name)
	}
	return p, nil
}

// Delete implements the Manager interface.
func (m *manager) Delete(ctx context.Context, name string) error {
	m.planMu.Lock()
	defer m.planMu.Unlock()
	delete(m.plans, name)
	return nil
}

// Add implements the Manager interface.
func (m *manager) Add(ctx context.Context, plan *config.Plan) error {
	m.planMu.Lock()
	defer m.planMu.Unlock()
	m.plans[plan.Name] = plan
	return nil
}

// Plans implements the Manager interface.
func (m *manager) Plans(ctx context.Context) ([]*config.Plan, error) {
	m.planMu.RLock()
	defer m.planMu.RUnlock()
	plans := []*config.Plan{}
	for _, plan := range m.plans {
		plans = append(plans, plan)
	}
	return plans, nil
}

// Execute implements the Executor interface.
func (m *manager) Execute(ctx context.Context, plan *config.Plan) error {
	return m.planExec.Execute(ctx, plan)
}
