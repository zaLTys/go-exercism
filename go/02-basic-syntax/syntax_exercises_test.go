package learn

import "testing"

func TestSyntaxDefaultString(t *testing.T) {
	if got := syntaxDefaultString(" ", "fallback"); got != "fallback" {
		t.Fatalf("got %q; want %q", got, "fallback")
	}
	if got := syntaxDefaultString(" hi ", "fallback"); got != "hi" {
		t.Fatalf("got %q; want %q", got, "hi")
	}
}

func TestSyntaxJoinWithComma(t *testing.T) {
	if got := syntaxJoinWithComma(nil); got != "" {
		t.Fatalf("got %q; want empty", got)
	}
	if got := syntaxJoinWithComma([]string{" a", "b ", " c "}); got != "a, b, c" {
		t.Fatalf("got %q; want %q", got, "a, b, c")
	}
}

func TestSyntaxHasPrefixInsensitive(t *testing.T) {
	if !syntaxHasPrefixInsensitive("Golang", "go") {
		t.Fatalf("expected true")
	}
	if syntaxHasPrefixInsensitive("Golang", "lang") {
		t.Fatalf("expected false")
	}
}
