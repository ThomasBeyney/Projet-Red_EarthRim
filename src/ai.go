package main

import (
	"fmt"
	"time"
)


func goblinPattern(goblin Monster, playerCharacter) {
	var damage int
	if goblin.Turn%3 == 0 {
		damage = goblin.attaque * 2
		fmt.Println(goblin.Name, "utilise une attaque puissante !")
	} else {
		damage = goblin.attaque
	}

	goblin.Turn++

	player.Hp -= damage
	if player.Hp < 0 {
		player.Hp = 0
	}

	fmt.Println(goblin.Name, "inflige", damage, "dégâts à", player.Name)
	fmt.Println("PV de", player.Name, ":", player.Hp, "/", player.MaxHp)

	if player.Hp == 0 {
		fmt.Println(player.Name, "est K.O.")
		isDead(player)
	}
}