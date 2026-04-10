package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- Control Flow in Go ---")

	// 1. BASIC IF / ELSE
	age := 18
	if age >= 18 {
		fmt.Println("1. If/Else: You are an adult.")
	} else {
		fmt.Println("1. If/Else: You are a minor.")
	}

	// 2. IF WITH INITIALIZATION (Very idiomatic Go)
	// You can declare a variable that is only available inside the if/else block.
	if length := 10; length > 5 {
		fmt.Printf("2. If with Init: Length %d is greater than 5\n", length)
	}
	// fmt.Println(length) // Error: length is not defined here!

	// 3. SWITCH STATEMENTS
	// Go's switch is different: it automatically 'breaks' after each case.
	day := "Monday"
	fmt.Print("3. Basic Switch: ")
	switch day {
	case "Monday":
		fmt.Println("Start of the work week.")
	case "Friday":
		fmt.Println("Almost the weekend!")
	default:
		fmt.Println("Just another day.")
	}

	// 4. SWITCH WITH MULTIPLE VALUES
	fmt.Print("4. Multiple Values in Switch: ")
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	// 5. EXPRESSIONLESS SWITCH (Conditionals in cases)
	// This is the cleanest way to write complex if/else chains in Go.
	num := 75
	fmt.Print("5. Expressionless Switch: ")
	switch {
	case num < 50:
		fmt.Println("Number is small.")
	case num >= 50 && num <= 100:
		fmt.Println("Number is medium.")
	default:
		fmt.Println("Number is large.")
	}

	// 6. THE 'FALLTHROUGH' KEYWORD
	// Go doesn't fall through by default. You must use 'fallthrough' explicitly.
	fmt.Println("6. Using Fallthrough:")
	switch 1 {
	case 1:
		fmt.Println("Case 1 matches.")
		fallthrough
	case 2:
		fmt.Println("Case 2 also runs because of fallthrough!")
	}

	/*
	SUMMARY:
	- 'if' doesn't need parentheses around the condition.
	- 'if' can have an initialization statement.
	- 'switch' cases break automatically (no 'break' needed).
	- Use expressionless switch for complex logic.
	- Use 'fallthrough' only if you explicitly want the next case to run.
	*/
}
