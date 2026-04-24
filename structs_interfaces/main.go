package main

import (
	"fmt"
	"math"
)

// 1. STRUCTS
// A struct is a collection of fields.
type Person struct {
	Name string
	Age  int
}

// 2. METHODS
// A method is a function with a special receiver argument.
// Value receiver: 'p' is a copy of the Person. Any changes won't affect the original.
func (p Person) Greet() {
	fmt.Printf("1. Structs & Methods: Hi, I'm %s and I'm %d years old.\n", p.Name, p.Age)
}

// Pointer receiver: 'p' is a pointer. Changes WILL affect the original.
// This is also more efficient for large structs as it avoids copying.
func (p *Person) HaveBirthday() {
	p.Age++
	fmt.Printf("   %s just had a birthday! New age: %d\n", p.Name, p.Age)
}

// 3. INTERFACES
// An interface defines a set of method signatures.
type Shape interface {
	Area() float64
}

// Circle implements Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Rectangle implements Shape
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 4. POLYMORPHISM
// This function takes ANY Shape.
func printArea(s Shape) {
	fmt.Printf("3. Interfaces: Area of the shape is %.2f\n", s.Area())
}

func main() {
	fmt.Println("--- Structs & Interfaces in Go ---")

	// 1. Initializing Structs
	alice := Person{Name: "Alice", Age: 25}
	alice.Greet()

	// 2. Methods with Pointer Receivers
	alice.HaveBirthday()

	// 3. Using Interfaces
	fmt.Println("\n2. Working with Interfaces:")
	c := Circle{Radius: 5}
	r := Rectangle{Width: 10, Height: 5}

	printArea(c)
	printArea(r)

	// 4. EMPTY INTERFACE (interface{})
	// This can hold values of ANY type.
	fmt.Println("\n4. Empty Interface (interface{}):")
	var genericVal interface{}

	genericVal = "Now I'm a string"
	fmt.Printf("   Value: %v | Type: %T\n", genericVal, genericVal)

	genericVal = 42
	fmt.Printf("   Value: %v | Type: %T\n", genericVal, genericVal)

	// 5. TYPE ASSERTION
	// Extracting the underlying value from an interface.
	s, ok := genericVal.(int)
	if ok {
		fmt.Printf("   Type Assertion success: %d is an int\n", s)
	}

	// 6. TYPE SWITCH
	fmt.Println("\n5. Type Switch:")
	checkType("Hello")
	checkType(3.14)
	checkType(true)

	/*
	SUMMARY:
	- Structs are for grouping data.
	- Methods add behavior to structs using receivers.
	- Pointer receivers (*T) allow mutation; Value receivers (T) are for read-only.
	- Interfaces represent decoupled behavior.
	- If a type has all methods an interface defines, it implements it (Implicitly!).
	- interface{} is Go's way of handling data of unknown type (use with caution).
	*/
}

func checkType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("   '%s' is a String\n", v)
	case int:
		fmt.Printf("   '%d' is an Integer\n", v)
	case float64:
		fmt.Printf("   '%.2f' is a Float\n", v)
	default:
		fmt.Printf("   Unknown type: %T\n", v)
	}
}
