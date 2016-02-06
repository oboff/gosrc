// main
package main

import (
	"fmt"
	"strings"
)

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	fmt.Println("Hello World!")
	addJpg := MakeAddSuffix(".jpg")
	fmt.Println(addJpg("my_file"))
	fmt.Println(addJpg("my_file2.bmp"))
}
