# Idiomatic Patterns — Theory Reference

## The Big Idea

Idiomatic Go is **predictable and readable**. It's less about clever tricks and more about consistent structure: useful zero values, small interfaces, explicit errors, channel-based orchestration, and `gofmt`-enforced formatting. These patterns are described in Effective Go and the Go Code Review Comments.

## Zero Value Usability

Design types so the **zero value is immediately usable** without initialization:

```go
// sync.Mutex — zero value is an unlocked mutex
var mu sync.Mutex
mu.Lock()

// bytes.Buffer — zero value is ready to write
var buf bytes.Buffer
buf.WriteString("hello")

// Your own type
type StringSet map[string]struct{}

func (s *StringSet) Add(v string) {
    if *s == nil {
        *s = make(StringSet)
    }
    (*s)[v] = struct{}{}
}

var set StringSet  // nil — but Add works
set.Add("x")
```

## Small Interfaces

```go
// GOOD: small, composable
type Saver interface{ Save() error }
type Loader interface{ Load() error }

// AVOID: fat interfaces that are hard to implement/mock
type Storage interface {
    Save() error
    Load() error
    Delete() error
    List() ([]string, error)
    // ...
}
```

> "The bigger the interface, the weaker the abstraction." — Rob Pike

## Accept Interfaces, Return Concrete Types

```go
// Good: accepts interface (flexible for callers)
func process(r io.Reader) error { ... }

// Good: returns concrete type (callers can use all methods)
func newDB(dsn string) (*sql.DB, error) { ... }

// Avoid: returning interface (limits callers unnecessarily)
func newDB(dsn string) (Database, error) { ... }  // less flexible
```

## Resource Cleanup with `defer`

```go
// C#
using var stream = File.OpenRead(path);
// use stream

// Go
f, err := os.Open(path)
if err != nil { return err }
defer f.Close()
// use f
```

## Functional Options Pattern

For structs with many optional config fields:

```go
type Server struct {
    timeout  time.Duration
    maxConns int
}

type Option func(*Server)

func WithTimeout(d time.Duration) Option {
    return func(s *Server) { s.timeout = d }
}

func NewServer(opts ...Option) *Server {
    s := &Server{timeout: 30 * time.Second, maxConns: 100}  // defaults
    for _, o := range opts {
        o(s)
    }
    return s
}

// usage
s := NewServer(WithTimeout(5*time.Second))
```

## Table-Driven Configuration

```go
var handlers = map[string]http.HandlerFunc{
    "/":       handleHome,
    "/health": handleHealth,
}
```

## Embedding for Composition

```go
type Logger struct{ prefix string }
func (l *Logger) Log(msg string) { fmt.Println(l.prefix, msg) }

type Service struct {
    Logger        // embedded — Service gets Log() promoted
    db *sql.DB
}

s := Service{Logger: Logger{prefix: "[svc]"}}
s.Log("started")   // promoted method
```

## Error Wrapping Convention

```go
// Add context to errors as they propagate up
func loadUser(id int) (*User, error) {
    u, err := db.QueryUser(id)
    if err != nil {
        return nil, fmt.Errorf("loadUser(%d): %w", id, err)
    }
    return u, nil
}
```

## Go Proverbs (heuristics)

- "Don't communicate by sharing memory; share memory by communicating."
- "Errors are values."
- "The bigger the interface, the weaker the abstraction."
- "Make the zero value useful."
- "A little copying is better than a little dependency."
- "Clear is better than clever."
- "Gofmt's style is no one's favorite, yet gofmt is everyone's favorite."

## Useful Links
- [Effective Go](https://go.dev/doc/effective_go)
- [Code Review Comments](https://go.dev/wiki/CodeReviewComments)
- [Share Memory By Communicating](https://go.dev/blog/codelab-share)
- [Go Proverbs](https://go-proverbs.github.io/)
