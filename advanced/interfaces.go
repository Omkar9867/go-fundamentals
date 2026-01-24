package advanced

import (
	"fmt"
	"math"
)

// Shape interface - defines behavior
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Rectangle implements Shape interface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle implements Shape interface
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Function that accepts any Shape
func printShapeInfo(s Shape) {
	fmt.Printf("   Area: %.2f\n", s.Area())
	fmt.Printf("   Perimeter: %.2f\n", s.Perimeter())
}

// RunInterfacesExample demonstrates interfaces in Go
func RunInterfacesExample() {
	fmt.Println("5️⃣  Interfaces")
	
	// Create shapes
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}
	
	// Use interface polymorphism
	fmt.Println("   Rectangle:")
	printShapeInfo(rect)
	
	fmt.Println("   Circle:")
	printShapeInfo(circle)
	
	// Slice of interface types
	shapes := []Shape{
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 5},
		Rectangle{Width: 6, Height: 8},
	}
	
	fmt.Println("   All shapes:")
	for i, shape := range shapes {
		fmt.Printf("   Shape %d - Area: %.2f\n", i+1, shape.Area())
	}
	
	// Empty interface (interface{}) can hold any type
	var anything interface{}
	anything = "Hello"
	fmt.Printf("   Anything (string): %v\n", anything)
	anything = 42
	fmt.Printf("   Anything (int): %v\n", anything)
	anything = true
	fmt.Printf("   Anything (bool): %v\n", anything)
	
	// Type assertion
	var i interface{} = "Hello, Go!"
	str, ok := i.(string)
	if ok {
		fmt.Printf("   Type assertion success: %s\n", str)
	}
	
	// Type switch
	checkType(21)
	checkType("Learning Go")
	checkType(3.14)
	checkType(true)
}

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("   Integer: %d\n", v)
	case string:
		fmt.Printf("   String: %s\n", v)
	case float64:
		fmt.Printf("   Float: %.2f\n", v)
	default:
		fmt.Printf("   Unknown type: %T\n", v)
	}
}
