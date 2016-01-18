/*
Modify the previous program to use a func expression.
*/

// main
package main

import (
	"fmt"
)

func half(i int) (d int, b bool) {
	d = i / 2
	b = i%2 == 0
	return
}

func printme(i int, b bool) {
	fmt.Printf("%d, %t\n", i, b)
}

func main() {
	//fmt.Println("Hello World!")
	a, b := half(1)
	c, d := half(2)
	printme(a, b)
	printme(c, d)

}
