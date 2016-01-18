/*
Write a function, foo, which can be called in all of these ways:
func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}
*/

// main
package main

import (
	"fmt"
)

func foo(n ...int) {
	for _, i := range n {
		fmt.Printf("%d, ", i)
	}
	fmt.Println()
}
func main() {
	fmt.Println("Hello World!")
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}
