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
			CharacterTurn(goblin, player)
		} else {
			damage = goblin.attaque
			CharacterTurn(goblin, player)
		}

		player.Hp -= damage
		if player.Hp < 0 {
			player.Hp = 0

		}

		fmt.Println(goblin.Name,"inflige à",player.Name,damage,"points de dégâts\n")
		time.Sleep(1 * time.Second)
		fmt.Println("PV de",player.Name,":",player.Hp,"/",player.MaxHp,"\n\n")
		time.Sleep(1 * time.Second)

		if player.Hp == 0 {
			time.Sleep(1 * time.Second)
			fmt.Println(player.Name," est K.O.\n")
			isDead(player)
			break
		}
	}
}
