package main 

import "fmt"

type Equipement struct {
    Tete   string 
    Torse  string 
    Pieds  string 
}

func EquipHat(c *Character) {
	for i, item := range c.Inventory {
		if item == "Chapeau de l’aventurier" {
			if c.Equipement.Tete != "" {
				c.Inventory = append(c.Inventory, c.Equipement.Tete)
			}
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			c.Equipement.Tete = item
			c.MaxHp += 10

			fmt.Println("Chapeau de l’aventurier équipé !")
			fmt.Println("Les points de vie max sont maintenant :", c.Hp, "/", c.MaxHp)
			return
		}
	}

	fmt.Println("Aucun chapeau à équiper dans l'inventaire !")
}


func EquipTors(c *Character) {
	for i, item := range c.Inventory {
		if item == "Tunique de l’aventurier" {
			if c.Equipement.Torse != "" {
				c.Inventory = append(c.Inventory, c.Equipement.Torse)
			}
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			c.Equipement.Torse = item
			c.MaxHp += 25

			fmt.Println("Tunique de l’aventurier équipé !")
			fmt.Println("Les points de vie max sont maintenant :", c.Hp, "/", c.MaxHp)
			return
		}
	}

	fmt.Println("Aucune tunique à équiper dans l'inventaire !")
}



func EquipPie(c *Character) {
	for i, item := range c.Inventory {
		if item == "Bottes de l’aventurier" {
			if c.Equipement.Pieds != "" {
				c.Inventory = append(c.Inventory, c.Equipement.Pieds)
			}
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			c.Equipement.Pieds = item
			c.MaxHp += 15

			fmt.Println("Bottes de l’aventurier équipé !")
			fmt.Println("Les points de vie max sont maintenant :", c.Hp, "/", c.MaxHp)
			return
		}
	}

	fmt.Println("Aucune bottes à équiper dans l'inventaire !")
}
