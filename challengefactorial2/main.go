// main
package main

import (
	"fmt"
)

func decrementor(n int) chan int {
	out := make(chan int)
	go func() {
		for i := n; i > 0; i-- {
			out <- i
		}
		close(out)
	}()
	return out
}

func multiplier(in chan int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for n := range in {
			total *= n
		}
		out <- total
		close(out)
	}()

	return out
}

func main() {
	fmt.Println("Hello World!")
	for i := 100; i > 0; i-- {
		for m := range multiplier(decrementor(i % 10)) {
			fmt.Println(i, ":", m)
		}
	}

}
