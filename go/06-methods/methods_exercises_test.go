package learn

import "testing"

func TestMethodsCartAddAndTotal(t *testing.T) {
	var c methodsCart // zero value
	c.Add("Pen", 1.50)
	c.Add("Notebook", 3.00)
	if c.ItemCount() != 2 {
		t.Fatalf("ItemCount()=%d; want 2", c.ItemCount())
	}
	want := 4.50
	if got := c.Total(); got != want {
		t.Fatalf("Total()=%v; want %v", got, want)
	}
}

func TestMethodsCartRemove(t *testing.T) {
	var c methodsCart
	c.Add("Pen", 1.50)
	c.Add("Notebook", 3.00)
	c.Add("Pen", 1.50)

	if !c.Remove("Pen") {
		t.Fatalf("Remove(\"Pen\")=false; want true")
	}
	if c.ItemCount() != 2 {
		t.Fatalf("after first Remove: ItemCount()=%d; want 2", c.ItemCount())
	}
	want := 4.50
	if got := c.Total(); got != want {
		t.Fatalf("after removing first Pen: Total()=%v; want %v", got, want)
	}
	if !c.Remove("Pen") {
		t.Fatalf("second Remove(\"Pen\")=false; want true — second Pen still in cart")
	}
	if c.Remove("Pen") {
		t.Fatalf("third Remove(\"Pen\")=true; want false — no more Pens")
	}
}

func TestMethodsCartEmpty(t *testing.T) {
	var c methodsCart
	if c.Total() != 0 || c.ItemCount() != 0 {
		t.Fatalf("empty cart: Total=%v, ItemCount=%d; want 0, 0", c.Total(), c.ItemCount())
	}
	if c.Remove("anything") {
		t.Fatalf("Remove on empty cart should return false")
	}
}
