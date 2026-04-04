package learn

// ============================================================
// GOTCHA EXERCISES — things that surprise developers from C#/.NET
// ============================================================

// GOTCHA 1: String indexing returns BYTES, not characters.
//
// ! In C#, s[i] gives a char (Unicode code point).
// ! In Go, s[i] gives a byte (uint8). For ASCII they're the same — but for
// ! any non-ASCII character (é, ñ, 中) this produces wrong results silently.
//
// Implement syntaxGotchaFirstRune: return the FIRST CHARACTER (rune) of s.
// Do NOT use s[0] — that returns a byte.
// Hint: a for range loop over a string yields (index, rune) pairs.
func syntaxGotchaFirstRune(s string) rune {
	// TODO: implement — return the first rune, not the first byte
	return 0
}

// GOTCHA 2: len(s) counts BYTES, not characters.
//
// ! len("café") == 5 (bytes), not 4 (characters).
// ! "é" is encoded as 2 bytes in UTF-8.
//
// Implement syntaxGotchaRuneCount: return the NUMBER OF CHARACTERS in s.
// Hint: convert to []rune and use len, or use utf8.RuneCountInString.
func syntaxGotchaRuneCount(s string) int {
	// TODO: implement — return character count, not byte count
	return 0
}

// GOTCHA 3: Short variable declaration := in inner scope creates a NEW variable,
// it does NOT update the outer one. This is "shadowing".
//
// ! The function below looks like it sets result to 42, but the inner :=
// ! creates a DIFFERENT variable. The outer result stays 0.
// ! This compiles without warning — Go's shadow check requires go vet or
// ! the -shadow flag.
//
// Implement syntaxGotchaNoShadow: return 42.
// Rules: use exactly ONE variable named result, declared with :=.
// You must NOT use an inner block — keep the assignment at the function level.
func syntaxGotchaNoShadow() int {
	// TODO: implement — just declare result := 42 and return it
	// (study the broken version below in the test file to see why shadowing bites)
	return 0
}
