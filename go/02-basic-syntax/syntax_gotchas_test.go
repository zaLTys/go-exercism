package learn

import "testing"

func TestSyntaxGotchaFirstRune(t *testing.T) {
	// ! s[0] returns byte 0xc3 for "école", NOT rune 'é' (233).
	// Prove this to yourself: uncomment the line below and see what you get.
	// t.Logf("byte at index 0: %d (want rune 'é'=%d)", "école"[0], 'é')

	if got := syntaxGotchaFirstRune("école"); got != 'é' {
		t.Fatalf("syntaxGotchaFirstRune(\"école\")=%v (%d); want 'é' (%d)",
			string(got), got, 'é')
	}
	if got := syntaxGotchaFirstRune("hello"); got != 'h' {
		t.Fatalf("syntaxGotchaFirstRune(\"hello\")=%v; want 'h'", string(got))
	}
}

func TestSyntaxGotchaRuneCount(t *testing.T) {
	// ! len("café") == 5 bytes, but there are only 4 characters.
	// Prove it: uncomment below and run with go test -v
	// t.Logf("len(\"café\") = %d bytes", len("café"))

	if got := syntaxGotchaRuneCount("café"); got != 4 {
		t.Fatalf("syntaxGotchaRuneCount(\"café\")=%d; want 4 characters (not bytes)", got)
	}
	if got := syntaxGotchaRuneCount("hello"); got != 5 {
		t.Fatalf("syntaxGotchaRuneCount(\"hello\")=%d; want 5", got)
	}
}

func TestSyntaxGotchaNoShadow(t *testing.T) {
	// Study this broken version to understand shadowing:
	//
	//   func broken() int {
	//       result := 0           // outer result
	//       if true {
	//           result := 42      // ! NEW variable — outer result unchanged!
	//           _ = result
	//       }
	//       return result         // still 0, not 42
	//   }
	//
	// Fix: use = instead of := inside the block, OR don't use a block at all.

	if got := syntaxGotchaNoShadow(); got != 42 {
		t.Fatalf("got %d; want 42 — check for variable shadowing", got)
	}
}
