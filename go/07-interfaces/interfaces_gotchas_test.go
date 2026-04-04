package learn

import "testing"

func TestInterfacesGotchaGetString(t *testing.T) {
	if got := interfacesGotchaGetString("hello", "fallback"); got != "hello" {
		t.Fatalf("got %q; want %q", got, "hello")
	}
	// ! Without comma-ok, this would panic
	if got := interfacesGotchaGetString(42, "fallback"); got != "fallback" {
		t.Fatalf("int input: got %q; want %q", got, "fallback")
	}
	if got := interfacesGotchaGetString(nil, "fallback"); got != "fallback" {
		t.Fatalf("nil input: got %q; want %q", got, "fallback")
	}
}

func TestInterfacesGotchaReturnNilError(t *testing.T) {
	// Success path — MUST return a true nil error interface
	err := interfacesGotchaReturnNilError(true)

	// ! This is THE classic Go gotcha — a typed nil wraps in interface => non-nil
	// If the function returns `var e *MyError; return e`, this test fails.
	if err != nil {
		t.Fatalf("success case: err != nil — got %v (type %T)\n"+
			"  HINT: you returned a *interfacesMyError(nil), not a bare nil.\n"+
			"  Fix: return nil  (not: var err *interfacesMyError; return err)", err, err)
	}

	// Failure path — should return a real error
	err = interfacesGotchaReturnNilError(false)
	if err == nil {
		t.Fatalf("failure case: expected non-nil error")
	}
}

func TestInterfacesGotchaIsNilError(t *testing.T) {
	if !interfacesGotchaIsNilError(nil) {
		t.Fatalf("nil error: got false; want true")
	}
	if interfacesGotchaIsNilError(interfacesGotchaReturnNilError(false)) {
		t.Fatalf("real error: got true; want false")
	}
}

func TestInterfacesGotchaDescribe(t *testing.T) {
	// nil interface — must not panic
	if got := interfacesGotchaDescribe(nil); got != "<nil>" {
		t.Fatalf("nil: got %q; want \"<nil>\"", got)
	}

	n := &interfacesNamed{name: "world"}
	if got := interfacesGotchaDescribe(n); got != "Named(world)" {
		t.Fatalf("got %q; want %q", got, "Named(world)")
	}
}
