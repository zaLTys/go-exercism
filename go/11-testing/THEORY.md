# Testing — Theory Reference

## The Big Idea

Go's standard `testing` package + `go test` command give you everything you need. The idiomatic style is **table-driven tests** (a slice of test cases iterated in a loop) with **subtests** (`t.Run`). No test frameworks or attributes needed.

## Basic Test Structure

```go
// mymath/math_test.go
package mymath

import "testing"

func TestAdd(t *testing.T) {
    if got := Add(1, 2); got != 3 {
        t.Fatalf("Add(1,2)=%d; want 3", got)
    }
}
```

Rules:
- Test file ends in `_test.go`.
- Test function starts with `Test` + uppercase letter.
- Takes `*testing.T` as its only argument.

| C# (xUnit) | Go |
|-----------|----|
| `[Fact]` attribute | Function named `TestXxx` |
| `Assert.Equal(want, got)` | `if got != want { t.Fatalf(...) }` |
| `Assert.Throws<T>()` | `if err == nil { t.Fatalf("expected error") }` |
| Test class constructor / `ITestOutputHelper` | `t.Log`, `t.Helper()` |

## Table-Driven Tests

```go
func TestDiv(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive", 10, 2, 5},
        {"zero divisor", 10, 0, 0},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got := div(tc.a, tc.b)
            if got != tc.want {
                t.Fatalf("div(%d,%d)=%d; want %d", tc.a, tc.b, got, tc.want)
            }
        })
    }
}
```

## `t.Run` — Subtests

```go
t.Run("subtest name", func(t *testing.T) {
    // runs as a named subtest
    // can be run individually: go test -run TestFoo/subtest_name
})
```

## Reporting Failures

| Function | Behavior |
|----------|---------|
| `t.Fatalf(format, ...)` | Log message and **stop** the test immediately |
| `t.Errorf(format, ...)` | Log message but **continue** the test |
| `t.Logf(format, ...)` | Log (only shown with `-v` or on failure) |
| `t.Helper()` | Mark as helper — errors point to caller, not helper |
| `t.Skip(...)` | Skip the test |

## Running Tests

```bash
go test ./...                         # all packages
go test -v ./...                      # verbose output
go test -run TestFoo ./...            # only tests matching "TestFoo"
go test -run TestFoo/subtest ./...    # specific subtest
go test -race ./...                   # with race detector
go test -count=1 ./...                # disable test caching
go test -bench=. ./...                # run benchmarks
go test -cover ./...                  # coverage report
```

## Benchmarks

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(1, 2)
    }
}
```

## Test Helpers

```go
func assertEqual(t *testing.T, got, want int) {
    t.Helper()   // important: errors blame the caller, not this function
    if got != want {
        t.Fatalf("got %d; want %d", got, want)
    }
}
```

## Test Main (Setup/Teardown)

```go
func TestMain(m *testing.M) {
    // setup
    code := m.Run()
    // teardown
    os.Exit(code)
}
```

## Useful Links
- [testing package docs](https://pkg.go.dev/testing)
- [Go Wiki: TableDrivenTests](https://go.dev/wiki/TableDrivenTests)
- [Go by Example: Testing](https://gobyexample.com/testing)
