package main

import "fmt"

func Marchand(c *Character) {
	fmt.Println("Bienvenue au Marchand")
	var rep int
	var reponse string
	var rep2 int
	for {
		fmt.Println("1- voulez vous achetez une potion de soin ?")
		fmt.Println("2- voulez vous achetez une potion de poison ?")
		fmt.Println("3- voulez vous achetez des livres de sort ?")
		fmt.Println("4- Partir")
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
		case 3:
		for {
			fmt.Println("Que voulez-vous acheter ?")
			fmt.Println("1- Acheter le livre de Boule de feu")
			fmt.Println("2- Revenir en arrière")
			fmt.Print("Choix : ")
			fmt.Scanln(&rep2)

			switch rep2 {
				case 1:
					fmt.Print("Voulez-vous vraiment acheter le livre de boule de feu ? oui/non ")
					fmt.Scanln(&reponse)

					if reponse == "oui" || reponse == "Oui" {
						if AddItem(c, "Livre de boule de feu") {
							fmt.Println("Vous obtenez un Livre de boule de feu")
						} else {
							fmt.Println("Inventaire plein ! Vous ne pouvez pas acheter.")
						}
					} else if reponse == "non" || reponse == "Non" {
						fmt.Println("Dommage, il pourrait s'avérer très utile.")
					} else {
						fmt.Println("Eh bien, ce n'était pas ma question.")
					}

				case 2:
					break 
				}
				
				if rep2 == 2 {
					break
				}
			}
		case 4 :
			fmt.Println("A bientot l'ami")
			return
		}
	}
}
