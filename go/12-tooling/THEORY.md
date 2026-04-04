# Tooling — Theory Reference

## The Big Idea

Go ships with a **cohesive, opinionated toolchain**. Formatting is not optional — `gofmt` is the cultural standard and CI pipelines check it. Static analysis (`go vet`) is built in. There's no separate linter framework to configure.

## Core Commands

| Command | Purpose | .NET equivalent |
|---------|---------|----------------|
| `go build ./...` | Compile | `dotnet build` |
| `go test ./...` | Run tests | `dotnet test` |
| `go test -race ./...` | Test with race detector | (no direct equivalent) |
| `go fmt ./...` | Format all code | `dotnet format` |
| `go vet ./...` | Static analysis / lint | Roslyn analyzers |
| `go doc fmt.Println` | Show docs for a symbol | (IDE docs) |
| `go mod tidy` | Sync deps | `dotnet restore` |
| `go get example.com/dep@v1` | Add/upgrade dep | `dotnet add package` |
| `go fix ./...` | Modernize deprecated APIs | (no direct equivalent) |

## `gofmt` / `go fmt`

- **All Go code uses `gofmt` format** — tabs for indentation, specific brace placement.
- `go fmt ./...` formats all `.go` files in the module.
- CI commonly runs `gofmt -l .` and fails if any file is not formatted.
- Configure your editor to run `gofmt` on save.

```bash
gofmt -w file.go          # format in-place
gofmt -l .                # list files that need formatting
go fmt ./...              # format all packages
```

## `go vet`

Catches suspicious code patterns the compiler doesn't:
- Mismatched `Printf` format verbs: `fmt.Printf("%d", "string")`
- Unreachable code
- Incorrect use of `sync/atomic`
- Struct tag errors

```bash
go vet ./...
```

## `go test` Flags

```bash
go test ./...                    # all packages
go test -v ./...                 # verbose: print all test names
go test -run TestFoo ./...       # filter by regex
go test -race ./...              # enable race detector
go test -cover ./...             # coverage percentage
go test -coverprofile=c.out ./.. # coverage file
go tool cover -html=c.out        # open coverage in browser
go test -bench=. -benchmem ./... # benchmarks + memory allocations
go test -count=1 ./...           # force re-run (no cache)
```

## `go doc`

```bash
go doc fmt.Println        # print doc for Println in fmt
go doc -all fmt           # all docs in the fmt package
go doc -src fmt.Println   # show source code
```

## Race Detector

```bash
go test -race ./...
go run -race main.go
go build -race -o app
```

The race detector instruments memory accesses and reports data races at runtime. Always use it in CI. It has ~2-20x slowdown in CPU time.

## `go/format` Package

Programmatic access to `gofmt` formatting from Go code:

```go
import "go/format"

formatted, err := format.Source([]byte(src))
```

## Editor Integration

- **VS Code**: Install the Go extension — it runs `gofmt`, `go vet`, `gopls` automatically.
- **GoLand**: Built-in Go support with all tools integrated.
- **Neovim**: `gopls` language server + conform.nvim for formatting.

## Use Cases

- **CI pipeline** — a minimal Go CI check:
  ```bash
  go fmt ./... && git diff --exit-code   # fail if any file needs formatting
  go vet ./...                            # fail on static analysis issues
  go test -race ./...                     # fail on test failures or races
  ```
- **Pre-commit hook**: run `go fmt` + `go vet` on staged files before each commit
- **Coverage gate**: `go test -coverprofile=c.out ./... && go tool cover -func=c.out | grep total` — check total coverage percentage in CI

## Common Mistakes & Caveats

- **`go vet` is not a full linter** — it only checks for a small set of well-defined bugs; consider adding `staticcheck` or `golangci-lint` to catch more issues (unused parameters, shadowed errors, style violations)
- **The race detector doesn't catch all races** — it only detects races that actually execute during the test run; 100% test coverage with `-race` still doesn't guarantee race-freedom of all code paths
- **Don't skip `go fmt` thinking "the IDE formats it"** — CI should still enforce it; different team members may have different editor configs or formatters
- **`go build` succeeds even with warnings** — Go has no warnings, only errors; use `go vet` to catch the things the compiler intentionally doesn't block
- **`go fix` is mostly safe but always review** — it rewrites code to use new APIs, but review the diff before committing, especially for large codebases

## Useful Links
- [gofmt blog](https://go.dev/blog/gofmt)
- [Command documentation](https://go.dev/doc/cmd)
- [go fix blog](https://go.dev/blog/gofix)
- [Race detector](https://go.dev/doc/articles/race_detector)
- [go/format package](https://pkg.go.dev/go/format)
