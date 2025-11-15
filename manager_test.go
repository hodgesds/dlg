package dlg

import (
	"context"
	"testing"

	"github.com/hodgesds/dlg/config"
	"github.com/hodgesds/dlg/executor"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestManagerGet(t *testing.T) {
	p, err := executor.NewPlan(executor.Params{}, nil)
	require.NoError(t, err)
	m := NewManager(p)
	_, err = m.Get(context.Background(), "foo")
	require.Error(t, err)

	err = m.Add(context.Background(), &config.Plan{Name: "test"})
	require.NoError(t, err)
	plan, err := m.Get(context.Background(), "test")
	require.NoError(t, err)
	require.NotNil(t, plan)
}

func TestManagerDelete(t *testing.T) {
	p, err := executor.NewPlan(executor.Params{}, nil)
	require.NoError(t, err)
	m := NewManager(p)
	err = m.Delete(context.Background(), "foo")
	require.NoError(t, err)
}

func TestManagerAdd(t *testing.T) {
	p, err := executor.NewPlan(executor.Params{}, nil)
	require.NoError(t, err)
	m := NewManager(p)
	err = m.Add(context.Background(), &config.Plan{Name: "test"})
	require.NoError(t, err)

	plans, err := m.Plans(context.Background())
	require.NoError(t, err)
	require.NotNil(t, plans)
}

// TestManagerGetNonExistent tests getting a non-existent plan.
func TestManagerGetNonExistent(t *testing.T) {
	p, err := executor.NewPlan(executor.Params{}, nil)
	require.NoError(t, err)
	m := NewManager(p)

	plan, err := m.Get(context.Background(), "nonexistent")
	assert.Error(t, err)
	assert.Nil(t, plan)
	assert.Contains(t, err.Error(), "no such plan")
}

// TestManagerAddMultiplePlans tests adding multiple plans.
func TestManagerAddMultiplePlans(t *testing.T) {
	p, err := executor.NewPlan(executor.Params{}, nil)
	require.NoError(t, err)
	m := NewManager(p)

	// Add multiple plans
	for i := 0; i < 5; i++ {
		planName := "test-plan-" + string(rune('0'+i))
		err := m.Add(context.Background(), &config.Plan{Name: planName})
		require.NoError(t, err)
	}

	// Verify all plans exist
	plans, err := m.Plans(context.Background())
	require.NoError(t, err)
	assert.Len(t, plans, 5)
}

// TestManagerUpdatePlan tests updating an existing plan.
func TestManagerUpdatePlan(t *testing.T) {
	p, err := executor.NewPlan(executor.Params{}, nil)
	require.NoError(t, err)
	m := NewManager(p)

	// Add initial plan
	initialPlan := &config.Plan{Name: "test", Description: "initial"}
	err = m.Add(context.Background(), initialPlan)
	require.NoError(t, err)

	// Update the plan
	updatedPlan := &config.Plan{Name: "test", Description: "updated"}
	err = m.Add(context.Background(), updatedPlan)
	require.NoError(t, err)

	// Verify the plan was updated
	plan, err := m.Get(context.Background(), "test")
	require.NoError(t, err)
	assert.Equal(t, "updated", plan.Description)
}

// TestManagerDeleteExisting tests deleting an existing plan.
func TestManagerDeleteExisting(t *testing.T) {
	p, err := executor.NewPlan(executor.Params{}, nil)
	require.NoError(t, err)
	m := NewManager(p)

	// Add a plan
	err = m.Add(context.Background(), &config.Plan{Name: "test"})
	require.NoError(t, err)

	// Verify it exists
	plan, err := m.Get(context.Background(), "test")
	require.NoError(t, err)
	require.NotNil(t, plan)

	// Delete the plan
	err = m.Delete(context.Background(), "test")
	require.NoError(t, err)

	// Verify it no longer exists
	plan, err = m.Get(context.Background(), "test")
	assert.Error(t, err)
	assert.Nil(t, plan)
}

// TestManagerPlansEmpty tests getting plans from an empty manager.
func TestManagerPlansEmpty(t *testing.T) {
	p, err := executor.NewPlan(executor.Params{}, nil)
	require.NoError(t, err)
	m := NewManager(p)

	plans, err := m.Plans(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, plans)
	assert.Len(t, plans, 0)
}

// TestManagerConcurrency tests concurrent access to the manager.
func TestManagerConcurrency(t *testing.T) {
	p, err := executor.NewPlan(executor.Params{}, nil)
	require.NoError(t, err)
	m := NewManager(p)

	// Run concurrent operations
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(idx int) {
			planName := "plan-" + string(rune('0'+idx))
			err := m.Add(context.Background(), &config.Plan{Name: planName})
			assert.NoError(t, err)
			_, err = m.Get(context.Background(), planName)
			assert.NoError(t, err)
			err = m.Delete(context.Background(), planName)
			assert.NoError(t, err)
			done <- true
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}
}

// TestManagerExecute tests the Execute method.
func TestManagerExecute(t *testing.T) {
	p, err := executor.NewPlan(executor.Params{}, nil)
	require.NoError(t, err)
	m := NewManager(p)

	// Create a simple plan with no stages
	plan := &config.Plan{Name: "test"}

	// Execute should delegate to the plan executor
	err = m.Execute(context.Background(), plan)
	require.NoError(t, err)
}
