package learn

// EXERCISES 1-3: Implement methodsStack with three methods:
//   - Push(v int)         — add v to the top
//   - Pop() (int, bool)   — remove and return top; false if empty
//   - Len() int           — current number of items
//
// Requirements:
//   - Zero value must be usable: var s methodsStack; s.Push(1) should work.
//   - Pop on empty stack => (0, false).

type methodsStack struct {
	// TODO: choose fields (hint: a slice of int works well)
}

// Push adds v to the top of the stack.
func (s *methodsStack) Push(v int) {
	// TODO: implement
}

// Pop removes and returns the top element.
// Returns (0, false) if the stack is empty.
func (s *methodsStack) Pop() (int, bool) {
	// TODO: implement
	return 0, false
}

// Len returns the number of elements in the stack.
func (s methodsStack) Len() int {
	// TODO: implement
	return 0
}
