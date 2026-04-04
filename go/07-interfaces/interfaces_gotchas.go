package learn

import "fmt"

// ============================================================
// GOTCHA EXERCISES — things that surprise developers from C#/.NET
// ============================================================

// GOTCHA 1: Type assertion WITHOUT comma-ok panics on wrong type.
//
// ! In C#, (Foo)obj throws InvalidCastException.
// ! In Go, obj.(Foo) panics with "interface conversion: ... is not Foo"
// ! The SAFE form is: v, ok := obj.(Foo) — ok is false instead of panic.
//
// Implement interfacesGotchaGetString: if v holds a string, return it.
// Otherwise return the fallback. Must NOT panic on any input type.
func interfacesGotchaGetString(v any, fallback string) string {
	// TODO: implement using the comma-ok type assertion form
	// ! v.(string) panics if v is not a string
	// Safe: s, ok := v.(string)
	return fallback
}

// GOTCHA 2: A nil POINTER wrapped in an interface is NOT nil.
//
// ! This is arguably the single most surprising thing in Go for newcomers.
// ! An interface holds two things: (type, pointer). An interface is only nil
// ! if BOTH are nil. A nil *MyError stored in an error interface is NOT nil
// ! because the type information is non-nil.
//
// The classic mistake:
//
//   type MyError struct{ msg string }
//   func (e *MyError) Error() string { return e.msg }
//
//   func badFunc() error {
//       var err *MyError = nil   // typed nil pointer
//       return err               // ! wraps nil pointer in error interface — NOT nil!
//   }
//
//   err := badFunc()
//   err == nil   // FALSE — the interface has type *MyError, data nil
//
// Implement interfacesGotchaReturnNilError: return nil (not a typed nil).
// Must return an error interface that == nil.
// The broken version is provided — fix it.
func interfacesGotchaReturnNilError(success bool) error {
	if success {
		// ! BUG: this returns a non-nil interface wrapping a nil pointer
		var err *interfacesMyError // typed nil
		return err                 // ! fix this — return nil directly
	}
	return &interfacesMyError{msg: "something went wrong"}
}

type interfacesMyError struct{ msg string }

func (e *interfacesMyError) Error() string { return e.msg }

// GOTCHA 3: Interface comparison — two interface values are equal only if
// both their dynamic TYPE and dynamic VALUE are equal.
//
// ! In C#, reference equality for interfaces is straightforward.
// ! In Go, comparing two interfaces with == compares (type, value) pairs.
// ! Two nils are equal; a nil interface != a non-nil interface with nil data.
//
// Implement interfacesGotchaIsNilError: return true ONLY if err is a genuine
// nil interface (not a non-nil interface wrapping a nil pointer).
// Hint: the simple err == nil check is ENOUGH for a real nil interface.
// The gotcha is returning a typed nil from functions — see GOTCHA 2 above.
func interfacesGotchaIsNilError(err error) bool {
	// TODO: implement — for this exercise, just return err == nil
	// (The real lesson is: NEVER return a typed nil from a function returning error)
	return false
}

// GOTCHA 4: Avoid fmt.Stringer pitfall — embedding interface in struct.
//
// ! If a struct embeds an interface type and that field is nil at runtime,
// ! calling any method on the embedded interface panics.
//
// This is usually a sign of "interface in struct" anti-pattern.
// Instead of embedding, accept the interface as a parameter.
//
// Implement interfacesGotchaDescribe: given a fmt.Stringer, return its String().
// If v is nil (the interface itself is nil), return "<nil>".
func interfacesGotchaDescribe(v fmt.Stringer) string {
	// TODO: implement — check for nil interface before calling v.String()
	return ""
}

// interfacesNamed is a local type to demonstrate Stringer with pointer receiver.
type interfacesNamed struct{ name string }

func (n *interfacesNamed) String() string { return fmt.Sprintf("Named(%s)", n.name) }
