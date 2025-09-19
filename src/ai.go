package main

import (
	"fmt"
	"time"
)


func goblinPattern(goblin *Monster, player *Character) {
	for turn := 1; player.Hp > 0; turn++ {
		if goblin.Hp <= 0 {
			return
		}

		var damage int
		if player.Turn%3 == 0 {
			damage = goblin.attaque * 2
			fmt.Println(" ")
			fmt.Println(goblin.Name, "utilise une attaque puissante !")
		} else {
			damage = goblin.attaque
		}

		player.Hp -= damage
		if player.Hp < 0 {
			player.Hp = 0
		}

		fmt.Println(" ")
		fmt.Println(goblin.Name, "inflige", damage, "dégâts à", player.Name)
		fmt.Println("PV de", player.Name, ":", player.Hp, "/", player.MaxHp)
		time.Sleep(1 * time.Second)

		if player.Hp == 0 {
			fmt.Println(" ")
			fmt.Println(player.Name, "est K.O.")
			time.Sleep(1 * time.Second)
			isDead(player)
			return
		}

		player.Turn += 1
		if player.Mana <= 0 {
			player.Mana = 0
			player.Mana += 15
		} else if player.Mana >= 35 {
			player.Mana = 50
		} else {
			player.Mana += 15
		}
		
		if !CharacterTurn(goblin, player) {
			return
		}
	}
}

func trollPattern(Troll *Monster, player *Character) {
	for turn := 1; player.Hp > 0; turn++ {
		if Troll.Hp <= 0 {
			return
		}

		var damage int
		if player.Turn%10 == 0 {
			damage = Troll.attaque * 10
			fmt.Println(" ")
			fmt.Println(Troll.Name, "utilise une attaque MORTEL !")
		} else {
			damage = Troll.attaque
		}

		player.Hp -= damage
		if player.Hp < 0 {
			player.Hp = 0
		}

		fmt.Println(" ")
		fmt.Println(Troll.Name, "inflige", damage, "dégâts à", player.Name)
		fmt.Println("PV de", player.Name, ":", player.Hp, "/", player.MaxHp)
		time.Sleep(1 * time.Second)

		if player.Hp == 0 {
			fmt.Println(" ")
			fmt.Println(player.Name, "est K.O.")
			time.Sleep(1 * time.Second)
			isDead(player)
			return
		}

		player.Turn += 1
		if player.Mana <= 0 {
			player.Mana = 0
			player.Mana += 15
		} else if player.Mana >= 35 {
			player.Mana = 50
		} else {
			player.Mana += 15
		}

		if !CombTroll(Troll, player) {
			return
		}
	}
}


func loupPattern(Loup *Monster, player *Character) {
	for turn := 1; player.Hp > 0; turn++ {
		if Loup.Hp <= 0 {
			return
		}

		var damage int
		if player.Turn%3 == 0 {
			damage = Loup.attaque
		} else {
			damage = Loup.attaque
		}

		player.Hp -= damage
		if player.Hp < 0 {
			player.Hp = 0
		}

		fmt.Println(" ")
		fmt.Println(Loup.Name, "inflige", damage, "dégâts à", player.Name)
		fmt.Println("PV de", player.Name, ":", player.Hp, "/", player.MaxHp)
		time.Sleep(1 * time.Second)

		if player.Hp == 0 {
			fmt.Println(" ")
			fmt.Println(player.Name, "est K.O.")
			time.Sleep(1 * time.Second)
			isDead(player)
			return
		}

		player.Turn += 1
		if player.Mana <= 0 {
			player.Mana = 0
			player.Mana += 15
		} else if player.Mana >= 35 {
			player.Mana = 50
		} else {
			player.Mana += 15
		}

		if !CombLoup(Loup, player) {
			return
		}
	}
}



func orcPattern(Orc *Monster, player *Character) {
	for turn := 1; player.Hp > 0; turn++ {
		if Orc.Hp <= 0 {
			return
		}

		var damage int
		if player.Turn%5 == 0 {
			damage = Orc.attaque * 3
			fmt.Println(" ")
			fmt.Println(Orc.Name, "utilise une attaque puissante !")
		} else {
			damage = Orc.attaque
		}

		player.Hp -= damage
		if player.Hp < 0 {
			player.Hp = 0
		}

		fmt.Println(" ")
		fmt.Println(Orc.Name, "inflige", damage, "dégâts à", player.Name)
		fmt.Println("PV de", player.Name, ":", player.Hp, "/", player.MaxHp)
		time.Sleep(1 * time.Second)

		if player.Hp == 0 {
			fmt.Println(" ")
			fmt.Println(player.Name, "est K.O.")
			time.Sleep(1 * time.Second)
			isDead(player)
			return
		}

		player.Turn += 1
		if player.Mana <= 0 {
			player.Mana = 0
			player.Mana += 15
		} else if player.Mana >= 35 {
			player.Mana = 50
		} else {
			player.Mana += 15
		}

		if !CombOrc(Orc, player) {
			return
		}
	}
}
