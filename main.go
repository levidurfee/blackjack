package main

import (
	"fmt"
)

func main() {
	fmt.Println("Blackjack")

	decks := NewDecks(1)

	dealer := NewDealer()
	player := NewPlayer("Levi")

	decks.Deal(&dealer)
	decks.Deal(&player)
	decks.Deal(&dealer)
	decks.Deal(&player)

	fmt.Println(dealer)
	fmt.Println(player)
}
