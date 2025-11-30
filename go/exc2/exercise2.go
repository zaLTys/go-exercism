package fanin

import "sync"

/*
Exercise 2 — Channel Fan-In (Merging Multiple Producers)
-------------------------------------------------------

Background:
-----------
This exercise teaches the "fan-in" pattern: merging multiple producer goroutines
into a single channel.

Task:
-----
Implement MergeGenerators(count, n int) []int:

1. Launch `n` goroutines.
2. Each goroutine sends integers 0..count-1 into its own channel.
3. Merge all n channels into a single channel.
4. Receive ALL values until all producers finish.
5. Return the collected integers.
6. Order does NOT matter.

*/

func MergeGenerators(count, n int) []int {
	chans := make([]<-chan int, 0, n)

	for i := 0; i < n; i++ {
		chans = append(chans, generator(count))
	}

	merged := merge(chans...)

	var result []int
	for v := range merged {
		result = append(result, v)
	}
	return result
}

func generator(count int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < count; i++ {
			ch <- i
		}
	}()

	return ch
}

func merge(chs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(chs))
	for _, ch := range chs {
		ch := ch
		go func() {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
