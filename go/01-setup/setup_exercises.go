package learn

import (
	"fmt"
	"strings"
	"strconv"
)

// EXERCISE 1: Implement setupHello.
// - If name is empty or whitespace => "Hello, world!"
// - Else => "Hello, <name>!"
func setupHello(name string) string {
	if strings.TrimSpace(name) != "" {
		return fmt.Sprintf("Hello, %s!", name)
	}
	return "Hello, world!"
}

// EXERCISE 2: Implement setupParseGoVersion.
// Parse "go version go1.22.3 darwin/arm64" or "go1.22.3".
// Return major=1, minor=22 for that example.
// Return ok=false if you can't parse.
func setupParseGoVersion(s string) (major int, minor int, ok bool) {
	for _, f := range strings.Fields(s) {
		if !strings.HasPrefix(f, "go") || len(f) <= 2 {
			continue
		}
		rest := f[2:]
		if rest == "" || rest[0] < '0' || rest[0] > '9' {
			continue
		}
		parts := strings.SplitN(rest, ".", 3)
		if len(parts) < 2 {
			continue
		}
		maj, err1 := strconv.Atoi(parts[0])
		min, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			continue
		}
		return maj, min, true
	}
	return 0, 0, false
}

// EXERCISE 3: Implement setupModulePathFromGoMod.
// From a go.mod file text, extract module path from the "module <path>" line.
func setupModulePathFromGoMod(goModText string) string {
	// TODO: implement
	return ""
}
