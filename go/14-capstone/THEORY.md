# Capstone Project — Theory Reference

## The Big Idea

This capstone integrates everything: **packages/modules, structs/methods, interfaces, explicit error handling, goroutines + channels, context cancellation, and table-driven tests**. The goal is a small but production-realistic worker-pool library.

## What You're Building

A `runner` package that:
1. Accepts a list of `Task` functions
2. Executes them concurrently with a bounded number of workers
3. Respects `context.Context` cancellation
4. Aggregates all errors with `errors.Join`
5. Is tested with table-driven tests

## Architecture

```
14-capstone/
  go.mod             (module example.com/go-capstone)
  THEORY.md
  runner/
    runner.go        ← implement this
    runner_test.go   ← tests (read-only)
```

## Key Concepts Used

### Worker Pool Pattern
```
        ┌──────────┐
jobs ──►│ worker 1 │──► results
        │ worker 2 │
        │ worker N │
        └──────────┘
```

Goroutines read from a `jobs` channel; results go to a `results` channel.
A separate goroutine feeds jobs (and respects context cancellation).
A `sync.WaitGroup` tracks when all workers finish so `results` can be closed.

### Context Cancellation
```go
select {
case <-ctx.Done():   // cancelled — stop scheduling
    return
case jobs <- i:      // send next job
}
```

### Error Aggregation
```go
errs := make([]error, len(tasks))
for res := range results {
    errs[res.i] = res.err
}
return errors.Join(errs...)  // nil entries are ignored by errors.Join
```

## The `runner.go` Interface

```go
package runner

import (
    "context"
    "errors"
    "fmt"
    "sync"
)

type Task func(context.Context) error

type Runner struct {
    Workers int
}

var ErrInvalidWorkers = errors.New("workers must be > 0")

// Run executes tasks with the bounded worker pool.
// Returns nil if all tasks succeed; returns joined errors otherwise.
// If ctx is cancelled, stops scheduling new tasks and returns ctx.Err()
// joined with any task errors already produced.
func (r Runner) Run(ctx context.Context, tasks []Task) error {
    // TODO: implement
    return nil
}

func (r Runner) String() string {
    return fmt.Sprintf("Runner(workers=%d)", r.Workers)
}
```

## Implementation Checklist

- [ ] Validate `r.Workers > 0` — return `ErrInvalidWorkers` if not
- [ ] Handle empty `tasks` slice — return `nil` immediately
- [ ] Create `jobs` and `results` channels
- [ ] Start `r.Workers` goroutines, each reading from `jobs`
- [ ] Start a goroutine to close `results` after all workers finish (`sync.WaitGroup`)
- [ ] Feed jobs from a goroutine, selecting on `ctx.Done()` to stop early
- [ ] Collect results into a slice (preserve index order)
- [ ] Return `errors.Join(errs...)` — or join with `ctx.Err()` if cancelled

## Why This Matches Real Go Practice

- **Worker pools** are the canonical Go concurrency pattern (Official Tour + pipelines blog)
- **`context.Context`** is the standard for cancellation across goroutines
- **`errors.Join`** is the standard for multi-error aggregation (Go 1.20+)
- **Table-driven tests** with named subtests are the idiomatic test style

## Useful Links
- [Go Concurrency Patterns: Pipelines](https://go.dev/blog/pipelines)
- [Go Concurrency Patterns: Context](https://go.dev/blog/context)
- [errors package](https://pkg.go.dev/errors)
- [sync package](https://pkg.go.dev/sync)
- [Organizing a Go module](https://go.dev/doc/modules/layout)
