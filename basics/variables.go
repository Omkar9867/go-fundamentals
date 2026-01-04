package basics

import "fmt"

// RunVariablesExample demonstrates variables, types, and constants in Go
func RunVariablesExample() {
	fmt.Println("1️⃣  Variables and Types")

	// Different ways to declare variables
	var name string = "Goku"
	var age int = 25
	var isLearning bool = true

	// Short declaration (most common)
	city := "Namaken"
	score := 95.5

	// Multiple variable declaration
	var (
		firstName = "Go"
		lastName  = "Language"
		year      = 2025
	)

	// Constants
	const PI = 3.14159
	const greeting = "Hello, World!"

	// Print examples
	fmt.Printf("   Name: %s\n", name)
	fmt.Printf("   Age: %d\n", age)
	fmt.Printf("   Learning: %t\n", isLearning)
	fmt.Printf("   City: %s\n", city)
	fmt.Printf("   Score: %.2f\n", score)
	fmt.Printf("   Full Name: %s %s (Year: %d)\n", firstName, lastName, year)
	fmt.Printf("   PI constant: %.5f\n", PI)
	fmt.Printf("   Greeting: %s\n", greeting)

	// Type conversion
	var integer int = 42
	var floating float64 = float64(integer)
	fmt.Printf("   Integer: %d, Float: %.2f\n", integer, floating)

	// Arrays and Slices
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	dynamicSlice := []string{"Go", "is", "awesome"}

	fmt.Printf("   Array: %v\n", numbers)
	fmt.Printf("   Slice: %v\n", dynamicSlice)

	// Maps (key-value pairs)
	scores := map[string]int{
		"Math":    90,
		"Science": 85,
		"English": 88,
	}
	fmt.Printf("   Scores map: %v\n", scores)
}
