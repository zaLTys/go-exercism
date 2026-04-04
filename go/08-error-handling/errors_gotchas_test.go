package learn

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrorsGotchaWrap(t *testing.T) {
	sentinel := errors.New("root cause")
	wrapped := errorsGotchaWrap(sentinel)

	if wrapped == nil {
		t.Fatalf("expected non-nil error")
	}
	// ! If %v was used instead of %w, this will fail.
	if !errors.Is(wrapped, sentinel) {
		t.Fatalf("errors.Is failed — %%w was not used for wrapping.\n"+
			"  Got message: %q\n"+
			"  HINT: fmt.Errorf(\"...: %%%%v\", err) loses the original; use %%%%w", wrapped.Error())
	}
	// Should also contain the context message
	if wrapped.Error() == sentinel.Error() {
		t.Fatalf("error was not wrapped — message should have context prefix")
	}
}

func TestErrorsGotchaIsWrapped(t *testing.T) {
	// Direct match
	if !errorsGotchaIsWrapped(errGotchaSentinel) {
		t.Fatalf("direct sentinel: want true")
	}

	// ! == comparison fails here; errors.Is succeeds
	onceWrapped := fmt.Errorf("layer1: %w", errGotchaSentinel)
	if !errorsGotchaIsWrapped(onceWrapped) {
		t.Fatalf("once-wrapped: want true — errors.Is should traverse the chain")
	}

	twiceWrapped := fmt.Errorf("layer2: %w", onceWrapped)
	if !errorsGotchaIsWrapped(twiceWrapped) {
		t.Fatalf("twice-wrapped: want true")
	}

	// Unrelated error — should be false
	if errorsGotchaIsWrapped(errors.New("other")) {
		t.Fatalf("unrelated error: want false")
	}
}

func TestErrorsGotchaParsePositive(t *testing.T) {
	n, err := errorsGotchaParsePositive("42")
	if err != nil || n != 42 {
		t.Fatalf("\"42\": got (%d, %v); want (42, nil)", n, err)
	}

	// Invalid input — must return wrapped error (not panic, not silent 0)
	_, err = errorsGotchaParsePositive("abc")
	if err == nil {
		t.Fatalf("\"abc\": expected error, got nil — don't discard parse errors")
	}
	if !containsAll(err.Error(), "abc") {
		t.Fatalf("error should mention the input %q; got: %v", "abc", err)
	}

	// Zero/negative — also an error
	_, err = errorsGotchaParsePositive("0")
	if err == nil {
		t.Fatalf("\"0\": expected error for non-positive number")
	}
	_, err = errorsGotchaParsePositive("-5")
	if err == nil {
		t.Fatalf("\"-5\": expected error for non-positive number")
	}
}

func TestErrorsGotchaDivide(t *testing.T) {
	n, err := errorsGotchaDivide(10, 2)
	if err != nil || n != 5 {
		t.Fatalf("10/2: got (%d, %v); want (5, nil)", n, err)
	}

	// ! Must return error, not panic
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("divide by zero caused a panic: %v — return an error instead of panicking", r)
			}
		}()
		_, err = errorsGotchaDivide(10, 0)
	}()
	if err == nil {
		t.Fatalf("divide by zero: expected error, got nil")
	}
}
