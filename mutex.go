package sync_patterns

import "sync"

// mutexNoSync demonstrates that the return value can be supplied before all go routines are finished
func mutexNoSync() int {
	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}
	return counter
}

// mutexWithWaitGroup demonstrates that even if we wait for all go routines, a race condition can occur
func mutexWithWaitGroup() int {
	n := 1000
	counter := 0
	var wg sync.WaitGroup

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			counter++
		}()
	}
	wg.Wait()
	return counter
}

// mutexSolution demonstrates a solution using mutex (atomic operation), which ensure that only one goroutine invokes the operation at time
func mutexSolution() int {
	n := 1000
	counter := 0
	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			defer mutex.Unlock()
			mutex.Lock()
			counter++
		}()
	}
	wg.Wait()
	return counter
}
