package pointers

import (
	"errors"
	"testing"
)

// Test STEP 1: Custom Error Type
func TestValidationError(t *testing.T) {
	err := &ValidationError{
		Field:   "Email",
		Message: "must contain @",
	}

	expected := "Email: must contain @"
	if err.Error() != expected {
		t.Errorf("Error() = %q, want %q", err.Error(), expected)
	}
}

// Test STEP 2: Email Validation
func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name      string
		user      User
		wantError bool
		errField  string
	}{
		{
			name:      "valid email",
			user:      User{Email: "alice@example.com"},
			wantError: false,
		},
		{
			name:      "empty email",
			user:      User{Email: ""},
			wantError: true,
			errField:  "Email",
		},
		{
			name:      "missing @ symbol",
			user:      User{Email: "notanemail"},
			wantError: true,
			errField:  "Email",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.ValidateEmail()
			if tt.wantError {
				if err == nil {
					t.Error("ValidateEmail() expected error, got nil")
					return
				}
				var valErr *ValidationError
				if !errors.As(err, &valErr) {
					t.Errorf("expected ValidationError, got %T", err)
					return
				}
				if valErr.Field != tt.errField {
					t.Errorf("Field = %q, want %q", valErr.Field, tt.errField)
				}
			} else {
				if err != nil {
					t.Errorf("ValidateEmail() unexpected error: %v", err)
				}
			}
		})
	}
}

// Test STEP 2: Password Validation
func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name      string
		user      User
		wantError bool
	}{
		{
			name:      "valid password",
			user:      User{Password: "password123"},
			wantError: false,
		},
		{
			name:      "exactly 8 characters",
			user:      User{Password: "12345678"},
			wantError: false,
		},
		{
			name:      "too short",
			user:      User{Password: "short"},
			wantError: true,
		},
		{
			name:      "empty password",
			user:      User{Password: ""},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.ValidatePassword()
			if tt.wantError && err == nil {
				t.Error("ValidatePassword() expected error, got nil")
			}
			if !tt.wantError && err != nil {
				t.Errorf("ValidatePassword() unexpected error: %v", err)
			}
		})
	}
}

// Test STEP 3: User Registration
func TestRegisterUser(t *testing.T) {
	tests := []struct {
		name      string
		email     string
		password  string
		wantError bool
		errField  string
	}{
		{
			name:      "valid registration",
			email:     "bob@example.com",
			password:  "securepass",
			wantError: false,
		},
		{
			name:      "invalid email",
			email:     "bademail",
			password:  "securepass",
			wantError: true,
			errField:  "Email",
		},
		{
			name:      "invalid password",
			email:     "bob@example.com",
			password:  "short",
			wantError: true,
			errField:  "Password",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := RegisterUser(tt.email, tt.password)
			if tt.wantError {
				if err == nil {
					t.Error("RegisterUser() expected error, got nil")
					return
				}
				if user != nil {
					t.Error("RegisterUser() expected nil user on error")
				}
				var valErr *ValidationError
				if errors.As(err, &valErr) {
					if valErr.Field != tt.errField {
						t.Errorf("Field = %q, want %q", valErr.Field, tt.errField)
					}
				}
			} else {
				if err != nil {
					t.Errorf("RegisterUser() unexpected error: %v", err)
					return
				}
				if user == nil {
					t.Error("RegisterUser() expected user, got nil")
					return
				}
				if user.Email != tt.email {
					t.Errorf("Email = %q, want %q", user.Email, tt.email)
				}
				if user.Password != tt.password {
					t.Errorf("Password = %q, want %q", user.Password, tt.password)
				}
			}
		})
	}
}

// Test STEP 4: Profile Update (Pointer Mutation)
func TestUpdateProfile(t *testing.T) {
	t.Run("successful update", func(t *testing.T) {
		user := &User{
			Email:    "old@example.com",
			Password: "password123",
			Name:     "Old Name",
		}

		err := user.UpdateProfile("new@example.com", "New Name")
		if err != nil {
			t.Errorf("UpdateProfile() unexpected error: %v", err)
		}

		if user.Email != "new@example.com" {
			t.Errorf("Email = %q, want %q", user.Email, "new@example.com")
		}
		if user.Name != "New Name" {
			t.Errorf("Name = %q, want %q", user.Name, "New Name")
		}
	})

	t.Run("invalid email rollback", func(t *testing.T) {
		user := &User{
			Email:    "valid@example.com",
			Password: "password123",
			Name:     "User Name",
		}

		originalEmail := user.Email
		err := user.UpdateProfile("invalidemail", "New Name")

		if err == nil {
			t.Error("UpdateProfile() expected error for invalid email")
			return
		}

		// Email should be rolled back to original
		if user.Email != originalEmail {
			t.Errorf("Email should be rolled back to %q, got %q", originalEmail, user.Email)
		}
	})
}

// Test STEP 5: Password Change
func TestChangePassword(t *testing.T) {
	t.Run("successful password change", func(t *testing.T) {
		user := &User{
			Email:    "user@example.com",
			Password: "oldpassword",
		}

		err := user.ChangePassword("oldpassword", "newpassword123")
		if err != nil {
			t.Errorf("ChangePassword() unexpected error: %v", err)
		}

		if user.Password != "newpassword123" {
			t.Errorf("Password = %q, want %q", user.Password, "newpassword123")
		}
	})

	t.Run("incorrect old password", func(t *testing.T) {
		user := &User{
			Email:    "user@example.com",
			Password: "correctpass",
		}

		err := user.ChangePassword("wrongpass", "newpassword123")
		if err == nil {
			t.Error("ChangePassword() expected error for wrong old password")
			return
		}

		var valErr *ValidationError
		if errors.As(err, &valErr) {
			if valErr.Field != "Password" {
				t.Errorf("Field = %q, want %q", valErr.Field, "Password")
			}
		}

		// Password should not change
		if user.Password != "correctpass" {
			t.Error("Password should not change when old password is incorrect")
		}
	})

	t.Run("new password too short", func(t *testing.T) {
		user := &User{
			Email:    "user@example.com",
			Password: "oldpassword",
		}

		err := user.ChangePassword("oldpassword", "short")
		if err == nil {
			t.Error("ChangePassword() expected error for short new password")
			return
		}

		// Password should not change
		if user.Password != "oldpassword" {
			t.Error("Password should not change when new password is invalid")
		}
	})
}

// Test STEP 6: Batch Operations
func TestProcessUsers(t *testing.T) {
	t.Run("all operations succeed", func(t *testing.T) {
		users := []*User{
			{Email: "user1@example.com", Password: "password1", Name: "User 1"},
			{Email: "user2@example.com", Password: "password2", Name: "User 2"},
		}

		errs := ProcessUsers(users, func(u *User) error {
			return u.UpdateProfile(u.Email, "Updated Name")
		})

		if len(errs) != 0 {
			t.Errorf("ProcessUsers() expected no errors, got %d", len(errs))
		}

		// Verify all names were updated
		for _, user := range users {
			if user.Name != "Updated Name" {
				t.Errorf("Name = %q, want %q", user.Name, "Updated Name")
			}
		}
	})

	t.Run("some operations fail", func(t *testing.T) {
		users := []*User{
			{Email: "user1@example.com", Password: "password1"},
			{Email: "user2@example.com", Password: "password2"},
			{Email: "user3@example.com", Password: "password3"},
		}

		errs := ProcessUsers(users, func(u *User) error {
			// This will fail because invalid email
			return u.UpdateProfile("invalidemail", "Name")
		})

		if len(errs) != 3 {
			t.Errorf("ProcessUsers() expected 3 errors, got %d", len(errs))
		}
	})

	t.Run("empty user list", func(t *testing.T) {
		var users []*User

		errs := ProcessUsers(users, func(u *User) error {
			return nil
		})

		if len(errs) != 0 {
			t.Errorf("ProcessUsers() expected no errors for empty list, got %d", len(errs))
		}
	})
}

// Benchmark: Pointer vs Value
func BenchmarkPointerReceiver(b *testing.B) {
	user := &User{Email: "test@example.com", Password: "password123"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = user.UpdateProfile("new@example.com", "New Name")
	}
}

func BenchmarkValueReceiver(b *testing.B) {
	user := User{Email: "test@example.com", Password: "password123"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = user.ValidateEmail()
	}
}
