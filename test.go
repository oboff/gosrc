package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	in := bufio.NewReader(os.Stdin)

	line, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	sl := strings.Split(line, " ")

	fmt.Printf("%q\n", sl)
	fmt.Printf("%q\n", line)

	for _, l := range sl[:len(sl)-1] {
		fmt.Println(l)
		x, err := strconv.ParseFloat(l, 64)
		if err != nil {
			fmt.Println(x)
		}
	}

	a, _ := strconv.ParseFloat("23.23", 64)
	fmt.Printf("%f\n", a)
}
