package sync_patterns

import (
	"runtime"
	"testing"
)

// Test_semaphoreNoGoroutineLimit demonstrates that there is no threshold for the number of running goroutines
func Test_semaphoreNoGoroutineLimit(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "should start 100 goroutines",
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nrg := runtime.NumGoroutine()
			if got := semaphoreNoGoroutineLimit() - nrg; got != tt.want {
				t.Errorf("semaphoreNoGoroutineLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test_semaphoreSolution demonstrates solution to limit number of parallel running goroutines
func Test_semaphoreSolution(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "should limit to 10 goroutines running in parallel",
			want: 13, // 3 goroutines are started independently
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := semaphoreSolution(); got != tt.want {
				t.Errorf("semaphoreSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
