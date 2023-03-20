package sync_patterns

import (
	"context"
	"runtime"
	"time"

	"golang.org/x/sync/semaphore"
)

// semaphoreNoGoroutineLimit demonstrates that there is no threshold for the number of running goroutines
func semaphoreNoGoroutineLimit() int {
	n := 100
	for i := 0; i < n; i++ {
		go func() {
			time.Sleep(time.Second)
		}()
	}
	return runtime.NumGoroutine()
}

// semaphoreSolution demonstrates solution to limit number of parallel running goroutines
func semaphoreSolution() int {
	n := 100
	limit := int64(10)
	sm := semaphore.NewWeighted(limit)
	ctx := context.Background()
	for i := 0; i < n; i++ {
		_ = sm.Acquire(ctx, 1)
		go func() {
			defer sm.Release(1)
			time.Sleep(time.Second)
		}()
	}
	return runtime.NumGoroutine()
}
