package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- Data Structures in Go ---")

	// 1. ARRAYS
	// Fixed size, defined at compile time.
	fmt.Println("\n1. Arrays (Fixed Size):")
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println("   Colors Array:", colors)
	fmt.Println("   Length of colors:", len(colors))

	// Array literal (shorthand)
	oddNums := [5]int{1, 3, 5, 7, 9}
	fmt.Println("   Odd Numbers:", oddNums)

	// 2. SLICES
	// Dynamic size, built on top of arrays. This is what you use 99% of the time.
	fmt.Println("\n2. Slices (Dynamic Size):")
	
	// Creating a slice from an array literal (leaving out the size)
	fruits := []string{"Apple", "Banana", "Cherry"}
	fmt.Println("   Fruits Slice:", fruits)

	// Appending to a slice
	fruits = append(fruits, "Dragonfruit", "Elderberry")
	fmt.Println("   After append:", fruits)
	fmt.Println("   Length:", len(fruits), "| Capacity:", cap(fruits))

	// Slicing a slice [low : high] (high is exclusive)
	subSlice := fruits[1:4]
	fmt.Println("   Sub-slice [1:4]:", subSlice)

	// Creating a slice with 'make' (useful for performance)
	// make([]type, length, capacity)
	mySlice := make([]int, 5, 10)
	fmt.Printf("   Make slice: len=%d, cap=%d, values=%v\n", len(mySlice), cap(mySlice), mySlice)

	// 3. MAPS
	// Key-value pairs (hash tables). Unordered.
	fmt.Println("\n3. Maps (Key-Value Pairs):")
	
	// Creating a map with make: make(map[keyType]valueType)
	userAges := make(map[string]int)
	userAges["Alice"] = 25
	userAges["Bob"] = 30
	userAges["Charlie"] = 35
	fmt.Println("   User Ages Map:", userAges)

	// Accessing values
	fmt.Println("   Alice's Age:", userAges["Alice"])

	// Map Literal
	inventory := map[string]int{
		"laptop": 10,
		"mouse":  50,
		"desk":   5,
	}
	fmt.Println("   Inventory:", inventory)

	// Deleting from a map
	delete(inventory, "mouse")
	fmt.Println("   Inventory after deleting mouse:", inventory)

	// Checking if a key exists
	val, exists := inventory["laptop"]
	if exists {
		fmt.Printf("   Laptop count: %d\n", val)
	} else {
		fmt.Println("   Laptop not found.")
	}

	/*
	SUMMARY:
	- Arrays have a FIXED size. Slices are DYNAMIC.
	- Slices are views into underlying arrays.
	- Use 'append' to add elements to a slice.
	- Maps are unordered hash tables for quick lookups.
	- Use 'make' to pre-allocate slices or maps for better performance.
	- Accessing a map returns an optional 'exists' boolean.
	*/
}
