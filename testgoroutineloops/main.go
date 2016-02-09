package main

import (
	"fmt"
)

type Collection interface {
	collection()
}

type IntSlice []int
type StringSlice []string

func (a IntSlice) collection()    {}
func (b StringSlice) collection() {}

func main() {
	a := IntSlice{10, 9, 3}
	b := StringSlice{"john", "bob", "alex"}
	it := func(iface interface{}) {
		switch iface := iface.(type) {
		case IntSlice:
			for _, x := range iface {
				fmt.Println(x)
			}
		case StringSlice:
			for _, x := range iface {
				fmt.Println(x)
			}
		default:
			fmt.Println("not a valid type")
		}
	}
	go it(a)
	go it(b)
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
