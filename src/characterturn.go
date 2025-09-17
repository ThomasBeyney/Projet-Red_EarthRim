package main

import "fmt"


func CharacterTurn(goblin *Monster, player *Character) {
	
	
	for player.Hp > 0 && goblin.Hp > 0 {
    	var choice int
    	fmt.Println("\n=== Interface de combat ===")
    	fmt.Println("1 - Statistique")
   		fmt.Println("2 - Attaquer")
    	fmt.Println("3 - Inventaire")
    	fmt.Println("4 - Quitter")
    	fmt.Print("Choix : ")
    	fmt.Scanln(&choice)

		switch choice {
			case 1:
				displayInfo(player)

			case 2:
    			fmt.Println("Vous attaquez le gobelin !")
    			goblin.Hp -= 5
    			if goblin.Hp < 0 {
        			goblin.Hp = 0
    			}
    			fmt.Println("Il reste", goblin.Hp, "/", goblin.MaxHp, "PV au gobelin.")

    			if goblin.Hp == 0 {
        			fmt.Println("Le gobelin est vaincu !")
        			return
    			}

				goblinPattern(goblin, player)

    			if player.Hp == 0 {
        			return
    			}

			case 3:
				var rep2 int
				for {
					fmt.Println("Que voulez-vous faire ?")
					fmt.Println("1 - Voir Inventaire")
					fmt.Println("2 - Revenir en arrière")
					fmt.Print("Choix : ")
					fmt.Scanln(&rep2)

					if rep2 == 2 {
						break
					}

					if len(player.Inventory) == 0 {
						fmt.Println("Votre inventaire est vide.")
						continue
					}

					fmt.Println("=== Inventaire ===")
					for i, item := range player.Inventory {
						fmt.Println(i+1, "-", item)
					}

					var itemChoice int
					fmt.Print("Quel objet voulez-vous utiliser ? ")
					fmt.Scanln(&itemChoice)

					if itemChoice < 1 || itemChoice > len(player.Inventory) {
						fmt.Println("Choix invalide.")
						continue
					}

					selectedItem := player.Inventory[itemChoice-1]

					switch selectedItem {
					case "Potion":
						takePot(player)
					case "Potion-de-poison":
						poison_pot(player)
					case "Livre de boule de feu":
						Spellbook(player)
					case "Chapeau de l’aventurier":
						EquipHat(player)
					case "Tunique de l’aventurier":
						EquipTors(player)
					case "Bottes de l’aventurier":
						EquipPie(player)
					case "Grand sac":
						upgradeInventorySlot(player)
					default:	
						fmt.Println("Cet objet ne peut pas être utilisé.")
					}
				}

			case 4:
				fmt.Println("Au revoir !")
				return

			default:
				fmt.Println("Choix invalide, veuillez réessayer.")
			}
		}
	//goblinPattern(goblin, player)
}