// Lesson 9: Pointers in Go
// Addresses, dereferencing, mutation, and common gotchas
package main

import "fmt"

func Pointers() {
	// BASICS: & takes an address, * dereferences it
	fmt.Println("=== Basics: &, *T, *p ===")
	x := 42
	p := &x // p is *int, holds address of x
	fmt.Println("p:", p, "| *p:", *p)
	*p = 100 // write through the pointer
	fmt.Println("x after *p = 100:", x)

	// NIL: zero value of any pointer is nil
	var nilP *int
	fmt.Println("nil pointer?", nilP == nil) // *nilP would panic
	fresh := new(int)                        // allocates, returns *int (0)
	fmt.Println("*new(int):", *fresh)

	// FUNCTIONS: Go is always pass-by-value
	fmt.Println("\n=== Functions ===")
	n := 10
	doubleValue(n)
	fmt.Println("after doubleValue:", n) // unchanged
	doublePointer(&n)
	fmt.Println("after doublePointer:", n) // 20

	// POINTER TO POINTER
	fmt.Println("\n=== **T ===")
	pp := &p // p still points at x
	**pp = 55
	fmt.Println("x via **pp:", **pp, "| x:", x)

	// ARRAYS: copied whole; use *[N]T to mutate in place
	fmt.Println("\n=== Arrays ===")
	arr := [3]int{1, 2, 3}
	modifyArrayCopy(arr)
	modifyArrayPtr(&arr)
	fmt.Println("arr after *[3]int:", arr)

	// SLICES: elements shared, but append needs the result back
	fmt.Println("\n=== Slices ===")
	s := []int{1, 2, 3}
	setFirst(s)
	fmt.Println("element write through []int:", s)
	appendLost(s)
	fmt.Println("append not returned:", s) // unchanged
	s = appendAndReturn(s, 4)
	fmt.Println("append returned:", s)
	appendViaPtr(&s, 5)
	fmt.Println("append via *[]int:", s)

	// STRUCTS: field access auto-dereferences
	fmt.Println("\n=== Structs & Receivers ===")
	acct := &PointerAccount{Owner: "Alice", Balance: 100}
	acct.Deposit(50) // pointer receiver mutates
	fmt.Printf("%+v | %s\n", *acct, acct.Summary())

	// GOTCHA: typed nil inside an interface
	fmt.Println("\n=== Typed nil in interface ===")
	var ta *PointerAccount
	var v any = ta
	fmt.Println("ptr nil:", ta == nil, "| iface nil:", v == nil) // false!
}

func doubleValue(n int)                    { n *= 2 } // copy: no effect
func doublePointer(n *int)                 { *n *= 2 }
func modifyArrayCopy(a [3]int)             { a[0] = 999 }
func modifyArrayPtr(a *[3]int)             { a[0] = 999 }
func setFirst(s []int)                     { s[0] = 999 }
func appendLost(s []int)                   { s = append(s, 4) }
func appendAndReturn(s []int, n int) []int { return append(s, n) }
func appendViaPtr(s *[]int, n int)         { *s = append(*s, n) }

type PointerAccount struct {
	Owner   string
	Balance int
}

func (a *PointerAccount) Deposit(n int) { a.Balance += n }
func (a PointerAccount) Summary() string {
	return fmt.Sprintf("%s has %d", a.Owner, a.Balance)
}
