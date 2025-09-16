package main

import "fmt"

type Character struct {
	Name      string
	Class     string
	Level     int
	Skill     []string
	MaxHp     int
	Hp        int
	Inventory []string
	Equipement []string
	Gold      int
}

func characterCreation() Character {
	var name string
	valid := false

	for !valid {
		fmt.Print("Quel est le nom de ton personnage ? ")
		fmt.Scanln(&name)

		valid = true
		for i := 0; i < len(name); i++ {
			if !(name[i] >= 'A' && name[i] <= 'Z') && !(name[i] >= 'a' && name[i] <= 'z') {
				fmt.Println("Le nom ne doit contenir que des lettres.")
				valid = false
				break
			}
		}
	}

	formattedName := ""
	if len(name) > 0 {
		first := name[0]
		if first >= 'a' && first <= 'z' {
			first = first - 32
		}
		formattedName += string(first)

		for i := 1; i < len(name); i++ {
			ch := name[i]
			if ch >= 'A' && ch <= 'Z' {
				ch = ch + 32
			}
			formattedName += string(ch)
		}
	}

	return initCharacter(formattedName)
}

func initCharacter(name string) Character {
	return Character{
		Name:      name,
		Class:     "Elf",
		Level:     1,
		Skill:     []string{"Coup de poing"},
		MaxHp:     100,
		Hp:        MaxHp/2,
		Inventory: []string{"Potion", "Epée"},
		Equipement: []string{},
		Gold:      100,
	}
}

func displayInfo(c *Character) {
	fmt.Println("\n--- Infos du personnage ---")
	fmt.Println("Nom :", c.Name)
	fmt.Println("Classe :", c.Class)
	fmt.Println("Niveau :", c.Level)
	fmt.Println("Sort actuel :", c.Skill)
	fmt.Println("Points de vie :", c.Hp, "/", c.MaxHp)
	fmt.Println("Inventaire :", c.Inventory)
	fmt.Println("Gold :", c.Gold)
	fmt.Println("--------------------------")
}
