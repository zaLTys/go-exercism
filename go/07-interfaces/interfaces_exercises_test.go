package learn

import (
	"strings"
	"testing"
)

func TestInterfacesTotalArea(t *testing.T) {
	shapes := []interfacesShape{
		interfacesSquare{Side: 2},   // area 4
		interfacesCircle{Radius: 1}, // area ~3.14159
	}
	got := interfacesTotalArea(shapes)
	if got < 7.1 || got > 7.2 {
		t.Fatalf("total area=%v; expected about 7.14..", got)
	}
}

func TestInterfacesDescribeAny(t *testing.T) {
	if got := interfacesDescribeAny(7); got != "int:7" {
		t.Fatalf("got %q; want %q", got, "int:7")
	}
	if got := interfacesDescribeAny("hi"); got != "string:hi" {
		t.Fatalf("got %q; want %q", got, "string:hi")
	}
	if got := interfacesDescribeAny(true); got != "unknown" {
		t.Fatalf("got %q; want %q", got, "unknown")
	}
}

func TestInterfacesReadAllUpper(t *testing.T) {
	got, err := interfacesReadAllUpper(strings.NewReader("go rocks"))
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if got != "GO ROCKS" {
		t.Fatalf("got %q; want %q", got, "GO ROCKS")
	}
}

