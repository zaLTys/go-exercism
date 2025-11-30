package fanin

import (
    "sort"
    "testing"
)

func TestMergeGeneratorsLength(t *testing.T) {
    count := 5
    n := 3
    got := MergeGenerators(count, n)

    want := count * n
    if len(got) != want {
        t.Fatalf("expected %d values, got %d", want, len(got))
    }
}

func TestMergeGeneratorsContent(t *testing.T) {
    count := 4
    n := 2
    got := MergeGenerators(count, n)

    counts := make(map[int]int)
    for _, v := range got {
        counts[v]++
    }

    for i := 0; i < count; i++ {
        if counts[i] != n {
            t.Errorf("expected %d occurrences of %d, got %d", n, i, counts[i])
        }
    }
}

func TestMergeGeneratorsDeterministicTotal(t *testing.T) {
    count := 10
    n := 5
    got := MergeGenerators(count, n)

    if len(got) != count*n {
        t.Fatalf("expected %d values, got %d", count*n, len(got))
    }

    sort.Ints(got)
}
