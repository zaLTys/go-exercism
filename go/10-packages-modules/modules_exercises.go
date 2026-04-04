package learn

import "fmt"

// EXERCISE 1: Implement modulesJoinImportPath.
// modulesJoinImportPath("example.com/m", "auth/token") => "example.com/m/auth/token"
// Handle empty subdir => return modulePath unchanged.
func modulesJoinImportPath(modulePath, subdir string) string {
	// TODO: implement
	return ""
}

// EXERCISE 2: Implement modulesWithMajorSuffix.
// If major >= 2, append "/v<major>" suffix.
// modulesWithMajorSuffix("example.com/m", 1) => "example.com/m"
// modulesWithMajorSuffix("example.com/m", 2) => "example.com/m/v2"
func modulesWithMajorSuffix(modulePath string, major int) string {
	// TODO: implement
	_ = fmt.Sprintf // hint: fmt.Sprintf is useful here
	return ""
}

// EXERCISE 3: Implement modulesParseRequireLine.
// Input:  "require example.com/dep v1.2.3"
// Output: mod="example.com/dep", ver="v1.2.3", ok=true
// Return ok=false if line doesn't match the expected format.
func modulesParseRequireLine(line string) (mod string, version string, ok bool) {
	// TODO: implement (hint: strings.Fields splits by whitespace)
	return "", "", false
}
