// Lesson 10: Interfaces in Go
// Abstraction and polymorphism
package main

import (
	"fmt"
	"math"
)

// DEFINING INTERFACES
// An interface defines behavior (methods) without implementation

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Stringer interface (from fmt package)
type Stringer interface {
	String() string
}

// IMPLEMENTING INTERFACES
// Types implement interfaces implicitly - no "implements" keyword

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle(%.2f x %.2f)", r.Width, r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle(radius: %.2f)", c.Radius)
}

type Triangle struct {
	A, B, C float64 // Sides
}

func (t Triangle) Area() float64 {
	// Heron's formula
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

// FUNCTION USING INTERFACE
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func totalArea(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

// EMPTY INTERFACE
func printAnything(v interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", v, v)
}

// Using any (alias for interface{} in Go 1.18+)
func printAny(v any) {
	fmt.Printf("Value: %v, Type: %T\n", v, v)
}

func Interfaces() {
	// USING INTERFACES
	fmt.Println("=== Basic Interface Usage ===")

	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}
	triangle := Triangle{A: 3, B: 4, C: 5}

	fmt.Println(rect)
	printShapeInfo(rect)

	fmt.Println(circle)
	printShapeInfo(circle)

	fmt.Println("Triangle")
	printShapeInfo(triangle)

	// INTERFACE SLICE
	fmt.Println("\n=== Slice of Interfaces ===")
	shapes := []Shape{rect, circle, triangle}

	for i, shape := range shapes {
		fmt.Printf("Shape %d: Area = %.2f\n", i+1, shape.Area())
	}

	fmt.Printf("Total area: %.2f\n", totalArea(shapes))

	// TYPE ASSERTION
	fmt.Println("\n=== Type Assertion ===")

	var s Shape = rect

	// Assert and convert to concrete type
	r, ok := s.(Rectangle)
	if ok {
		fmt.Printf("It's a Rectangle with width %.2f\n", r.Width)
	}

	// This would panic if assertion fails without ok check
	// c := s.(Circle) // Would panic!

	// TYPE SWITCH
	fmt.Println("\n=== Type Switch ===")
	for _, shape := range shapes {
		switch v := shape.(type) {
		case Rectangle:
			fmt.Printf("Rectangle: %v\n", v)
		case Circle:
			fmt.Printf("Circle with radius %.2f\n", v.Radius)
		case Triangle:
			fmt.Printf("Triangle with sides %.2f, %.2f, %.2f\n", v.A, v.B, v.C)
		default:
			fmt.Println("Unknown shape")
		}
	}

	// EMPTY INTERFACE
	fmt.Println("\n=== Empty Interface ===")
	printAnything(42)
	printAnything("hello")
	printAnything(true)
	printAnything(rect)

	// Slice of any type
	mixed := []any{1, "two", 3.0, true, rect}
	for _, item := range mixed {
		printAny(item)
	}

	// INTERFACE COMPOSITION
	fmt.Println("\n=== Interface Composition ===")

	type Reader interface {
		Read() string
	}

	type Writer interface {
		Write(string)
	}

	// Composed interface
	type ReadWriter interface {
		Reader
		Writer
	}

	// CHECKING INTERFACE IMPLEMENTATION
	fmt.Println("\n=== Interface Check ===")

	var _ Shape = Rectangle{} // Compile-time check
	var _ Shape = Circle{}
	// var _ Shape = "string" // Would not compile - string doesn't implement Shape
}
