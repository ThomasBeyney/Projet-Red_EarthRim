package main 

type Monster struct {
	Name      string
	MaxHp     int
	Hp        int
	attaque   int
}

func initGoblin() Monster {
	return Character{
		Name:      "Gobelin d'entrainement",
		MaxHp:     40,
		Hp:        40,
		attaque:   5,
	}
}

