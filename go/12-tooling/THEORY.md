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

## Useful Links
- [gofmt blog](https://go.dev/blog/gofmt)
- [Command documentation](https://go.dev/doc/cmd)
- [go fix blog](https://go.dev/blog/gofix)
- [Race detector](https://go.dev/doc/articles/race_detector)
- [go/format package](https://pkg.go.dev/go/format)
