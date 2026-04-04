package learn

type typesLineItem struct {
	Product  string
	Price    float64
	Quantity int
}

type typesPerson struct {
	Name string
	Age  int
}

// EXERCISE 1: Implement typesOrderTotal.
// Calculate the total cost of all line items (Price * Quantity for each).
// Must handle nil slices safely (nil => 0.0).
//
// Real-world context: this is the core of any e-commerce checkout —
// iterate a cart of items and compute the total.
func typesOrderTotal(items []typesLineItem) float64 {
	// TODO: implement
	return 0
}

// EXERCISE 2: Implement typesInvertMap.
// Input:  map[string]int{"a":1,"b":1,"c":2}
// Output: map[int][]string{1:["a","b"], 2:["c"]}
// (values are sorted in the test for determinism — you don't need to sort here)
//
// Real-world context: building a reverse index — e.g. given product→category,
// produce category→[]products. This pattern appears constantly in search,
// tagging, and grouping systems.
func typesInvertMap(in map[string]int) map[int][]string {
	// TODO: implement (hint: remember nil maps can't be assigned into)
	return nil
}

// EXERCISE 3: Implement typesGroupByAgeDecade.
// Example: Age 29 -> "20s", Age 30 -> "30s".
// Hint: the decade key can be built with fmt.Sprintf("%ds", (age/10)*10)
//       e.g. age=29 => (29/10)*10 = 20 => "20s"
//       You'll need to import "fmt" for this.
//
// Real-world context: segmenting users by demographic buckets for analytics
// dashboards or targeted notifications.
func typesGroupByAgeDecade(people []typesPerson) map[string][]string {
	// TODO: implement
	return nil
}
