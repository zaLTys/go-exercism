package learn

import (
	"reflect"
	"sort"
	"testing"
)

func TestTypesSumInts(t *testing.T) {
	if got := typesSumInts(nil); got != 0 {
		t.Fatalf("got %d; want 0", got)
	}
	if got := typesSumInts([]int{1, 2, 3}); got != 6 {
		t.Fatalf("got %d; want 6", got)
	}
}

func TestTypesInvertMap(t *testing.T) {
	in := map[string]int{"b": 1, "a": 1, "c": 2}
	got := typesInvertMap(in)
	for _, v := range got {
		sort.Strings(v)
	}
	want := map[int][]string{1: {"a", "b"}, 2: {"c"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v; want %#v", got, want)
	}
}

func TestTypesGroupByAgeDecade(t *testing.T) {
	people := []typesPerson{
		{Name: "A", Age: 29},
		{Name: "B", Age: 30},
		{Name: "C", Age: 31},
	}
	got := typesGroupByAgeDecade(people)
	for _, v := range got {
		sort.Strings(v)
	}
	want := map[string][]string{
		"20s": {"A"},
		"30s": {"B", "C"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v; want %#v", got, want)
	}
}
