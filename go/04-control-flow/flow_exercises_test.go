package learn

import "testing"

func TestFlowGCD(t *testing.T) {
	if got := flowGCD(54, 24); got != 6 {
		t.Fatalf("got %d; want 6", got)
	}
	if got := flowGCD(0, 5); got != 5 {
		t.Fatalf("got %d; want 5", got)
	}
}

func TestFlowFizzBuzz(t *testing.T) {
	if got := flowFizzBuzz(5); got != "1,2,Fizz,4,Buzz" {
		t.Fatalf("got %q; want %q", got, "1,2,Fizz,4,Buzz")
	}
}

func TestFlowFirstEven(t *testing.T) {
	if v, ok := flowFirstEven([]int{1, 3, 4, 6}); !ok || v != 4 {
		t.Fatalf("got (%d,%v); want (4,true)", v, ok)
	}
	if v, ok := flowFirstEven([]int{1, 3, 5}); ok || v != 0 {
		t.Fatalf("got (%d,%v); want (0,false)", v, ok)
	}
}
