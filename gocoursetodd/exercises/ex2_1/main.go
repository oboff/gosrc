/*
Write a function which takes an integer. The function will have two returns.
The first return should be the argument divided by 2.
The second return should be a bool that let’s us know
whether or not the argument was even. For example:
half(1) returns (0, false)
half(2) returns (1, true)
*/

// main
package main

import (
	"fmt"
)

func half(i int) {
	d := i / 2
	b := i%2 == 0

	fmt.Printf("(%d, %t)\n", d, b)
}

func main() {
	fmt.Println("Hello World!")
	half(1)
	half(2)
}
