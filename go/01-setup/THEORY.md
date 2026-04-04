# Setup — Theory Reference

## The Big Idea

Go's workflow is intentionally simple: **install → create module → write code → run/test**. There is no solution file, no project file, no NuGet. Everything is driven by the `go` command.

## .NET → Go Command Mapping

| .NET / C#                        | Go equivalent                        |
|----------------------------------|--------------------------------------|
| `dotnet new console -n MyApp`    | `mkdir myapp && go mod init example.com/myapp` |
| `dotnet run`                     | `go run .`                           |
| `dotnet test`                    | `go test ./...`                      |
| `dotnet build`                   | `go build ./...`                     |
| `dotnet restore`                 | `go mod tidy`                        |
| `dotnet --info`                  | `go version`                         |

## Key Concepts

### Module (`go.mod`)
- Created with `go mod init <module-path>` — the module path is also the base for all import paths.
- The file pins the Go version and lists dependencies (like a `.csproj` / `packages.config`).

```
module example.com/myapp

go 1.22

require (
    example.com/somedep v1.2.3
)
```

### Package
- Every `.go` file starts with `package <name>`.
- Files in the same directory share a package.
- The `main` package with a `main()` function is the entry point for executables.

### Entry point vs library
```go
// executable
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

```go
// library package (no main)
package mylib

func Greet(name string) string { ... }
```

### Running and testing
```bash
go run .                  # run main package in current dir
go test ./...             # test all packages recursively
go test -v ./...          # verbose output
go test -run TestFoo ./.. # run only tests matching "TestFoo"
```

## Go Version String Format
`go version go1.22.3 darwin/arm64` — the version part is `go<major>.<minor>.<patch>`.

## Quick Tips for .NET Developers
- No `Main(string[] args)` ceremony — just `func main()` in `package main`.
- No class required — top-level functions and variables are fine.
- `fmt.Println` ≈ `Console.WriteLine`.
- `strings.TrimSpace` ≈ `string.Trim()`.

## Useful Links
- [Install Go](https://go.dev/doc/install)
- [Getting Started Tutorial](https://go.dev/doc/tutorial/getting-started)
- [Create a Module Tutorial](https://go.dev/doc/tutorial/create-module)
- [go.mod reference](https://go.dev/doc/modules/gomod-ref)
