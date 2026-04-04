# Go Curriculum — Exercise Reference

> For the full project overview see the [root README](../README.md).

## How to Use Each Topic

```bash
cd 01-setup

go test ./...          # all tests fail — that's the starting point

# 1. Read THEORY.md
# 2. Implement TODOs in *_exercises.go
go test ./...          # make the exercises pass

# 3. Implement TODOs in *_gotchas.go  (gotcha exercises)
go test ./...          # make the gotcha exercises pass

go fmt ./...           # format before moving on
```

## Curriculum

| # | Folder | Key Concepts | Files |
|---|--------|-------------|-------|
| 01 | [01-setup](./01-setup/) | `go mod init`, `go run`, `go test`, module paths | 2 exercise + 0 gotcha |
| 02 | [02-basic-syntax](./02-basic-syntax/) | Declarations, imports, `:=`, `gofmt`, string bytes vs runes | 2 exercise + 1 gotcha |
| 03 | [03-types](./03-types/) | Structs, slices, maps, zero values, pointers | 2 exercise + 1 gotcha |
| 04 | [04-control-flow](./04-control-flow/) | `if`, `for`, `switch`, `range`, `defer` | 2 exercise + 1 gotcha |
| 05 | [05-functions](./05-functions/) | Multiple returns, variadics, closures, error returns | 2 exercise + 1 gotcha |
| 06 | [06-methods](./06-methods/) | Receivers, pointer vs value, method sets | 2 exercise + 1 gotcha |
| 07 | [07-interfaces](./07-interfaces/) | Implicit satisfaction, type assertions, nil interface gotcha | 2 exercise + 1 gotcha |
| 08 | [08-error-handling](./08-error-handling/) | `error`, `%w` wrapping, `errors.Is/As/Join`, `panic` | 2 exercise + 1 gotcha |
| 09 | [09-concurrency](./09-concurrency/) | Goroutines, channels, `select`, `context`, actor pattern | 2 exercise + 0 gotcha |
| 10 | [10-packages-modules](./10-packages-modules/) | Import paths, `go.mod`, `internal/`, versioning | 2 exercise + 0 gotcha |
| 11 | [11-testing](./11-testing/) | Table-driven tests, subtests, `t.Helper`, benchmarks | 2 exercise + 0 gotcha |
| 12 | [12-tooling](./12-tooling/) | `go fmt`, `go vet`, `go doc`, `-race`, `go/format` | 2 exercise + 0 gotcha |
| 13 | [13-idiomatic-patterns](./13-idiomatic-patterns/) | Zero values, small interfaces, functional options | 2 exercise + 0 gotcha |
| 14 | [14-capstone](./14-capstone/) | Worker pool, context cancellation, error aggregation | capstone |

## What Each THEORY.md Contains

Every topic's `THEORY.md` is structured the same way:

1. **The Big Idea** — one paragraph on the core concept
2. **.NET parallel** — what you'd reach for in C# and why Go differs
3. **Syntax reference** — copy-paste-ready code blocks with `// !` markers on surprising lines
4. **Use Cases** — real scenarios where this feature applies
5. **Common Mistakes & Caveats** — the things that bite .NET developers specifically
6. **Useful Links** — official Go docs, blog posts, Tour sections

## Mental Model Shifts at a Glance

| .NET / C# | Go | Topic |
|-----------|-----|-------|
| `try/catch/throw` | `value, err` returns | 08 |
| `class` + inheritance | `struct` + composition | 03, 06 |
| `class X : IFoo` | implicit interface satisfaction | 07 |
| `Task` / `async/await` | goroutines + channels | 09 |
| `CancellationToken` | `context.Context` | 09 |
| `new List<T>()` | `[]T{}` or `append` | 03 |
| `Dictionary<K,V>` | `map[K]V` | 03 |
| `dotnet test` + xUnit | `go test` + `TestXxx` functions | 11 |
| `NuGet` + `.csproj` | `go.mod` + `go get` | 10 |
