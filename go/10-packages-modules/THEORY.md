# Packages & Modules — Theory Reference

## The Big Idea

Go code is organized into **packages** (a directory of `.go` files) and **modules** (a versioned set of packages with a `go.mod`). The import path is a real identifier in the tooling — it's how the `go` command finds and downloads code.

## Package

- Every `.go` file starts with `package <name>`.
- All files in the same directory must have the same package name.
- Package name ≈ directory name (by convention).
- **Exported** identifiers start with an uppercase letter: `Println`, `Error`.
- **Unexported** identifiers start with lowercase: `helper`, `errInternal`.

```go
// mymath/math.go
package mymath

// Exported — visible outside the package
func Add(a, b int) int { return a + b }

// Unexported — internal only
func square(x int) int { return x * x }
```

## Module

- A module is a tree of packages rooted at a `go.mod` file.
- Created with: `go mod init <module-path>`
- Module path = base for all import paths within the module.

```
// go.mod
module example.com/myapp

go 1.22

require (
    github.com/some/dep v1.2.3
)
```

| C# | Go |
|----|-----|
| `.csproj` / `solution.sln` | `go.mod` |
| `<PackageReference Include="X" Version="1.0"/>` | `require X v1.0` in `go.mod` |
| `dotnet restore` | `go mod tidy` |
| `dotnet add package X` | `go get X` |
| Assembly `namespace` | Package (directory-based) |

## Import Paths

```go
import (
    "fmt"                          // stdlib — no module path prefix
    "strings"
    "example.com/myapp/mymath"     // local package — module path + subdir
    "github.com/user/repo/pkg"     // external dependency
)
```

The import alias is the package name (last element), not the full path:
```go
import "math/rand"
rand.Intn(100)  // NOT math/rand.Intn
```

Custom alias:
```go
import mrand "math/rand"
mrand.Intn(100)
```

## Recommended Module Layout

```
mymodule/
  go.mod
  go.sum              # auto-generated checksum file
  cmd/
    myapp/
      main.go         # main package for the executable
  internal/           # packages here cannot be imported by outside modules
    db/
      db.go
  api/
    api.go
  README.md
```

`internal/` is enforced by the Go toolchain — packages within `internal/` can only be imported by code in the parent of `internal/`.

## Semantic Import Versioning

```
v1: example.com/mymod
v2: example.com/mymod/v2   ← major version in path for v2+
v3: example.com/mymod/v3
```

- v0 and v1: no version suffix in the import path.
- v2+: must add `/v2`, `/v3`, etc. to the module path and all imports.

## Key Commands

```bash
go mod init example.com/myapp    # create go.mod
go mod tidy                       # add missing, remove unused deps
go get example.com/dep@v1.2.3    # add/update a dependency
go list ./...                     # list all packages in module
go build ./...                    # build everything
go test ./...                     # test everything
```

## Useful Links
- [How to Write Go Code](https://go.dev/doc/code)
- [Go Modules Reference](https://go.dev/ref/mod)
- [go.mod reference](https://go.dev/doc/modules/gomod-ref)
- [Organizing a Go module](https://go.dev/doc/modules/layout)
- [Using Go Modules](https://go.dev/blog/using-go-modules)
