// main
package main

import (
	"fmt"
)

func init() {
	go func() {
		for i := 0; i < 45; i++ {
			fmt.Println("1: ", i)
		}
	}()
	go func() {
		for i := 100; i < 145; i++ {
			fmt.Println("2: ", i)
		}
	}()
	go func() {
		for i := 0; i < 45; i++ {
			fmt.Println("3: ", i)
		}
	}()
	go func() {
		for i := 100; i < 145; i++ {
			fmt.Println("4: ", i)
		}
	}()
}

func main() {
	fmt.Println("Hello World!")
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
