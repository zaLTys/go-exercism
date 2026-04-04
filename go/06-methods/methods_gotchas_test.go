package learn

import "testing"

func TestMethodsGotchaIncrementMap(t *testing.T) {
	m := map[string]methodsCounter{
		"hits": {count: 5},
	}

	// ! m["hits"].Inc() would be a compile error.
	// The implementation must copy, modify, store back.
	got := methodsGotchaIncrementMap(m, "hits")
	if got != 6 {
		t.Fatalf("got %d; want 6", got)
	}
	// The map value itself must be updated
	if m["hits"].count != 6 {
		t.Fatalf("map value not updated: m[\"hits\"].count=%d; want 6", m["hits"].count)
	}
}

func TestMethodsGotchaStringerCheck(t *testing.T) {
	// ! Pointer: *methodsNamed satisfies fmt.Stringer (pointer receiver method)
	ptr := &methodsNamed{name: "Alice"}
	if got := methodsGotchaStringerCheck(ptr); got != "Named(Alice)" {
		t.Fatalf("ptr: got %q; want %q", got, "Named(Alice)")
	}

	// ! Value: methodsNamed does NOT satisfy fmt.Stringer
	// The String() method is on *methodsNamed, not methodsNamed.
	// When stored as `any`, Go cannot auto-take the address, so it's not a Stringer.
	val := methodsNamed{name: "Bob"}
	if got := methodsGotchaStringerCheck(val); got != "<not a Stringer>" {
		t.Fatalf("value: got %q; want \"<not a Stringer>\" — value type lacks pointer receiver method", got)
	}

	// Non-stringer
	if got := methodsGotchaStringerCheck(42); got != "<not a Stringer>" {
		t.Fatalf("int: got %q; want \"<not a Stringer>\"", got)
	}
}
