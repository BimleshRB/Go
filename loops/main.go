package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- Loops in Go ---")
	// IMPORTANT: Go only has ONE looping construct: the 'for' loop.
	// There are no 'while' or 'do-while' keywords in Go. 
	// The 'for' loop is used in different ways to achieve the same results.

	// 1. CLASSIC FOR LOOP
	// Syntax: for initialization; condition; post { ... }
	fmt.Println("\n1. Classic For Loop:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 2. WHILE-STYLE LOOP
	// In Go, you just omit the initialization and post statements.
	fmt.Println("\n2. While-style Loop (using 'for'):")
	count := 1
	for count <= 5 {
		fmt.Printf("%d ", count)
		count++
	}
	fmt.Println()

	// 3. INFINITE LOOP
	// Omit all conditions to loop forever. Use 'break' to stop.
	fmt.Println("\n3. Infinite Loop (with break):")
	n := 1
	for {
		if n > 5 {
			break // Exit the loop
		}
		fmt.Printf("%d ", n)
		n++
	}
	fmt.Println()

	// 4. CONTINUE STATEMENT
	// Skips the current iteration and goes to the next one.
	fmt.Println("\n4. Using 'continue' (skipping even numbers):")
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			continue // Skip 2 and 4
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 5. FOR RANGE (Arrays, Slices, Maps, Strings)
	// This is the most common way to iterate over collections.
	fmt.Println("\n5. For-Range (Iterating over a slice):")
	fruits := []string{"Apple", "Banana", "Cherry"}
	for index, value := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// You can use '_' if you don't need the index
	fmt.Println("\nIterating over a string (Rune by Rune):")
	for _, char := range "Go🚀" {
		fmt.Printf("%c ", char)
	}
	fmt.Println()

	// 6. LABELED LOOPS (Advanced)
	// Labels allow you to control outer loops from within inner loops.
	// They are most useful in nested loops to avoid using 'flag' variables.

	fmt.Println("\n6a. Labeled Break (Exiting an outer loop):")
MyLabel:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				fmt.Println("   Exiting 'MyLabel' entirely at i=2, j=2")
				break MyLabel // Stops both loops
			}
			fmt.Printf("(%d,%d) ", i, j)
		}
		fmt.Println()
	}

	fmt.Println("\n6b. Labeled Continue (Skipping to next iteration of outer loop):")
Outer:
	for i := 1; i <= 3; i++ {
		fmt.Printf("Outer %d: ", i)
		for j := 1; j <= 3; j++ {
			if i == 2 {
				fmt.Print("Skipping rest of Outer 2... ")
				continue Outer // Jumps back to the 'for i' line
			}
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}

	/* 
	PRO TIP ON LABELS:
	- Labels are always follow by a colon (e.g., LabelName:)
	- They are typically placed immediately before a 'for', 'switch', or 'select' statement.
	- By convention, labels use CamelCase (MyLabel) or UPPER_CASE (OUTER), but never lowercase.
	- They help avoid complex 'boolean flag' logic like 'if shouldBreak { break }'.
	*/

	/*
	SUMMARY:
	- Go ONLY has 'for'.
	- 'for' can be used as a classic loop, a while loop, or an infinite loop.
	- 'range' is the idiomatic way to loop through collections.
	- 'break' exits a loop; 'continue' skips to the next iteration.
	- Labels can be used with break/continue for nested logic.
	*/
}
