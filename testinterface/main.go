// main
package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

func printString(value interface{}) {
	switch str := value.(type) {
	case string:
		fmt.Println(str)
	case Stringer:
		fmt.Println(str.String())
	}
}

type Person struct {
	name string
	age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("Name: %v, Age: %v", p.name, p.age)
}

func main() {
	fmt.Println("Hello World!")
	j := &Person{name: "John", age: 24}
	fmt.Println(j.String())
	fmt.Printf("%T\t%[1]t\n", j)
	printString("C'est ci n'est pas un string")
	printString(j)
}
