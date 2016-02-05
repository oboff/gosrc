package main

import "fmt"

func main() {
	defer println()
	var f = Adder()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
}

func Adder() func(int) int {
	var x int
	return func(b int) int {
		x += b
		return x
	}
}
