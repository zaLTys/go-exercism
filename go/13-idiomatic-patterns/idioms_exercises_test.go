package learn

import (
	"io"
	"strings"
	"testing"
)

func TestIdiomsStringSetZeroValueUsable(t *testing.T) {
	var s idiomsStringSet // nil map
	if s.Len() != 0 || s.Has("x") {
		t.Fatalf("unexpected initial state")
	}
	s.Add("x")
	s.Add("y")
	if !s.Has("x") || s.Len() != 2 {
		t.Fatalf("set not behaving as expected: %#v", s)
	}
}

func TestIdiomsReadAllAndClose(t *testing.T) {
	rc := io.NopCloser(strings.NewReader("hi"))
	b, err := idiomsReadAllAndClose(rc)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if string(b) != "hi" {
		t.Fatalf("got %q; want %q", string(b), "hi")
	}
}

func TestIdiomsPreferConcreteReturn(t *testing.T) {
	in := strings.NewReader("a\nb\n")
	got, err := idiomsPreferConcreteReturn(in)
	if err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if len(got) != 2 || got[0] != "a" || got[1] != "b" {
		t.Fatalf("got %#v; want [\"a\",\"b\"]", got)
	}
}
