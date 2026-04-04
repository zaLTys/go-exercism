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
//
// Suggested approach (3 steps):
//  1. Build a cleaned string: loop over s, keep only letters/digits, lowercased.
//     Use: (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') to test ASCII alphanumeric.
//     Use: c + ('a' - 'A') to lowercase an uppercase ASCII letter.
//  2. Compare the cleaned string with its reverse:
//     reverse by building a second string reading the cleaned one backwards.
//  3. Return cleaned == reversed.
func testingIsPalindromeASCII(s string) bool {
	// TODO: implement
	return false
}

// EXERCISE 3: Implement testingSplitCSVSimple.
// Split by commas, trim spaces from each field, allow empty fields.
// "a, b,,c" => ["a","b","","c"]
func testingSplitCSVSimple(s string) []string {
	// TODO: implement (hint: strings.Split(s, ",") then strings.TrimSpace each element)
	return nil
}
