package learn

import "testing"

func TestFuncsFormatAddress(t *testing.T) {
	t.Run("valid address", func(t *testing.T) {
		got, err := funcsFormatAddress("123 Main St", "Springfield", "62704")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := "123 Main St, Springfield 62704"
		if got != want {
			t.Fatalf("got %q; want %q", got, want)
		}
	})
	t.Run("empty street", func(t *testing.T) {
		_, err := funcsFormatAddress("", "Springfield", "62704")
		if err == nil {
			t.Fatalf("expected error for empty street")
		}
	})
	t.Run("whitespace city", func(t *testing.T) {
		_, err := funcsFormatAddress("123 Main St", "  ", "62704")
		if err == nil {
			t.Fatalf("expected error for whitespace-only city")
		}
	})
	t.Run("empty zip", func(t *testing.T) {
		_, err := funcsFormatAddress("123 Main St", "Springfield", "")
		if err == nil {
			t.Fatalf("expected error for empty zip")
		}
	})
}

func TestFuncsConcatNonEmpty(t *testing.T) {
	t.Run("no args", func(t *testing.T) {
		if got := funcsConcatNonEmpty(); got != "" {
			t.Fatalf("got %q; want empty", got)
		}
	})
	t.Run("mixed with blanks", func(t *testing.T) {
		got := funcsConcatNonEmpty("hello", "", "  ", "world")
		if got != "hello world" {
			t.Fatalf("got %q; want %q", got, "hello world")
		}
	})
	t.Run("single part", func(t *testing.T) {
		if got := funcsConcatNonEmpty("solo"); got != "solo" {
			t.Fatalf("got %q; want %q", got, "solo")
		}
	})
	t.Run("all blank", func(t *testing.T) {
		if got := funcsConcatNonEmpty("", " ", "  "); got != "" {
			t.Fatalf("got %q; want empty", got)
		}
	})
}

func TestFuncsMakeIDGenerator(t *testing.T) {
	next := funcsMakeIDGenerator("order")
	if got := next(); got != "order-1" {
		t.Fatalf("first call: got %q; want %q", got, "order-1")
	}
	if got := next(); got != "order-2" {
		t.Fatalf("second call: got %q; want %q", got, "order-2")
	}
	if got := next(); got != "order-3" {
		t.Fatalf("third call: got %q; want %q", got, "order-3")
	}

	other := funcsMakeIDGenerator("req")
	if got := other(); got != "req-1" {
		t.Fatalf("separate generator: got %q; want %q — each generator should have its own counter", got, "req-1")
	}
}
