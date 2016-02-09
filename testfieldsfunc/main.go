package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your string:")
	line, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	f := func(c rune) bool {
		b := !unicode.IsNumber(c) && '.' != c //!unicode.IsPunct(c)
		return b
	}
	ff := strings.FieldsFunc(line, f)
	fmt.Printf("ff: %v\n", ff)
	fff := []float64{}
	for _, n := range ff {
		x, err := strconv.ParseFloat(n, 64)
		if err == nil {
			fff = append(fff, x)
		}
	}
	fmt.Printf("fff: %v\n", fff)

}
