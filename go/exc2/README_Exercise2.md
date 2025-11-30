# 📘 Exercise 2 — Channel Fan-In (Merging Multiple Producers)

## Concept
Fan-in merges multiple channels into a single output channel.

```
A ----B ----- > MERGED → consumer
C ----/
```

## Task
Implement:

```go
func MergeGenerators(count, n int) []int
```

Each of the `n` goroutines must output integers `0..count-1`.  
Merge all outputs into one channel and return all numbers in any order.

## Hints
- Use a WaitGroup.
- Each producer should close its own channel.
- The merged output channel should close only after all producers are done.

## Goal
Learn the fan-in pattern using goroutines and channels.
