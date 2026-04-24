package main

import (
	"fmt"
)

// 1. BASIC FUNCTION
// Simple function with no parameters and no returns
func greet() {
	fmt.Println("1. Basic Function: Hello from functions!")
}

// 2. PARAMETERS AND SINGLE RETURN
// Types come AFTER the variable name.
func add(a int, b int) int {
	return a + b
}

// 3. MULTIPLE RETURN VALUES
// Go functions can return more than one value. This is very common for error handling.
func swap(x, y string) (string, string) {
	return y, x
}

// 4. NAMED RETURN VALUES
// You can name the return variables in the function signature.
// They are treated as variables defined at the top of the function.
func divide(dividend, divisor float64) (quotient float64, err error) {
	if divisor == 0 {
		err = fmt.Errorf("cannot divide by zero")
		return // Naked return: returns 'quotient' and 'err' currently in scope
	}
	quotient = dividend / divisor
	return // Naked return
}

// 5. VARIADIC FUNCTIONS
// A function that takes any number of arguments of a specific type.
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	fmt.Println("--- Functions in Go ---")

	// 1. Calling basic function
	greet()

	// 2. Calling with parameters
	result := add(10, 20)
	fmt.Printf("2. Parameters: 10 + 20 = %d\n", result)

	// 3. Multiple returns
	a, b := swap("Hello", "World")
	fmt.Printf("3. Multiple Returns: swapped 'Hello World' -> '%s %s'\n", a, b)

	// 4. Named returns & Error handling pattern
	q, err := divide(10, 2)
	if err != nil {
		fmt.Printf("4. Named Returns: Error: %v\n", err)
	} else {
		fmt.Printf("4. Named Returns: 10 / 2 = %.1f\n", q)
	}

	// 5. Variadic functions
	fmt.Println("5. Variadic Functions:")
	fmt.Printf("   Sum(1, 2) = %d\n", sum(1, 2))
	fmt.Printf("   Sum(1, 2, 3, 4) = %d\n", sum(1, 2, 3, 4))
	
	// You can also pass a slice to a variadic function using '...'
	myNums := []int{10, 20, 30}
	fmt.Printf("   Sum(slice...) = %d\n", sum(myNums...))

	// 6. ANONYMOUS FUNCTIONS & CLOSURES
	// You can define a function inside another function.
	fmt.Println("6. Anonymous Functions & Closures:")
	
	// Anonymous function assigned to a variable
	multiply := func(x, y int) int {
		return x * y
	}
	fmt.Printf("   Anonymous: 5 * 6 = %d\n", multiply(5, 6))

	// Closure: A function that 'closes over' variables in its scope.
	nextInt := intSeq()
	fmt.Printf("   Closure 1: %d\n", nextInt()) // 1
	fmt.Printf("   Closure 2: %d\n", nextInt()) // 2
	fmt.Printf("   Closure 3: %d\n", nextInt()) // 3

	/*
	SUMMARY:
	- Function types come AFTER the name.
	- Functions can return multiple values.
	- Named returns allow for 'naked' returns (use sparingly in large functions).
	- Variadic functions use '...' for flexible arguments.
	- Closures allow functions to maintain state.
	*/
}

// Helper for Closure example
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
