package learn

import (
	"bytes"
	"testing"
)

func TestToolingFormatGoSource(t *testing.T) {
	in := "package p\n\nfunc f(){return}\n"
	out, err := toolingFormatGoSource(in)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if !bytes.Contains([]byte(out), []byte("func f()")) {
		t.Fatalf("formatted output missing expected substring; got:\n%s", out)
	}
}

func TestToolingLooksGofmted(t *testing.T) {
	src := "package p\n\nfunc f() {}\n"
	if !toolingLooksGofmted(src) {
		t.Fatalf("expected true for already formatted source")
	}
	if toolingLooksGofmted("package p\n\nfunc f(){}\n") {
		t.Fatalf("expected false for unformatted source")
	}
}

func TestToolingCountTestFunctions(t *testing.T) {
	src := `
package p

import "testing"

func TestA(t *testing.T) {}
func TestB(t *testing.T) {}
func helper() {}
`
	if got := toolingCountTestFunctions(src); got != 2 {
		t.Fatalf("got %d; want 2", got)
	}
}
