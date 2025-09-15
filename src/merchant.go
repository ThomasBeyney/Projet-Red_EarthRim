package main

import "fmt"

func Marchand(c *Character) {
	fmt.Println("Bienvenue au Marchand")
	var reponse string
	fmt.Print("Voulait vous achetez une potion ?  oui/non ")
	fmt.Scanln(&reponse)

	if reponse == "oui" || reponse == "Oui" {
		fmt.Println("Vous obtenez une potion")
		c.Inventory = append(c.Inventory, "Potion")
	} else if reponse == "non" || reponse == "Non" {
		fmt.Println("Dommage, elle pourrait ètre utile")
	} else {
		fmt.Println("Et bien, ce n'était pas ma question")
	}
}
