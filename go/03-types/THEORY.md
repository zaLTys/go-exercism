# Types — Theory Reference

## The Big Idea

Go is **statically typed**. The "zero value" concept is fundamental: every type has a useful default. Data modeling uses **structs** (not classes), **slices** (not arrays), and **maps**. There are no generics on built-in types — slice/map syntax is built into the language.

## Primitives & Zero Values

| Type | Zero value | Notes |
|------|-----------|-------|
| `int`, `int64`, `float64` | `0` | Size-specific variants exist |
| `string` | `""` | Immutable, UTF-8 bytes |
| `bool` | `false` | |
| `*T` (pointer) | `nil` | |
| slice `[]T` | `nil` | nil slice is safe to `append` to |
| map `map[K]V` | `nil` | nil map panics on write — must initialize |

```go
var i int     // 0
var s string  // ""
var b bool    // false
var p *int    // nil
var sl []int  // nil — but append(sl, 1) is safe!
```

## Structs

```go
type Person struct {
    Name string
    Age  int
}

// Initialize
p := Person{Name: "Paul", Age: 30}
p2 := Person{"Paul", 30}  // positional (fragile — avoid for exported types)

// Access
fmt.Println(p.Name)
p.Age = 31
```

**No inheritance.** Use **embedding** for composition:
```go
type Employee struct {
    Person               // embedded — Employee gets Person's fields
    Department string
}
e := Employee{Person: Person{Name: "Paul", Age: 30}, Department: "Eng"}
fmt.Println(e.Name)  // promoted field
```

## Slices

```go
// literal
nums := []int{1, 2, 3}

// make (length, capacity)
s := make([]int, 5)      // [0 0 0 0 0]
s2 := make([]int, 0, 10) // empty, capacity 10

// append — returns NEW slice (may reallocate)
nums = append(nums, 4, 5)

// range
for i, v := range nums {
    fmt.Println(i, v)
}

// slice of slice
sub := nums[1:3]  // [2, 3] — shares memory
```

| C# | Go |
|----|-----|
| `new List<int> { 1, 2, 3 }` | `[]int{1, 2, 3}` |
| `list.Add(x)` | `list = append(list, x)` |
| `list[i]` | `list[i]` |
| `list.Count` | `len(list)` |

## Maps

```go
// literal
scores := map[string]int{"a": 1, "b": 2}

// make — required for empty map before writing
m := make(map[string]int)
m["key"] = 42

// safe lookup — second return is "exists"
v, ok := m["key"]
if !ok {
    fmt.Println("not found")
}

// delete
delete(m, "key")

// range
for k, v := range m {
    fmt.Println(k, v)
}
```

**Key rule:** A nil map panics on write. Always initialize: `m := make(map[K]V)` or `m := map[K]V{}`.

| C# | Go |
|----|-----|
| `new Dictionary<string,int>()` | `make(map[string]int)` |
| `dict["key"] = val` | `m["key"] = val` |
| `dict.TryGetValue("key", out val)` | `val, ok := m["key"]` |
| `dict.ContainsKey("key")` | `_, ok := m["key"]; ok` |
| `dict.Remove("key")` | `delete(m, "key")` |

## Pointers

```go
x := 42
p := &x      // p is *int
*p = 100     // dereference to set
fmt.Println(x)  // 100

// Pointers to structs: auto-dereference
person := &Person{Name: "Paul"}
fmt.Println(person.Name)  // no need for (*person).Name
```

## Useful Links
- [Tour: Types](https://go.dev/tour/moretypes)
- [Go maps in action](https://go.dev/blog/maps)
- [Effective Go — zero values](https://go.dev/doc/effective_go#zero-value)
