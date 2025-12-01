package concurpatterns

import (
	"context"
	"sync"
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
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// STEP 2: Create cancellable context
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// STEP 3: Create result channel
	type result struct {
		value string
		err   error
	}

	results := make(chan result, len(sources))
	// STEP 4: Create WaitGroup
	var wg sync.WaitGroup
	wg.Add(len(sources))

	// STEP 5: Launch worker goroutines (Fan-Out)
	for _, src := range sources {
		go func() {
			defer wg.Done()
			value, err := src(ctx)
			results <- result{value, err}
		}()
	}

	// STEP 6: Close results when workers finish
	go func() { wg.Wait(); close(results) }()

	// STEP 7: Collect results (Fan-In)
	var got []string
	for res := range results {
		if res.err != nil {
			cancel()
			return nil, res.err
		}
		got = append(got, res.value)
	}

	return got, nil
}
