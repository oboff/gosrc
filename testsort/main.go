// main
package main

import (
	"fmt"
	"reflect"
	"sort"
)

type Person struct {
	Last  string
	First string
	Age   int
}

type People []Person

func (p People) Len() int {
	return len(p)
}
func (p People) Less(i, j int) bool {
	if p[i].Last+p[i].First == p[j].Last+p[j].First {
		return p[i].Age < p[j].Age
	} else {
		return p[i].Last+p[i].First < p[j].Last+p[j].First
	}
}
func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func Sorter(i interface{}) {
	switch i.(type) {
	case []string:
		r, ok := i.([]string) //type assertion
		if !ok {
			break
		}
		sort.StringSlice(r).Sort()
	case []int:
		r, ok := i.([]int)
		if !ok {
			break
		}
		sort.Sort(sort.IntSlice(r))
	case People:
		r, ok := i.(People)
		if !ok {
			break
		}
		sort.Sort(sort.Interface(r))

	}
}

func main() {

	s := []string{"Eddy", "John", "Aldo", "Brenda", "Julia", "Cristina"}
	fmt.Println(s)
	//sort.StringSlice(s).Sort()
	Sorter(s)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))

	n := []int{23, 2, 17, 59, 3, 42}
	fmt.Println(n)
	//sort.StringSlice(s).Sort()
	Sorter(n)
	fmt.Println(n)
	fmt.Println(reflect.TypeOf(n))

	p := People{Person{"Smith", "John", 30},
		Person{"Johnson", "Michael", 40},
		Person{"Smith", "Julie", 12},
		Person{"Smith", "Julie", 20},
		Person{"Smith", "Julie", 100},
		Person{"Johnson", "Alex", 64}}
	fmt.Println(p)
	Sorter(p)
	fmt.Println(p)
	fmt.Printf("%v\n", reflect.TypeOf(p))
}
