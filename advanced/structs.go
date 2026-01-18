package advanced

import "fmt"

// Person struct - custom data type
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
}

// Method with value receiver
func (p Person) GetFullName() string {
	return p.FirstName + " " + p.LastName
}

// Method with pointer receiver (can modify the struct)
func (p *Person) HaveBirthday() {
	p.Age++
}

// RunStructsExample demonstrates structs and methods in Go
func RunStructsExample() {
	fmt.Println("4️⃣  Structs and Methods")

	// Creating a struct
	person1 := Person{
		FirstName: "Omkar",
		LastName:  "Sharma",
		Age:       25,
		Email:     "omkar@example.com",
	}

	// Accessing fields
	fmt.Printf("   Name: %s\n", person1.GetFullName())
	fmt.Printf("   Age: %d\n", person1.Age)
	fmt.Printf("   Email: %s\n", person1.Email)

	// Using method with pointer receiver
	fmt.Println("   🎂 Having a birthday...")
	person1.HaveBirthday()
	fmt.Printf("   New Age: %d\n", person1.Age)

	// Short struct initialization
	person2 := Person{"Raj", "Kumar", 30, "raj@example.com"}
	fmt.Printf("   Person 2: %s (%d years old)\n", person2.GetFullName(), person2.Age)

	// Pointer to struct
	person3 := &Person{
		FirstName: "Priya",
		LastName:  "Singh",
		Age:       28,
	}
	fmt.Printf("   Person 3: %s\n", person3.GetFullName())

	// Anonymous struct
	car := struct {
		Brand string
		Model string
		Year  int
	}{
		Brand: "Toyota",
		Model: "Camry",
		Year:  2024,
	}
	fmt.Printf("   Car: %s %s (%d)\n", car.Brand, car.Model, car.Year)
}
