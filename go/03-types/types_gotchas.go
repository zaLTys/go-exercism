package learn

// ============================================================
// GOTCHA EXERCISES — things that surprise developers from C#/.NET
// ============================================================

// GOTCHA 1: append() returns a NEW slice — forgetting to reassign is a silent bug.
//
// ! In C#, list.Add(x) modifies the list in place.
// ! In Go, append(s, x) may return a DIFFERENT backing array.
// ! If you don't assign the result back, the original is unchanged.
//
// Broken example (compiles fine, silently wrong):
//   append(nums, x)       // result discarded — nums is unchanged!
//   _ = append(nums, x)   // same bug, just quieter
//
// Implement typesGotchaAppend: append x to nums and return the result.
func typesGotchaAppend(nums []int, x int) []int {
	// TODO: implement — don't forget to RETURN the result of append
	return nums // ! BUG: x is never added; replace this line
}

// GOTCHA 2: Sub-slices SHARE the backing array with the original.
//
// ! In C#, arr[1..3] (range operator) creates an independent copy.
// ! In Go, nums[1:3] is a VIEW into the same memory.
// ! Modifying the sub-slice also modifies the original.
//
// Broken example:
//   sub := nums[start:end]
//   sub[0] = 99             // ! also changes nums[start] — same memory!
//
// Implement typesGotchaIndependentSlice: return a COPY of nums[start:end]
// that does NOT share memory with nums.
// Hint: append([]int(nil), nums[start:end]...) produces an independent slice.
func typesGotchaIndependentSlice(nums []int, start, end int) []int {
	// TODO: implement — return an independent copy, not a shared sub-slice
	return nums[start:end] // ! BUG: shares backing array; fix this
}

// GOTCHA 3: Reading from a nil map is safe; WRITING to a nil map panics.
//
// ! var m map[string]int  ← this is nil, not empty
// ! m["x"]               ← safe: returns zero value (0)
// ! m["x"] = 1           ← PANIC: assignment to entry in nil map
//
// The distinction: nil map is "read-only empty"; an initialized map allows writes.
//
// Implement typesGotchaSafeRead: return the value at key from m.
// Must not panic even if m is nil or key is absent — return 0 in those cases.
func typesGotchaSafeRead(m map[string]int, key string) int {
	// TODO: implement — reading from nil map is already safe, just do it
	return 0
}

// GOTCHA 4: Map iteration order is deliberately randomized.
//
// ! In C# Dictionary iteration order is undefined but often consistent.
// ! In Go, the runtime RANDOMIZES map iteration order on purpose to prevent
// ! programs from depending on it.
//
// Implement typesGotchaMapKeys: return the keys of m sorted alphabetically.
// You MUST sort them — don't rely on range order.
// Hint: range m into a []string, then use sort.Strings.
func typesGotchaMapKeys(m map[string]int) []string {
	// TODO: implement — collect keys, sort them, return
	return nil
}
