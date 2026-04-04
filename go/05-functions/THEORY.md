# Functions — Theory Reference

## The Big Idea

Functions in Go are **first-class values**. The most important difference from C#: **multiple return values** are the norm (especially `(value, error)` pairs). No exceptions — errors are just returned. Closures exist and are used heavily (but not for giant expression pipelines).

## Multiple Return Values

```go
func divMod(a, b int) (int, int) {
    return a / b, a % b
}

q, r := divMod(10, 3)   // q=3, r=1
_, r := divMod(10, 3)   // ignore quotient with _
```

| C# | Go |
|----|-----|
| `(int q, int r) DivMod(int a, int b) => (a/b, a%b)` | `func divMod(a, b int) (int, int)` |
| `out` parameters | just return multiple values |
| `throw new Exception(...)` | `return 0, errors.New("...")` |

## Named Return Values

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return   // "naked return" — returns x and y
}
```

Prefer named returns for documentation purposes on short functions. Avoid naked returns in long functions (reduces readability).

## Error Return Pattern

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// caller checks the error
result, err := divide(10, 0)
if err != nil {
    fmt.Println("error:", err)
}
```

## Variadic Functions

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

sum()           // 0
sum(1, 2, 3)    // 6

// spread a slice
nums := []int{1, 2, 3}
sum(nums...)    // spread operator
```

`fmt.Println`, `fmt.Sprintf`, `append` are all variadic.

## Closures

```go
func makeAdder(x int) func(int) int {
    return func(y int) int {
        return x + y   // closes over x
    }
}

add5 := makeAdder(5)
fmt.Println(add5(3))  // 8
```

```go
// counter closure
func makeCounter() func() int {
    n := 0
    return func() int {
        n++
        return n
    }
}

next := makeCounter()
next()  // 1
next()  // 2
```

## Functions as Values

```go
// function type
type transform func(int) int

func apply(nums []int, f transform) []int {
    out := make([]int, len(nums))
    for i, v := range nums {
        out[i] = f(v)
    }
    return out
}

doubled := apply([]int{1, 2, 3}, func(n int) int { return n * 2 })
```

## `init` Function

Each package can have one or more `init()` functions — they run before `main()` and before any other code in the package. Use sparingly.

## Use Cases

- **Middleware / decorator**: a function that wraps another function to add logging, timing, or auth:
  ```go
  func withLogging(fn func() error) func() error {
      return func() error {
          log.Println("starting")
          err := fn()
          log.Println("done:", err)
          return err
      }
  }
  ```
- **Retry with backoff**: accept a `func() error` as an argument, call it up to N times
- **Functional options**: accept `...Option` where `type Option func(*Config)` — covered in detail in topic 13

## Common Mistakes & Caveats

- **Never silently ignore errors** — `result, _ := parse(s)` buries bugs; at minimum log them:
  ```go
  result, err := parse(s)
  if err != nil {
      return 0, fmt.Errorf("parse failed: %w", err)
  }
  ```
- **Closure captures variable by reference** — in a loop before Go 1.22, closures share the loop variable:
  ```go
  for i := 0; i < 3; i++ {
      fns = append(fns, func() int { return i })  // ! all return 3 — they share ONE variable i
  }
  // Fix: create a new binding per iteration
  for i := 0; i < 3; i++ {
      i := i  // ! new i variable scoped to this iteration
      fns = append(fns, func() int { return i })  // each closure has its own i ✓
  }
  ```
- **Calling a nil function panics** — `var f func(); f()` panics with "nil function"; check `if f != nil` before calling
- **Named returns + defer = subtle bugs** — a deferred function can modify named return values:
  ```go
  func risky() (result int, err error) {
      defer func() {
          if err != nil { result = -1 }  // ! overwrites result AFTER "return" already set it
      }()
      ...
  }
  ```
  This is sometimes intentional (adding context on error) but surprises people who don't expect it.
- **Never silently discard errors** — `result, _ = parse(s)` compiles fine but hides real bugs:
  ```go
  n, _ := strconv.Atoi(s)   // ! if s is "abc", n=0 silently — always handle the error
  ```

## Useful Links
- [Tour: Functions](https://go.dev/tour/basics/4)
- [Tour: Closures](https://go.dev/tour/moretypes/25)
- [Go Spec: Function types](https://go.dev/ref/spec#Function_types)
