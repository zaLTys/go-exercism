# 📘 Exercise 4 — Go Concurrency Best Practices (Worker Pools, Context, Cancellation, Error Handling)

This exercise is a **big jump forward** — it mirrors real-world Go concurrency patterns used in production systems such as:
- background job processors  
- HTTP handlers with cancellation  
- distributed workers  
- data pipelines  
- services that must not leak goroutines  

You will build a **cancellable worker pool** that processes integers, returns squared values, and safely stops on error.

---

# 🧠 Concepts Covered

This exercise intentionally combines several best practices that experienced Go developers use.

---

## ✅ 1. Worker Pool (Fan-Out)

Instead of one goroutine processing everything, you run N workers that all pull tasks from the **same jobs channel**:

```
      jobs channel
           |
      +----+----+
      v         v
  worker 1   worker 2   ... worker N
      \         |         /
       +--------+--------+
                |
         results channel
```

A worker pool allows:
- parallel processing
- controlled concurrency
- good throughput

Each worker:
1. Reads from `jobs <-chan int`
2. Performs work
3. Sends the result to `results chan<- Result`

---

## ✅ 2. `sync.WaitGroup` to Wait for All Workers

Workers run concurrently, so the main goroutine needs to know *when all workers are done*:

```go
var wg sync.WaitGroup
wg.Add(workers)

for i := 0; i < workers; i++ {
    go func() {
        defer wg.Done()
        // do work
    }()
}

wg.Wait()  // ensures all workers finished
```

This is the standard pattern for controlled goroutine shutdown.

---

## ✅ 3. Context Cancellation (Clean Shutdown)

Context lets the caller cancel the entire processing pipeline:

```go
ctx, cancel := context.WithCancel(ctx)
```

Workers must check:

```go
select {
case <-ctx.Done():
    return  // stop early
default:
}
```

This ensures:
- all workers exit immediately on cancellation  
- no leaked goroutines  
- wasted work is avoided  

Cancellation happens:
- when user passes a cancelled context  
- when timeout occurs  
- when *your code* calls `cancel()` after encountering an error  

---

## ✅ 4. Structured Error Propagation

Instead of letting workers return errors directly, you use a **results channel**:

```go
type Result struct {
    Value int
    Err   error
}
```

Each worker reports:
- `Result{Value: square}` on success  
- `Result{Err: someError}` on failure  

The main goroutine:
- reads from the result channel  
- returns the **first error encountered**  
- triggers `cancel()`  
- drains remaining results to avoid leaks  

---

## ✅ 5. Avoiding Goroutine Leaks (Very Important!)

**Most concurrency bugs in Go come from goroutines that never exit.**

This exercise forces you to close channels in correct order:

### Correct shutdown sequence:
1. Main goroutine closes the **jobs** channel  
2. Workers finish → they call `wg.Done()`  
3. After `wg.Wait()`, main goroutine closes **results** channel  
4. Main goroutine reads all results until results channel is closed  

### Why?
- Workers must stop reading when jobs is closed  
- Main goroutine must stop reading when results is closed  
- Workers must never be blocked on sending  
- Nobody should write to a closed channel  

This is the correct, leak-free lifecycle.

---

## 🎯 Your Task

Implement:

```go
func ProcessNumbers(ctx context.Context, nums []int, workers int) ([]int, error)
```

Processing rules:

| Input | Output |
|-------|---------|
| `n >= 0` | Return `n*n` |
| `n < 0`  | Return error `"negative number"` and cancel EVERYTHING |

### Behavior requirements:

### 🟢 When all numbers are ≥ 0:
- All squares are returned (order does not matter)
- `error == nil`

### 🔴 When any number is < 0:
- Stop immediately  
- Cancel all workers  
- Return only results for tasks processed **before the error**  

Example:  

Input: `[2, -1, 5, 7]`  
Output:  
- `results = [4]`  
- `error = "negative number"`  

Because `-1` triggers cancellation.

### ⛔ When context is already cancelled:
- Return error immediately  
- No results  

---

# 🛠 Step-by-Step Implementation Guide

Follow these steps in order. Each step builds on the previous one. The actual exercise implements `FetchAll`, not `ProcessNumbers` - the steps below are adapted for the correct function.

---

## Step 1: Check if Context is Already Cancelled

**What:** Before doing any work, check if the context is already done.

**Why:** Avoid wasting resources if the caller has already cancelled the operation.

**How:**
```go
select {
case <-ctx.Done():
    return nil, ctx.Err()
default:
}
```

**💡 Tip:** This is a non-blocking check. If `ctx.Done()` is ready, we return immediately. Otherwise, we continue.

**⚠️ Common Mistake:** Forgetting this check means your function will start work even when the context is cancelled.

---

## Step 2: Create a Cancellable Context

**What:** Wrap the incoming context with your own cancel function.

**Why:** We need the ability to cancel all workers if one encounters an error.

**How:**
```go
ctx, cancel := context.WithCancel(ctx)
defer cancel()
```

**💡 Tip:** Always `defer cancel()` to ensure cleanup happens even if you return early or panic.

**⚠️ Common Mistake:** Not calling `cancel()` can leak resources. The `defer` ensures it's always called.

---

## Step 3: Create Result Channel and Result Type

**What:** Define a struct to hold either a value or an error, then create a buffered channel.

**Why:** Workers need to send results back, and we need to handle both success and failure cases.

**How:**
```go
type result struct {
    value string
    err   error
}

results := make(chan result, len(sources))
```

**💡 Tip:** The buffer size equals `len(sources)` so workers never block when sending results.

**⚠️ Common Mistake:** Unbuffered channel = workers block forever waiting for receiver. This causes goroutine leaks!

---

## Step 4: Create WaitGroup

**What:** Set up a WaitGroup to track when all goroutines finish.

**Why:** We need to know when it's safe to close the results channel.

**How:**
```go
var wg sync.WaitGroup
wg.Add(len(sources))
```

**💡 Tip:** Call `Add()` before launching goroutines to avoid race conditions.

**⚠️ Common Mistake:** Calling `Add()` inside the goroutine can cause the main thread to call `Wait()` before `Add()` runs.

---

## Step 5: Launch Worker Goroutines (Fan-Out)

**What:** Start one goroutine per source to fetch data concurrently.

**Why:** This is the "fan-out" pattern - parallel execution for better performance.

**How:**
```go
for _, src := range sources {
    src := src  // IMPORTANT: Capture loop variable!
    go func() {
        defer wg.Done()  // Always signal completion
        
        val, err := src(ctx)
        results <- result{value: val, err: err}
    }()
}
```

**💡 Tip:** `src := src` creates a new variable for each iteration. Without this, all goroutines share the same variable!

**⚠️ Common Mistake:** Forgetting `src := src` means all goroutines call the LAST source in the slice.

**⚠️ Common Mistake:** Forgetting `defer wg.Done()` causes `wg.Wait()` to block forever (deadlock).

---

## Step 6: Close Results Channel When Workers Finish

**What:** Launch a goroutine that waits for all workers, then closes the results channel.

**Why:** Closing the channel signals to the collection loop that no more results are coming.

**How:**
```go
go func() {
    wg.Wait()       // Wait for all workers to finish
    close(results)  // Then close the channel
}()
```

**💡 Tip:** This runs in a separate goroutine so the main function can proceed to collecting results.

**⚠️ Common Mistake:** Closing `results` BEFORE `wg.Wait()` → workers panic when sending to closed channel.

**⚠️ Common Mistake:** Not closing `results` → the collection loop blocks forever (deadlock).

---

## Step 7: Collect Results (Fan-In)

**What:** Read from the results channel until it closes, handling both successes and errors.

**Why:** This is the "fan-in" pattern - gathering all concurrent results into one place.

**How:**
```go
var (
    collected []string
    firstErr  error
)

for res := range results {
    if res.err != nil {
        if firstErr == nil {
            firstErr = res.err  // Remember first error
            cancel()            // Cancel remaining workers
        }
        continue  // Keep draining the channel!
    }
    
    // Only collect successful results if we haven't seen an error
    if firstErr == nil {
        collected = append(collected, res.value)
    }
}

return collected, firstErr
```

**💡 Tip:** The `range` loop exits when the channel is closed (Step 6).

**💡 Tip:** We continue draining even after an error to prevent goroutine leaks.

**⚠️ Common Mistake:** Returning immediately on error → workers block trying to send → goroutine leak!

**⚠️ Common Mistake:** Not calling `cancel()` when error occurs → remaining workers keep running unnecessarily.

---

# 🧪 Testing Your Implementation

## Running Tests

```bash
# Run all tests
go test

# Run with race detector (highly recommended!)
go test -race

# Run specific test
go test -run TestFetchAllSuccess

# Verbose output
go test -v
```

## Test Scenarios Explained

### ✅ TestFetchAllSuccess
**Scenario:** All sources return data successfully  
**Expected:** 
- Return all 3 results
- No error
- All results collected correctly

### ❌ TestFetchAllError  
**Scenario:** One source returns an error during execution  
**Expected:**
- Function returns an error
- Context is cancelled (stopping other workers)
- No panic or deadlock

### 🚫 TestFetchAllContextCancelled
**Scenario:** Context is already cancelled before `FetchAll` is called  
**Expected:**
- Return error immediately
- No goroutines launched
- Fast failure

---

# ⚠️ Common Pitfalls - Quick Reference

| Mistake | Consequence | Fix |
|---------|-------------|-----|
| Unbuffered results channel | Workers block forever | `make(chan result, len(sources))` |
| Not closing results channel | Collection loop never exits | Close after `wg.Wait()` |
| Closing results too early | Workers panic on send | Close only after all `wg.Done()` |
| Not draining on error | Workers block on send | `continue` reading until channel closes |
| Forgetting `defer wg.Done()` | Deadlock on `wg.Wait()` | Always use `defer wg.Done()` |
| Not capturing loop variable | All goroutines use last value | `src := src` inside loop |
| Not calling `cancel()` | Resource leak | `defer cancel()` |

---

# 🐛 Debugging Tips

## Check for Deadlocks
If your tests hang forever, you likely have a deadlock. Common causes:
- Forgot to close the results channel
- Forgot `defer wg.Done()` in a worker
- Workers are blocked trying to send to an unbuffered channel

## Check for Goroutine Leaks
```bash
# Run with race detector
go test -race
```

Add this to verify goroutines complete:
```go
// At the end of your test
runtime.GC()
time.Sleep(100 * time.Millisecond)
fmt.Println("Goroutines:", runtime.NumGoroutine())
```

## Test with Print Statements
Temporarily add logging to understand execution flow:
```go
fmt.Println("Launching worker", i)
fmt.Println("Worker received:", val, err)
fmt.Println("Collected result")
```

## Test Context Cancellation
Verify your code respects cancellation:
```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
defer cancel()
```

---

# ⭐ What You'll Learn

After completing this exercise, you'll understand:

✓ **Fan-Out/Fan-In Pattern** - The foundation of concurrent processing  
✓ **Context Cancellation** - How to stop work gracefully  
✓ **WaitGroups** - Coordinating multiple goroutines safely  
✓ **Channel Patterns** - Buffered channels and proper closing  
✓ **Error Propagation** - Handling errors in concurrent code  
✓ **Goroutine Lifecycle** - Preventing leaks and ensuring cleanup  

This is one of the most important patterns in production Go services.

---

# 🚀 Next Steps & Advanced Challenges

Once you've completed the basic implementation, try these extensions:

## 1. Add Timeout Support
```go
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
```
What happens if sources take too long?

## 2. Limit Concurrent Workers
Instead of one goroutine per source, use a worker pool:
- Create N workers (e.g., 3 workers)
- Feed sources through a channel
- Workers process from the channel

## 3. Collect Partial Results on Error
Currently, we might lose some successful results when an error occurs. Modify the code to return ALL successfully completed results even when there's an error.

## 4. Add Retry Logic
If a source fails, retry it up to 3 times before giving up.

## 5. Add Metrics
Track:
- How many sources succeeded vs failed
- Average response time
- Which source was slowest

## 6. Implement First-Win Pattern
Return as soon as ANY source succeeds (useful for trying multiple mirrors of the same data).

---

# 📚 Additional Resources

- [Go Blog: Concurrency Patterns](https://go.dev/blog/pipelines)
- [Effective Go: Concurrency](https://go.dev/doc/effective_go#concurrency)
- [Go by Example: Worker Pools](https://gobyexample.com/worker-pools)

---

**Good luck — this one is a milestone toward professional-level Go concurrency mastery!** 🚀
