# Error Handling — Theory Reference

## The Big Idea

In Go, **errors are values** — they are returned explicitly, not thrown. The caller is responsible for checking `err != nil`. This is normal control flow, not exceptional control flow. `panic`/`recover` exist but are reserved for truly unrecoverable situations.

## Basic Error Pattern

```go
// C#
try {
    var text = File.ReadAllText(path);
    return text;
} catch (IOException ex) {
    throw new InvalidOperationException($"read failed: {path}", ex);
}

// Go
b, err := os.ReadFile(path)
if err != nil {
    return "", fmt.Errorf("read failed %q: %w", path, err)
}
return string(b), nil
```

**`%w` wraps the error** so callers can unwrap with `errors.Is`/`errors.As`.

## Creating Errors

```go
// sentinel error (package-level variable)
var ErrNotFound = errors.New("not found")

// formatted error
err := fmt.Errorf("user %d not found", userID)

// wrapped error (preserves cause for Is/As inspection)
err := fmt.Errorf("load user: %w", ErrNotFound)
```

## Checking Errors

```go
// errors.Is — checks the error chain for a specific sentinel
if errors.Is(err, ErrNotFound) {
    // handle not found
}

// errors.As — checks the chain for a specific type
var fieldErr *FieldError
if errors.As(err, &fieldErr) {
    fmt.Println("field:", fieldErr.Field)
}
```

| C# | Go |
|----|-----|
| `catch (NotFoundException ex)` | `errors.As(err, &notFoundErr)` |
| `catch (Exception ex) when (ex is IOE)` | `errors.Is(err, ErrIO)` |
| `ex.InnerException` | `errors.Unwrap(err)` |
| `throw new Agg(ex1, ex2)` | `errors.Join(err1, err2)` |

## Custom Error Types

```go
type ValidationError struct {
    Field string
    Msg   string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

// usage
return &ValidationError{Field: "age", Msg: "must be >= 0"}

// inspection
var ve *ValidationError
if errors.As(err, &ve) {
    fmt.Println(ve.Field)
}
```

## Error Wrapping Chain

```go
var ErrBase = errors.New("base")

// wrapping
wrapped := fmt.Errorf("layer2: %w", fmt.Errorf("layer1: %w", ErrBase))

errors.Is(wrapped, ErrBase)  // true — traverses the chain
```

## `errors.Join` (Go 1.20+)

```go
err1 := errors.New("e1")
err2 := errors.New("e2")
combined := errors.Join(err1, err2)

errors.Is(combined, err1)  // true
errors.Is(combined, err2)  // true
```

## `defer`, `panic`, `recover`

```go
// panic — for truly unrecoverable situations (not for normal errors)
if len(config) == 0 {
    panic("config must not be empty")
}

// recover — must be called inside a deferred function
func safeRun(f func()) (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panicked: %v", r)
        }
    }()
    f()
    return nil
}
```

**Rule:** Don't use `panic` for normal error handling. Use it only for programmer errors (impossible states), then let `recover` at the boundary turn it into an error if needed.

## Use Cases

- **API layer wrapping DB errors**: `return fmt.Errorf("getUser(%d): %w", id, err)` — callers see context, and `errors.Is(err, sql.ErrNoRows)` still works through the chain
- **Validation with multiple errors**: collect field errors into a slice, then `errors.Join(errs...)` — callers get one error that contains all failures
- **HTTP handler translating errors**: check `errors.Is(err, ErrNotFound)` → 404, `errors.As(err, &validErr)` → 400, everything else → 500
- **Defer for guaranteed cleanup**: `defer rows.Close()` after a DB query — runs even if the processing loop returns an error

## Common Mistakes & Caveats

- **Never ignore errors** — `result, _ = riskyOp()` hides real bugs; if you truly don't care, document why with a comment
- **Don't use `fmt.Errorf` without `%w` if callers need to inspect the cause**:
  ```go
  return fmt.Errorf("failed: %v", err)   // ! %v loses the original — errors.Is returns false
  return fmt.Errorf("failed: %w", err)   // ✓ %w wraps — errors.Is traverses the chain
  ```
- **Avoid over-wrapping error chains** — each layer adds its own context; don't wrap the same error 5 times with the same message
- **Don't use `panic` for expected error conditions** — panic is for programmer mistakes (index out of bounds, impossible state), not for user input or IO errors. If a caller can reasonably recover, return an error.
- **Sentinel errors should be exported, unexported ones kept private**:
  ```go
  var ErrNotFound = errors.New("not found")   // exported — callers can check with errors.Is
  var errInternal = errors.New("internal")    // unexported — callers can't depend on it
  ```
- **Value vs pointer receiver on error types**: for small structs, a value receiver `func (e MyError) Error() string` is fine. Pointer receivers on errors can cause the nil-interface-wrapping-nil-pointer gotcha.

## Useful Links
- [Tour: Errors](https://go.dev/tour/methods/19)
- [Error handling and Go](https://go.dev/blog/error-handling-and-go)
- [Errors are values](https://go.dev/blog/errors-are-values)
- [errors package docs](https://pkg.go.dev/errors)
- [Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)
