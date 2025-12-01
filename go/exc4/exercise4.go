package concurpatterns

import (
	"context"
)

/*
Exercise 4 — Go Concurrency Best Practices
-------------------------------------------
See README_Exercise4 (2).md for detailed step-by-step instructions.

Pattern: Fan-Out, Fan-In with Error Handling
- Fan-Out: Launch one goroutine per source (concurrent fetching)
- Fan-In: Collect all results into a single slice
- Early Exit: If any source errors, cancel all others
*/

// Source is a function that fetches a string from some source.
// It respects context cancellation and may return an error.
type Source func(ctx context.Context) (string, error)

// FetchAll fetches data from multiple sources concurrently.
// Returns all results if successful, or first error encountered.
// See README for complete implementation guide.
func FetchAll(ctx context.Context, sources []Source) ([]string, error) {
	// STEP 1: Check if context is already cancelled
	// TODO: Implement early return for cancelled context

	// STEP 2: Create cancellable context
	// TODO: ctx, cancel := context.WithCancel(ctx)
	// TODO: defer cancel()

	// STEP 3: Create result channel
	type result struct {
		value string
		err   error
	}
	// TODO: results := make(chan result, len(sources))

	// STEP 4: Create WaitGroup
	// TODO: var wg sync.WaitGroup
	// TODO: wg.Add(len(sources))

	// STEP 5: Launch worker goroutines (Fan-Out)
	// TODO: for _, src := range sources { ... }

	// STEP 6: Close results when workers finish
	// TODO: go func() { wg.Wait(); close(results) }()

	// STEP 7: Collect results (Fan-In)
	// TODO: Implement result collection with error handling

	return nil, nil
}
