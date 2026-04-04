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

## Real-World Variants of This Pattern

This worker pool is a simplified version of patterns you'll see constantly in production Go:

- **HTTP batch processor**: fan out N API calls, collect results, return aggregated response
- **Database bulk importer**: bounded goroutine pool reads rows, writes to DB, collects insert errors
- **File processor**: scan a directory tree, process each file concurrently, report all failures at the end
- **Build system**: `go build ./...` itself uses a worker pool internally for parallel compilation

## Common Mistakes & Caveats

- **Not closing the `jobs` channel causes workers to block forever** — the `for i := range jobs` loop only exits when `jobs` is closed; the job-feeding goroutine must always close it, even on cancellation (`defer close(jobs)`)
- **Not closing `results` causes the collector to block forever** — `results` should be closed by the goroutine that waits for all workers (`wg.Wait()` then `close(results)`)
- **Race on `stoppedEarly`** — if the job-feeder goroutine sets `stoppedEarly = true` while the collector is reading it, you have a data race; use an `atomic` or a channel signal instead:
  ```go
  // Safe alternative: read ctx.Err() after collecting all results
  if ctx.Err() != nil {
      return errors.Join(ctx.Err(), combined)
  }
  ```
- **Buffering channels affects throughput** — unbuffered `jobs` means each worker handoff is synchronous; a buffered channel (size = len(tasks)) allows the feeder to run ahead without blocking, but uses more memory
- **`errors.Join` ignores nil entries** — you don't need to filter them out; `errors.Join(nil, err, nil)` returns just `err`

## Useful Links
- [Go Concurrency Patterns: Pipelines](https://go.dev/blog/pipelines)
- [Go Concurrency Patterns: Context](https://go.dev/blog/context)
- [errors package](https://pkg.go.dev/errors)
- [sync package](https://pkg.go.dev/sync)
- [Organizing a Go module](https://go.dev/doc/modules/layout)
