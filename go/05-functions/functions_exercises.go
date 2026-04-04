package learn

import "errors"

var errDivideByZero = errors.New("divide by zero")

// EXERCISE 1: Implement funcsDivMod.
// Return error on b==0. Otherwise return q=a/b and r=a%b.
func funcsDivMod(a, b int) (q int, r int, err error) {
	// TODO: implement
	return 0, 0, nil
}

// EXERCISE 2: Implement funcsSumVariadic.
// Sum all provided numbers. No args => 0.
func funcsSumVariadic(nums ...int) int {
	// TODO: implement
	return 0
}

// EXERCISE 3: Implement funcsMakeCounter closure.
// Returns a function that yields 1, 2, 3, ... on successive calls.
func funcsMakeCounter() func() int {
	// TODO: implement (hint: close over a variable in the outer function)
	return func() int { return 0 }
}
