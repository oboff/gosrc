// main
package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Hello World!")
	b := big.NewInt(1)
	c := big.NewInt(10)
	for i := 0; i < 100000; i++ {
		d := new(big.Int)
		d.Mul(b, c)
		b = d
	}
	fmt.Printf("%v\n", b)

	a1 := big.NewInt(0)
	b1 := big.NewInt(0)

	fmt.Println(a1 == b1)
}
