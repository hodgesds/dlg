package ratelimiter

import (
	"context"
	"errors"
	"time"

	"github.com/jonboulle/clockwork"
	"golang.org/x/time/rate"
)

// RateLimiter is used for ratelimiting in eithers ops/sec or bytes/sec.
type RateLimiter interface {
	Wait(context.Context) error
	WaitBytes(context.Context, int) error
	Reset()
}

// NewLimiter returns a new RateLimiter.
func NewLimiter() RateLimiter {
	return &limiter{}
}

type limiter struct {
	c clockwork.Clock
	l *rate.Limiter
}

func (l *limiter) Wait(ctx context.Context) error {
	if l.l == nil {
		return nil
	}
	r := l.l.ReserveN(time.Now(), 1)
	if !r.OK() {
		return errors.New("invalid reservation")
	}
	limits.Inc()
	time.Sleep(r.Delay())
	return nil
}

func (l *limiter) WaitBytes(ctx context.Context, b int) error {
	if l.l == nil {
		return nil
	}
	r := l.l.ReserveN(time.Now(), b)
	if !r.OK() {
		return errors.New("invalid reservation")
	}
	limits.Inc()
	time.Sleep(r.Delay())
	return nil
}

func (l *limiter) Reset() {
	//limits.Reset()
}
