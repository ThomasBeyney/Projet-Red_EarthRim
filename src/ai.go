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

		fmt.Printf("%s inflige à %s %d de dégâts\n", goblin.Name, player.Name, damage)
		fmt.Printf("PV de %s : %d/%d\n\n", player.Name, player.Hp, player.MaxHp)

		if player.Hp == 0 {
			fmt.Printf("%s est K.O.\n", player.Name)
			break
		}
	}
}
