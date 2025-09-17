package main

import "fmt"

func Spellbook(c *Character) {
	if c.Skillmax > 1 {
		fmt.Println("vous avez déja appris ce sort")
		return
	}
	for i, item := range c.Inventory {
		if item == "Livre de boule de feu" {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			c.Skill = append(c.Skill, "Boule de feu")
			fmt.Println("Vous sentez la magie vous envahir !")
			fmt.Println("Soudenement vous savez comment lancer une boule de feu")
			fmt.Println("Incroyable !")
			c.Skillmax++
			return
		}
	}
	fmt.Println("Aucune Livre de boule de feu dans l'inventaire")
	}
