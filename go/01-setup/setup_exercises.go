package learn

// EXERCISE 1: Implement setupHello.
// - If name is empty or whitespace => "Hello, world!"
// - Else => "Hello, <name>!"
func setupHello(name string) string {
	// TODO: implement
	return ""
}

// EXERCISE 2: Implement setupParseGoVersion.
// Parse "go version go1.22.3 darwin/arm64" or "go1.22.3".
// Return major=1, minor=22 for that example.
// Return ok=false if you can't parse.
func setupParseGoVersion(s string) (major int, minor int, ok bool) {
	// TODO: implement
	return 0, 0, false
}

// EXERCISE 3: Implement setupModulePathFromGoMod.
// From a go.mod file text, extract module path from the "module <path>" line.
func setupModulePathFromGoMod(goModText string) string {
	// TODO: implement
	return ""
}
