// main
package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter your name: ")
	var name string
	fmt.Scanf("%s", &name)
	fmt.Println("Hello, ", name)
}
