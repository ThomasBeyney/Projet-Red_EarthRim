package main

import "fmt"

func Marchand(c *Character) {
	fmt.Println("Bienvenue au Marchand")
	var rep int
	var reponse string
	for {
		fmt.Println("1- voulez vous achetez une potion de soin ?")
		fmt.Println("2- voulez vous achetez une potion de poison ?")
		fmt.Println("3- Partir")
		fmt.Print("Choix : ")
		fmt.Scanln(&rep)
		switch rep {
		case 1 :
			fmt.Print("Voulait vous vraiment achetez une potion de soin ?  oui/non ")
			fmt.Scanln(&reponse)

			if reponse == "oui" || reponse == "Oui" {
                if AddItem(c, "Potion") {
                    fmt.Println("Vous obtenez une potion")
                } else {
                    fmt.Println("Inventaire plein ! Vous ne pouvez pas acheter.")
				}
			} else if reponse == "non" || reponse == "Non" {
				fmt.Println("Dommage, elle pourrait ètre utile")
			} else {
				fmt.Println("Et bien, ce n'était pas ma question")
			}
		case 2 :
			fmt.Print("Voulait vous vraiment achetez une potion de poison ? oui/non ")
			fmt.Scanln(&reponse)

			if reponse == "oui" || reponse == "Oui" {
                if AddItem(c, "Potion-de-poison") {
                    fmt.Println("Vous obtenez une potion de poison")
                } else {
                    fmt.Println("Inventaire plein ! Vous ne pouvez pas acheter.")
				}
			} else if reponse == "non" || reponse == "Non" {
				fmt.Println("Dommage, elle pourrait ètre utile, ou mortel")
			} else {
				fmt.Println("Et bien, ce n'était pas ma question")
			}
		case 3 :
			fmt.Println("A bientot l'ami")
			return
		}
	}
}
