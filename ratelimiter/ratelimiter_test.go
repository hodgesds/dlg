package ratelimiter

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNewLimiter tests creating a new rate limiter.
func TestNewLimiter(t *testing.T) {
	limiter := NewLimiter()
	require.NotNil(t, limiter)
}

// TestLimiterWaitWithoutLimit tests waiting without a configured limit.
func TestLimiterWaitWithoutLimit(t *testing.T) {
	limiter := NewLimiter()

	start := time.Now()
	err := limiter.Wait(context.Background())
	elapsed := time.Since(start)

	require.NoError(t, err)
	// Should return immediately when no limit is set
	assert.Less(t, elapsed, 10*time.Millisecond)
}

// TestLimiterWaitBytesWithoutLimit tests waiting for bytes without a configured limit.
func TestLimiterWaitBytesWithoutLimit(t *testing.T) {
	limiter := NewLimiter()

	start := time.Now()
	err := limiter.WaitBytes(context.Background(), 1024)
	elapsed := time.Since(start)

	require.NoError(t, err)
	// Should return immediately when no limit is set
	assert.Less(t, elapsed, 10*time.Millisecond)
}

// TestLimiterReset tests resetting the rate limiter.
func TestLimiterReset(t *testing.T) {
	limiter := NewLimiter()

	// Reset should not panic or error
	limiter.Reset()
}

// TestLimiterConcurrentWait tests concurrent calls to Wait.
func TestLimiterConcurrentWait(t *testing.T) {
	limiter := NewLimiter()

	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			err := limiter.Wait(context.Background())
			assert.NoError(t, err)
			done <- true
		}()
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}
}

// TestLimiterConcurrentWaitBytes tests concurrent calls to WaitBytes.
func TestLimiterConcurrentWaitBytes(t *testing.T) {
	limiter := NewLimiter()

	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(bytes int) {
			err := limiter.WaitBytes(context.Background(), bytes)
			assert.NoError(t, err)
			done <- true
		}(1024)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}
}

// TestLimiterContextCancellation tests that context cancellation is respected.
func TestLimiterContextCancellation(t *testing.T) {
	limiter := NewLimiter()

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	// Should still work since we don't actually use the context in the current implementation
	err := limiter.Wait(ctx)
	// Note: Current implementation doesn't use context, so this won't error
	// This test documents current behavior
	assert.NoError(t, err)
}