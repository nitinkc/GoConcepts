# Final Knowledge Check

A mixed quiz covering the whole course. Aim for 100% before calling yourself done. :material-trophy:

## Section 1 — Language Core

<quiz>
Which is a valid short variable declaration?
- [ ] `var x := 5`
- [x] `x := 5`
- [ ] `let x = 5`
- [ ] `int x = 5`
</quiz>

<quiz>
What is the zero value of a `string`?
- [ ] `nil`
- [ ] `null`
- [x] `""`
- [ ] `0`
</quiz>

<quiz>
Does Go support `while` loops?
- [ ] Yes, `while cond {}`
- [x] No — `for cond {}` is the while-equivalent
- [ ] Only inside `main`
- [ ] Only with `goto`
</quiz>

## Section 2 — Data Structures

<quiz>
What does `make([]int, 3, 10)` create?
- [ ] Length 10, capacity 3
- [x] Length 3, capacity 10
- [ ] A compile error
- [ ] A 3×10 matrix
</quiz>

<quiz>
Writing to a `nil` map causes a…
- [ ] Silent no-op
- [x] Panic
- [ ] Compile error
- [ ] Automatic allocation
</quiz>

<quiz>
To mutate a struct inside a method, use a…
- [ ] value receiver
- [x] pointer receiver
- [ ] global variable
- [ ] `mut` keyword
</quiz>

## Section 3 — Advanced

<quiz>
Go interfaces are implemented…
- [ ] With an `implements` clause
- [x] Implicitly, by satisfying the method set
- [ ] Via inheritance
- [ ] Through generics
</quiz>

<quiz>
How should you return a failure from a function?
- [ ] `throw new Error(...)`
- [ ] `panic` for everything
- [x] Return `(value, error)` with `error` as the last return
- [ ] Set a global `errno`
</quiz>

<quiz>
What is the Go proverb about shared memory?
- [ ] "Lock everything, always"
- [x] "Don't communicate by sharing memory; share memory by communicating"
- [ ] "Goroutines are threads"
- [ ] "Channels are optional"
</quiz>
