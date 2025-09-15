package main

import "fmt"

func isDead(c *Character) {
	if c.Hp <= 0 {
		fmt.Println("vous etes completement dead")
		fmt.Println("Mais vous ressuscité, (pas entierement il ne faut pas abuser)")
		c.Hp = c.MaxHp / 2 
		fmt.Println("Maintenant vos point de vie sont de : ", c.Hp )
	}
}
