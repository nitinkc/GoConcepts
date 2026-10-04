// Entry point — run a lesson by name or number.
//
//	go run . basics        (or: go run . 1)
//	go run . scan          (interactive stdin demo)
//	go run .               (list all lessons)
package main

import (
	"fmt"
	"os"
)

var lessons = map[string]func(){
	"1": Basics, "basics": Basics,
	"2": Variables, "variables": Variables,
	"scan": ScanInput,
	"3": DataTypes, "datatypes": DataTypes,
	"4": ControlFlow, "controlflow": ControlFlow,
	"5": Functions, "functions": Functions,
	"6": ArraysSlices, "slices": ArraysSlices,
	"7": Maps, "maps": Maps,
	"8": Structs, "structs": Structs,
	"9": Pointers, "pointers": Pointers,
	"10": Interfaces, "interfaces": Interfaces,
	"11": Errors, "errors": Errors,
	"12": Concurrency, "concurrency": Concurrency,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <lesson>")
		fmt.Println("Lessons:")
		for k := range lessons {
			fmt.Println(" ", k)
		}
		return
	}
	if run, ok := lessons[os.Args[1]]; ok {
		run()
	} else {
		fmt.Println("Unknown lesson:", os.Args[1])
	}
}
