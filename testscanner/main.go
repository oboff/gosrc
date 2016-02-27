// main
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Int64SliceParse() (int64s []int64) {
	fmt.Println("[int]Enter text:")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "" {
			break
		}
		strs := strings.FieldsFunc(input.Text(), func(c rune) bool { return !unicode.IsDigit(c) })
		for _, str := range strs {
			p, err := strconv.ParseInt(str, 10, 64)
			if err == nil {
				int64s = append(int64s, p)
			}
			fmt.Println(p, err)
		}
		fmt.Printf("%#v\n", int64s)
	}
	return
}

func Float64SliceParse() (flt64s []float64) {
	fmt.Println("[float]Enter text:")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "" {
			break
		}
		strs := strings.FieldsFunc(input.Text(), func(c rune) bool { return !unicode.IsDigit(c) && '.' != c })
		for _, str := range strs {
			p, err := strconv.ParseFloat(str, 64)
			if err == nil {
				flt64s = append(flt64s, p)
			}
			fmt.Println(p, err)
		}
		fmt.Printf("%#v\n", flt64s)
	}
	return
}

func main() {
Main:
	for {
		fmt.Println("Parse floats('f') or ints('i')?")
		selection := bufio.NewScanner(os.Stdin)
		if selection.Scan() {
			switch selection.Text() {
			case "":
				break Main
			case "i":
				Int64SliceParse()
			case "f":
				Float64SliceParse()
			default:
				fmt.Println("Selection unknown, try again...")
			}
		}
	}

}
