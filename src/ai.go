package main

import "fmt"

func goblinPattern(goblin *Monster, player *Character) {
	for turn := 1; player.Hp > 0; turn++ {
		var damage int
		if turn%3 == 0 {
			damage = goblin.attaque * 2
		} else {
			damage = goblin.attaque
		}

		player.Hp -= damage
		if player.Hp < 0 {
			player.Hp = 0

		}

		fmt.Println(goblin.Name,"inflige à",player.Name,damage,"points de dégâts\n")
		fmt.Println("PV de",player.Name,":",player.Hp,"/",player.MaxHp,"\n\n")

		if player.Hp == 0 {
			fmt.Println(player.Name," est K.O.\n")
			isDead(player)
			break
		}
	}
}
