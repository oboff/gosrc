package main

import (
	"fmt"
	"strings"
)

func main() {
	my_str := "orlando\nis my name\ndon't you forget it"
	fmt.Printf("%q\n", my_str)
	fmt.Printf("%q", strings.Split(my_str, "\n"))
}
