package main

import "fmt"

func accessInventory(c *Character) {
	fmt.Println("Items dans l'inventaire :")
	for _, item := range c.Inventory {
		fmt.Println("-", item)
	}
}

func takePot(c *Character) {
	for i, item := range c.Inventory {
		if item == "Potion de soin" {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			c.Hp += 50
			if c.Hp > c.MaxHp {
				c.Hp = c.MaxHp
			}
			fmt.Println("Potion utilisée !")
			fmt.Println("Points de vie :", c.Hp, "/", c.MaxHp)
			return
		}
	}
	fmt.Println("Aucune potion dans l'inventaire !")
}


func AddItem(c *Character, item string) bool {
	if len(c.Inventory) < 10 {
		c.Inventory = append(c.Inventory, item)
		fmt.Println("Vous avez ajouté :", item)
		return true
	} else {
		fmt.Println("Inventaire plein ! Impossible d’ajouter", item)
		return false
	}
}
