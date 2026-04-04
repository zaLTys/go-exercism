package learn

import (
	"errors"
	"fmt"
)

// ============================================================
// GOTCHA EXERCISES — things that surprise developers from C#/.NET
// ============================================================

// GOTCHA 1: %v breaks the errors.Is chain; %w preserves it.
//
// ! fmt.Errorf("context: %v", err) creates a new error with the MESSAGE of err.
// ! The original err is LOST — errors.Is(wrapped, err) returns false.
//
// ! fmt.Errorf("context: %w", err) WRAPS the error — errors.Is traverses it.
//
// Implement errorsGotchaWrap: wrap err with the message "operation failed: <err>"
// using %w so that errors.Is still works on the result.
func errorsGotchaWrap(err error) error {
	// TODO: implement using fmt.Errorf with %w
	return fmt.Errorf("operation failed: %v", err) // ! BUG: %v breaks errors.Is; change to %w
}

// GOTCHA 2: errors.Is traverses the CHAIN; == does not.
//
// ! In C#, you might check: ex.InnerException == specificEx
// ! In Go, you should NOT use == to compare errors once wrapping is involved.
// ! errors.Is(err, target) walks the Unwrap chain and checks each level.
//
// Implement errorsGotchaIsWrapped: return true if err or any wrapped error
// in its chain IS the sentinel errGotchaSentinel.
// Must work even if err has been wrapped multiple times.
var errGotchaSentinel = errors.New("sentinel")

func errorsGotchaIsWrapped(err error) bool {
	// TODO: implement — do NOT use err == errGotchaSentinel; use errors.Is
	return err == errGotchaSentinel // ! BUG: fails for wrapped errors; fix it
}

// GOTCHA 3: Ignoring errors silently is a Go anti-pattern.
//
// ! In C# you can sometimes ignore exceptions with a bare try{} catch{}.
// ! In Go, the compiler doesn't force you to check errors — it's your discipline.
// ! The blank identifier _ explicitly discards; no identifier silently discards.
//
// Broken examples:
//   fmt.Fprintln(w, msg)       // return values (n, err) silently discarded
//   os.Remove(tmpFile)         // error discarded — file might not be deleted
//   result, _ = parse(input)   // _ is explicit but still risky without a comment
//
// Implement errorsGotchaParsePositive: parse s as an int and return it.
// Requirements:
//   - If parsing fails, return 0 and a wrapped error (include s in the message).
//   - If the number is <= 0, return 0 and a descriptive error.
//   - If both succeed, return the number and nil.
// Do NOT use _ to discard the parse error — handle it explicitly.
func errorsGotchaParsePositive(s string) (int, error) {
	// TODO: implement — handle the parse error explicitly, don't discard it
	return 0, nil
}

// GOTCHA 4: panic is NOT a substitute for error returns.
//
// ! Coming from C#, you might reach for panic like you'd throw an exception.
// ! In Go, panic is for PROGRAMMER errors (impossible states, violated invariants).
// ! For expected error conditions (bad input, IO failure), return an error.
//
// The rough rule:
//   panic  →  "this should never happen; if it does, it's a bug in my code"
//   error  →  "this might happen; the caller should handle it"
//
// Implement errorsGotchaDivide: return a/b.
// If b == 0, return an error (do NOT panic).
// ! A common mistake from exception-based languages: panic("divide by zero")
func errorsGotchaDivide(a, b int) (int, error) {
	// TODO: implement — return error for b==0, do not panic
	if b == 0 {
		panic("divide by zero") // ! BAD: use an error return instead; fix this
	}
	return a / b, nil
}
