package main

import "fmt"

type Character interface {
	Attack(c BaseCharacter)
	GetName() string
}

type BaseCharacter struct {
	Name string
	HP   int
	Dmg  int
}

type Goblin struct {
}

func (g Goblin) Spitball() {

}

func (p BaseCharacter) Attack(e BaseCharacter) {}
func (p BaseCharacter) GetName() string {
	return p.Name
}

func Game(p, e BaseCharacter) {
	p.Attack(e)
	fmt.Println(p.GetName())
	e.Attack(p)
	fmt.Println(e.GetName())
}

func main() {
	fmt.Println("Hello, playground")
	p := BaseCharacter{"Hero", 100, 10}
	e := Goblin{BaseCharacter{"Goblin", 50, 5}}
	Game(p, e)
}
