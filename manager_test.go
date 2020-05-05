package dlg

import (
	"context"
	"testing"

	"github.com/hodgesds/dlg/config"
	"github.com/hodgesds/dlg/executor"
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
