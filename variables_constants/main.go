// variables and constants in go
package main

import "fmt"

func main() {
	// --- VARIABLES ---
	// Variables are used to store data that can change during execution.

	// 1. Explicitly declared variable
	var name string = "Bimlesh"
	fmt.Println("Name:", name)

	// 2. Type Inference (Go figures out the type)
	var age = 25
	fmt.Printf("Age: %d (Type: %T)\n", age, age)

	// 3. Shorthand syntax := (Only inside functions)
	city := "San Francisco"
	fmt.Println("City:", city)

	// --- BASIC TYPES ---

	// String: Sequence of characters in quotes
	var message string = "Hello Go!"

	// Integers: Whole numbers (int, int8, int16, int32, int64)
	var standardInt int = 100
	var largeInt int64 = 9223372036854775807

	// Unsigned Integers: Only positive numbers (uint, uint8, uint16, etc.)
	var positiveOnly uint = 50

	// Floating Point: Decimal numbers (float32, float64)
	var pi float64 = 3.14159

	// Boolean: true or false
	var isGoFun bool = true

	// --- SPECIAL TYPES & ALIASES ---

	// Byte: Alias for uint8 (used for raw data/ASCII)
	var asciiChar byte = 'A' // Dec: 65

	// Rune: Alias for int32 (represents a Unicode code point / character)
	var emoji rune = '🚀'

	// Complex Numbers: complex64, complex128
	var c complex128 = complex(5, 7) // 5 + 7i

	// Printing all variables to satisfy Go's "unused variable" rule
	fmt.Println("--- All Variable Examples ---")
	fmt.Printf("Message: %s, Int: %d, LargeInt: %d, Positive: %d\n", message, standardInt, largeInt, positiveOnly)
	fmt.Printf("Pi: %.2f, IsGoFun: %v, Byte: %c, Rune: %c\n", pi, isGoFun, asciiChar, emoji)
	fmt.Println("Complex number:", c)

	// --- CONSTANTS ---
	// Constants are values that CANNOT change. Use 'const' instead of 'var'.
	const Country = "India"
	const Pi = 3.14

	fmt.Println("--- Constants ---")
	fmt.Println("Constant Country:", Country)
	fmt.Println("Constant Pi:", Pi)

	// constants cannot be changed
	// Multime constants at once->
	// untyped constants example
	// const (
	// 	a = 1
	// 	b = 2
	// 	c = 3
	// )
	// typed constants example
	// const (
	// 	d int = 1
	// 	e int = 2
	// 	f int = 3
	// )

	/*
		Summary of Types:
		- int, int8, int16, int32, int64: Integers
		- uint, uint8, uint16, uint32, uint64: Unsigned (Positive) Integers
		- float32, float64: Decimal Numbers
		- complex64, complex128: Complex Numbers
		- bool: true/false
		- string: Text
		- byte: Alias for uint8
		- rune: Alias for int32 (Unicode character)
	*/
}
