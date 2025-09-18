package main

import (
    "fmt"
    "time"
)

func displayLore(lore string, delayMs int) bool {
    for _, r := range lore {
        fmt.Printf("%c", r)
        time.Sleep(time.Duration(delayMs) * time.Millisecond)
    }
    fmt.Println()
    return true
}

func main() {
    result := displayLore("Mon texte, voici un grand texte qui ne sert pas a grand chose mais c'est pas grave, en vrai c'est juste un test.", 30)
    fmt.Println("Fin du lore ?", result)
}