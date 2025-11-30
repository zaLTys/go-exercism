# 📘 **README --- Exercise 1: Introduction to Goroutines & Channels**

## 🚀 Welcome

This exercise introduces two of Go's most important concurrency tools:

-   **goroutines** --- lightweight concurrent functions\
-   **channels** --- typed communication pathways between goroutines

Understanding these two concepts is the foundation for all Go
concurrency patterns.

This exercise guides you step by step, with short examples and hints to
help you complete `RunPingPong`.

------------------------------------------------------------------------

# 🧠 What Are Goroutines?

A **goroutine** is a function running *concurrently* (not necessarily in
parallel) with the rest of your program.

You start a goroutine by prefixing a function call with the `go`
keyword:

``` go
go doSomething()
```

Goroutines are extremely lightweight and scheduled by the Go runtime.

------------------------------------------------------------------------

# 📮 What Are Channels?

A **channel** is a messaging pipe that goroutines use to communicate.

``` go
ch := make(chan int)
ch <- 10
v := <-ch
close(ch)
```

Channels synchronize goroutines by default.

------------------------------------------------------------------------

# 🔗 Goroutines + Channels Together

``` go
ch := make(chan string)

go func() {
    ch <- "hello"
}()

msg := <-ch
fmt.Println(msg)
```

------------------------------------------------------------------------

# ⏳ Waiting for Goroutines

Using a WaitGroup:

``` go
var wg sync.WaitGroup
wg.Add(2)
go func(){ defer wg.Done() }()
go func(){ defer wg.Done() }()
wg.Wait()
```

------------------------------------------------------------------------

# 🎯 Task

Implement:

``` go
func RunPingPong(count int) []string
```

Your function must:

1.  Create a channel.
2.  Launch two goroutines:
    -   one sends "ping" count times
    -   the other sends "pong" count times
3.  Receive exactly 2×count messages.
4.  Close the channel after both goroutines are done.
5.  Return a slice of all messages.

Order does **not** matter.

------------------------------------------------------------------------

# 📌 Hints

-   Use a `sync.WaitGroup` or a done channel.
-   Do NOT use `time.Sleep`.
-   You know how many messages you expect, so loop that many times.

------------------------------------------------------------------------

# ✔️ Success Criteria

If all tests pass, you have learned:

-   starting goroutines\
-   basic channel communication\
-   synchronization\
-   non-deterministic scheduling\
-   collecting concurrent results
