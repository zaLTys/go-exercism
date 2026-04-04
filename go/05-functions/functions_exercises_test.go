package learn

import "testing"

func TestFuncsDivMod(t *testing.T) {
	_, _, err := funcsDivMod(10, 0)
	if err == nil {
		t.Fatalf("expected error for b==0")
	}
	q, r, err := funcsDivMod(10, 3)
	if err != nil || q != 3 || r != 1 {
		t.Fatalf("got (q=%d,r=%d,err=%v); want (3,1,nil)", q, r, err)
	}
}

func TestFuncsSumVariadic(t *testing.T) {
	if got := funcsSumVariadic(); got != 0 {
		t.Fatalf("got %d; want 0", got)
	}
	if got := funcsSumVariadic(1, 2, 3); got != 6 {
		t.Fatalf("got %d; want 6", got)
	}
}

func TestFuncsMakeCounter(t *testing.T) {
	next := funcsMakeCounter()
	if next() != 1 || next() != 2 || next() != 3 {
		t.Fatalf("counter did not increment as expected")
	}
}
