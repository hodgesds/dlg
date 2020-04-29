package stage

import (
	"context"
	"errors"
	"sync"

	"github.com/hodgesds/dlg/config"
	"github.com/hodgesds/dlg/executor"
	"go.uber.org/multierr"
)

var (
	// ErrNoStageExecutor is returned when a stage has no configured
	// executor.
	ErrNoStageExecutor = errors.New("no executor for stage")
)

type stageExecutor struct {
	dns   executor.DNS
	etcd  executor.ETCD
	http  executor.HTTP
	redis executor.Redis
	sql   executor.SQL
	udp   executor.UDP
}

// Params is used for configuring a Stage executor.
type Params struct {
	DNS   executor.DNS
	ETCD  executor.ETCD
	HTTP  executor.HTTP
	Redis executor.Redis
	SQL   executor.SQL
	UDP   executor.UDP
}

// New returns a new Stage executor.
func New(p Params) executor.Stage {
	return &stageExecutor{
		dns:   p.DNS,
		etcd:  p.ETCD,
		http:  p.HTTP,
		redis: p.Redis,
		sql:   p.SQL,
		udp:   p.UDP,
	}
}

// Execute implements the Stage interface.
func (e *stageExecutor) Execute(ctx context.Context, s *config.Stage) error {
	if err := s.Validate(); err != nil {
		return err
	}

	var (
		// exCtx is the context for this execution, since a stage can
		// be repeated multiple times with a timeout a copy of the
		// original context must be used.
		exCtx  context.Context
		cancel func()
	)
	if s.Timeout != nil {
		exCtx, cancel = context.WithTimeout(ctx, *s.Timeout)
	} else {
		exCtx, cancel = context.WithCancel(ctx)
	}
	defer cancel()

	if s.HTTP != nil {
		if e.http == nil {
			return ErrNoStageExecutor
		}
		if err := e.http.Execute(exCtx, s.HTTP); err != nil {
			return err
		}
	}
	if s.Redis != nil {
		if e.redis == nil {
			return ErrNoStageExecutor
		}
		if err := e.redis.Execute(exCtx, s.Redis); err != nil {
			return err
		}
	}
	if s.ETCD != nil {
		if e.etcd == nil {
			return ErrNoStageExecutor
		}
		if err := e.etcd.Execute(exCtx, s.ETCD); err != nil {
			return err
		}
	}
	if s.SQL != nil {
		if e.sql == nil {
			return ErrNoStageExecutor
		}
		if err := e.sql.Execute(exCtx, s.SQL); err != nil {
			return err
		}
	}
	if s.UDP != nil {
		if e.udp == nil {
			return ErrNoStageExecutor
		}
		if err := e.udp.Execute(exCtx, s.UDP); err != nil {
			return err
		}
	}
	if s.DNS != nil {
		if e.dns == nil {
			return ErrNoStageExecutor
		}
		if err := e.dns.Execute(exCtx, s.DNS); err != nil {
			return err
		}
	}

	// Execute any children.
	if len(s.Children) > 1 && s.Concurrent {
		if err := e.execParallel(exCtx, s.Children); err != nil {
			return err
		}
		if s.Repeat > 0 {
			s.Repeat--
			return e.Execute(ctx, s)
		}
	}

	for _, child := range s.Children {
		if err := e.Execute(exCtx, child); err != nil {
			return err
		}
	}

	if s.Repeat > 0 {
		s.Repeat--
		return e.Execute(ctx, s)
	}

	return nil
}

func (e *stageExecutor) execParallel(ctx context.Context, stages []*config.Stage) error {
	var (
		wg  sync.WaitGroup
		mu  sync.Mutex
		err error
	)
	for _, stage := range stages {
		wg.Add(1)
		go func(stage *config.Stage) {
			err2 := e.Execute(ctx, stage)
			if err2 != nil {
				mu.Lock()
				err = multierr.Append(err, err2)
				mu.Unlock()
			}
			wg.Done()
		}(stage)
	}
	wg.Wait()
	return err
}
