package learn

import (
	"reflect"
	"sort"
	"testing"
)

func TestTypesOrderTotal(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		if got := typesOrderTotal(nil); got != 0 {
			t.Fatalf("nil: got %v; want 0", got)
		}
	})
	t.Run("single item", func(t *testing.T) {
		items := []typesLineItem{{"Pen", 1.50, 4}}
		if got := typesOrderTotal(items); got != 6.0 {
			t.Fatalf("got %v; want 6.0", got)
		}
	})
	t.Run("multiple items", func(t *testing.T) {
		items := []typesLineItem{
			{"Notebook", 4.00, 2},
			{"Pen", 1.50, 3},
		}
		got := typesOrderTotal(items)
		want := 4.00*2 + 1.50*3
		if got != want {
			t.Fatalf("got %v; want %v", got, want)
		}
	})
	t.Run("empty slice", func(t *testing.T) {
		if got := typesOrderTotal([]typesLineItem{}); got != 0 {
			t.Fatalf("empty: got %v; want 0", got)
		}
	})
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
