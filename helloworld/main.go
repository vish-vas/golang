package main

import (
	"fmt"
)

func main() {
	cards := []string{newCard(), "pop"}
	cards = append(cards, "new pop")
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of Hearts"
}
