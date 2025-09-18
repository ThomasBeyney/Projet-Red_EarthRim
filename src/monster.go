package main 

type Monster struct {
	Name      string
	MaxHp     int
	Hp        int
	attaque   int
}

func initGoblin() Monster {
	return Monster{
		Name:      "Gobelin d'entrainement",
		MaxHp:     40,
		Hp:        40,
		attaque:   5,
	}
}

func initTroll() Monster {
	return Monster{
		Name:    "Troll",
		MaxHp:   80,
		Hp:      80,
		attaque: 10,
	}
}

func initLoup() Monster {
	return Monster{
		Name:    "Loup",
		MaxHp:   25,
		Hp:      25,
		attaque: 3,
	}
}

func initOrc() Monster {
	return Monster{
		Name:    "Orc",
		MaxHp:   50,
		Hp:      50,
		attaque: 5,
	}
}