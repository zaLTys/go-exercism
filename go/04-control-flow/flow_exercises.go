package learn

// EXERCISE 1: Implement flowClassifyHTTPStatus.
// Classify an HTTP status code into a human-readable category using a switch statement:
//   1xx => "informational"
//   2xx => "success"
//   3xx => "redirect"
//   4xx => "client_error"
//   5xx => "server_error"
//   anything else => "unknown"
//
// Real-world context: every HTTP API framework needs a function like this —
// mapping status codes to categories for logging, metrics, and error handling.
// Hint: switch on code / 100 to get the leading digit.
func flowClassifyHTTPStatus(code int) string {
	// TODO: implement using switch
	return ""
}

// EXERCISE 2: Implement flowBuildGradeReport.
// Given a slice of exam scores (0-100), produce a comma-separated string of letter grades.
//   90-100 => "A", 80-89 => "B", 70-79 => "C", 60-69 => "D", below 60 => "F"
// Example: [95, 82, 67] => "A,B,D"
// Empty or nil slice => ""
//
// Real-world context: batch-processing records and building a summary report —
// a pattern used in grading systems, analytics pipelines, and CSV exports.
// Hint: you'll need the "strings" package for strings.Builder.
func flowBuildGradeReport(scores []int) string {
	// TODO: implement (hint: strings.Builder + for range + switch or if/else)
	return ""
}

// EXERCISE 3: Implement flowFirstMatch.
// Return the first string in items that starts with the given prefix, and true.
// If none match, return ("", false).
//
// Real-world context: searching a list of options, filenames, or config keys
// for the first entry matching a prefix — a common pattern in CLI tools and
// autocomplete systems.
// Hint: strings.HasPrefix(s, prefix).
func flowFirstMatch(items []string, prefix string) (string, bool) {
	// TODO: implement
	return "", false
}
