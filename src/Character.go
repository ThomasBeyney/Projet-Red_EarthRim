package main

import "fmt"

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
		Inventory: []string{},
	}
}

func main() {
	Player := initCharacter()
	fmt.Println("Bienvenue", Player.Name,"tu a ",Player.Hp,"/",Player.MaxHp,"hp")
}
