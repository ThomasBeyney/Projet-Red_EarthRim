package main

import "fmt"

func Forgeron(c *Character) {
	fmt.Println("Bienvenue chez le Forgeron")
	var rep int
	var reponse string
	for {
		fmt.Println("Que shouatait vous fabriquer ?")
		fmt.Println("1- Le chapeau de l’aventurier ?")
		fmt.Println("2- La tunique de l’aventurier?")
		fmt.Println("3- Les bottes de l’aventurier")
		fmt.Println("4- Partir")
		fmt.Print("Choix : ")
		fmt.Scanln(&rep)
		switch rep {
		case 1 :
			fmt.Print("Voulait vous vraiment fabriquer le chapeau de l’aventurier  ?  oui/non ")
			fmt.Scanln(&reponse)

			if reponse == "oui" || reponse == "Oui" {
			    if c.Gold >= 5 {
                if AddItem(c, "Chapeau de l’aventurier") {
					fmt.Println("Vous payer 5 d'or")
                    fmt.Println("Vous obtenez le chapeau de l’aventurier")
					c.Gold -= 5
                } else {
                    fmt.Println("Inventaire plein ! Vous ne pouvez pas le fabriquer.")
				}} else {
					fmt.Println("Vous n'avez pas assez d'or")
				}
			}  
			} else if reponse == "non" || reponse == "Non" {
				fmt.Println("Dommage, il pourrait vous ètre utile")
			} else {
				fmt.Println("Et bien, ce n'était pas ma question")
			}
		case 2 :
			fmt.Print("Voulait vous vraiment fabriquer la tunique de l’aventurier ? oui/non ")
			fmt.Scanln(&reponse)

			if reponse == "oui" || reponse == "Oui" {
				if c.Gold >= 5 {
               	   if AddItem(c, "Tunique de l’aventurier") {
					  fmt.Println("Vous payer 5 d'or")
                      fmt.Println("Vous obtenez la tunique de l’aventurier")
					  c.Gold -= 5
                   } else {
                    fmt.Println("Inventaire plein ! Vous ne pouvez pas la fabriquer.")
				   }} else {
						fmt.Println("Vous n'avez pas assez d'or")
			    }
			} else if reponse == "non" || reponse == "Non" {
				fmt.Println("Dommage, elle pourrait vous ètre utile")
			} else {
				fmt.Println("Et bien, ce n'était pas ma question")
			}
		case 3 :
			fmt.Print("Voulait vous vraiment fabriquer les bottes de l’aventurier ? oui/non ")
			fmt.Scanln(&reponse)

			if reponse == "oui" || reponse == "Oui" {
                if AddItem(c, "Bottes de l’aventurier") {
					fmt.Println("Vous payer 5 d'or")
                    fmt.Println("Vous obtenez les bottes de l’aventurier")
					c.Gold -= 5
                } else {
                    fmt.Println("Inventaire plein ! Vous ne pouvez pas les fabriquer.")
				}} else {
						fmt.Println("Vous n'avez pas assez d'or")
			    }
			} else if reponse == "non" || reponse == "Non" {
				fmt.Println("Dommage, elle pourrait vous ètre utiles")
			} else {
				fmt.Println("Et bien, ce n'était pas ma question")
			}
		case 4 :
			fmt.Println("A bientot l'ami")
			return
		}
	}
}
