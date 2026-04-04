# Basic Syntax — Theory Reference

## The Big Idea

Go syntax is **intentionally small**. No semicolons (the formatter inserts them). No parentheses around `if`/`for` conditions. Formatting is enforced by `gofmt` — all Go code looks structurally the same.

## Package & Import

```go
package main          // every file starts with this

import "fmt"          // single import
import (              // grouped import (preferred)
    "fmt"
    "strings"
)
```

- Unused imports are **compile errors** (not warnings).
- Package name = last element of import path: `import "math/rand"` → use as `rand.Intn()`.

## Variable Declaration

```go
// var keyword — explicit type or inferred
var x int        // zero value: 0
var s string     // zero value: ""
var b bool       // zero value: false

// short declaration — inside functions only
x := 42
s := "hello"

// const
const Pi = 3.14159
const MaxRetries = 3
```

| C# | Go |
|----|-----|
| `int x = 42;` | `x := 42` or `var x int = 42` |
| `var x = 42;` | `x := 42` |
| `const int Max = 10;` | `const Max = 10` |
| `string.IsNullOrWhiteSpace(s)` | `strings.TrimSpace(s) == ""` |

## Function Syntax

```go
func add(a, b int) int {
    return a + b
}

// multiple returns
func divmod(a, b int) (int, int) {
    return a / b, a % b
}
```

## Control Flow (brief — covered in depth in 04)

```go
if x > 0 {           // no parentheses!
    fmt.Println("positive")
}

for i := 0; i < 10; i++ {    // the only loop keyword
    fmt.Println(i)
}
```

## Formatting Rules

- `gofmt` is the standard — run it on save.
- Tabs for indentation (not spaces).
- Opening `{` always on the same line as the statement.
- No trailing commas needed in single-line calls.

```go
// WRONG — won't compile (opening brace must be on same line)
func bad()
{
}

// RIGHT
func good() {
}
```

## .NET → Go Quick Mapping

| C# | Go |
|----|-----|
| `Console.WriteLine(...)` | `fmt.Println(...)` |
| `string.Join(", ", parts)` | `strings.Join(parts, ", ")` |
| `s.Trim()` | `strings.TrimSpace(s)` |
| `s.ToLower()` | `strings.ToLower(s)` |
| `s.StartsWith("go")` | `strings.HasPrefix(s, "go")` |
| `string.IsNullOrWhiteSpace(s)` | `strings.TrimSpace(s) == ""` |

## Use Cases

- **Text utilities**: trimming, joining, splitting strings — the `strings` package covers almost everything you need without custom helpers
- **Config parsing**: reading `KEY=VALUE` lines from a file using `strings.Split`, `strings.TrimSpace`, and `strings.HasPrefix`
- **Building a greeting service**: combine variable declarations, string formatting with `fmt.Sprintf`, and simple `if` logic

## Common Mistakes & Caveats

- **`:=` only works inside functions** — package-level variables must use `var x = ...`; using `:=` at package level is a syntax error
- **Unused variables are compile errors** — `x := 5` then never using `x` won't compile; use `_` to explicitly discard: `_ = x`
- **Variable shadowing with `:=`** — a new `:=` inside an inner block creates a new variable, not an update:
  ```go
  x := 1
  if true {
      x := 2   // NEW variable — outer x is still 1!
      _ = x
  }
  ```
- **Strings are immutable UTF-8 byte sequences** — indexing `s[i]` gives a `byte` (uint8), not a character; use `[]rune(s)` for character-level access
- **`fmt.Println` adds spaces between args; `fmt.Printf` does not** — `fmt.Printf("hello %s\n", name)` is what you want for formatted output

## Useful Links
- [Tour of Go — Basics](https://go.dev/tour/basics)
- [Go Spec](https://go.dev/ref/spec)
- [gofmt blog](https://go.dev/blog/gofmt)
- [strings package](https://pkg.go.dev/strings)
