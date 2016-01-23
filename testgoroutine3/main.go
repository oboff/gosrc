// main
package main

import (
	"fmt"
)

func foobar(s string) {
	for i := 0; i < 100; i++ {
		fmt.Println(s, ":", i)
	}
}

func main() {
	go foobar("foo")
	go foobar("bar")
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
