package basics

import "fmt"

// RunFunctionsExample demonstrates different types of functions in Go
func RunFunctionsExample() {
	fmt.Println("2️⃣  Functions")

	// Simple function call
	greet("Omkar")

	// Function with return value
	sum := add(10, 20)
	fmt.Printf("   Sum: %d\n", sum)

	// Multiple return values
	result, difference := calculate(50, 30)
	fmt.Printf("   Sum: %d, Difference: %d\n", result, difference)

	// Named return values
	area := rectangleArea(5, 10)
	fmt.Printf("   Rectangle Area: %d\n", area)

	// Variadic functions (variable number of arguments)
	total := sumAll(1, 2, 3, 4, 5)
	fmt.Printf("   Sum of all: %d\n", total)

	// Anonymous function (lambda)
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Printf("   Multiply: %d\n", multiply(6, 7))

	// Closures
	counter := makeCounter()
	fmt.Printf("   Counter: %d\n", counter())
	fmt.Printf("   Counter: %d\n", counter())
	fmt.Printf("   Counter: %d\n", counter())
}

// Simple function with parameter
func greet(name string) {
	fmt.Printf("   Hello, %s!\n", name)
}

// Function with return value
func add(a int, b int) int {
	return a + b
}

// Function with multiple return values
func calculate(a, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

// Function with named return values
func rectangleArea(length, width int) (area int) {
	area = length * width
	return // naked return
}

// Variadic function
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Closure - function that returns a function
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
