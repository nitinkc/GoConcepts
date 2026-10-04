# Lesson 9: Pointers

## Big Picture

```mermaid
flowchart LR
    subgraph Stack/Memory
        X["x (int)<br/>value: 42<br/>addr: 0x1040"]
    end
    P["p (*int) = &x"] -->|"stores 0x1040"| X
    DP["*p = 100"] -->|"dereference & write"| X
    N["var p *int"] -->|"zero value"| NIL[("nil")]
```

`&` takes an address, `*` dereferences it. Passing a pointer lets a function mutate the caller's value; passing a struct pointer avoids a copy. A `nil` pointer dereference panics.

## Concepts

### Pointer Basics
```go
x := 42
p := &x       // Get address
fmt.Println(*p) // Dereference (42)
*p = 100      // Modify through pointer
```

### Zero Value
```go
var p *int    // nil
p = new(int)  // Allocates, returns pointer
```

### Functions
```go
// Pass by reference
func modify(p *int) {
    *p = 100  // Changes original
}

num := 42
modify(&num)  // num is now 100
```

### Pointer to a Slice

```go
func addNumber(slice *[]int, num int) {
    *slice = append(*slice, num)
}

numbers := []int{1, 2}
addNumber(&numbers, 3) // numbers is now [1 2 3]
```

In `slice *[]int`, `[]int` is a slice of integers and the leading `*` makes the parameter a pointer to that slice. The caller passes its slice's address with `&numbers`, and `*slice` dereferences the pointer to access or replace the caller's slice value.

A slice is a small descriptor containing a pointer to an underlying array, a length, and a capacity. `append` returns an updated descriptor and may allocate a new underlying array, so its result must be assigned:

```go
*slice = append(*slice, num) // Correct: updates the caller's slice
_ = append(*slice, num)     // Result discarded: caller's length is unchanged
```

The blank identifier `_` explicitly discards a value. Although the second line compiles, it is not a valid way to add an element to the caller's slice because the new slice descriptor is lost. An often simpler and more idiomatic alternative is to return the updated slice:

```go
func addNumber(slice []int, num int) []int {
    return append(slice, num)
}

numbers = addNumber(numbers, 3)
```

### Structs
```go
person := &Person{Name: "Alice"}
person.Name = "Bob"  // Auto-dereference
```

### When to Use
- Modify function parameters
- Avoid copying large structs
- Represent "no value" (nil)

## Code Walkthrough

=== "09_pointers.go"

    ```go
    --8<-- "code/09_pointers.go"
    ```

!!! example "Try it yourself"
    ```bash
    cd code
    go run . pointers
    ```

!!! tip "Exercises"
    1. Write `swap(a, b *int)` and verify both caller variables change.
    2. Print `&x` twice — does a variable's address stay stable?
    3. Dereference a nil pointer inside a controlled run — observe the panic.

## Check Your Understanding

<quiz>
What does `*p` mean when `p` is a `*int`?
- [ ] Address of `p`
- [x] Value stored at the address `p` points to
- [ ] Multiplication
- [ ] Pointer type declaration only
</quiz>

<quiz>
Go is pass-by-... what?
- [x] Pass-by-value (always — pointers are copied too)
- [ ] Pass-by-reference
- [ ] Pass-by-name
- [ ] Depends on the type
</quiz>
