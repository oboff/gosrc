// main
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	ns := []int64{}
	fmt.Println("Enter text:")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "" {
			break
		}
		i, err := strconv.ParseInt(input.Text(), 10, 64)
		if err != nil {
			break
		}
		ns = append(ns, i)
		fmt.Printf("%#v\n", ns)
	}

}
