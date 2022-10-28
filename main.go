package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Printf("Blackjack v(%s) b(%s)\n\n", Version, Build)

	var numberOfDecks int
	flag.IntVar(&numberOfDecks, "decks", 1, "Specify the number of decks")
	flag.Parse()

	fmt.Printf("Using %d decks\n", numberOfDecks)

	decks := NewDecks(numberOfDecks)

	dealer := NewDealer()
	player := NewPlayer("Levi")

	decks.Deal(&player.Hands[0])
	decks.Deal(&dealer.Hands[0])
	decks.Deal(&player.Hands[0])
	decks.Deal(&dealer.Hands[0])

	fmt.Println(dealer)
	fmt.Println(player)
}
