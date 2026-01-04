package basics

import "fmt"

// RunControlStructuresExample demonstrates if/else, loops, and switch statements
func RunControlStructuresExample() {
	fmt.Println("3️⃣  Control Structures")

	// If-else statements
	age := 20
	if age >= 18 {
		fmt.Println("   ✓ You are an adult")
	} else {
		fmt.Println("   ✗ You are a minor")
	}

	// If with short statement
	if score := 85; score >= 90 {
		fmt.Println("   Grade: A")
	} else if score >= 80 {
		fmt.Println("   Grade: B")
	} else {
		fmt.Println("   Grade: C")
	}

	// For loop (the only loop in Go!)
	fmt.Print("   Counting: ")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	nums := []string{"Go", "do", "your", "work", "now"}
	for i := 0; i <= 4; i++ {
		fmt.Printf("%s", nums[i]+" ")
	}
	fmt.Println()

	// While-style loop
	counter := 0
	fmt.Print("   While loop: ")
	for counter < 3 {
		fmt.Printf("%d ", counter)
		counter++
	}
	fmt.Println()

	// Infinite loop with break
	fmt.Print("   Break example: ")
	n := 0
	for {
		if n >= 3 {
			break
		}
		fmt.Printf("%d ", n)
		n++
	}
	fmt.Println()

	// Continue statement
	fmt.Print("   Skip even numbers: ")
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Range over slice
	fruits := []string{"Apple", "Banana", "Cherry"}
	fmt.Println("   Fruits:")
	for index, fruit := range fruits {
		fmt.Printf("      %d: %s\n", index, fruit)
	}

	// Range over map
	scores := map[string]int{"Math": 90, "Science": 85}
	fmt.Println("   Scores:")
	for subject, score := range scores {
		fmt.Printf("      %s: %d\n", subject, score)
	}

	// Switch statement
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("   😞 Start of the week")
	case "Friday":
		fmt.Println("   😊 Almost weekend!")
	case "Saturday", "Sunday":
		fmt.Println("   🎉 Weekend!")
	default:
		fmt.Println("   📅 Midweek day")
	}

	// Switch with expression
	number := 15
	switch {
	case number < 10:
		fmt.Println("   Number is less than 10")
	case number >= 10 && number < 20:
		fmt.Println("   Number is between 10 and 20")
	default:
		fmt.Println("   Number is 20 or greater")
	}
}
