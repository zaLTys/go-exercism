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

## Useful Links
- [Tour: Flow control](https://go.dev/tour/flowcontrol)
- [Go Spec: Statements](https://go.dev/ref/spec#Statements)
- [Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)
