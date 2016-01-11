// main
package main

import (
	"fmt"
	"os"
)

type User struct {
	First, Last string
	ID          int
}

type Group struct {
	Members []*User
}

func NewUser(f string, l string, id int) *User {
	fmt.Fprintf(os.Stdout, "[New user created]\tID: %v, Name: %v %v\n", id, f, l)
	return &User{f, l, id}
}

func (g *Group) PrintGroup() {
	for _, user := range g.Members {
		fmt.Fprintf(os.Stdout, "{ID: %v}\n\tFirst Name: %v\n\tLast Name: %v\n", user.ID, user.First, user.Last)
	}
}

func main() {
	fmt.Println("Hello World!")
	//users := []*User{&User{"John", "Doe", 32}, &User{"Mary", "Jane", 21}}

	//	for _, u := range users {
	//		fmt.Printf("%v\n", *u)
	//	}

	group := &Group{Members: []*User{&User{"John", "Doe", 32}, &User{"Mary", "Jane", 21}}}
	group.PrintGroup()
	user := NewUser("Larry", "Paige", 95)
	group.Members = append(group.Members, user)
	fmt.Println(group.Members[2])

}
