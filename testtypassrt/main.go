// main
package main

import (
	"fmt"
	"time"
)

func toMapStrIface(y interface{}) {
	z, ok := y.(map[string]interface{})
	if ok {
		z["updated_at"] = time.Now()
	}
}

type Writer interface {
	Write()
}

type IntWriter int

func (i IntWriter) Write() {
	fmt.Println("Hi, I'm a Writer")
}

type Animal interface {
	Move()
}

type Legs int

type Dog struct {
	legs Legs
}

func (d *Dog) Move() {
	if d.legs == 4 {
		fmt.Println("Yes, this is dog.")
	}
}

func NewDog() *Dog {
	return &Dog{legs: 4}

}

func main() {
	fmt.Println("Hello World!")
	foo := map[string]interface{}{
		"Jo": 42,
	}
	toMapStrIface(foo)
	fmt.Println(foo)

	var w Writer
	w = new(IntWriter)
	w.Write()

	var a Animal
	a = NewDog()
	a.Move()

}
