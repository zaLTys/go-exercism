package learn

import (
	"errors"
	"sync"
)

var errWorkers = errors.New("workers must be > 0")

// EXERCISE 1: Implement concurrencySquareAll.
// Process nums using at most `workers` goroutines concurrently.
// Preserve input order in output.
// Return error if workers <= 0.
func concurrencySquareAll(nums []int, workers int) ([]int, error) {
	// TODO: implement worker pool with goroutines + channels
	// Hint: send (index, value) jobs; gather (index, result) pairs and
	//       write into an output slice to preserve order.
	return nil, nil
}

// EXERCISE 2: Implement concurrencySelectFirstString.
// Return the first available value from channel a or b (use select).
func concurrencySelectFirstString(a, b <-chan string) string {
	// TODO: implement using select
	return ""
}

// EXERCISE 3: Implement concurrencyCounterOwnedByGoroutine.
// Return three funcs: inc(), value(), stop().
//   - inc()   increments and returns the new counter value
//   - value() returns the current counter value without incrementing
//   - stop()  terminates the background goroutine (must be safe to call multiple times)
func concurrencyCounterOwnedByGoroutine() (inc func() int, value func() int, stop func()) {
	// TODO: implement using one goroutine that owns the state.
	// Use channels for commands (increment, get value, stop).
	// Use sync.Once so stop() is idempotent.
	_ = sync.Once{}
	return func() int { return 0 }, func() int { return 0 }, func() {}
}
