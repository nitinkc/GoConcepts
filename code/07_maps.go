// Lesson 7: Maps in Go
// Key-value data structures
package main

import "fmt"

func Maps() {
	// CREATING MAPS
	fmt.Println("=== Creating Maps ===")

	// Using make
	ages := make(map[string]int)
	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 35
	fmt.Println("Ages:", ages)

	// Map literal
	capitals := map[string]string{
		"USA":    "Washington D.C.",
		"France": "Paris",
		"Japan":  "Tokyo",
	}
	fmt.Println("Capitals:", capitals)

	// ACCESSING VALUES
	fmt.Println("\n=== Accessing Values ===")
	fmt.Println("Alice's age:", ages["Alice"])
	fmt.Println("Capital of Japan:", capitals["Japan"])

	// Check if key exists (comma ok idiom)
	age, exists := ages["David"]
	if exists {
		fmt.Println("David's age:", age)
	} else {
		fmt.Println("David not found")
	}

	// Shorter version
	if capital, ok := capitals["Germany"]; ok {
		fmt.Println("Capital of Germany:", capital)
	} else {
		fmt.Println("Germany not in map")
	}

	// MODIFYING MAPS
	fmt.Println("\n=== Modifying Maps ===")

	// Add new key
	ages["David"] = 28
	fmt.Println("After adding David:", ages)

	// Update existing key
	ages["Alice"] = 26
	fmt.Println("After updating Alice:", ages)

	// Delete key
	delete(ages, "Charlie")
	fmt.Println("After deleting Charlie:", ages)

	// ITERATING
	fmt.Println("\n=== Iteration ===")
	fmt.Println("All capitals:")
	for country, capital := range capitals {
		fmt.Printf("  %s: %s\n", country, capital)
	}

	// Just keys
	fmt.Print("Countries: ")
	for country := range capitals {
		fmt.Print(country, " ")
	}
	fmt.Println()

	// LENGTH
	fmt.Println("\n=== Map Length ===")
	fmt.Println("Number of people:", len(ages))
	fmt.Println("Number of countries:", len(capitals))

	// MAPS WITH STRUCT VALUES
	fmt.Println("\n=== Maps with Structs ===")
	type Person struct {
		Name string
		Age  int
		City string
	}

	people := map[string]Person{
		"emp001": {Name: "Alice", Age: 25, City: "NYC"},
		"emp002": {Name: "Bob", Age: 30, City: "LA"},
	}
	fmt.Println("Employee emp001:", people["emp001"])
	fmt.Println("emp001's city:", people["emp001"].City)

	// NESTED MAPS
	fmt.Println("\n=== Nested Maps ===")
	users := map[string]map[string]string{
		"user1": {
			"name":  "Alice",
			"email": "alice@example.com",
		},
		"user2": {
			"name":  "Bob",
			"email": "bob@example.com",
		},
	}
	fmt.Println("User1 email:", users["user1"]["email"])

	// Adding to nested map
	users["user3"] = map[string]string{
		"name":  "Charlie",
		"email": "charlie@example.com",
	}
	fmt.Println("Users:", users)

	// MAP AS SET
	fmt.Println("\n=== Map as Set ===")
	uniqueWords := make(map[string]bool)
	words := []string{"apple", "banana", "apple", "cherry", "banana"}

	for _, word := range words {
		uniqueWords[word] = true
	}

	fmt.Print("Unique words: ")
	for word := range uniqueWords {
		fmt.Print(word, " ")
	}
	fmt.Println()
}
