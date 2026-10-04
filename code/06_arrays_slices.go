// Lesson 6: Arrays and Slices
package main

import (
	"fmt"
	"strings"
)

func ArraysSlices() {
	demoArrays()
	demoSlices()
	demoSliceIteration()
	demoFirstNames()
}

func demoArrays() {
	// ARRAYS - Fixed size
	fmt.Println("=== Arrays ===")

	// Declare array with size
	var numbers [5]int
	fmt.Println("Empty array:", numbers) // Zero values

	// Initialize array
	numbers[0] = 10
	numbers[1] = 20
	fmt.Println("After assignment:", numbers)

	// Array literal
	fruits := [3]string{"apple", "banana", "cherry"}
	fmt.Println("Fruits:", fruits)
	fmt.Println("Length:", len(fruits))

	// Let compiler count
	colors := [...]string{"red", "green", "blue", "yellow"}
	fmt.Println("Colors:", colors)
	fmt.Println("Length:", len(colors))
}

func demoSlices() {
	// SLICES - Dynamic size
	fmt.Println("=== Slices ===")

	// Create a base array and slices from it
	colors := [...]string{"red", "green", "blue", "yellow"}
	allColors := colors[:]    // All elements
	someColors := colors[1:3] // Elements 1 and 2
	fmt.Println("All colors:", allColors)
	fmt.Println("Some colors (1:3):", someColors)

	// Slice literal
	primes := []int{2, 3, 5, 7, 11}
	fmt.Println("Primes:", primes)

	// Make slice with length and capacity
	slice := make([]int, 3, 5) // length=3, capacity=5
	fmt.Printf("Slice: %v, len=%d, cap=%d\n", slice, len(slice), cap(slice))

	// SLICE OPERATIONS
	fmt.Println("\n=== Slice Operations ===")

	// Append
	slice = append(slice, 10)
	slice = append(slice, 20, 30) // Multiple values
	fmt.Println("After append:", slice)

	// Append slice to slice
	more := []int{40, 50}
	slice = append(slice, more...)
	fmt.Println("After appending slice:", slice)

	// Copy
	source := []int{1, 2, 3}
	dest := make([]int, len(source))
	copy(dest, source)
	fmt.Println("Copied slice:", dest)

	// SLICE TRICKS
	fmt.Println("\n=== Slice Tricks ===")

	data := []int{1, 2, 3, 4, 5}

	// Remove element at index 2
	removed := append(data[:2], data[3:]...)
	fmt.Println("Remove index 2:", removed)

	// Insert element
	data = []int{1, 2, 4, 5}
	data = append(data[:2], append([]int{3}, data[2:]...)...)
	fmt.Println("Insert 3 at index 2:", data)

	// 2D SLICES
	fmt.Println("\n=== 2D Slices ===")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matrix:")
	for _, row := range matrix {
		fmt.Println(" ", row)
	}
	fmt.Println("Element [1][2]:", matrix[1][2])
}

func demoSliceIteration() {
	// Iteration examples for slices
	fmt.Println("=== Slice Iteration Examples ===")

	primes := []int{2, 3, 5, 7, 11}

	// For loop with index
	fmt.Print("By index: ")
	for i := 0; i < len(primes); i++ {
		fmt.Print(primes[i], " ")
	}
	fmt.Println()

	// For-range
	fmt.Println("With range: ")
	for index, value := range primes {
		fmt.Printf("Index : %v :: value : %v \n", index, value)
	}
	fmt.Println()

	// Range with index
	fmt.Println("Index and value:")
	for i, v := range primes[:3] {
		fmt.Printf("  [%d] = %d\n", i, v)
	}

	// Additional iteration patterns
	fmt.Println("\n=== Additional Iteration Patterns ===")

	// Using range to modify a slice (careful: value is a copy)
	nums := []int{10, 20, 30}
	fmt.Println("Before modify attempt:", nums)
	for _, v := range nums {
		v = v + 1 // does not modify original
	}
	fmt.Println("After modify attempt (unchanged):", nums)

	// Proper way to modify in-place using index
	for i := range nums {
		nums[i] = nums[i] + 1
	}
	fmt.Println("After in-place modify:", nums)

	// Iterating over a slice of structs
	type Point struct{ X, Y int }
	points := []Point{{1, 2}, {3, 4}}
	for i, p := range points {
		fmt.Printf("Point %d: %#v\n", i, p)
	}

	// Safe iteration while deleting: iterate and build a new slice
	vals := []int{1, 2, 3, 4, 5}
	keep := vals[:0] // reuse backing array
	for _, v := range vals {
		if v%2 == 1 { // keep odd
			keep = append(keep, v)
		}
	}
	fmt.Println("Kept odds:", keep)
}

func demoFirstNames() {
	// Declare an array of 3 strings
	var fullNames [3]string

	// Assign values to the array
	fullNames[0] = "Alice Smith"
	fullNames[1] = "Bob Johnson"
	fullNames[2] = "Charlie Brown"

	// Print the full names
	for i, name := range fullNames {
		fmt.Println("Full Name", i+1, ":", name)
	}

	//Slice to hold first names
	firstNames := []string{}
	for _, name := range fullNames {
		firstName := strings.Fields(name)[0]
		// Append the first name to the slice
		firstNames = append(firstNames, firstName)
	}

	fmt.Println(firstNames)
}
