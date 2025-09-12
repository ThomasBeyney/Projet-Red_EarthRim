package main

import "fmt"

func accessInventory(inventory []string) {
	fmt.Println("Items dans l'inventaire :")
	for _, item := range inventory {
		fmt.Println("-", item)
	}
}

func main() {
	inventory := []string{"Épée", "Potion"}
	accessInventory(inventory)
}
 func takePot() {

	
 }
 