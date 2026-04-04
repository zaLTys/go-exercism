package learn

import (
	"context"
	"testing"
	"time"
)

func TestConcurrencySquareAll(t *testing.T) {
	got, err := concurrencySquareAll([]int{1, 2, 3, 4}, 2)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	want := []int{1, 4, 9, 16}
	if !equalIntSlices(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}
	_, err = concurrencySquareAll([]int{1}, 0)
	if err == nil {
		t.Fatalf("expected error for workers<=0")
	}
}

func TestConcurrencySelectFirstString(t *testing.T) {
	a := make(chan string, 1)
	b := make(chan string, 1)
	b <- "B"
	got := concurrencySelectFirstString(a, b)
	if got != "B" {
		t.Fatalf("got %q; want %q", got, "B")
	}
}

func TestConcurrencyCounterOwnedByGoroutine(t *testing.T) {
	inc, val, stop := concurrencyCounterOwnedByGoroutine()
	defer stop()
	if val() != 0 {
		t.Fatalf("want 0")
	}
	if inc() != 1 || inc() != 2 {
		t.Fatalf("counter not incrementing")
	}
	if val() != 2 {
		t.Fatalf("want 2")
	}
	stop() // safe to call multiple times
}

// withTimeout creates a context with deadline — useful for preventing test hangs.
func withTimeout(t *testing.T, d time.Duration) (context.Context, context.CancelFunc) {
	t.Helper()
	return context.WithTimeout(context.Background(), d)
}

func equalIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
