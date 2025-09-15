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
		if item == "Potion" {
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


