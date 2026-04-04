package runner

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

// Task is a function that does work and reports an error.
type Task func(context.Context) error

// Runner executes a list of Tasks with a bounded number of concurrent workers.
type Runner struct {
	Workers int
}

// ErrInvalidWorkers is returned when Workers <= 0.
var ErrInvalidWorkers = errors.New("workers must be > 0")

type result struct {
	i   int
	err error
}

// Run executes tasks with a worker limit.
// If ctx is canceled, it stops scheduling new tasks and returns ctx.Err()
// joined with any task errors already produced.
//
// TODO: implement this function.
// Hints:
//  1. Return ErrInvalidWorkers if r.Workers <= 0.
//  2. Return nil immediately for empty task list.
//  3. Create jobs (chan int) and results (chan result) channels.
//  4. Start r.Workers goroutines that read job indices from jobs and send results.
//  5. Start a goroutine that waits for all workers (WaitGroup) then closes results.
//  6. Feed job indices to jobs in a goroutine, selecting on ctx.Done() to cancel early.
//  7. Collect results into []error indexed by task position.
//  8. Return errors.Join(errs...) and ctx.Err() if stopped early.
func (r Runner) Run(ctx context.Context, tasks []Task) error {
	// TODO: implement
	_, _ = sync.WaitGroup{}, result{}
	return nil
}

// String returns a human-readable description of the Runner.
func (r Runner) String() string {
	return fmt.Sprintf("Runner(workers=%d)", r.Workers)
}
