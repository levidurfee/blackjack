package main

import (
	"flag"
	"fmt"

	"git.x6c.co/go/blackjack/internal/blackjack"
	"git.x6c.co/go/blackjack/internal/version"
)

const (
	DefaultHandIndex = 0
)

const (
	PromptQuestion = "What would you like to do?"
	PromptOptions  = "(h)it (s)stand (q)uit"
	HasBustText    = "has bust!"
)

const (
	HitChar   = "h"
	StandChar = "s"
	QuitChar  = "q"
)

func main() {
	fmt.Printf("Blackjack v(%s) b(%s)\n\n", version.Version, version.Build)

	var numberOfDecks int
	flag.IntVar(&numberOfDecks, "decks", 4, "Specify the number of decks")

	var playerName string
	flag.StringVar(&playerName, "name", "Player 1", "Player name")

	flag.Parse()

	fmt.Printf("Using %d decks\n", numberOfDecks)

	decks := blackjack.NewDecks(numberOfDecks)

	for {
		fmt.Printf("%d cards left\n", len(decks.Cards))
		if !decks.HasEnoughCards() {
			return
		}

		dealer := blackjack.NewDealer()
		player := blackjack.NewPlayer(playerName)

		decks.Deal(&player.Hands[DefaultHandIndex])
		decks.Deal(&dealer.Hands[DefaultHandIndex])
		decks.Deal(&player.Hands[DefaultHandIndex])
		decks.Deal(&dealer.Hands[DefaultHandIndex])

		var cmd string
		for {
			if dealer.Hands[DefaultHandIndex].Bust() {
				fmt.Println(dealer.Name, HasBustText)
				break
			}

			if player.Hands[DefaultHandIndex].Bust() {
				fmt.Println(player.Name, HasBustText)
				break
			}

			fmt.Println(dealer)
			fmt.Println(player)

			fmt.Println(PromptQuestion)
			fmt.Println(PromptOptions)
			fmt.Scanln(&cmd)

			switch cmd {
			case HitChar:
				decks.Deal(&player.Hands[DefaultHandIndex])
			case QuitChar:
				return
			}
			if cmd == StandChar {
				break
			}
		}
	}

}
