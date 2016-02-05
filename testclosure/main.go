// main
package main

import (
	"fmt"
)

func Adder() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}

func main() {
	var f = Adder()
	fmt.Println("f(2):", f(2))
	fmt.Println("f(3)", f(3))
	fmt.Println("f(300)", f(300))
}
