// main
package main

import (
	"fmt"
)

type Attacker interface {
	Attack(m Monster)
}

type Defender interface {
	Defend(m Monster)
}

type Weapon struct {
	Name string
	Dmg  int
}

type Armor struct {
	Name string
	Def  int
}
type Equip struct {
	Weapons []Weapon
	Armors  []Armor
	//Inventory []Items
}

type Monster struct {
	Name          string
	Str, Dex, Con int
	HP            int
	AC            int
	Equip
}

func (m Monster) Attack(n Monster) {

}

func (m Monster) Defend(n Monster) {

}

func (m Monster) GetHP() int {
	return m.Con
}

func (m Monster) GetAC() int {
	return m.Armors[0].Def
}

func main() {
	p := Monster{"Hero", 10, 10, 10, 10, 10,
		Equip{[]Weapon{Weapon{"Sword", 10}},
			[]Armor{Armor{"Plate", 4}}}}

	fmt.Println(p)
}
