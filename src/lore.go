package main

import (
    "fmt"
    "time"
)

func displayLore(lore string, delayMs int) {
    for _, r := range lore {
        fmt.Printf("%c", r)
        time.Sleep(time.Duration(delayMs) * time.Millisecond)
    }
    fmt.Println()
	characterCreation()
}
