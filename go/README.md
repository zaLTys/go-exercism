# Go Learning Curriculum for .NET Developers

A structured, hands-on Go learning path designed for experienced .NET/C# developers.
Each topic includes a quick-read theory reference and concrete exercises with unit tests.

## How to Use

1. **Read the theory** — open `THEORY.md` in the topic folder for a concise reference.
2. **Run the tests** — `cd` into the topic folder and run `go test ./...` (they will all fail at first).
3. **Implement the stubs** — fill in the `// TODO: implement` bodies in the `.go` files.
4. **Make the tests pass** — run `go test ./...` again; fix until all tests are green.
5. **Format your code** — run `go fmt ./...` before moving on.

## Curriculum

| # | Topic | Est. Time | Key Concepts |
|---|-------|-----------|--------------|
| [01-setup](./01-setup/) | Setup & Workflow | 1–2h | `go mod init`, `go run`, `go test`, module paths |
| [02-basic-syntax](./02-basic-syntax/) | Basic Syntax | 2–3h | `package`, `import`, `:=`, `var`, `const`, `gofmt` |
| [03-types](./03-types/) | Types | 4–6h | Structs, slices, maps, zero values, pointers |
| [04-control-flow](./04-control-flow/) | Control Flow | 2–3h | `if`, `for`, `switch`, `range`, `defer` |
| [05-functions](./05-functions/) | Functions | 2–4h | Multiple returns, variadics, closures, error returns |
| [06-methods](./06-methods/) | Methods | 2–3h | Receivers, pointer vs value, zero-value usability |
| [07-interfaces](./07-interfaces/) | Interfaces | 3–5h | Implicit satisfaction, type assertions, type switches |
| [08-error-handling](./08-error-handling/) | Error Handling | 3–5h | `error`, wrapping (`%w`), `errors.Is/As/Join`, `panic`/`recover` |
| [09-concurrency](./09-concurrency/) | Concurrency | 6–10h | Goroutines, channels, `select`, `context`, race detector |
| [10-packages-modules](./10-packages-modules/) | Packages & Modules | 3–5h | Import paths, `go.mod`, `internal/`, semantic versioning |
| [11-testing](./11-testing/) | Testing | 3–5h | `testing` package, table-driven tests, subtests, `go test` flags |
| [12-tooling](./12-tooling/) | Tooling | 2–4h | `go fmt`, `go vet`, `go doc`, `go fix`, `-race` |
| [13-idiomatic-patterns](./13-idiomatic-patterns/) | Idiomatic Patterns | 4–8h | Zero values, small interfaces, composition, Go proverbs |
| [14-capstone](./14-capstone/) | Capstone Project | — | Worker pool, context cancellation, error aggregation |

**Total estimated time:** 25–40 hours (~2–3 weeks at 1–2 hours/day)

## Quick Start

```bash
# Enter a topic
cd 01-setup

# Run tests (will fail — that's expected)
go test ./...

# Implement the stubs in setup_exercises.go, then re-run
go test ./...

# Format your code
go fmt ./...
```

## Learning Schedule (from the PDF plan)

```
Week 1 — Foundations
  Day 1: Setup (01)
  Day 2-3: Basic Syntax + Types (02, 03)
  Day 4-5: Control Flow + Functions (04, 05)

Week 2 — Core Go Model
  Day 1: Methods (06)
  Day 2-3: Interfaces (07)
  Day 4-5: Error Handling (08)

Week 3 — Concurrency + Ecosystem + Capstone
  Day 1-3: Concurrency (09)
  Day 4: Packages & Modules (10)
  Day 5-6: Testing + Tooling + Idiomatic Patterns (11, 12, 13)
  Day 7: Capstone (14)
```

## Key .NET → Go Mental Model Shifts

| .NET / C# | Go |
|-----------|-----|
| Exceptions (`try/catch`) | Errors as values (`value, err`) |
| Classes + inheritance | Structs + composition + embedding |
| Explicit interface implementation (`class X : IFoo`) | Implicit interface satisfaction |
| `Task` / `async/await` / `CancellationToken` | Goroutines + channels + `context.Context` |
| NuGet + `.csproj` | `go.mod` + `go get` |
| `dotnet test` with xUnit attributes | `go test` with `TestXxx` functions |

## Resources

- [A Tour of Go](https://go.dev/tour/) — interactive exercises
- [Effective Go](https://go.dev/doc/effective_go) — idioms and style
- [Go by Example](https://gobyexample.com/) — quick annotated examples
- [Go Blog](https://go.dev/blog/) — authoritative articles
- [Go Spec](https://go.dev/ref/spec) — language reference
- [pkg.go.dev](https://pkg.go.dev/) — standard library and package docs
