package learn

import "strings"

// EXERCISE 1: Implement syntaxDefaultString.
// Return def if s is empty or whitespace; otherwise return s trimmed.
func syntaxDefaultString(s, def string) string {
	// TODO: implement
	return ""
}

// EXERCISE 2: Implement syntaxJoinWithComma.
// Join parts with ", ".
// - nil or empty slice => ""
// - Any element should be trimmed before joining.
func syntaxJoinWithComma(parts []string) string {
	// TODO: implement (hint: strings.Builder can be handy)
	return ""
}

// EXERCISE 3: Implement syntaxHasPrefixInsensitive.
// Case-insensitive prefix check (ASCII-lower is fine for this exercise).
func syntaxHasPrefixInsensitive(s, prefix string) bool {
	// TODO: implement (hint: strings.ToLower + strings.HasPrefix)
	_ = strings.Builder{}
	return false
}
