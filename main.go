package main

import (
	"flag"
	"fmt"
)

const (
	DefaultHandIndex = 0
)

func main() {
	fmt.Printf("Blackjack v(%s) b(%s)\n\n", Version, Build)

	var numberOfDecks int
	flag.IntVar(&numberOfDecks, "decks", 1, "Specify the number of decks")
	flag.Parse()

	fmt.Printf("Using %d decks\n", numberOfDecks)

	decks := NewDecks(numberOfDecks)

	for {
		fmt.Printf("%d cards left\n", len(decks.Cards))
		if len(decks.Cards) < 4 {
			return
		}

		dealer := NewDealer()
		player := NewPlayer("Levi")

		decks.Deal(&player.Hands[DefaultHandIndex])
		decks.Deal(&dealer.Hands[DefaultHandIndex])
		decks.Deal(&player.Hands[DefaultHandIndex])
		decks.Deal(&dealer.Hands[DefaultHandIndex])

		var cmd string
		for {
			fmt.Println(dealer)
			fmt.Println(player)

			fmt.Println("What would you like to do?")
			fmt.Println("(h)it (s)stand (q)uit")
			fmt.Scanln(&cmd)

			switch cmd {
			case "h":
				decks.Deal(&player.Hands[0])
			case "q":
				return
			}
			if cmd == "s" {
				break
			}
		}
	}

}
