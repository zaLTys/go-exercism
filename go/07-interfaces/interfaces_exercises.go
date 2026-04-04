package learn

import (
	"io"
	"math"
)

type interfacesShape interface {
	Area() float64
}

type interfacesSquare struct{ Side float64 }
type interfacesCircle struct{ Radius float64 }

// EXERCISE 1: Implement Area() so both types satisfy interfacesShape.
//
// interfacesSquare.Area() = Side * Side
// interfacesCircle.Area() = math.Pi * Radius * Radius

func (s interfacesSquare) Area() float64 {
	// TODO: implement
	return 0
}

func (c interfacesCircle) Area() float64 {
	// TODO: implement
	_ = math.Pi
	return 0
}

// interfacesTotalArea sums the areas of all shapes.
func interfacesTotalArea(shapes []interfacesShape) float64 {
	// TODO: implement
	return 0
}

// EXERCISE 2: Implement interfacesDescribeAny using a type switch.
// Rules:
//   - int    => "int:<n>"
//   - string => "string:<value>"
//   - default => "unknown"
func interfacesDescribeAny(v any) string {
	// TODO: implement type switch
	return ""
}

// EXERCISE 3: Implement interfacesReadAllUpper.
// Read everything from r, return the uppercase string.
func interfacesReadAllUpper(r io.Reader) (string, error) {
	// TODO: implement (hint: io.ReadAll + strings.ToUpper)
	return "", nil
}
