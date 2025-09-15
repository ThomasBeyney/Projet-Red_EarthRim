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
		Inventory: []string{"Potion", "Epée"},
	}
}

func displayInfo(c *Character) {
	fmt.Println("\n--- Infos du personnage ---")
	fmt.Println("Nom :", c.Name)
	fmt.Println("Classe :", c.Class)
	fmt.Println("Niveau :", c.Level)
	fmt.Println("Points de vie :", c.Hp, "/", c.MaxHp)
	fmt.Println("Inventaire :", c.Inventory)
	fmt.Println("--------------------------")
}
