package main

import (
	"flag"
	"fmt"

	"git.x6c.co/go/blackjack/pkg/blackjack"
	"git.x6c.co/go/blackjack/pkg/strategy"
	"git.x6c.co/go/blackjack/pkg/version"
)

const (
	DefaultHandIndex = 0
)

const (
	PromptQuestion = "What would you like to do?"
	PromptOptions  = "(h)it (s)stand (q)uit"
	HasBustText    = "has bust!"
	AnyKeyText     = "Hit any key to continue..."
)

const (
	HitChar   = "h"
	StandChar = "s"
	QuitChar  = "q"
)

const (
	AnsiClearScreen = "\033[H\033[2J"
)

func main() {
	fmt.Print(AnsiClearScreen)
	fmt.Printf("Blackjack v(%s) b(%s)\n\n", version.Version, version.Build)

	var numberOfDecks int
	flag.IntVar(&numberOfDecks, "decks", 4, "Specify the number of decks")

	var playerName string
	flag.StringVar(&playerName, "name", "Player 1", "Player name")

	flag.Parse()

	fmt.Printf("Using %d decks\n", numberOfDecks)
	fmt.Println(AnyKeyText)

	decks := blackjack.NewDecks(numberOfDecks)
	fmt.Scanln()

	for {
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
			fmt.Print(AnsiClearScreen)
			fmt.Printf("%d cards left\n", len(decks.Cards))

			fmt.Println(dealer)
			fmt.Println(player)

			if dealer.Hands[DefaultHandIndex].Bust() {
				fmt.Println(dealer.Name, HasBustText)
				fmt.Println(AnyKeyText)
				fmt.Scanln()
				break
			}

			if player.Hands[DefaultHandIndex].Bust() {
				fmt.Println(player.Name, HasBustText)
				fmt.Println(AnyKeyText)
				fmt.Scanln()
				break
			}

			if player.Hands[DefaultHandIndex].IsHard() {
				fmt.Println("Hard hand")
				total := player.Hands[DefaultHandIndex].Total()
				dealerCard := dealer.Hands[DefaultHandIndex].Cards[0].Name
				fmt.Println(strategy.Hard.Get(total, dealerCard))
			}

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
