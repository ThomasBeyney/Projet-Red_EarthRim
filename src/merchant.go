package main

import "fmt"

func Marchand(c *Character) {
	fmt.Println("Bienvenue au Marchand")
	var reponse string
	fmt.Print("Voulait vous achetez une potion ?")
	fmt.Scanln(&reponse)

	if reponse == "oui" || reponse == "Oui" {
		fmt.Println("Vous obtenez une potion")
	} else {
		fmt.Println("Dommage, elle pourrait ètre utile")
	}
}
