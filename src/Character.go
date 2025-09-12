package main

import (
	"fmt"
)

type Character struct {
	Name      string
	Class     string
	Level     int
	MaxHp     int
	Hp        int
	Inventory []string
}

func initCharacter() Character {
	var name string
	fmt.Print("Quel est le nom de ton personnage ? ")
	fmt.Scanln(&name)

	return Character{
		Name:      name,
		Class:     "Elf",
		Level:     1,
		MaxHp:     100,
		Hp:        40,
		Inventory: []string{"Potion","Epée"},
	}
}

func displayInfo() {
	Player := initCharacter()
	fmt.Println("Vos Hp actuel sont :",Player.Hp)
	fmt.Println("Votre level actuel est :",Player.Level)
}

func main() {
	displayInfo()
}
