package pointers

import (
	"fmt"
	"strings"
)

/*
Exercise 5 — Pointers and Error Handling
-----------------------------------------
See README_Exercise5.md for detailed concepts and step-by-step instructions.

Goal: Build a User Management System demonstrating:
- Custom error types
- Pointer vs value receivers
- Pointer mutations
- Error handling patterns
- Batch operations with error collection
*/

// User represents a user account in the system.
type User struct {
	Email    string
	Password string
	Name     string
}

// STEP 1: Define Custom Error Type
// TODO: Create a ValidationError struct with:
//       - Field string (which field failed)
//       - Message string (what went wrong)
// TODO: Implement Error() string method to satisfy error interface
//       Format: "Field: Message" (e.g., "Email: must contain @")

type ValidationError struct {
	Field   string
	Message string
}

// TODO: Implement Error() method
//
//	func (e *ValidationError) Error() string {
//	    return fmt.Sprintf("%s: %s", e.Field, e.Message)
//	}
func (err *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", err.Field, err.Message)
}

// STEP 2: Implement Validation Methods (Value Receivers)
// These methods only READ the user, so they use value receivers

// ValidateEmail checks if the email is valid.
// Returns ValidationError if empty or doesn't contain "@".
// TODO: Implement this method
func (user User) ValidateEmail() error {
	if user.Email == "" {
		return &ValidationError{Field: "Email", Message: "Is empty"}
	}
	if !strings.Contains(user.Email, "@") {
		return &ValidationError{Field: "Email", Message: "@ is missing"}
	}
	return nil
}

// ValidatePassword checks if the password meets requirements.
// Returns ValidationError if less than 8 characters.
// TODO: Implement this method
func (user User) ValidatePassword() error {
	if len(user.Password) < 8 {
		return &ValidationError{Field: "Password", Message: "is too short"}
	}
	return nil
}

// STEP 3: Implement User Registration 
// TODO: Create RegisterUser function that:
//       1. Creates a User with given email and password
//       2. Validates email using ValidateEmail()
//       3. Validates password using ValidatePassword()
//       4. Returns (*User, error) - pointer because we'll modify it later
//       5. Returns (nil, error) if validation fails

func RegisterUser(email, password string) (*User, error) {
	user := User{
		Email:    email,
		Password: password,
		Name:     email,
	}

	// TODO: Validate email
	// TODO: Validate password
	// TODO: Return pointer to user if valid, or nil with error
	err := user.ValidateEmail()
	if err != nil {
		return nil, err
	}
	err = user.ValidatePassword()
	if err != nil {
		return nil, err
	}
	return &user, err

}

// STEP 4: Implement Profile Update (Pointer Receiver - Mutation!)
// TODO: Create UpdateProfile method with POINTER receiver
//       This modifies the original user, so we need a pointer!

// UpdateProfile updates the user's email and name.
// Uses pointer receiver because it modifies the user.
// TODO: Implement this method
// func (u *User) UpdateProfile(newEmail, newName string) error {
//     // TODO: Save old email in case we need to rollback
//     // TODO: Update u.Email and u.Name
//     // TODO: Validate the new email
//     // TODO: If validation fails, revert email and return error
//     // TODO: Return nil on success
// }

// STEP 5: Implement Password Change (Pointer Receiver)
// TODO: Create ChangePassword method with POINTER receiver
//       Combines pointers (mutation) and error handling (validation)

// ChangePassword changes the user's password after verifying the old one.
// TODO: Implement this method
// func (u *User) ChangePassword(oldPassword, newPassword string) error {
//     // TODO: Check if oldPassword matches u.Password
//     //       If not, return ValidationError{Field: "Password", Message: "incorrect old password"}
//     // TODO: Check if newPassword is at least 8 characters
//     //       If not, return ValidationError{Field: "Password", Message: "must be at least 8 characters"}
//     // TODO: Update u.Password to newPassword
//     // TODO: Return nil
// }

// STEP 6: Batch Operations with Error Collection
// TODO: Create ProcessUsers function that processes multiple users

// ProcessUsers applies an operation to each user and collects errors.
// Returns a slice of all errors encountered (empty if all succeeded).
// TODO: Implement this function
// func ProcessUsers(users []*User, operation func(*User) error) []error {
//     // TODO: Create empty error slice
//     // TODO: Loop through users
//     // TODO: Apply operation to each user
//     // TODO: If operation returns error, append to errors slice
//     // TODO: Return collected errors
// }

// Example usage (for your reference, don't modify):
func ExampleUsage() {
	// This won't work until you implement the functions!

	// Register a user
	// user, err := RegisterUser("alice@example.com", "securepass123")
	// if err != nil {
	//     fmt.Println("Registration failed:", err)
	//     return
	// }

	// Update profile
	// err = user.UpdateProfile("alice.smith@example.com", "Alice Smith")
	// if err != nil {
	//     fmt.Println("Update failed:", err)
	// }

	// Change password
	// err = user.ChangePassword("securepass123", "newsecurepass456")
	// if err != nil {
	//     fmt.Println("Password change failed:", err)
	// }

	fmt.Println("Complete the TODOs to make this work!")
}
