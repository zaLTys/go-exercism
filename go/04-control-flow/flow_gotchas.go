package learn

import "errors"

// ============================================================
// GOTCHA EXERCISES — things that surprise developers from C#/.NET
// ============================================================

// GOTCHA 1: defer arguments are evaluated IMMEDIATELY, not when the defer runs.
//
// ! In C# try/finally, the finally block runs with the current variable values.
// ! In Go, defer evaluates its ARGUMENTS at the defer site, not at return time.
//
// This function demonstrates the timing. Read it carefully:
//
//   func deferTiming() string {
//       msg := "original"
//       defer log(msg)     // ! "original" is captured HERE, not at return
//       msg = "changed"
//       return msg
//   }
//
// The defer will log "original" even though msg was changed to "changed".
// Exception: if a deferred CLOSURE (func(){...}) is used, it captures the
// variable by reference and sees the final value.
//
// Implement flowGotchaDeferValue: uses a deferred function (closure) to capture
// the FINAL value of result, and returns what the defer captured via *captured.
// After the function returns, *captured should equal the final return value.
func flowGotchaDeferValue(values []int, captured *string) string {
	// TODO: implement:
	// 1. Build result by joining values as "1,2,3" (use fmt.Sprintf or strconv)
	// 2. defer a CLOSURE that sets *captured = result
	// 3. return result
	//
	// ! If you use defer fmt.Sprintf(...) directly, it captures at call time.
	// ! A closure func() { *captured = result } captures the variable,
	// ! so it always sees the final value of result.
	return ""
}

// GOTCHA 2: Named return values can be modified by deferred functions.
//
// ! This is powerful but easy to misread. The defer runs AFTER "return expr"
// ! sets the named return, so the defer can still change what actually comes back.
//
//   func doubleOnError() (result int, err error) {
//       defer func() {
//           if err != nil { result = -1 }  // ! overwrites the return value
//       }()
//       result = 42
//       return result, someErr             // defer still fires and sets result=-1
//   }
//
// Implement flowGotchaNamedReturn: return n squared.
// BUT: if n < 0, a deferred function should set the result to 0 instead.
// The deferred function must check the named return `result` and zero it if negative.
func flowGotchaNamedReturn(n int) (result int) {
	// TODO: implement
	// 1. defer func() that zeroes result if result < 0
	// 2. result = n * n
	// 3. return  (naked return — returns the named result)
	return
}

// GOTCHA 3: defer inside a loop defers ALL calls to function return, not loop end.
//
// ! In C# using inside a loop closes each resource at the end of each iteration.
// ! In Go, defer fires at FUNCTION return — if you defer f.Close() in a loop,
// ! ALL files stay open until the function returns.
//
// Broken pattern (holds all resources open simultaneously):
//   for _, path := range paths {
//       f, _ := os.Open(path)
//       defer f.Close()   // ! All files stay open until function returns!
//   }
//
// Fix: wrap the body in an inner function (anonymous or named):
//   for _, path := range paths {
//       func() {
//           f, _ := os.Open(path)
//           defer f.Close()   // Now defers at inner-func return (end of iteration)
//       }()
//   }
//
// Implement flowGotchaProcessEach: call process(item) for each item.
// Each call to process must happen BEFORE the next one starts (sequential).
// If any process() call returns an error, collect it; continue with the rest.
// Return all errors as a single combined error (or nil if all succeeded).
// Do NOT use defer for resource management here — focus on the error collection.
func flowGotchaProcessEach(items []string, process func(string) error) error {
	// TODO: implement — collect errors, join them, return combined or nil
	// Hint: append non-nil errors to []error, check len at the end,
	//       use errors.Join to combine (or return the single error directly)
	_ = errors.New // hint: errors package is available
	return nil
}
