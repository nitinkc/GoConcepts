// Lesson 9: Pointers in Go
// Addresses, dereferencing, mutation, and production trade-offs
package main

import "fmt"

func Pointers() {
	pointerBasics()
	pointerZeroValues()
	pointersAsArguments()
	pointerToPointer()
	pointersAndArrays()
	pointersAndSlices()
	pointersAndStructs()
	pointerReceivers()
	pointerCollections()
	pointerInterfaceTrap()
	pointerProductionNotes()
}

func pointerBasics() {
	fmt.Println("=== 1. &, *T, and *p ===")

	x := 42
	p := &x

	fmt.Println("x value:", x)
	fmt.Println("&x address:", &x)
	fmt.Println("p stores address:", p)
	fmt.Println("*p reads value:", *p)

	*p = 100
	fmt.Println("x after *p = 100:", x)

	q := p
	*q = 200
	fmt.Println("x after changing copied pointer q:", x)
}

func pointerZeroValues() {
	fmt.Println("\n=== 2. nil and new ===")

	var p *int
	fmt.Println("zero-value pointer is nil:", p == nil)

	if p != nil {
		fmt.Println(*p)
	}

	allocated := new(int)
	fmt.Println("new(int) points to zero value:", *allocated)
	*allocated = 7
	fmt.Println("value after assignment:", *allocated)
}

func pointersAsArguments() {
	fmt.Println("\n=== 3. Function Arguments ===")

	n := 10
	doubleValue(n)
	fmt.Println("after doubleValue:", n)

	doublePointer(&n)
	fmt.Println("after doublePointer:", n)

	swap(&n, newInt(99))
	fmt.Println("after swapping n with 99:", n)
}

func pointerToPointer() {
	fmt.Println("\n=== 4. Pointer to Pointer ===")

	n := 5
	p := &n
	pp := &p

	fmt.Println("n, *p, **pp:", n, *p, **pp)
	**pp = 10
	fmt.Println("n after **pp = 10:", n)

	replacement := 50
	*pp = &replacement
	fmt.Println("p now points to replacement:", *p)
}

func pointersAndArrays() {
	fmt.Println("\n=== 5. Arrays and Pointers ===")

	values := [3]int{1, 2, 3}
	modifyArrayCopy(values)
	fmt.Println("after modifying array copy:", values)

	modifyArrayPointer(&values)
	fmt.Println("after modifying *[3]int:", values)

	first := &values[0]
	*first = 111
	fmt.Println("after changing pointer to one element:", values)
}

func pointersAndSlices() {
	fmt.Println("\n=== 6. Slices and Pointer-to-Slice ===")

	values := []int{1, 2, 3}
	modifySliceElement(values)
	fmt.Println("element mutation through []int:", values)

	appendWithoutReturning(values)
	fmt.Println("append without returning slice:", values)

	values = appendAndReturn(values, 4)
	fmt.Println("append and return:", values)

	appendThroughPointer(&values, 5)
	fmt.Println("append through *[]int:", values)
}

func pointersAndStructs() {
	fmt.Println("\n=== 7. Struct Pointers ===")

	account := &PointerAccount{Owner: "Alice", Balance: 100}
	fmt.Printf("struct value: %+v\n", *account)
	fmt.Println("field access auto-dereferences:", account.Owner)

	account.Balance += 50
	fmt.Printf("after field update: %+v\n", *account)

	copyOfAccount := *account
	copyOfAccount.Owner = "Copied Alice"
	fmt.Printf("original: %+v, copy: %+v\n", *account, copyOfAccount)
}

func pointerReceivers() {
	fmt.Println("\n=== 8. Method Receivers ===")

	account := PointerAccount{Owner: "Bob", Balance: 80}
	account.Deposit(20)
	fmt.Println("pointer receiver changed balance:", account.Balance)
	fmt.Println("value receiver summary:", account.Summary())
}

func pointerCollections() {
	fmt.Println("\n=== 9. Collections of Pointers ===")

	accounts := []PointerAccount{{Owner: "A"}, {Owner: "B"}}
	for i := range accounts {
		account := &accounts[i]
		account.Balance = (i + 1) * 100
	}
	fmt.Println("address slice elements by index:", accounts)

	byID := map[string]*PointerAccount{
		"primary": {Owner: "Carol", Balance: 300},
	}
	byID["primary"].Deposit(25)
	fmt.Println("pointer stored in map:", byID["primary"].Balance)
}

func pointerInterfaceTrap() {
	fmt.Println("\n=== 10. Typed nil in an Interface ===")

	var account *PointerAccount
	var value any = account

	fmt.Println("pointer is nil:", account == nil)
	fmt.Println("interface holding typed nil is nil:", value == nil)
}

func pointerProductionNotes() {
	fmt.Println("\n=== 11. Production Guidance ===")
	fmt.Println("Use pointers for mutation, optional values, identity, or costly copies.")
	fmt.Println("Prefer values for small immutable data and simpler ownership.")
	fmt.Println("Pointers shared by goroutines need synchronization.")
	fmt.Println("Go has no safe pointer arithmetic; unsafe is for specialized code only.")
}

func doubleValue(n int) {
	n *= 2
}

func doublePointer(n *int) {
	if n != nil {
		*n *= 2
	}
}

func newInt(value int) *int {
	return &value
}

func swap(a, b *int) {
	if a != nil && b != nil {
		*a, *b = *b, *a
	}
}

func modifyArrayCopy(values [3]int) {
	values[0] = 999
}

func modifyArrayPointer(values *[3]int) {
	if values != nil {
		values[0] = 999
	}
}

func modifySliceElement(values []int) {
	if len(values) > 0 {
		values[0] = 999
	}
}

func appendWithoutReturning(values []int) {
	values = append(values, 4)
	fmt.Println("local slice inside function:", values)
}

func appendAndReturn(values []int, number int) []int {
	return append(values, number)
}

func appendThroughPointer(values *[]int, number int) {
	if values != nil {
		*values = append(*values, number)
	}
}

type PointerAccount struct {
	Owner   string
	Balance int
}

func (a *PointerAccount) Deposit(amount int) {
	if a != nil {
		a.Balance += amount
	}
}

func (a PointerAccount) Summary() string {
	return fmt.Sprintf("%s has %d", a.Owner, a.Balance)
}
