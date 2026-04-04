# Control Flow — Theory Reference

## The Big Idea

Go has a **deliberately small** set of control-flow constructs: `if`, `for` (the only loop keyword — no `while`, no `do`), and `switch`. No parentheses around conditions. `defer` also belongs here — it queues calls to run when the function returns.

## `if` Statement

```go
if x > 0 {
    fmt.Println("positive")
} else if x < 0 {
    fmt.Println("negative")
} else {
    fmt.Println("zero")
}

// if with init statement (scope limited to block)
if err := doSomething(); err != nil {
    return err
}
```

## `for` — The Only Loop

```go
// classic C-style
for i := 0; i < 10; i++ { ... }

// while equivalent
for x < 100 { x *= 2 }

// infinite loop
for {
    if done { break }
}

// range over slice
for i, v := range nums { ... }
for _, v := range nums { ... }   // ignore index
for i := range nums { ... }      // index only

// range over map (order not guaranteed)
for k, v := range m { ... }

// range over string (yields rune, not byte)
for i, r := range "hello" { ... }

// range over channel
for msg := range ch { ... }
```

| C# | Go |
|----|-----|
| `foreach (var x in xs)` | `for _, x := range xs` |
| `while (cond) { }` | `for cond { }` |
| `for (int i=0; i<n; i++)` | `for i := 0; i < n; i++` |
| `break` / `continue` | `break` / `continue` (same) |

## `switch`

```go
switch kind {
case "A":
    return 1
case "B", "C":    // multiple values per case
    return 2
default:
    return 0
}

// switch with no condition (acts like if/else chain)
switch {
case x < 0:
    return "negative"
case x == 0:
    return "zero"
default:
    return "positive"
}

// type switch (used with interfaces)
switch v := val.(type) {
case int:
    fmt.Println("int:", v)
case string:
    fmt.Println("string:", v)
}
```

**Key difference from C#:** No implicit fallthrough! Each case breaks automatically. Use `fallthrough` keyword to fall through explicitly (rare).

## `defer`

```go
f, err := os.Open(path)
if err != nil { return err }
defer f.Close()   // runs when surrounding function returns, LIFO order

// multiple defers run in LIFO (last-in, first-out)
defer fmt.Println("1")
defer fmt.Println("2")   // prints "2" then "1"
```

`defer` ≈ C# `using` / `try/finally` — used for cleanup.

## `break` with Labels

```go
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i+j > 2 {
            break outer   // breaks both loops
        }
    }
}
```

## Use Cases

- **Retry loop**: `for attempts := 0; attempts < maxRetries; attempts++` with a `break` on success
- **State machine**: `switch state { case "idle": ... case "running": ... }` — clean and exhaustive
- **Resource cleanup**: `defer f.Close()` immediately after `os.Open` — runs even if later code panics
- **Early return guard**: `if err := validate(input); err != nil { return err }` — keeps the happy path unindented

## Common Mistakes & Caveats

- **`defer` in a loop opens resource leaks** — each iteration schedules a new defer; they all fire at function return, not at loop end:
  ```go
  for _, path := range paths {
      f, _ := os.Open(path)
      defer f.Close()   // ! ALL files stay open until function returns — not end of iteration
  }
  // Fix: wrap in an inner function so defer fires per iteration
  for _, path := range paths {
      func() {
          f, _ := os.Open(path)
          defer f.Close()  // now defers at inner-func return ✓
      }()
  }
  ```
- **`defer` arguments are evaluated immediately** — the value is captured at the `defer` line, not when it runs:
  ```go
  x := 1
  defer fmt.Println(x)   // ! captures x=1 NOW — prints 1, even though x changes below
  x = 2
  // ! If you want to capture the final value, use a closure:
  defer func() { fmt.Println(x) }()   // captures x by reference — prints 2
  ```
- **`switch` has no implicit fallthrough** — unlike C/C#, each case exits automatically; use `fallthrough` explicitly (rare)
- **Loop variable capture in goroutines (pre-Go 1.22)** — `go func() { fmt.Println(i) }()` inside a range loop captures the variable by reference, not value; by the time the goroutine runs, `i` is the final value. Fix: `i := i` before the `go` statement, or pass as argument `go func(n int) { ... }(i)`
- **`for range` on a nil slice is safe** — it simply iterates zero times; no nil check needed

## Useful Links
- [Tour: Flow control](https://go.dev/tour/flowcontrol)
- [Go Spec: Statements](https://go.dev/ref/spec#Statements)
- [Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)
