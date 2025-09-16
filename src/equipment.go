package main 

type Equipment struct {
    Tete   string 
    Torse  string 
    Pieds  string 
}

func Equip(c *Character) {
	for i, item := range c.Inventory {
		switch item {
		case "Chapeau de l'aventurier":
			if c.Equipment.Head != "" {
				c.Inventory = append(c.Inventory, c.Equipment.Head)
			}
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			// Équipe le nouvel objet et ajoute bonus HP
			c.Equipment.Head = item
			c.MaxHp += 10
			fmt.Println("Chapeau de l'aventurier équipé !")
			fmt.Println("Les points de vie max ont augmenté :", c.Hp, "/", c.MaxHp)
			return