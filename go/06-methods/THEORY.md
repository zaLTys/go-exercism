# Methods — Theory Reference

## The Big Idea

In Go, a **method** is just a function with a **receiver** parameter. No classes — methods live on types (structs, type aliases, etc.). The critical design decision you make: **value receiver vs pointer receiver**.

## Syntax

```go
type Counter struct{ n int }

// pointer receiver — can mutate the struct
func (c *Counter) Inc() int {
    c.n++
    return c.n
}

// value receiver — works on a copy, can't mutate
func (c Counter) Value() int {
    return c.n
}

c := Counter{}
c.Inc()      // Go auto-takes address: (&c).Inc()
c.Value()    // works on both *Counter and Counter
```

## Value vs Pointer Receiver — Rules

| Use **pointer** receiver when... | Use **value** receiver when... |
|----------------------------------|-------------------------------|
| Method mutates the struct | Method only reads |
| Struct is large (avoid copy) | Type is small / cheap to copy |
| You want consistent interface satisfaction | Simple read-only computation |

**Rule of thumb:** If any method on a type uses a pointer receiver, use pointer receivers for all methods on that type (for consistency with interface satisfaction).

| C# | Go |
|----|-----|
| `class Counter { private int _n; public int Inc() => ++_n; }` | `type Counter struct{ n int }` + `func (c *Counter) Inc() int` |
| Instance method | Method with receiver |
| `new Counter()` | `Counter{}` or `&Counter{}` |

## Zero Value Usability Pattern

Design types so the zero value is useful:

```go
type Stack struct {
    items []int   // nil slice — safe to append to
}

func (s *Stack) Push(v int) {
    s.items = append(s.items, v)
}

// zero value Stack{} is immediately usable:
var s Stack
s.Push(1)
```

## Methods on Non-Struct Types

You can define methods on any named type in the same package:

```go
type Celsius float64
type Fahrenheit float64

func (c Celsius) ToF() Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}
```

## Auto-Dereferencing

```go
p := &Counter{}
p.Inc()       // same as (*p).Inc() — Go auto-dereferences
p.Value()     // also fine
```

## Method Sets (Interface Satisfaction)

- Value `T`: has methods with value receivers only
- Pointer `*T`: has methods with both value AND pointer receivers

This is critical for interface satisfaction:

```go
type Incr interface { Inc() int }

var i Incr = &Counter{}   // OK — *Counter has Inc() (pointer receiver)
var i Incr = Counter{}    // COMPILE ERROR — Counter does not have Inc()
```

## Useful Links
- [Tour: Methods](https://go.dev/tour/methods)
- [Go Wiki: MethodSets](https://go.dev/wiki/MethodSets)
- [Effective Go: Methods](https://go.dev/doc/effective_go#methods)
