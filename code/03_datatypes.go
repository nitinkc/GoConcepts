// Lesson 3: Data Types
package main

import "fmt"

func DataTypes() {
	demoInts()
	demoFloats()
	demoStrings()
	demoRunes()
	demoBools()
	demoTypeConversion()
}

func demoInts() {
	fmt.Println("=== Integers ===")
	var i int = 42        // Platform dependent (32 or 64 bit)
	var i8 int8 = 127     // -128 to 127
	var i16 int16 = 32767 // -32768 to 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807

	fmt.Printf("int: %d, int8: %d, int16: %d\n", i, i8, i16)
	fmt.Printf("int32: %d, int64: %d\n", i32, i64)

	// Unsigned integers (no negative values)
	var ui uint = 42
	var ui8 uint8 = 255 // 0 to 255 (also called byte)
	fmt.Printf("uint: %d, uint8/byte: %d\n", ui, ui8)
}

func demoFloats() {
	fmt.Println("\n=== Floating Point ===")
	var f32 float32 = 3.14
	var f64 float64 = 3.141592653589793
	fmt.Printf("float32: %f, float64: %.15f\n", f32, f64)
}

func demoStrings() {
	fmt.Println("\n=== Strings ===")
	var greeting string = "Hello, Go!"
	multiLine := `This is a
multi-line string
using backticks`

	fmt.Println(greeting)
	fmt.Println(multiLine)
	fmt.Println("String length:", len(greeting))

	// String concatenation
	first := "Hello"
	second := "World"
	combined := first + " " + second
	fmt.Println("Combined:", combined)
}

func demoRunes() {
	fmt.Println("=== Runes ===")
	var r rune = 'A' // rune is alias for int32
	var emoji rune = '😀'
	fmt.Printf("Rune 'A': %c (value: %d)\n", r, r)
	fmt.Printf("Emoji: %c (value: %d)\n", emoji, emoji)
}

func demoBools() {
	fmt.Println("\n=== Booleans ===")
	var isTrue bool = true
	var isFalse bool = false
	fmt.Printf("isTrue: %t, isFalse: %t\n", isTrue, isFalse)

	// Boolean operations
	fmt.Printf("AND: %t, OR: %t, NOT: %t\n", isTrue && isFalse, isTrue || isFalse, !isTrue)
}

func demoTypeConversion() {
	fmt.Println("=== Type Conversion ===")
	var intNum int = 42
	var floatNum float64 = float64(intNum) // Must be explicit
	var backToInt int = int(floatNum)
	fmt.Printf("int: %d -> float64: %f -> int: %d\n", intNum, floatNum, backToInt)
}
