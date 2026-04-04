package learn

import (
	"errors"
	"io"
)

// idiomsStringSet is a set of strings backed by a map.
type idiomsStringSet map[string]struct{}

// EXERCISE 1: Implement methods on idiomsStringSet.
//   - Add(s string)  — must work even when the receiver is a nil map
//   - Has(s string) bool
//   - Len() int
//
// Hint: Add needs a pointer receiver (*idiomsStringSet) to initialize a nil map.

func (s *idiomsStringSet) Add(v string) {
	// TODO: implement
	// If *s is nil, initialize it: *s = make(idiomsStringSet)
	// Then (*s)[v] = struct{}{}
}

func (s idiomsStringSet) Has(v string) bool {
	// TODO: implement
	return false
}

func (s idiomsStringSet) Len() int {
	// TODO: implement
	return 0
}

// EXERCISE 2: Implement idiomsReadAllAndClose.
// Read everything from rc and always close it.
// If both readErr and closeErr are non-nil, return errors.Join(readErr, closeErr).
func idiomsReadAllAndClose(rc io.ReadCloser) ([]byte, error) {
	// TODO: implement
	_ = errors.Join(nil, nil)
	return nil, nil
}

// EXERCISE 3: Implement idiomsPreferConcreteReturn.
// Accept io.Reader (interface), return []string (concrete).
// Split the content by newlines, ignoring any empty trailing line.
func idiomsPreferConcreteReturn(r io.Reader) ([]string, error) {
	// TODO: implement (hint: io.ReadAll to get bytes, strings.Split on "\n",
	//       then filter out the final empty element if present)
	return nil, nil
}
