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
	InventoryMax int
	limInventory int
	Equipement Equipement
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

	classe, PVmax, PVactuel := choixClasse()
	return initCharacter(formattedName, classe, PVmax, PVactuel)
}

func choixClasse() (string, int, int) {
	var Classe string
	var PVmax int
	var PVactuel int
	var rep int
	valid := false

	for !valid {
		fmt.Println("Quelle classe choisissez vous ?")
		fmt.Println("1- Humain")
		fmt.Println("2- Elf")
		fmt.Println("3- Nain")
		fmt.Print("Choix : ")
		fmt.Scanln(&rep)

		switch rep {
		case 1:
			Classe = "Humain"
			PVmax = 100
			valid = true
		case 2:
			Classe = "Elf"
			PVmax = 80
			valid = true
		case 3:
			Classe = "Nain"
			PVmax = 120
			valid = true
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
	PVactuel = PVmax / 2
	return Classe, PVmax, PVactuel
}

func initCharacter(name string, Classe string, PVmax int, PVactuel int) Character {
	return Character{
		Name:      name,
		Class:     Classe,
		Level:     1,
		Skill:     []string{"Coup de poing"},
		MaxHp:     PVmax,
		Hp:        PVactuel,
		Inventory: []string{"Potion", "Epée"},
		InventoryMax: 10,
		limInventory: 0,
		Equipement: Equipement{
			Tete: "",
			Torse: "",
			Pieds: "",
		},
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
	fmt.Println("Equipement ;",c.Equipement, ",")
	fmt.Println("Gold :", c.Gold)
	fmt.Println("--------------------------")
}
