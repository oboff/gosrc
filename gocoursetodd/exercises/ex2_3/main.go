/*
Write a function with one variadic parameter that finds the greatest number in a list of numbers.
*/
// main
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func max(n []int64) int64 {
	var max int64
	var low int64
	low = 1<<63 - 1
	low = -1 * low
	for _, num := range n {
		if num > low {
			low = num
			max = num
		}
	}
	return max
}

func printMax(n int64) {
	fmt.Printf("The max number is: %d\n", n)
}

func main() {
	ns := []int64{}
	fmt.Print("Enter list of numbers to find max: ")
Main:
	for {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			if input.Text() == "" {
				break Main
			}
			i, err := strconv.ParseInt(input.Text(), 10, 64)
			if err != nil {
				printMax(max(ns))
				break Main
			}
			ns = append(ns, i)
			fmt.Printf("%#v\n", ns)
		}
	}
	printMax(max(ns))
}
