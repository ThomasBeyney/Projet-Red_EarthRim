package main

import "fmt"

func Marchand(c *Character) {
	fmt.Println("Bienvenue au Marchand")
	var rep int
	var reponse string
	var rep2 int
	var rep3 int
	for {
		fmt.Println("1- voulez vous achetez une potion de soin ?")
		fmt.Println("2- voulez vous achetez une potion de poison ?")
		fmt.Println("3- voulez vous achetez des livres de sort ?")
		fmt.Println("4- Voulez achetez des ressources ?")
		fmt.Println("5- Partir")
		fmt.Print("Choix : ")
		fmt.Scanln(&rep)
		switch rep {
		case 1 :
			fmt.Print("Voulait vous vraiment achetez une potion de soin ? (coute 3 d'or)  oui/non ")
			fmt.Scanln(&reponse)

			if reponse == "oui" || reponse == "Oui" {
				if c.Gold >= 3 {
					if AddItem(c, "Potion") {
						fmt.Println("Vous payer 3 d'or")
						fmt.Println("Vous obtenez une potion")
						c.Gold -= 3
					} else {
						fmt.Println("Inventaire plein ! Vous ne pouvez pas acheter.")
					}} else {
						fmt.Println("Vous n'avez pas assez d'or")
					}
			} else if reponse == "non" || reponse == "Non" {
				fmt.Println("Dommage, elle pourrait ètre utile")
			} else {
				fmt.Println("Et bien, ce n'était pas ma question")
			}
		case 2 :
			fmt.Print("Voulait vous vraiment achetez une potion de poison ? (coute 6 d'or) oui/non ")
			fmt.Scanln(&reponse)

			if reponse == "oui" || reponse == "Oui" {
                if c.Gold >= 6 {
					if AddItem(c, "Potion-de-poison") {
						fmt.Println("Vous payer 6 d'or")
						fmt.Println("Vous obtenez une potion de poison")
						c.Gold -= 6
					} else {
						fmt.Println("Inventaire plein ! Vous ne pouvez pas acheter.")
					}} else {
						fmt.Println("Vous n'avez pas assez d'or")
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
					fmt.Print("Voulez-vous vraiment acheter le livre de boule de feu ? (coute 25 d'or) oui/non ")
					fmt.Scanln(&reponse)

					if reponse == "oui" || reponse == "Oui" {
						if c.Gold >= 25 {
							if AddItem(c, "Livre de boule de feu") {
								fmt.Println("Vous payer 25 d'or")
								fmt.Println("Vous obtenez un Livre de boule de feu")
								c.Gold -= 25
							} else {
								fmt.Println("Inventaire plein ! Vous ne pouvez pas acheter.")
							}} else { 
								fmt.Println("Vous n'avez pas assez d'or")
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
			for {
				fmt.Println("Que voulez vous achetez ?")
				fmt.Println("1-  Fourrure de Loup ")
				fmt.Println("2-  Peau de Troll")
				fmt.Println("3- Cuir de Sanglier")
				fmt.Println("4- Plume de Corbeau")
				fmt.Println("5- Revenir en arrière")
				fmt.Scanln(&rep3)
				
				switch rep3 {
				case 1:
					fmt.Println("Voulez vous vraiment achetez une Fourrure de loup ? (pour 4 d'or) oui/non")
					fmt.Scanln(&reponse)

					if reponse == "oui" || reponse == "Oui" {
						if c.Gold >= 4 {
							if AddItem(c, "Fourrure de loup") {
								fmt.Println("Vous payer 4 d'or")
								fmt.Println("Vous obtenez une Fourrure de loup")
								c.Gold -= 4
							} else {
								fmt.Println("Inventaire plein ! Vous ne pouvez pas acheter.")
							}} else {
								fmt.Println("Vous n'avez pas assez d'or")
							}
						} else if reponse == "non" || reponse == "Non" {
							fmt.Println("Dommage, elle pourrait s'avérer très utile.")
						} else {
							fmt.Println("Eh bien, ce n'était pas ma question.")
						}
						
				case 2:
					fmt.Println("Voulez vous vraiment achetez une Peau de troll ? (pour 7 d'or) oui/non")
					fmt.Scanln(&reponse)

					if reponse == "oui" || reponse == "Oui" {
						if c.Gold >= 7 {
							if AddItem(c, "Peau de troll") {
								fmt.Println("Vous payer 7 d'or")
								fmt.Println("Vous obtenez une Peau de troll")
								c.Gold -= 7
							} else {
								fmt.Println("Inventaire plein ! Vous ne pouvez pas acheter.")
							}} else {
								fmt.Println("Vous n'avez pas assez d'or")
							}
						} else if reponse == "non" || reponse == "Non" {
							fmt.Println("Dommage, elle pourrait s'avérer très utile.")
						} else {
							fmt.Println("Eh bien, ce n'était pas ma question.")
						}

				case 3: 
					fmt.Println("Voulez vous vraiment achetez une Cuir de sanglier ? (pour 3 d'or) oui/non")
					fmt.Scanln(&reponse)

					if reponse == "oui" || reponse == "Oui" {
						if c.Gold >= 3 {
							if AddItem(c, "Cuir de sanglier") {
								fmt.Println("Vous payer 3 d'or")
								fmt.Println("Vous obtenez une Cuir de sanglier")
								c.Gold -= 3
							} else {
								fmt.Println("Inventaire plein ! Vous ne pouvez pas acheter.")
							}} else {
								fmt.Println("Vous n'avez pas assez d'or")
							}
						} else if reponse == "non" || reponse == "Non" {
							fmt.Println("Dommage, elle pourrait s'avérer très utile.")
						} else {
							fmt.Println("Eh bien, ce n'était pas ma question.")
						}
						
				case 4:
					fmt.Println("Voulez vous vraiment achetez une Plume de corbeau ? (pour 1 d'or) oui/non")
					fmt.Scanln(&reponse)

					if reponse == "oui" || reponse == "Oui" {
						if c.Gold >= 1 {
							if AddItem(c, "Plume de corbeau") {
								fmt.Println("Vous payer 1 d'or")
								fmt.Println("Vous obtenez une Plume de corbeau")
								c.Gold -= 1
							} else {
								fmt.Println("Inventaire plein ! Vous ne pouvez pas acheter.")
							}} else {
								fmt.Println("Vous n'avez pas assez d'or")
							}
						} else if reponse == "non" || reponse == "Non" {
							fmt.Println("Dommage, elle pourrait s'avérer très utile.")
						} else {
							fmt.Println("Eh bien, ce n'était pas ma question.")
						}

				case 5:
					break
				}
				if rep3 == 5 {
					break
				}
			}
		case 5 :
			fmt.Println("A bientot l'ami !")
			return
		}
	}
}
