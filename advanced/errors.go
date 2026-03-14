package advanced

import (
	"errors"
	"fmt"
)

// RunErrorHandlingExample demonstrates error handling in Go
func RunErrorHandlingExample() {
	fmt.Println("6️⃣  Error Handling")
	
	// Basic error handling
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Division result: %.2f\n", result)
	}
	
	// Error case
	_, err = divide(10, 0)
	if err != nil {
		fmt.Printf("   ⚠️  Error caught: %v\n", err)
	}
	
	// Custom errors
	age := 15
	err = validateAge(age)
	if err != nil {
		fmt.Printf("   ⚠️  Validation error: %v\n", err)
	}
	
	age = 25
	err = validateAge(age)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   ✓ Age %d is valid\n", age)
	}
	
	// Multiple error checks
	user, err := createUser("Omkar", "omkar@example.com")
	if err != nil {
		fmt.Printf("   Error creating user: %v\n", err)
	} else {
		fmt.Printf("   ✓ User created: %s (%s)\n", user.Name, user.Email)
	}
	
	// Error with invalid email
	_, err = createUser("John", "invalid-email")
	if err != nil {
		fmt.Printf("   ⚠️  Error: %v\n", err)
	}
	
	// Panic and recover (for critical errors)
	fmt.Println("   Testing panic recovery...")
	safeDivide(10, 0)
	fmt.Println("   ✓ Program continued after panic!")
}

// Function that returns error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

// Custom error creation
func validateAge(age int) error {
	if age < 18 {
		return fmt.Errorf("age %d is below minimum required age of 18", age)
	}
	return nil
}

// User type for example
type User struct {
	Name  string
	Email string
}

// Function with validation and error
func createUser(name, email string) (*User, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if !contains(email, "@") {
		return nil, errors.New("invalid email format")
	}
	return &User{Name: name, Email: email}, nil
}

// Helper function
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// Panic and recover
func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("   ⚠️  Recovered from panic: %v\n", r)
		}
	}()
	
	if b == 0 {
		panic("cannot divide by zero!")
	}
	fmt.Printf("   Result: %d\n", a/b)
}
