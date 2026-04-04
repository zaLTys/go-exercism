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

## Actor Pattern — Goroutine Owns State via Channels

Instead of using a mutex to protect shared state, one goroutine can **own** the data and communicate via channels. Other goroutines send requests and receive responses. This avoids data races by design.

```go
type cmd struct {
    kind    string
    reply   chan int
}

func newCounter() (inc func() int, val func() int, stop func()) {
    ch := make(chan cmd)

    go func() {
        n := 0
        for c := range ch {
            switch c.kind {
            case "inc":
                n++
                c.reply <- n
            case "val":
                c.reply <- n
            }
        }
    }()

    inc = func() int {
        r := make(chan int, 1)
        ch <- cmd{"inc", r}
        return <-r
    }
    val = func() int {
        r := make(chan int, 1)
        ch <- cmd{"val", r}
        return <-r
    }
    stop = func() { close(ch) }  // closing channel ends the goroutine's range loop

    return inc, val, stop
}
```

Key insight: the goroutine is the **only** one reading/writing `n`. No mutex needed.
Use `sync.Once` when `stop()` must be safe to call multiple times (closing an already-closed channel panics).

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

## Use Cases

- **Fan-out / fan-in**: scatter N items to a worker pool, collect results in order — the classic pipeline pattern
- **Parallel HTTP requests with rate limiting**: bounded goroutine pool (semaphore via buffered channel) to fetch URLs concurrently without hammering the server
- **Background periodic task**: `for { select { case <-ticker.C: doWork(); case <-ctx.Done(): return } }`
- **Cache owned by one goroutine**: actor pattern with channels — no mutex needed, state never leaves the owning goroutine

## Common Mistakes & Caveats

- **Goroutine leaks** — a goroutine blocked on a channel send/receive that nobody ever reads from runs forever. Always give goroutines a way to exit:
  ```go
  // ! BAD — goroutine leaks if nobody reads results; runs until process exit
  go func() { results <- compute() }()

  // GOOD — goroutine respects cancellation ✓
  go func() {
      select {
      case results <- compute():
      case <-ctx.Done():
      }
  }()
  ```
- **Deadlock**: if all goroutines are blocked (e.g. unbuffered channel with no receiver), Go's runtime detects it and panics:
  ```
  fatal error: all goroutines are asleep - deadlock!
  ```
  Common cause: `ch := make(chan int); ch <- 1` with no goroutine reading from `ch`.
- **Closing a channel twice panics** — only the sender should close; if multiple goroutines might close, use `sync.Once`:
  ```go
  var once sync.Once
  close := func() { once.Do(func() { close(ch) }) }
  ```
- **Don't use goroutines for everything** — goroutines add synchronization complexity; for small, fast, sequential work they're net overhead. Use concurrency when tasks are truly independent and parallel execution helps.
- **`sync.Mutex` vs channels**: use a mutex when protecting a shared data structure; use channels when coordinating work or ownership between goroutines. Both are valid — match the tool to the problem.

## Useful Links
- [Tour: Goroutines & Channels](https://go.dev/tour/concurrency)
- [Go Concurrency Patterns: Pipelines](https://go.dev/blog/pipelines)
- [Go Concurrency Patterns: Context](https://go.dev/blog/context)
- [The Go Memory Model](https://go.dev/ref/mem)
- [Race Detector](https://go.dev/doc/articles/race_detector)
