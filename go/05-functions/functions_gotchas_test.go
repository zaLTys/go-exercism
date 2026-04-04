package learn

import (
	"errors"
	"fmt"
	"testing"
)

func TestFuncsGotchaAdders(t *testing.T) {
	adders := funcsGotchaAdders([]int{1, 2, 3})
	if len(adders) != 3 {
		t.Fatalf("expected 3 adders, got %d", len(adders))
	}
	// ! Each adder must use its OWN captured value, not the last one (3).
	for i, add := range adders {
		want := (i + 1) + 10 // adders[0](10)=11, adders[1](10)=12, adders[2](10)=13
		if got := add(10); got != want {
			t.Fatalf("adders[%d](10)=%d; want %d — closure captured wrong value of n", i, got, want)
		}
	}
}

func TestFuncsGotchaSafeCall(t *testing.T) {
	// ! Calling a nil function panics — the guard prevents this.
	if got := funcsGotchaSafeCall(nil, 5); got != 0 {
		t.Fatalf("nil fn: got %d; want 0", got)
	}
	double := func(x int) int { return x * 2 }
	if got := funcsGotchaSafeCall(double, 5); got != 10 {
		t.Fatalf("double(5): got %d; want 10", got)
	}
}

func TestFuncsGotchaWrapError(t *testing.T) {
	sentinel := errors.New("original")

	// Success case — no wrapping
	err := funcsGotchaWrapError(func() error { return nil })
	if err != nil {
		t.Fatalf("expected nil; got %v", err)
	}

	// Error case — defer must wrap the error
	err = funcsGotchaWrapError(func() error { return sentinel })
	if err == nil {
		t.Fatalf("expected non-nil error")
	}
	// ! The caller gets a WRAPPED error, but errors.Is still finds the original.
	if !errors.Is(err, sentinel) {
		t.Fatalf("errors.Is failed — wrapping broke the chain; got: %v", err)
	}
	// The message should contain the wrapper context
	if err.Error() == sentinel.Error() {
		t.Fatalf("error was not wrapped — should have extra context prefix; got: %v", err)
	}
	fmt.Printf("  [info] wrapped error message: %q\n", err.Error())
}
