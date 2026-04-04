# Interfaces — Theory Reference

## The Big Idea

Go interfaces are satisfied **implicitly** — if a type has all the required methods, it implements the interface. No `implements` keyword. This enables decoupling without tight coupling to a specific library's interface type.

## Implicit Satisfaction

```go
// Define the interface
type Shape interface {
    Area() float64
}

// Square satisfies Shape — implicitly, no declaration needed
type Square struct{ Side float64 }
func (s Square) Area() float64 { return s.Side * s.Side }

// Circle also satisfies Shape
type Circle struct{ Radius float64 }
func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }

// Use the interface
func totalArea(shapes []Shape) float64 {
    var total float64
    for _, s := range shapes {
        total += s.Area()
    }
    return total
}
```

| C# | Go |
|----|-----|
| `interface IShape { double Area(); }` | `type Shape interface { Area() float64 }` |
| `class Square : IShape { ... }` | `type Square struct{ ... }` + `func (s Square) Area() float64` |
| Explicit declaration | Implicit satisfaction |

## Design Principle: Accept Interfaces, Return Concrete Types

```go
// Good: accepts interface (flexible)
func save(w io.Writer, data []byte) error { ... }

// Prefer: return concrete type, not interface
func newFile(path string) (*os.File, error) { ... }
```

## Small Interfaces are Idiomatic

```go
// Standard library examples — small and composable
type Reader interface { Read(p []byte) (n int, err error) }
type Writer interface { Write(p []byte) (n int, err error) }
type ReadWriter interface { Reader; Writer }  // composition
```

Prefer 1–2 method interfaces. Large interfaces are harder to implement and mock.

## Type Assertions

```go
var s Shape = Square{Side: 5}

// Type assertion — panics if wrong type
sq := s.(Square)

// Safe assertion — returns (value, bool)
sq, ok := s.(Square)
if ok {
    fmt.Println(sq.Side)
}
```

## Type Switch

```go
func describe(v any) string {
    switch x := v.(type) {
    case int:
        return fmt.Sprintf("int:%d", x)
    case string:
        return fmt.Sprintf("string:%s", x)
    default:
        return "unknown"
    }
}
```

## `any` (alias for `interface{}`)

Since Go 1.18, `any` is the preferred alias for `interface{}`:

```go
func printAnything(v any) { fmt.Println(v) }
```

## Nil Interface Gotcha

```go
// A nil *MyError is NOT a nil interface value!
var err *MyError = nil
var iface error = err
fmt.Println(iface == nil)  // false! — the interface has type info
```

Always return `nil` directly for nil errors, not a typed nil pointer:
```go
// WRONG
func bad() error {
    var err *MyError
    return err   // non-nil interface wrapping nil pointer!
}

// RIGHT
func good() error {
    return nil
}
```

## Key Standard Library Interfaces

| Interface | Package | Methods |
|-----------|---------|---------|
| `error` | builtin | `Error() string` |
| `io.Reader` | io | `Read([]byte) (int, error)` |
| `io.Writer` | io | `Write([]byte) (int, error)` |
| `io.Closer` | io | `Close() error` |
| `fmt.Stringer` | fmt | `String() string` |

## Working with `io.Reader`

The most common way to fully consume an `io.Reader` is `io.ReadAll`:

```go
import "io"

data, err := io.ReadAll(r)   // reads everything into []byte
if err != nil {
    return "", err
}
text := string(data)
```

`strings.NewReader("hello")` creates an `io.Reader` from a string — useful in tests.
`strings.ToUpper(s)` converts a string to uppercase — like C#'s `s.ToUpper()`.

## Use Cases

- **Dependency injection**: accept `io.Writer` instead of `*os.File` — callers can pass a file, a buffer, or a network connection; tests can pass `bytes.Buffer`
- **Test doubles (mocking)**: define a small interface, implement it with a fake in tests — no mocking framework needed:
  ```go
  type EmailSender interface { Send(to, body string) error }

  type fakeSender struct{ sent []string }
  func (f *fakeSender) Send(to, body string) error {
      f.sent = append(f.sent, to)
      return nil
  }
  ```
- **Plugin extensibility**: `sort.Interface`, `http.Handler`, `io.Reader` — the standard library is built entirely on small interfaces

## Common Mistakes & Caveats

- **Interface pollution** — don't define an interface for every type up front "just in case". Define interfaces at the point of use, in the *consuming* package, not next to the concrete type. The concrete type shouldn't know the interface exists.
- **Type assertion without the comma-ok form panics**:
  ```go
  sq := s.(Square)        // ! panics if s is not a Square — interface conversion panic
  sq, ok := s.(Square)   // safe — ok is false instead of panic ✓
  ```
- **Nil interface ≠ nil concrete pointer** — the most surprising thing in Go:
  ```go
  var err *MyError = nil   // typed nil pointer
  var iface error = err    // ! wraps (type=*MyError, data=nil) into interface
  iface == nil             // ! FALSE — the interface is NOT nil because type is set
  
  // Always return bare nil from error-returning functions:
  return nil       // ✓ iface == nil is true
  return (*MyError)(nil)  // ! iface == nil is false — don't do this
  ```
- **Overusing `any`** — `func process(v any)` loses type safety; if all callers pass the same type, use that type; if there are multiple types with common behaviour, use an interface; only use `any` when truly heterogeneous
- **Large interfaces are fragile** — an interface with 10 methods is hard to implement, hard to mock in tests, and hard to compose. Split into small, single-purpose interfaces.

## Useful Links
- [Tour: Interfaces](https://go.dev/tour/methods/9)
- [Tour: Type switches](https://go.dev/tour/methods/16)
- [Go Wiki: MethodSets](https://go.dev/wiki/MethodSets)
- [Code Review Comments](https://go.dev/wiki/CodeReviewComments)
