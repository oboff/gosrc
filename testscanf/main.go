// main
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World!")
	var name string
	fmt.Println("Enter name: ")
	fmt.Scanf("%s", &name)
	fmt.Println(name)
	fmt.Fprintln(os.Stdout, "Boffill")
}
