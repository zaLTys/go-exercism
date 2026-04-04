package learn

import "testing"

func TestMethodsStackPushPop(t *testing.T) {
	var s methodsStack // zero value
	s.Push(10)
	s.Push(20)
	if s.Len() != 2 {
		t.Fatalf("Len()=%d; want 2", s.Len())
	}
	v, ok := s.Pop()
	if !ok || v != 20 {
		t.Fatalf("Pop()=(%d,%v); want (20,true)", v, ok)
	}
	v, ok = s.Pop()
	if !ok || v != 10 {
		t.Fatalf("Pop()=(%d,%v); want (10,true)", v, ok)
	}
	v, ok = s.Pop()
	if ok || v != 0 {
		t.Fatalf("Pop()=(%d,%v); want (0,false)", v, ok)
	}
}

func TestMethodsStackLenEmpty(t *testing.T) {
	var s methodsStack
	if s.Len() != 0 {
		t.Fatalf("Len()=%d; want 0", s.Len())
	}
}

func TestMethodsStackInterleave(t *testing.T) {
	var s methodsStack
	s.Push(1)
	s.Push(2)
	_, _ = s.Pop()
	s.Push(3)
	v, ok := s.Pop()
	if !ok || v != 3 {
		t.Fatalf("got (%d,%v); want (3,true)", v, ok)
	}
	v, ok = s.Pop()
	if !ok || v != 1 {
		t.Fatalf("got (%d,%v); want (1,true)", v, ok)
	}
}
