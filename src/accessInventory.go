package main

import "fmt"

func accessInventory(inventory []string) string {
	fmt.Println("Items dans l'inventaire :")
	for _, items := range inventory {
		fmt.Println(items)
	}
}
