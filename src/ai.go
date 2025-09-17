package main

import (
	"fmt"
	"time"
)


func goblinPattern(goblin *Monster, player *Character) {
	for turn := 1; player.Hp > 0; turn++ {
		var damage int
		if turn%3 == 0 {
			damage = goblin.attaque * 2
			fmt.Println(goblin.Name, "utilise une attaque puissante !")
		} else {
			damage = goblin.attaque
		}


		player.Hp -= damage
		if player.Hp < 0 {
			player.Hp = 0
		}

		fmt.Println(goblin.Name, "inflige", damage, "dégâts à", player.Name)
		fmt.Println("PV de", player.Name, ":", player.Hp, "/", player.MaxHp)
		time.Sleep(1 * time.Second)

		if player.Hp == 0 {
			fmt.Println(player.Name, "est K.O.")
			time.Sleep(1 * time.Second)
			isDead(player)
		}
	CharacterTurn(goblin, player)
	}
}
