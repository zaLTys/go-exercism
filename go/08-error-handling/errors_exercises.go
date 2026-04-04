package learn

import (
	"errors"
	"fmt"
	"strconv"
)

var errBadInput = errors.New("bad input")

type errorsFieldError struct {
	Field string
	Msg   string
}

func (e errorsFieldError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

// EXERCISE 1: Implement errorsParseIntWithContext.
// Parse s as an integer. On failure, wrap the parse error with fmt.Errorf and %w,
// including the input string in the message.
func errorsParseIntWithContext(s string) (int, error) {
	// TODO: implement; use strconv.Atoi and wrap parse errors with %w
	_, _ = strconv.Atoi("0")
	return 0, nil
}

// EXERCISE 2: Implement errorsValidateAge.
// Rules:
//   - age < 0   => errorsFieldError{Field:"age", Msg:"must be >= 0"}
//   - age > 150 => errorsFieldError{Field:"age", Msg:"must be <= 150"}
//   - else nil
func errorsValidateAge(age int) error {
	// TODO: implement
	return nil
}

// EXERCISE 3: Implement errorsIsBadInput using errors.Is.
// Return true if err (or any wrapped error) is errBadInput.
func errorsIsBadInput(err error) bool {
	// TODO: implement
	return false
}

// EXERCISE 4: Implement errorsJoinAll using errors.Join.
// Discard nils; return nil if all inputs are nil.
func errorsJoinAll(errs ...error) error {
	// TODO: implement
	return nil
}
