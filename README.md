# Synchronization Patterns in Go

Based on an [article by Noam Yadgar](https://code-pilot.me/synchronization-patterns-in-go) and extended to include the comprehensive tests.

## Topics Covered
- [Mutex](mutex.go) - short for "mutual exclusion," is a synchronization primitive used in computer programming to ensure that only one thread can access a shared resource or a critical section at a time. Mutexes help prevent race conditions and other concurrency-related issues in multithreaded applications (def by ChatGPT).
- [Semaphore](semaphore.go)