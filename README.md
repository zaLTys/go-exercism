# Go Learning Path
A hands-on Go curriculum built for developers who want to get productive in Go fast. Structured around the mental model shifts that matter most — explicit errors, implicit interfaces, goroutines, and the Go toolchain.

## Structure

```
go/
├── 01-setup/                 Install, modules, go run / go test
├── 02-basic-syntax/          Declarations, imports, gofmt
├── 03-types/                 Structs, slices, maps, zero values, pointers
├── 04-control-flow/          if / for / switch / range / defer
├── 05-functions/             Multiple returns, variadics, closures
├── 06-methods/               Receivers, pointer vs value, method sets
├── 07-interfaces/            Implicit satisfaction, type assertions
├── 08-error-handling/        errors.Is/As/Join, wrapping, panic/recover
├── 09-concurrency/           Goroutines, channels, select, context
├── 10-packages-modules/      go.mod, import paths, internal/, versioning
├── 11-testing/               Table-driven tests, subtests, go test flags
├── 12-tooling/               go fmt, go vet, go doc, race detector
├── 13-idiomatic-patterns/    Zero values, small interfaces, Go proverbs
└── 14-capstone/              Worker pool integrating all topics
```

Each topic folder contains:

| File | Purpose |
|------|---------|
| `THEORY.md` | Concise reference: concept, .NET parallel, syntax, use cases, caveats |
| `*_exercises.go` | Stub implementations — fill in the `// TODO` bodies |
| `*_exercises_test.go` | Tests that drive the exercises (read-only) |
| `*_gotchas.go` | Exercises targeting common mistakes and surprising behaviours |
| `*_gotchas_test.go` | Tests for the gotcha exercises |

## How to Work Through It

```bash
cd go/01-setup

# Tests fail by design — stubs return zero values
go test ./...

# Open *_exercises.go, read the THEORY.md, implement the TODOs
go test ./...   # make them green

# Then do the gotcha exercises in *_gotchas.go
go test ./...

# Format before moving on
go fmt ./...
```

## Topics at a Glance

| # | Topic | Est. time | The key .NET → Go shift |
|---|-------|-----------|------------------------|
| 01 | Setup | 1–2h | `go mod init` replaces `.csproj`; `go test` replaces `dotnet test` |
| 02 | Basic Syntax | 2–3h | No semicolons, no class boilerplate, `gofmt` is non-negotiable |
| 03 | Types | 4–6h | Slices/maps are not classes; zero values do real work; `append` returns a new slice |
| 04 | Control Flow | 2–3h | `for` is the only loop; `defer` replaces `using`/`try/finally` |
| 05 | Functions | 2–4h | Multiple returns replace `out` params and exceptions; closures capture by reference |
| 06 | Methods | 2–3h | No classes — methods are functions with a receiver; pointer vs value matters for interfaces |
| 07 | Interfaces | 3–5h | Satisfied implicitly (no `implements`); nil interface ≠ nil concrete value |
| 08 | Error Handling | 3–5h | Errors are values; `%w` wraps; `panic` is not `throw` |
| 09 | Concurrency | 6–10h | Goroutines + channels + `context.Context`; goroutine leaks are real |
| 10 | Packages & Modules | 3–5h | Import path = identity; circular imports are a build error |
| 11 | Testing | 3–5h | `TestXxx` functions, table-driven style, `t.Helper()` matters |
| 12 | Tooling | 2–4h | `go vet` + `go fmt` are part of the culture, not optional |
| 13 | Idiomatic Patterns | 4–8h | Small interfaces, useful zero values, functional options, Go proverbs |
| 14 | Capstone | — | Worker pool with cancellation, error aggregation, table-driven tests |

**Total: ~25–40 hours** (~2–3 weeks at 1–2 hours/day)

## Resources

- [A Tour of Go](https://go.dev/tour/) — interactive, official starting point
- [Effective Go](https://go.dev/doc/effective_go) — idiomatic style guide
- [Go by Example](https://gobyexample.com/) — annotated code snippets
- [pkg.go.dev](https://pkg.go.dev/) — standard library docs
- [Go Blog](https://go.dev/blog/) — deep dives on errors, concurrency, modules
