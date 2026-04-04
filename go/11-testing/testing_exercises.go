package learn

import "strings"

// EXERCISE 1: Implement testingNormalizeWhitespace.
// Convert any run of whitespace (spaces, tabs, newlines) to a single space,
// and trim leading/trailing whitespace.
// Example: "a\tb\nc" => "a b c"
func testingNormalizeWhitespace(s string) string {
	// TODO: implement (hint: strings.Fields splits on any whitespace,
	//       then strings.Join reassembles with single spaces)
	_ = strings.Builder{}
	return ""
}

// EXERCISE 2: Implement testingIsPalindromeASCII.
// Only consider ASCII letters and digits; ignore case and all other characters.
// "A man, a plan, a canal: Panama" => true
func testingIsPalindromeASCII(s string) bool {
	// TODO: implement (hint: filter to alphanumeric, lowercase, then check palindrome)
	return false
}

// EXERCISE 3: Implement testingSplitCSVSimple.
// Split by commas, trim spaces from each field, allow empty fields.
// "a, b,,c" => ["a","b","","c"]
func testingSplitCSVSimple(s string) []string {
	// TODO: implement (hint: strings.Split(s, ",") then strings.TrimSpace each element)
	return nil
}
