// Lesson 5: Functions
package main

import "fmt"

// Basic function - no parameters, no return
func sayHello() {
	fmt.Println("Hello!")
}

// Function with parameters, return void
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Function with return value
func add(a int, b int) int {
	return a + b
}

// Multiple parameters of same type - shorthand
func multiply(a, b int) int {
	return a * b
}

// Multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Named return values
func rectangle(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return // naked return - returns named values
}

// Variadic function - accepts any number of arguments
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Function as a value
var mathOp func(int, int) int

// Higher-order function - takes a function as parameter
func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// Function returning a function (closure)
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func Functions() {
	fmt.Println("=== Basic Functions ===")
	sayHello()
	greet("Gopher")

	fmt.Println("\n=== Functions with Return ===")
	result := add(5, 3)
	fmt.Println("5 + 3 =", result)
	fmt.Println("4 * 6 =", multiply(4, 6))

	fmt.Println("\n=== Multiple Returns ===")
	quotient, err := divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", quotient)
	}

	// Ignoring a return value with _
	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\n=== Named Returns ===")
	area, perimeter := rectangle(5, 3)
	fmt.Printf("Rectangle 5x3: Area=%.2f, Perimeter=%.2f\n", area, perimeter)

	fmt.Println("\n=== Variadic Functions ===")
	fmt.Println("Sum(1,2,3):", sum(1, 2, 3))
	fmt.Println("Sum(1,2,3,4,5):", sum(1, 2, 3, 4, 5))

	// Pass slice to variadic using ...
	nums := []int{10, 20, 30}
	fmt.Println("Sum(slice):", sum(nums...))

	fmt.Println("\n=== Functions as Values ===")
	mathOp = add
	fmt.Println("Using mathOp (add):", mathOp(10, 5))
	mathOp = multiply
	fmt.Println("Using mathOp (multiply):", mathOp(10, 5))

	fmt.Println("\n=== Higher-Order Functions ===")
	fmt.Println("Apply add:", applyOperation(10, 5, add))
	fmt.Println("Apply multiply:", applyOperation(10, 5, multiply))

	// Anonymous function
	fmt.Println("Apply anonymous:", applyOperation(10, 5, func(a, b int) int {
		return a - b
	}))

	fmt.Println("\n=== Closures ===")
	double := multiplier(2)
	triple := multiplier(3)
	fmt.Println("Double 5:", double(5))
	fmt.Println("Triple 5:", triple(5))
}
