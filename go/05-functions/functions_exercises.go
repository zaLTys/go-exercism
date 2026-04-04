package learn

import "errors"

var errEmptyField = errors.New("required field is empty")

// EXERCISE 1: Implement funcsFormatAddress.
// Format a mailing address as "street, city zip".
// Return errEmptyField if any field is empty or whitespace-only.
//
// Real-world context: formatting addresses for shipping labels, invoices,
// or display in a UI. Validates required fields and returns a structured error.
// Hint: you'll need "strings" for TrimSpace and "fmt" for Sprintf.
func funcsFormatAddress(street, city, zip string) (string, error) {
	// TODO: implement
	return "", nil
}

// EXERCISE 2: Implement funcsConcatNonEmpty.
// Join all non-empty parts with a space separator.
// Empty or whitespace-only parts are skipped.
//   funcsConcatNonEmpty("hello", "", " ", "world") => "hello world"
//   funcsConcatNonEmpty() => ""
//
// Real-world context: building display names, breadcrumbs, or log messages
// from optional parts — skip blanks, join with a separator.
// Hint: you'll need the "strings" package.
func funcsConcatNonEmpty(parts ...string) string {
	// TODO: implement
	return ""
}

// EXERCISE 3: Implement funcsMakeIDGenerator closure.
// Returns a function that yields "prefix-1", "prefix-2", "prefix-3", ...
// on successive calls. Each call increments an internal counter.
//
// Real-world context: generating unique IDs for orders, requests, or
// log correlation tokens. The closure captures state without a struct.
// Hint: close over a counter variable; use "fmt" for Sprintf.
func funcsMakeIDGenerator(prefix string) func() string {
	// TODO: implement (hint: close over a counter variable)
	return func() string { return "" }
}
