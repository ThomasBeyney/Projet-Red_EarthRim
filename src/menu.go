package main

import "fmt"

func main() {
	player := initCharacter()
	fmt.Println("Bienvenue", player.Name, "tu as", player.Hp, "/", player.MaxHp, "hp")

	var choice int
	for {
		fmt.Println("\n=== Menu Principal ===")
		fmt.Println("1 - Afficher les informations du personnage")
		fmt.Println("2 - Accéder au contenu de l'inventaire")
		fmt.Println("3 - Utiliser une potion de soin")
		fmt.Println("4 - Utiliser une potion de poison")
		fmt.Println("5 - Marchand")
		fmt.Println("6 - Quitter")
		fmt.Print("Choix : ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			displayInfo(&player)
			fmt.Println("--------------------------")
		case 2:
			accessInventory(&player)
		case 3:
			takePot(&player)
		case 4:
			poison_pot(&player)
		case 5:
			Marchand(&player)
		case 6:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}
