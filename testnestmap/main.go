package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	mymap := make(map[string]map[string][]int)
	to := make(map[string][]int)
	to["stuff"] = []int{10, 20, 30}
	fmt.Println(to)
	mymap["bigger"] = to
	fmt.Println(mymap)
	fmt.Println(mymap["bigger"]["stuff"])
	mymap["bigger"]["stuff"] = append(mymap["bigger"]["stuff"], 50)
	fmt.Println(mymap["bigger"]["stuff"])
	mymap["bigger"]["stuff"] = append(mymap["bigger"]["stuff"], 50)
	fmt.Println(mymap["bigger"]["stuff"])
	fmt.Println(mymap)
}
