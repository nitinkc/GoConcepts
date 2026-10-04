// Lesson 11: Error Handling in Go
// Go's approach to handling errors
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// CUSTOM ERROR TYPES

// Simple error using errors.New
var ErrDivisionByZero = errors.New("division by zero")
var ErrNegativeNumber = errors.New("negative number not allowed")

// Custom error type
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on %s: %s", e.Field, e.Message)
}

// Error with additional context
type HTTPError struct {
	StatusCode int
	Message    string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Message)
}

// FUNCTIONS THAT RETURN ERRORS

func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, ErrNegativeNumber
	}
	// Simple approximation
	if n == 0 {
		return 0, nil
	}
	x := n
	for i := 0; i < 10; i++ {
		x = (x + n/x) / 2
	}
	return x, nil
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{
			Field:   "age",
			Message: "must be positive",
		}
	}
	if age > 150 {
		return &ValidationError{
			Field:   "age",
			Message: "unrealistic value",
		}
	}
	return nil
}

func fetchUser(id int) (string, error) {
	if id <= 0 {
		return "", &HTTPError{StatusCode: 400, Message: "invalid user ID"}
	}
	if id > 1000 {
		return "", &HTTPError{StatusCode: 404, Message: "user not found"}
	}
	return fmt.Sprintf("User%d", id), nil
}

// WRAPPING ERRORS (Go 1.13+)

func readConfig(filename string) (string, error) {
	_, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read config file: %w", err)
	}
	return "config data", nil
}

func Errors() {
	// BASIC ERROR HANDLING
	fmt.Println("=== Basic Error Handling ===")

	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	result, err = safeDivide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// CHECKING SPECIFIC ERRORS
	fmt.Println("\n=== Checking Specific Errors ===")

	_, err = sqrt(-4)
	if err == ErrNegativeNumber {
		fmt.Println("Cannot calculate sqrt of negative number")
	}

	// Using errors.Is (Go 1.13+)
	_, err = safeDivide(5, 0)
	if errors.Is(err, ErrDivisionByZero) {
		fmt.Println("Caught division by zero!")
	}

	// CUSTOM ERROR TYPES
	fmt.Println("\n=== Custom Error Types ===")

	err = validateAge(-5)
	if err != nil {
		fmt.Println(err)

		// Type assertion to get more info
		if ve, ok := err.(*ValidationError); ok {
			fmt.Printf("  Field: %s\n", ve.Field)
		}
	}

	err = validateAge(200)
	if err != nil {
		fmt.Println(err)
	}

	// Using errors.As (Go 1.13+)
	_, err = fetchUser(9999)
	var httpErr *HTTPError
	if errors.As(err, &httpErr) {
		fmt.Printf("HTTP Error - Status: %d, Message: %s\n",
			httpErr.StatusCode, httpErr.Message)
	}

	// MULTIPLE ERROR CHECKS
	fmt.Println("\n=== Multiple Operations ===")

	// Common pattern: check error immediately
	num, err := strconv.Atoi("42")
	if err != nil {
		fmt.Println("Failed to parse number")
	} else {
		fmt.Println("Parsed number:", num)
	}

	// Handle invalid input
	_, err = strconv.Atoi("not a number")
	if err != nil {
		fmt.Println("Parse error:", err)
	}

	// ERROR WRAPPING
	fmt.Println("\n=== Error Wrapping ===")

	_, err = readConfig("nonexistent.json")
	if err != nil {
		fmt.Println("Error:", err)

		// Unwrap to check original error
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("  (File does not exist)")
		}
	}

	// DEFER, PANIC, RECOVER
	fmt.Println("\n=== Panic and Recover ===")

	// Recover from panic
	result2 := safeOperation()
	fmt.Println("Safe operation result:", result2)

	// IGNORING ERRORS (not recommended, but sometimes done)
	fmt.Println("\n=== Ignoring Errors ===")
	// Use blank identifier to explicitly ignore
	_, _ = safeDivide(10, 0) // Explicitly ignored
	fmt.Println("Continued after ignored error")
}

// Using defer and recover for panic handling
// A deferred function runs when its surrounding function is about to return.
// Multiple deferred functions run in LIFO order: the last one deferred runs first.
func safeOperation() (result string) {
	defer func() {
		if r := recover(); r != nil {
			result = fmt.Sprintf("Recovered from: %v", r)
		}
	}()

	// This would normally crash the program
	panic("something went wrong!")
}

// Example of panic (for truly exceptional cases)
func mustParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("mustParseInt: %v", err))
	}
	return n
}
