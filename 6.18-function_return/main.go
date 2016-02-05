// main
package main

import (
	"fmt"
)

func main() {
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))
	fmt.Printf("Call Add2 for 5 gives: %v\n", p2(5))
}

// Add2 ...
func Add2() func(int) int {
	return func(b int) int {
		return b + 2
	}
}

// Adder ...
func Adder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
