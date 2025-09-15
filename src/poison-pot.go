package main

import (
	"fmt"
	"time"
)

func poison_pot(c *Character) {
	for i, item := range c.Inventory {
		if item == "Potion-de-poison" {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			c.Hp -= 10
			time.Sleep(1 * time.Second)
			c.Hp -= 10
			time.Sleep(1 * time.Second)
			c.Hp -= 10
			if c.Hp < 0 {
				c.Hp = 0
			}
			fmt.Println("Potion-de-poison utilisée !")
			fmt.Println("Vous vous santez un peut bizzard")
			fmt.Println("Points de vie :", c.Hp, "/", c.MaxHp)
			return
		}
	}
	fmt.Println("Aucune potion dans l'inventaire !")
}
