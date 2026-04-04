package learn

// ============================================================
// GOTCHA EXERCISES — things that surprise developers from C#/.NET
// ============================================================

// GOTCHA 1: Closures capture VARIABLES by reference, not by value.
//
// ! In C# lambdas also capture by reference, but loop variable capture is
// ! a well-known issue there too. In Go it's the same problem.
//
// Broken example — all closures return the last value of n:
//
//   funcs := make([]func() int, len(nums))
//   for i, n := range nums {         // n is ONE variable, reused each iteration
//       funcs[i] = func() int {
//           return n                  // ! ALL closures share the same n variable
//       }                             // by the time they run, n is the last value
//   }
//
// Note: In Go 1.22+ loop variables are per-iteration copies, so this specific
// bug is fixed — but closures capturing OTHER outer variables still behave this way.
//
// Implement funcsGotchaAdders: for each n in nums, return a closure that
// adds n to its argument. Each closure must capture its OWN copy of n.
// Hint: either pass n as a function argument, or assign n := n inside the loop.
func funcsGotchaAdders(nums []int) []func(int) int {
	// TODO: implement — ensure each closure captures an independent copy of n
	fns := make([]func(int) int, len(nums))
	for i, n := range nums {
		// ! If you write: fns[i] = func(x int) int { return x + n }
		// ! you capture the loop variable n by reference.
		// ! In Go <1.22 all closures would return x + last_n.
		// Fix: n := n  (creates a new variable in this scope)
		_ = n
		fns[i] = func(x int) int { return x } // ! BUG: fix this
	}
	return fns
}

// GOTCHA 2: Calling a nil function panics — always guard before calling.
//
// ! In C#, calling a null delegate throws NullReferenceException.
// ! In Go, calling a nil func panics with "runtime error: invalid memory address".
// ! There is no automatic nil check — you must do it yourself.
//
// Implement funcsGotchaSafeCall: call fn(x) if fn is not nil, returning the result.
// If fn is nil, return 0 instead of panicking.
func funcsGotchaSafeCall(fn func(int) int, x int) int {
	// TODO: implement — check fn != nil before calling
	return fn(x) // ! PANIC if fn is nil — guard this
}

// GOTCHA 3: Named returns + defer can change what a function actually returns.
//
// ! This is sometimes useful (e.g. adding cleanup on error return), but it
// ! makes functions harder to read because "return err" might not be the final word.
//
// Implement funcsGotchaWrapError: call fn() and return its error.
// Use a named return `err error` and a deferred function that, if err != nil,
// wraps it with additional context: err = fmt.Errorf("funcsGotchaWrapError: %w", err).
// The CALLER gets the wrapped error even though fn() returned the original.
func funcsGotchaWrapError(fn func() error) (err error) {
	// TODO: implement
	// 1. defer func() { if err != nil { err = fmt.Errorf("...%w", err) } }()
	// 2. err = fn()
	// 3. return
	return fn() // ! this bypasses the defer wrapping — restructure it
}
