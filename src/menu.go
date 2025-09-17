package main

import "fmt"

func main() {
	player := characterCreation()
	goblin := initGoblin()
	fmt.Println("Bienvenue", player.Name, "tu as", player.Hp, "/", player.MaxHp, "hp")

	var choice int
	for {
		fmt.Println("\n=== Menu Principal ===")
		fmt.Println("1 - Afficher les informations du personnage")
		fmt.Println("2 - Accéder au contenu de l'inventaire")
		fmt.Println("3 - Utiliser un objet")
		fmt.Println("4 - Forgeron")
		fmt.Println("5 - Marchand")
		fmt.Println("6 - Combat")
		fmt.Println("7 - Quitter")
		fmt.Print("Choix : ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			displayInfo(&player)
			fmt.Println("--------------------------")
		case 2:
			accessInventory(&player)
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
            takePot(&player)
        case "Potion-de-poison":
            poison_pot(&player)
		case "Livre de boule de feu": 
			Spellbook(&player)
		case "Chapeau de l’aventurier":
			EquipHat(&player)
		case "Tunique de l’aventurier":
			EquipTors(&player)
		case "Bottes de l’aventurier":
			EquipPie(&player)
		case "Grand sac":
			upgradeInventorySlot(&player)
        default:
            fmt.Println("Cet objet ne peut pas être utilisé.")
        }
		}
		case 4:
			Forgeron(&player)
		case 5:
			Marchand(&player)
		case 6:
			CharacterTurn(&goblin, &player)
		case 7:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}
