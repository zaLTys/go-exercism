package learn

// EXERCISES 1-4: Implement methodsCart with four methods:
//   - Add(product string, price float64)  — add an item to the cart
//   - Remove(product string) bool         — remove first matching item; false if not found
//   - Total() float64                     — sum of all item prices
//   - ItemCount() int                     — number of items in the cart
//
// Requirements:
//   - Zero value must be usable: var c methodsCart; c.Add("pen", 1.50) should work.
//   - Remove removes only the FIRST item matching the product name.
//
// Real-world context: this is the core of any e-commerce cart — adding items,
// removing items, and computing totals. It exercises pointer vs value receivers,
// struct design, and the zero-value usability pattern.

type methodsCartItem struct {
	Product string
	Price   float64
}

type methodsCart struct {
	// TODO: choose fields (hint: a slice of methodsCartItem works well)
}

// Add appends an item to the cart.
func (c *methodsCart) Add(product string, price float64) {
	// TODO: implement
}

// Remove deletes the first item matching the product name.
// Returns false if the product was not found.
func (c *methodsCart) Remove(product string) bool {
	// TODO: implement
	return false
}

// Total returns the sum of all item prices in the cart.
func (c methodsCart) Total() float64 {
	// TODO: implement
	return 0
}

// ItemCount returns the number of items in the cart.
func (c methodsCart) ItemCount() int {
	// TODO: implement
	return 0
}
