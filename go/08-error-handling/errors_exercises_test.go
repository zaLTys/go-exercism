package learn

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrorsParseIntWithContext(t *testing.T) {
	_, err := errorsParseIntWithContext("nope")
	if err == nil {
		t.Fatalf("expected error")
	}
	if got := err.Error(); !containsAll(got, "nope") {
		t.Fatalf("error message %q should mention input", got)
	}
}

func TestErrorsValidateAge(t *testing.T) {
	if err := errorsValidateAge(30); err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if err := errorsValidateAge(-1); err == nil || err.Error() != "age: must be >= 0" {
		t.Fatalf("got %v", err)
	}
	if err := errorsValidateAge(151); err == nil || err.Error() != "age: must be <= 150" {
		t.Fatalf("got %v", err)
	}
}

func TestErrorsIsBadInput(t *testing.T) {
	wrapped := fmt.Errorf("oops: %w", errBadInput)
	if !errorsIsBadInput(wrapped) {
		t.Fatalf("expected true")
	}
	if errorsIsBadInput(errors.New("other")) {
		t.Fatalf("expected false")
	}
}

func TestErrorsJoinAll(t *testing.T) {
	if err := errorsJoinAll(nil, nil); err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
	e1 := errors.New("e1")
	e2 := errors.New("e2")
	err := errorsJoinAll(e1, nil, e2)
	if err == nil {
		t.Fatalf("expected non-nil")
	}
	if !errors.Is(err, e1) || !errors.Is(err, e2) {
		t.Fatalf("joined error should match components via errors.Is")
	}
}

// containsAll checks that s contains all given substrings.
func containsAll(s string, parts ...string) bool {
	for _, p := range parts {
		if !strContains(s, p) {
			return false
		}
	}
	return true
}

func strContains(s, sub string) bool {
	if len(sub) == 0 {
		return true
	}
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
