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





func CombTroll(Troll *Monster, player *Character) bool {
	
	
	for player.Hp > 0 && Troll.Hp > 0 {
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
    			fmt.Println("Vous attaquez le Troll !")
    			Troll.Hp -= 5
    			if Troll.Hp < 0 {
        			Troll.Hp = 0
    			}
    			fmt.Println("Il reste", Troll.Hp, "/", Troll.MaxHp, "PV au Troll.")

    			if Troll.Hp == 0 {
        			fmt.Println("Le troll est vaincu !")
					player.Exp += 100
					if player.Exp >= 100 {
						player.Exp = 0
						player.Level += 1
						player.MaxHp += 15
						fmt.Println(" ")
						fmt.Println("Bravo ! Vous gagnez un Niveau")
					} else {
						fmt.Println(" ")
						fmt.Println("Vous gagnez 100 points d'expériences")
					}
        			return false
    			}

				trollPattern(Troll, player)

    			if player.Hp == 0 {
        			return false
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
				return false

			default:
				fmt.Println("Choix invalide, veuillez réessayer.")
			}
		return false
		}
	return false
}





func CombLoup(Loup *Monster, player *Character) bool {
	
	
	for player.Hp > 0 && Loup.Hp > 0 {
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
    			fmt.Println("Vous attaquez le Loup !")
    			Loup.Hp -= 5
    			if Loup.Hp < 0 {
        			Loup.Hp = 0
    			}
    			fmt.Println("Il reste", Loup.Hp, "/", Loup.MaxHp, "PV au Loup.")

    			if Loup.Hp == 0 {
        			fmt.Println("Le loup est vaincu !")
					player.Exp += 15
					if player.Exp >= 100 {
						player.Exp = 0
						player.Level += 1
						player.MaxHp += 15
						fmt.Println(" ")
						fmt.Println("Bravo ! Vous gagnez un Niveau")
					} else {
						fmt.Println(" ")
						fmt.Println("Vous gagnez 15 points d'expériences")
					}
        			return false
    			}

				loupPattern(Loup, player)

    			if player.Hp == 0 {
        			return false
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
				return false

			default:
				fmt.Println("Choix invalide, veuillez réessayer.")
			}
		return false
		}
	return false
}




func CombOrc(Orc *Monster, player *Character) bool {
	
	
	for player.Hp > 0 && Orc.Hp > 0 {
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
    			fmt.Println("Vous attaquez l'Orc !")
    			Orc.Hp -= 5
    			if Orc.Hp < 0 {
        			Orc.Hp = 0
    			}
    			fmt.Println("Il reste", Orc.Hp, "/", Orc.MaxHp, "PV au Orc.")

    			if Orc.Hp == 0 {
        			fmt.Println("L'Orc est vaincu !")
					player.Exp += 45
					if player.Exp >= 100 {
						player.Exp = 0
						player.Level += 1
						player.MaxHp += 15
						fmt.Println(" ")
						fmt.Println("Bravo ! Vous gagnez un Niveau")
					} else {
						fmt.Println(" ")
						fmt.Println("Vous gagnez 45 points d'expériences")
					}
        			return false
    			}

				orcPattern(Orc, player)

    			if player.Hp == 0 {
        			return false
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
				return false

			default:
				fmt.Println("Choix invalide, veuillez réessayer.")
			}
		return false
		}
	return false
}
