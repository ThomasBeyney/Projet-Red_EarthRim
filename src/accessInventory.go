package main

import "fmt"

func accessInventory(c *Character) {
	fmt.Println("Items dans l'inventaire :")
	for _, item := range c.Inventory {
		fmt.Println("-", item)
	}
}

func takePot(c *Character) {
	for i, item := range c.Inventory {
		if item == "Potion" {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			c.Hp += 50
			if c.Hp > c.MaxHp {
				c.Hp = c.MaxHp
			}
			fmt.Println("Potion utilisée !")
			fmt.Println("Points de vie :", c.Hp, "/", c.MaxHp)
			return
		}
	}
	fmt.Println("Aucune potion dans l'inventaire !")
}


func AddItem(c *Character, item string) bool {
	if len(c.Inventory) < c.InventoryMax {
		c.Inventory = append(c.Inventory, item)
		fmt.Println("Vous avez ajouté :", item)
		return true
	} else {
		fmt.Println("Inventaire plein ! Impossible d'ajouter", item)
		return false
	}
}



func upgradeInventorySlot(c *Character) {
    if c.LimInventory >= 3 {
        fmt.Println("Vous avez déjà atteint le nombre maximum d'augmentations.")
        return
    }

    for i, item := range c.Inventory {
        if item == "Grand sac" {
            c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
            c.InventoryMax += 10
            c.LimInventory++
            fmt.Println("Capacité d'inventaire augmentée de +10 !")
            fmt.Println("Nouvelle capacité max :", c.InventoryMax)
            return
        }
    }

    fmt.Println("Aucun 'Grand sac' trouvé dans votre inventaire.")
}
