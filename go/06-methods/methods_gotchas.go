package learn

import "fmt"

// ============================================================
// GOTCHA EXERCISES — things that surprise developers from C#/.NET
// ============================================================

// GOTCHA 1: Cannot call a POINTER RECEIVER method on a non-addressable value.
//
// ! Map VALUES are not addressable in Go — you cannot take their address.
// ! So if a type has pointer receiver methods, you can't call them on a map value directly.
//
// Broken example:
//
//   type Counter struct{ n int }
//   func (c *Counter) Inc() { c.n++ }
//
//   m := map[string]Counter{}
//   m["x"].Inc()   // COMPILE ERROR: cannot take address of m["x"]
//
// Workaround: copy the value out, modify, put it back:
//   c := m["x"]
//   c.Inc()
//   m["x"] = c
//
// Implement methodsGotchaIncrementMap: increment the "count" field of the Counter
// stored at key in m, and return the new count.
// The methodsCounter type has a pointer receiver method Inc() — you cannot call
// it directly on m[key]. Use the copy-modify-store workaround.

type methodsCounter struct{ count int }

func (c *methodsCounter) Inc() int {
	c.count++
	return c.count
}

func methodsGotchaIncrementMap(m map[string]methodsCounter, key string) int {
	// TODO: implement the copy-modify-store pattern
	// ! m[key].Inc() won't compile — m[key] is not addressable
	return 0
}

// GOTCHA 2: Pointer receiver methods are NOT in the method set of the VALUE type.
//
// ! This means a *T satisfies interfaces requiring pointer receiver methods,
// ! but T (value) does NOT — even though Go auto-takes address in some cases.
//
// Practically: when you store a value (not pointer) in an interface variable,
// pointer receiver methods become unreachable.
//
// Example:
//   type Stringer interface{ String() string }
//   type Foo struct{ name string }
//   func (f *Foo) String() string { return f.name }  // pointer receiver
//
//   var s Stringer = Foo{"bar"}   // COMPILE ERROR — Foo doesn't have String()
//   var s Stringer = &Foo{"bar"}  // OK — *Foo has String()
//
// Implement methodsGotchaStringerCheck: given an any value, return its String()
// result if it implements fmt.Stringer; otherwise return "<not a Stringer>".
// ! To see the gotcha: try passing a methodsNamed value vs a *methodsNamed pointer.

type methodsNamed struct{ name string }

// String is defined on *methodsNamed (pointer receiver)
func (n *methodsNamed) String() string {
	return fmt.Sprintf("Named(%s)", n.name)
}

func methodsGotchaStringerCheck(v any) string {
	// TODO: implement using a type assertion to fmt.Stringer
	// Hint: s, ok := v.(fmt.Stringer)
	return ""
}
