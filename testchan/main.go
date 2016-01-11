package main

import (
	"fmt"
	"math/big"
)

func fibonacci(c chan *big.Int) {
	// Initialize two big ints with the first two numbers in the sequence.
	a := big.NewInt(0)
	c <- a
	b := big.NewInt(1)
	c <- b

	// Initialize limit as 10^99, the smallest integer with 100 digits.
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(99), nil)

	// Loop while a is smaller than 1e100.
	for a.Cmp(&limit) < 0 {
		c <- a
		// Compute the next Fibonacci number, storing it in a.
		a.Add(a, b)
		// Swap a and b so that b is the next number in the sequence.
		a, b = b, a
	}
	close(c)
}

func main() {
	c := make(chan *big.Int, 1)
	go fibonacci(c)
	for x := range c {
		fmt.Println(x)
	}

}
