package learn

import (
	"errors"
	"fmt"
	"testing"
)

func TestFlowGotchaDeferValue(t *testing.T) {
	var captured string
	result := flowGotchaDeferValue([]int{1, 2, 3}, &captured)

	if result == "" {
		t.Fatalf("result is empty — did you build the string?")
	}
	// ! The key assertion: a closure-based defer sees the FINAL value of result.
	if captured != result {
		t.Fatalf("captured=%q but result=%q — defer closure should capture the final value", captured, result)
	}
}

func TestFlowGotchaNamedReturn(t *testing.T) {
	// Positive: n*n returned normally
	if got := flowGotchaNamedReturn(4); got != 16 {
		t.Fatalf("flowGotchaNamedReturn(4)=%d; want 16", got)
	}
	// ! n is negative: n*n would be positive (e.g. (-3)^2=9), but the
	// ! defer should check the named result and zero it out.
	// This demonstrates that defer can override what "return" sets.
	if got := flowGotchaNamedReturn(-3); got != 0 {
		t.Fatalf("flowGotchaNamedReturn(-3)=%d; want 0 — defer should have zeroed the result", got)
	}
	if got := flowGotchaNamedReturn(0); got != 0 {
		t.Fatalf("flowGotchaNamedReturn(0)=%d; want 0", got)
	}
}

func TestFlowGotchaProcessEach(t *testing.T) {
	called := []string{}
	process := func(s string) error {
		called = append(called, s)
		if s == "bad" {
			return fmt.Errorf("item %q failed", s)
		}
		return nil
	}

	// All succeed — no error
	err := flowGotchaProcessEach([]string{"a", "b", "c"}, process)
	if err != nil {
		t.Fatalf("expected nil; got %v", err)
	}
	if len(called) != 3 {
		t.Fatalf("process called %d times; want 3", len(called))
	}

	// One fails — should still process others, return error
	called = called[:0]
	err = flowGotchaProcessEach([]string{"a", "bad", "c"}, process)
	if err == nil {
		t.Fatalf("expected error for 'bad' item")
	}
	// ! Must have processed ALL items, not stopped at the first error
	if len(called) != 3 {
		t.Fatalf("process called %d times; want 3 — must not stop on first error", len(called))
	}

	// Multiple failures — all errors should be present
	err = flowGotchaProcessEach([]string{"bad", "bad"}, func(string) error {
		return errors.New("fail")
	})
	if err == nil {
		t.Fatalf("expected combined error")
	}
}
