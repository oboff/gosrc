// main
package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 0)

	c := []int{}
	d := make([]int, len(a))

	b = append(b, a...)
	func(xs []int, n int) []int {
		for i, _ := range xs {
			xs[i] += n
		}
		return xs
	}(b, 2)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%#v\n", c)
	fmt.Printf("%#v\n", d)
}
