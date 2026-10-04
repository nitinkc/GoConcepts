// Lesson 1: Go Basics
package main // Every Go program starts with a package declaration

import "fmt" // Import the fmt package for printing

func Basics() {
	// Print "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Println adds a newline at the end
	fmt.Println("Welcome to Go programming!")

	// Print without newline using Print
	fmt.Print("This ")
	fmt.Print("is ")
	fmt.Println("on the same line")

	// Printf allows formatted output
	fmt.Printf("Go is %s and %s!\n", "simple", "powerful")
}
