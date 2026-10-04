// Lesson 9: Pointers in Go
// Memory addresses and references
package main

import "fmt"

func Pointers() {
	// POINTER BASICS
	fmt.Println("=== Pointer Basics ===")

	x := 42
	var p *int = &x // p is a pointer to int, & gets address

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Pointer p:", p)
	fmt.Println("Value at p:", *p) // Dereferencing

	// Modify through pointer
	*p = 100
	fmt.Println("x after *p = 100:", x)

	// ZERO VALUE
	fmt.Println("\n=== Zero Value ===")
	var nilPointer *int
	fmt.Println("Nil pointer:", nilPointer)
	fmt.Println("Is nil?:", nilPointer == nil)

	// NEW FUNCTION
	fmt.Println("\n=== Using new() ===")
	ptr := new(int) // Allocates memory, returns pointer
	fmt.Println("new(int):", ptr)
	fmt.Println("Value:", *ptr) // Zero value
	*ptr = 50
	fmt.Println("After assignment:", *ptr)

	// POINTERS AND FUNCTIONS
	fmt.Println("\n=== Pointers in Functions ===")

	// Without pointer (pass by value)
	num := 10
	doubleValue(num)
	fmt.Println("After doubleValue:", num) // Still 10

	// With pointer (pass by reference)
	doublePointer(&num)
	fmt.Println("After doublePointer:", num) // Now 20

	// POINTERS AND STRUCTS
	fmt.Println("\n=== Pointers and Structs ===")

	type Pointee struct {
		Name string
		Age  int
	}

	// Creating pointer to struct
	person := &Pointee{Name: "Alice", Age: 25}
	fmt.Printf("Person: %+v\n", *person)

	// Accessing fields (Go auto-dereferences)
	fmt.Println("Name:", person.Name) // Same as (*person).Name
	person.Age = 26
	fmt.Printf("After update: %+v\n", *person)

	// POINTER TO POINTER
	fmt.Println("\n=== Pointer to Pointer ===")
	a := 5
	pa := &a
	ppa := &pa

	fmt.Println("a:", a)
	fmt.Println("*pa:", *pa)
	fmt.Println("**ppa:", **ppa)

	**ppa = 10
	fmt.Println("a after **ppa = 10:", a)

	// ARRAYS VS SLICES
	fmt.Println("\n=== Arrays vs Slices ===")

	// Arrays are copied
	arr := [3]int{1, 2, 3}
	modifyArray(arr)
	fmt.Println("Array after function:", arr) // Unchanged

	// Slices contain pointer to underlying array
	slice := []int{1, 2, 3}
	modifySlice(slice)
	fmt.Println("Slice after function:", slice) // Changed!

	// WHEN TO USE POINTERS
	fmt.Println("\n=== Practical Example ===")

	// Large struct - use pointer to avoid copying
	data := &LargeData{}
	processLargeData(data) // Efficient - no copy

	// Returning pointer from function
	config := createConfig()
	fmt.Printf("Config: %+v\n", *config)
}

// Function without pointer - cannot modify original
func doubleValue(n int) {
	n = n * 2 // Only modifies local copy
}

// Function with pointer - can modify original
func doublePointer(n *int) {
	*n = *n * 2 // Modifies original
}

// Arrays are passed by value (copied)
func modifyArray(arr [3]int) {
	arr[0] = 999 // Only modifies copy
}

// Slices contain pointer to array
func modifySlice(slice []int) {
	slice[0] = 999 // Modifies original array
}

// Efficient for large structs
type LargeData struct {
	Data [1000]int
}

func processLargeData(d *LargeData) {
	d.Data[0] = 42
}

// Return pointer from function
type Config struct {
	Host string
	Port int
}

func createConfig() *Config {
	return &Config{
		Host: "localhost",
		Port: 8080,
	}
}
