# Concurrency — Theory Reference

## The Big Idea

Go concurrency is built on **goroutines** (cheap concurrent functions) and **channels** (typed communication pipes). `select` multiplexes channel operations. `context.Context` is the standard cancellation/timeout mechanism. The race detector (`go test -race`) is part of the normal workflow.

## Goroutines

```go
// C#: Task.Run(() => doWork())
// Go: go doWork()

go func() {
    fmt.Println("running in goroutine")
}()

// goroutines are very cheap — millions can run simultaneously
```

## Channels

```go
// create
ch := make(chan int)       // unbuffered — sync handoff
ch := make(chan int, 10)   // buffered — up to 10 items without blocking

// send (blocks until receiver is ready for unbuffered)
ch <- 42

// receive (blocks until value available)
v := <-ch

// close (signals no more sends)
close(ch)

// receive from closed channel returns zero value + false
v, ok := <-ch
if !ok { /* channel closed */ }

// range over channel (until closed)
for v := range ch {
    fmt.Println(v)
}
```

| C# | Go |
|----|-----|
| `Channel<T>` / `BlockingCollection<T>` | `make(chan T)` |
| `await Task.WhenAll(...)` | `wg.Wait()` + goroutines |
| `CancellationToken` | `context.Context` |
| `Task.Run(() => ...)` | `go func() { ... }()` |

## `select` — Multiplex Channels

```go
select {
case v := <-ch1:
    fmt.Println("from ch1:", v)
case v := <-ch2:
    fmt.Println("from ch2:", v)
case <-ctx.Done():
    return ctx.Err()
default:
    // non-blocking: runs if no channel is ready
}
```

## `sync.WaitGroup` — Wait for Goroutines

```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        doWork(id)
    }(i)
}

wg.Wait()  // blocks until all Done() calls
```

## `context.Context` — Cancellation & Timeouts

```go
// create with cancel
ctx, cancel := context.WithCancel(context.Background())
defer cancel()   // always call cancel to release resources

// create with timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// pass ctx through call chain
go worker(ctx)

// goroutine checks for cancellation
func worker(ctx context.Context) {
    select {
    case <-ctx.Done():
        return  // cancelled
    case job := <-jobs:
        process(job)
    }
}
```

## Worker Pool Pattern

```go
jobs := make(chan int, len(items))
results := make(chan int, len(items))

// start workers
for w := 0; w < numWorkers; w++ {
    go func() {
        for j := range jobs {
            results <- j * j
        }
    }()
}

// send jobs
for _, item := range items {
    jobs <- item
}
close(jobs)

// collect results
for range items {
    fmt.Println(<-results)
}
```

## Race Detector

```bash
go test -race ./...       # detect data races in tests
go run -race main.go      # detect in a running program
```

Always run with `-race` in CI. A data race is undefined behavior — it will corrupt data silently.

## Common Mistakes

| Mistake | Fix |
|---------|-----|
| Goroutine leak — goroutine never exits | Pass `context.Context`; close channels to signal done |
| Writing to closed channel | Only the sender should close; use `sync.Once` |
| Loop variable capture in goroutine (pre-Go 1.22) | Pass as argument `go f(i)` |
| Mutex + defer `wg.Done()` out of order | `defer wg.Done()` immediately after `wg.Add(1)` |

## Useful Links
- [Tour: Goroutines & Channels](https://go.dev/tour/concurrency)
- [Go Concurrency Patterns: Pipelines](https://go.dev/blog/pipelines)
- [Go Concurrency Patterns: Context](https://go.dev/blog/context)
- [The Go Memory Model](https://go.dev/ref/mem)
- [Race Detector](https://go.dev/doc/articles/race_detector)
