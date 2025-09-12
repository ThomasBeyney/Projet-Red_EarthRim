package main

import (
	"fmt"
	"strings"
)

func Menu() {
	for {
		fmt.Println(" MENU ")
		fmt.Println("1. Afficher le personnage")
		fmt.Println("2. Voir l’inventaire")
		fmt.Println("3. Quitter")
		fmt.Print("Choisissez une option : ")

		var input string
		fmt.Scanln(&input)                 // lire l'entrée
		input = strings.TrimSpace(input)   // enlever les espaces éventuels

		switch input {
		case "1":
			fmt.Println("Vous avez choisi d afficher le personnage !")
		case "2":
			fmt.Println("Vous avez choisi de voir l inventaire !")
		case "3":
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Option invalide")
		}
	}
}
