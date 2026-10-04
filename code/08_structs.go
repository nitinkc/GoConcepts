// Lesson 8: Structs
package main

import "fmt"

// Define a struct type
type Person struct {
	Name    string
	Age     int
	Email   string
	IsAdmin bool
}

// Struct with embedded struct
type Address struct {
	Street  string
	City    string
	Country string
	ZipCode string
}

type Employee struct {
	Person  // Embedded struct (composition)
	ID      string
	Salary  float64
	Address Address // Named field
}

// Method on struct - value receiver
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I'm %s and I'm %d years old", p.Name, p.Age)
}

// Method with pointer receiver (can modify struct)
func (p *Person) Birthday() {
	p.Age++
}

// Constructor function (Go convention)
func NewPerson(name string, age int, email string) *Person {
	return &Person{
		Name:  name,
		Age:   age,
		Email: email,
	}
}

// Method on Employee
func (e Employee) GetFullInfo() string {
	return fmt.Sprintf("%s (ID: %s) - %s, %s",
		e.Name, e.ID, e.Address.City, e.Address.Country)
}

func Structs() {
	// CREATING STRUCTS
	fmt.Println("=== Creating Structs ===")

	// Method 1: Zero value
	var person1 Person
	fmt.Printf("Zero value: %+v\n", person1)

	// Method 2: Struct literal
	person2 := Person{
		Name:  "Alice",
		Age:   25,
		Email: "alice@example.com",
	}
	fmt.Printf("Literal: %+v\n", person2)

	// Method 3: Positional (not recommended)
	person3 := Person{"Bob", 30, "bob@example.com", false}
	fmt.Printf("Positional: %+v\n", person3)

	// Method 4: Using constructor
	person4 := NewPerson("Charlie", 35, "charlie@example.com")
	fmt.Printf("Constructor: %+v\n", *person4)

	// ACCESSING FIELDS
	fmt.Println("\n=== Accessing Fields ===")
	fmt.Println("Name:", person2.Name)
	fmt.Println("Age:", person2.Age)

	// Modify fields
	person2.Age = 26
	fmt.Println("Updated age:", person2.Age)

	// METHODS
	fmt.Println("\n=== Methods ===")
	fmt.Println(person2.Greet())

	// Pointer receiver method
	fmt.Println("Age before birthday:", person2.Age)
	person2.Birthday() // Go automatically takes pointer
	fmt.Println("Age after birthday:", person2.Age)

	// EMBEDDED STRUCTS
	fmt.Println("\n=== Embedded Structs ===")
	emp := Employee{
		Person: Person{
			Name:  "David",
			Age:   40,
			Email: "david@company.com",
		},
		ID:     "EMP001",
		Salary: 75000,
		Address: Address{
			Street:  "123 Main St",
			City:    "San Francisco",
			Country: "USA",
			ZipCode: "94102",
		},
	}

	// Access embedded fields directly
	fmt.Println("Employee name:", emp.Name) // Promoted from Person
	fmt.Println("Employee city:", emp.Address.City)
	fmt.Println("Full info:", emp.GetFullInfo())

	// Call promoted method
	fmt.Println(emp.Greet())

	// COMPARING STRUCTS
	fmt.Println("\n=== Comparing Structs ===")
	p1 := Person{Name: "Test", Age: 20}
	p2 := Person{Name: "Test", Age: 20}
	p3 := Person{Name: "Test", Age: 21}

	fmt.Println("p1 == p2:", p1 == p2) // true
	fmt.Println("p1 == p3:", p1 == p3) // false

	// ANONYMOUS STRUCTS
	fmt.Println("\n=== Anonymous Structs ===")
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Printf("Config: %+v\n", config)

	// SLICE OF STRUCTS
	fmt.Println("\n=== Slice of Structs ===")
	team := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}

	for _, member := range team {
		fmt.Printf("- %s (%d)\n", member.Name, member.Age)
	}
}
