// Lesson 4: Control Flow in Go
// Learn if/else, switch, and loops
package main

import "fmt"

func ControlFlow() {
	// IF-ELSE STATEMENTS
	fmt.Println("=== If-Else ===")
	age := 20

	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	// If with initialization statement
	if score := 85; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}
	// Note: 'score' is not accessible here

	// SWITCH STATEMENTS
	fmt.Println("\n=== Switch ===")
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Start of work week")
	case "Friday":
		fmt.Println("TGIF!")
	case "Saturday", "Sunday":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Midweek")
	}

	// Switch without expression (like if-else chain)
	temperature := 25
	switch {
	case temperature < 0:
		fmt.Println("Freezing!")
	case temperature < 15:
		fmt.Println("Cold")
	case temperature < 25:
		fmt.Println("Nice")
	default:
		fmt.Println("Hot!")
	}

	// FOR LOOP (Go's only loop construct)
	fmt.Println("\n=== For Loops ===")

	// Traditional for loop
	fmt.Print("Traditional: ")
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// While-style loop
	fmt.Print("While-style: ")
	count := 0
	for count < 5 {
		fmt.Print(count, " ")
		count++
	}
	fmt.Println()

	// Infinite loop with break
	fmt.Print("With break: ")
	num := 0
	for {
		if num >= 5 {
			break
		}
		fmt.Print(num, " ")
		num++
	}
	fmt.Println()

	// Continue statement
	fmt.Print("Skip evens: ")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// For-range loop (preview - more in arrays/slices)
	fmt.Println("\n=== For-Range ===")
	message := "Go!"
	for index, char := range message {
		fmt.Printf("Index %d: %c\n", index, char)
	}
}
