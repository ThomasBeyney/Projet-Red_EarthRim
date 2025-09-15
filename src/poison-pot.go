package main

import (
	"fmt"
	"time"
)

func poison_pot(c *Character) {
	for i, item := range c.Inventory {
		if item == "Potion-de-poison" {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			fmt.Println("Potion-de-poison utilisée !")
			fmt.Println("Vous vous santez un peut bizzard")
			fmt.Println("vous perdez 10 points de vie, aïe")
			c.Hp -= 10
			isDead(c)
			time.Sleep(1 * time.Second)
			fmt.Println("vous reperdez 10 points de vie")
			c.Hp -= 10
			isDead(c)
			time.Sleep(1 * time.Second)
			fmt.Println("les effets du poison disparaisse mais vous perdez encore 10 points de vie")
			c.Hp -= 10
			isDead(c)
			if c.Hp < 0 {
				c.Hp = 0
			}
			fmt.Println("Points de vie :", c.Hp, "/", c.MaxHp)
			return
		}
	}
	fmt.Println("Aucune potion dans l'inventaire !")
}
