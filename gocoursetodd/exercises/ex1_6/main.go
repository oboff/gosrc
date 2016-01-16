// main
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		switch {
		case (i%3 == 0) && (i%5 == 0):
			fmt.Printf("%d : FizzBuzz\n", i)
		case i%3 == 0:
			fmt.Printf("%d : Fizz\n", i)
		case i%5 == 0:
			fmt.Printf("%d : Buzz\n", i)
		default:
			fmt.Println(i, ":")
		}
	}

}
