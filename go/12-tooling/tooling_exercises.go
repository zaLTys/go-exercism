package learn

import "go/format"

// EXERCISE 1: Implement toolingFormatGoSource.
// Return the gofmt-formatted version of src using the go/format package.
func toolingFormatGoSource(src string) (string, error) {
	// TODO: implement (hint: format.Source takes []byte, returns []byte)
	_, _ = format.Source([]byte("package p"))
	return "", nil
}

// EXERCISE 2: Implement toolingLooksGofmted.
// Return true if formatting src with toolingFormatGoSource produces identical output.
func toolingLooksGofmted(src string) bool {
	// TODO: implement
	return false
}

// EXERCISE 3: Implement toolingCountTestFunctions.
// Count occurrences of "func TestXxx(" in a Go source string.
// Simple string scanning is fine (no need to parse the AST).
func toolingCountTestFunctions(src string) int {
	// TODO: implement (hint: strings.Index in a loop, or strings.Count)
	return 0
}
