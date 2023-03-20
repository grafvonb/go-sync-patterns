package sync_patterns

import (
	"runtime"
	"testing"
)

// Test_mutexNoSync demonstrates that the return value can be supplied before all go routines are finished
func Test_mutexNoSync(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "1: should calculate 1000, but never get it",
			want: 1000,
		},
		{
			name: "2: should calculate 1000, but never get it",
			want: 1000,
		},
		{
			name: "3: should calculate 1000, but never get it",
			want: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ngr := runtime.NumGoroutine()
			if got := mutexNoSync(); got != tt.want {
				t.Logf("go routines still running: %d", runtime.NumGoroutine()-ngr)
				t.Logf("mutexNoSync() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test_mutexWithWaitGroup demonstrates that even if we wait for all go routines, a race condition can occur
func Test_mutexWithWaitGroup(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "1: should calculate 1000, but never get it",
			want: 1000,
		},
		{
			name: "2: should calculate 1000, but never get it",
			want: 1000,
		},
		{
			name: "3: should calculate 1000, but never get it",
			want: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ngr := runtime.NumGoroutine()
			if got := mutexWithWaitGroup(); got != tt.want {
				t.Logf("go routines still running: %d", runtime.NumGoroutine()-ngr)
				t.Logf("mutexWithWaitGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test_mutexSolution demonstrates a solution using mutex (atomic operation), which ensure that only one goroutine invokes the operation at time
func Test_mutexSolution(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "1: should calculate and get 1000",
			want: 1000,
		},
		{
			name: "2: should calculate and get 1000",
			want: 1000,
		},
		{
			name: "3: should calculate and get 1000",
			want: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mutexSolution(); got != tt.want {
				t.Errorf("mutexSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
