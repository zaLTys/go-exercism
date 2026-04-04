package learn

import (
	"reflect"
	"testing"
)

func TestTypesGotchaAppend(t *testing.T) {
	// ! This test catches the "forgot to assign append result" bug.
	original := []int{1, 2, 3}
	result := typesGotchaAppend(original, 4)
	if len(result) != 4 || result[3] != 4 {
		t.Fatalf("got %v; want [1 2 3 4] — did you forget to return append(nums, x)?", result)
	}
}

func TestTypesGotchaIndependentSlice(t *testing.T) {
	nums := []int{10, 20, 30, 40, 50}
	got := typesGotchaIndependentSlice(nums, 1, 4) // [20 30 40]

	if !reflect.DeepEqual(got, []int{20, 30, 40}) {
		t.Fatalf("got %v; want [20 30 40]", got)
	}

	// ! Key assertion: modifying the returned slice must NOT affect the original.
	got[0] = 999
	if nums[1] == 999 {
		t.Fatalf("modifying the returned slice changed nums[1] — they share memory! use append([]int(nil), ...) to copy")
	}
}

func TestTypesGotchaSafeRead(t *testing.T) {
	// ! nil map read — should NOT panic
	var nilMap map[string]int
	if got := typesGotchaSafeRead(nilMap, "x"); got != 0 {
		t.Fatalf("nil map read returned %d; want 0", got)
	}

	// normal read
	m := map[string]int{"a": 1}
	if got := typesGotchaSafeRead(m, "a"); got != 1 {
		t.Fatalf("got %d; want 1", got)
	}
	// missing key returns zero value
	if got := typesGotchaSafeRead(m, "z"); got != 0 {
		t.Fatalf("got %d; want 0 for missing key", got)
	}
}

func TestTypesGotchaMapKeys(t *testing.T) {
	m := map[string]int{"banana": 2, "apple": 5, "cherry": 1}
	got := typesGotchaMapKeys(m)
	want := []string{"apple", "banana", "cherry"}
	// ! If you just collected keys with range and didn't sort, this will
	// ! sometimes pass and sometimes fail — that's exactly the gotcha.
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v; want %v — did you sort the keys?", got, want)
	}
}
