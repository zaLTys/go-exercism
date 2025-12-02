# Exercise 5 — Pointers and Error Handling

**Learning Objectives:**
- Master Go pointers (addresses, dereferencing, mutations)
- Understand when to use pointers vs values
- Implement custom error types
- Practice error wrapping and handling
- Work with pointer receivers vs value receivers

---

## Part 1: Understanding Pointers

### What is a Pointer?

A **pointer** is a variable that stores the memory address of another variable.

```go
x := 42        // x is an integer with value 42
p := &x        // p is a pointer to x (stores x's memory address)
fmt.Println(p) // prints something like: 0xc0000140a0
```

### Key Concepts

#### 1. The `&` Operator (Address-of)
Gets the memory address of a variable:
```go
name := "Alice"
ptr := &name  // ptr holds the address of name
```

#### 2. The `*` Operator (Two Uses!)

**Use 1: Type Declaration** - declares a pointer type:
```go
var p *int  // p is a pointer to an int
```

**Use 2: Dereferencing** - accesses the value at the pointer's address:
```go
x := 10
p := &x      // p points to x
*p = 20      // dereference p and change the value
// Now x == 20!
```

### When to Use Pointers

| Use Pointers When | Use Values When |
|-------------------|-----------------|
| You need to modify the original variable | The data is small (int, bool, etc.) |
| The struct is large (avoid copying) | You want immutability |
| You need to represent "nil" or "absence" | The function is read-only |
| You're working with methods that mutate state | You're doing simple calculations |

---

## Part 2: Receivers (Value vs Pointer)

Methods can have either **value receivers** or **pointer receivers**.

### Value Receiver
```go
func (u User) GetName() string {
    return u.Name  // reads only, doesn't modify
}
```
- Receives a **copy** of the struct
- Cannot modify the original
- Good for read-only operations

### Pointer Receiver
```go
func (u *User) SetName(name string) {
    u.Name = name  // modifies the original!
}
```
- Receives a **pointer** to the struct
- Can modify the original
- Required for mutations

> **Rule of Thumb:** If any method needs a pointer receiver, use pointer receivers for ALL methods on that type (consistency).

---

## Part 3: Error Handling in Go

### Built-in Errors
```go
err := errors.New("something went wrong")
return nil, err
```

### Custom Error Types

Create custom errors by implementing the `error` interface:

```go
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}
```

**Benefits:**
- More descriptive errors
- Can carry additional context
- Enables type checking with `errors.As()`

### Error Wrapping

Use `fmt.Errorf` with `%w` to wrap errors:

```go
if err != nil {
    return fmt.Errorf("failed to register user: %w", err)
}
```

**Why wrap?**
- Preserves the original error
- Adds context to the error chain
- Enables unwrapping with `errors.Unwrap()` and `errors.Is()`

---

## Part 4: The Exercise

You'll build a **User Management System** with these features:

1. ✅ **Custom Error Types** - Create meaningful validation errors
2. ✅ **User Validation** - Validate email and password requirements
3. ✅ **User Registration** - Register new users with error handling
4. ✅ **Profile Updates** - Use pointers to modify user data
5. ✅ **Password Changes** - Combine pointers and validation
6. ✅ **Batch Operations** - Process multiple users with error collection

---

## Step-by-Step Guide

### STEP 1: Define Custom Error Types

Create `ValidationError` struct with:
- `Field string` - which field failed validation
- `Message string` - what went wrong

Implement the `Error() string` method to satisfy the `error` interface.

**Hint:**
```go
type ValidationError struct {
    // Your fields here
}

func (e *ValidationError) Error() string {
    // Return formatted string
}
```

---

### STEP 2: Implement Validation Methods

Create two validation methods with **value receivers** (read-only):

**`(u User) ValidateEmail() error`**
- Return `ValidationError` if email is empty or doesn't contain "@"
- Return `nil` if valid

**`(u User) ValidatePassword() error`**
- Return `ValidationError` if password is less than 8 characters
- Return `nil` if valid

**Why value receivers?** These methods only *read* the User; they don't modify it.

---

### STEP 3: Implement RegisterUser

Create a function that:
1. Validates email using `ValidateEmail()`
2. Validates password using `ValidatePassword()`
3. If either validation fails, return `nil` and the error
4. If both pass, return a **pointer** to a new User

**Signature:**
```go
func RegisterUser(email, password string) (*User, error)
```

**Why return a pointer?** 
- We'll modify this user later
- Pointers are efficient for struct returns
- Can return `nil` to represent "no user"

---

### STEP 4: Implement UpdateProfile (Pointer Mutation)

Create a method with a **pointer receiver**:

**`(u *User) UpdateProfile(newEmail, newName string) error`**

1. Update `u.Email` with `newEmail`
2. Update `u.Name` with `newName`
3. Validate the new email using `ValidateEmail()`
4. If validation fails, **revert the changes** and return the error
5. If valid, return `nil`

**Key Learning:** Pointer receivers let you modify the original struct!

---

### STEP 5: Implement ChangePassword

Create a method with a **pointer receiver**:

**`(u *User) ChangePassword(oldPassword, newPassword string) error`**

1. Check if `oldPassword` matches `u.Password`
   - If not, return `ValidationError` with field "Password" and message "incorrect old password"
2. Validate `newPassword` length (≥8 characters)
   - If invalid, return appropriate `ValidationError`
3. Update `u.Password` to `newPassword`
4. Return `nil` on success

**Combines:** Pointers (mutation) + Error handling (validation)

---

### STEP 6: Implement ProcessUsers (Batch Operations)

Create a function that processes a slice of user pointers:

**`func ProcessUsers(users []*User, operation func(*User) error) []error`**

1. Create an empty `[]error` slice to collect errors
2. For each user pointer in the slice:
   - Call `operation(user)`
   - If it returns an error, append it to the errors slice
3. Return all collected errors

**Use Case:** Apply an operation to many users, collect all failures.

**Example:**
```go
errors := ProcessUsers(users, func(u *User) error {
    return u.UpdateProfile("new@email.com", "NewName")
})
```

---

## Common Pitfalls

### 🚫 Pitfall 1: Forgetting to Dereference
```go
func (u *User) SetName(name string) {
    u.Name = name  // ✅ Correct (implicit dereference)
    (*u).Name = name  // ✅ Also correct (explicit)
}
```

### 🚫 Pitfall 2: Nil Pointer Dereference
```go
var u *User
u.Name = "Alice"  // ❌ PANIC! u is nil
```

Always check for nil before dereferencing:
```go
if u != nil {
    u.Name = "Alice"  // ✅ Safe
}
```

### 🚫 Pitfall 3: Copying Pointers in Loops
```go
for _, user := range users {
    processedUsers = append(processedUsers, &user)  // ❌ Wrong!
}
```

The loop variable is reused! Use index or copy:
```go
for i := range users {
    processedUsers = append(processedUsers, &users[i])  // ✅ Correct
}
```

---

## Testing Your Implementation

Run the tests:
```bash
cd c:\Users\povil\Exercism\go\exc5
go test -v
```

All tests should pass when you complete the 6 steps!

---

## Key Takeaways

1. **Pointers store addresses** - Use `&` to get address, `*` to dereference
2. **Use pointers to modify** - Pointer receivers enable mutations
3. **Custom errors add context** - Implement `error` interface for rich errors
4. **Wrap errors** - Use `fmt.Errorf` with `%w` to maintain error chains
5. **Validate early** - Check inputs before processing
6. **Handle errors explicitly** - Go doesn't have exceptions; check every error

Good luck! 🚀
