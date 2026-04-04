package learn

import "testing"

// This file uses table-driven tests + subtests — the idiomatic Go testing style.
// Study the structure as much as the exercises themselves.

func TestTestingNormalizeWhitespace(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"already clean", "a b", "a b"},
		{"tabs/newlines", "a\tb\nc", "a b c"},
		{"leading/trailing", " a b ", "a b"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := testingNormalizeWhitespace(tc.in); got != tc.want {
				t.Fatalf("got %q; want %q", got, tc.want)
			}
		})
	}
}

func TestTestingIsPalindromeASCII(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"racecar", true},
		{"RaceCar", true},
		{"A man, a plan, a canal: Panama", true},
		{"not one", false},
	}
	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			if got := testingIsPalindromeASCII(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}

func TestTestingSplitCSVSimple(t *testing.T) {
	in := "a, b,,c"
	want := []string{"a", "b", "", "c"}
	got := testingSplitCSVSimple(in)
	if len(got) != len(want) {
		t.Fatalf("len=%d; want %d, got=%v", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("i=%d got %q want %q (got=%v)", i, got[i], want[i], got)
		}
	}
}
