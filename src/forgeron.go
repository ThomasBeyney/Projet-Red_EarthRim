package main

import "fmt"

func Forgeron(c *Character) {
	fmt.Println("Bienvenue chez le Forgeron")
	var rep int
	var reponse string

	for {
		fmt.Println("Que souhaitez-vous fabriquer ?")
		fmt.Println("1- Le chapeau de l’aventurier")
		fmt.Println("2- La tunique de l’aventurier")
		fmt.Println("3- Les bottes de l’aventurier")
		fmt.Println("4- Partir")
		fmt.Print("Choix : ")
		fmt.Scanln(&rep)

		switch rep {
		case 1:
			fmt.Print("Voulez-vous vraiment fabriquer le chapeau de l’aventurier (pour 5 pièces d'or, une plume de corbeau et un cuir de sanglier) ? oui/non : ")
			fmt.Scanln(&reponse)

			if reponse == "oui" || reponse == "Oui" {

				// pour la Plume de corbeau
				plumeIndex := -1
				for i, item := range c.Inventory {
					if item == "Plume de corbeau" {
						plumeIndex = i
						break
					}
				}
				if plumeIndex == -1 {
					fmt.Println("Vous n'avez pas de Plume de Corbeau pour fabriquer ce chapeau !")
					break
				}
				c.Inventory = append(c.Inventory[:plumeIndex], c.Inventory[plumeIndex+1:]...)
				fmt.Println("Plume de Corbeau utilisée !")

				// pour le Cuir de sanglier
				cuirIndex := -1
				for i, item := range c.Inventory {
					if item == "Cuir de sanglier" {
						cuirIndex = i
						break
					}
				}
				if cuirIndex == -1 {
					fmt.Println("Vous n'avez pas de Cuir de sanglier pour fabriquer ce chapeau !")
					break
				}
				c.Inventory = append(c.Inventory[:cuirIndex], c.Inventory[cuirIndex+1:]...)
				fmt.Println("Cuir de sanglier utilisé !")

				if c.Gold >= 5 {
					if AddItem(c, "Chapeau de l’aventurier") {
						fmt.Println("Vous payez 5 d'or")
						fmt.Println("Vous obtenez le chapeau de l’aventurier")
						c.Gold -= 5
					} else {
						fmt.Println("Inventaire plein ! Vous ne pouvez pas le fabriquer.")
					}
				} else {
					fmt.Println("Vous n'avez pas assez d'or")
				}

			} else if reponse == "non" || reponse == "Non" {
				fmt.Println("Dommage, il pourrait vous être utile")
			} else {
				fmt.Println("Et bien, ce n'était pas ma question")
			}

		case 2:
			fmt.Print("Voulez-vous vraiment fabriquer la tunique de l’aventurier (pour 5 pièces d'or, 2 Fourrure de loup et une peau de Troll) ? oui/non : ")
			fmt.Scanln(&reponse)

			if reponse == "oui" || reponse == "Oui" {

				//pour la Fourrure de loup
				FourrureIndex := -1
				for i, item := range c.Inventory {
					if item == "Fourrure de loup" {
						FourrureIndex = i
						break
					}
				}
				if FourrureIndex == -1 {
					fmt.Println("Vous n'avez pas suffisement de fourrure de loup pour fabriquer cette tunique !")
					break
				}
				c.Inventory = append(c.Inventory[:FourrureIndex], c.Inventory[FourrureIndex+1:]...)
				fmt.Println("Fourrure de loup utilisée !")

				for i, item := range c.Inventory {
					if item == "Fourrure de loup" {
						FourrureIndex = i
						break
					}
				}
				if FourrureIndex == -1 {
					fmt.Println("Vous n'avez pas suffisement de fourrure de loup pour fabriquer cette tunique !")
					break
				}
				c.Inventory = append(c.Inventory[:FourrureIndex], c.Inventory[FourrureIndex+1:]...)
				fmt.Println("Fourrure de loup utilisée !")

				// pour la Peau de troll
				PeauTrollIndex := -1
				for i, item := range c.Inventory {
					if item == "Peau de troll" {
						PeauTrollIndex = i
						break
					}
				}
				if PeauTrollIndex == -1 {
					fmt.Println("Vous n'avez pas de Peau de Troll pour fabriquer cette tunique !")
					break
				}
				c.Inventory = append(c.Inventory[:PeauTrollIndex], c.Inventory[PeauTrollIndex+1:]...)
				fmt.Println("Peau de troll utilisé !")

				if c.Gold >= 5 {
					if AddItem(c, "Tunique de l’aventurier") {
						fmt.Println("Vous payez 5 d'or")
						fmt.Println("Vous obtenez la tunique de l’aventurier")
						c.Gold -= 5
					} else {
						fmt.Println("Inventaire plein ! Vous ne pouvez pas le fabriquer.")
					}
				} else {
					fmt.Println("Vous n'avez pas assez d'or")
				}
			} else if reponse == "non" || reponse == "Non" {
				fmt.Println("Dommage, elle pourrait vous être utile")
			} else {
				fmt.Println("Et bien, ce n'était pas ma question")
			}

		case 3:
			fmt.Print("Voulez-vous vraiment fabriquer les bottes de l’aventurier (pour 5 pièces d'or, une Fourrure de loup et une Peau de troll) ? oui/non : ")
			fmt.Scanln(&reponse)

			if reponse == "oui" || reponse == "Oui" {

				//pour la Fourrure de loup
				FourrureIndex := -1
				for i, item := range c.Inventory {
					if item == "Fourrure de loup" {
						FourrureIndex = i
						break
					}
				}
				if FourrureIndex == -1 {
					fmt.Println("Vous n'avez pas de Fourrure de loup pour fabriquer ce chapeau !")
					break
				}
				c.Inventory = append(c.Inventory[:FourrureIndex], c.Inventory[FourrureIndex+1:]...)
				fmt.Println("Fourrure de loup utilisée !")

				//pour le Cuir de sanglier
				cuirIndex := -1
				for i, item := range c.Inventory {
					if item == "Cuir de sanglier" {
						cuirIndex = i
						break
					}
				}
				if cuirIndex == -1 {
					fmt.Println("Vous n'avez pas de Cuir de sanglier pour fabriquer ces bottes !")
					break
				}
				c.Inventory = append(c.Inventory[:cuirIndex], c.Inventory[cuirIndex+1:]...)
				fmt.Println("Cuir de sanglier utilisé !")

				if c.Gold >= 5 {
					if AddItem(c, "Bottes de l’aventurier") {
						fmt.Println("Vous payez 5 d'or")
						fmt.Println("Vous obtenez les bottes de l’aventurier")
						c.Gold -= 5
					} else {
						fmt.Println("Inventaire plein ! Vous ne pouvez pas les fabriquer.")
					}
				} else {
					fmt.Println("Vous n'avez pas assez d'or")
				}
			} else if reponse == "non" || reponse == "Non" {
				fmt.Println("Dommage, elles pourraient vous être utiles")
			} else {
				fmt.Println("Et bien, ce n'était pas ma question")
			}

		case 4:
			fmt.Println("À bientôt l'ami !")
			return
		}
	}
}
