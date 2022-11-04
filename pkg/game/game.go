package game

import (
	"errors"
	"fmt"

	"git.x6c.co/go/blackjack/pkg/blackjack"
	"git.x6c.co/go/blackjack/pkg/strategy"
)

const (
	cls       = "\033[H\033[2J"
	handIndex = blackjack.DefaultHandIndex
)

var (
	dealer = blackjack.NewDealer()
	player = blackjack.NewPlayer("Levi")
	deck   = blackjack.NewDecks(8)
)

func clear() {
	fmt.Print(cls)
}

func deal() {
	player.Hands = blackjack.NewHand()
	dealer.Hands = blackjack.NewHand()

	deck.Deal(&player.Hands[handIndex])
	deck.Deal(&dealer.Hands[handIndex])

	deck.Deal(&player.Hands[handIndex])
	deck.Deal(&dealer.Hands[handIndex])
}

func printPlayers() {
	clear()
	fmt.Print(dealer)
	fmt.Print(player)
}

func Run() {
	for {
		if !deck.HasEnoughCards() {
			return
		}

		deal()
		printPlayers()

	actionloop:
		for {
			printPlayers()
			if player.Hands[handIndex].Bust() {
				break actionloop
			}
			if player.Hands[handIndex].IsBlackjack() {
				break actionloop
			}
			action := strategy.Evaluate(player.Hands[handIndex], dealer.Hands[handIndex].DealerUpCard())

			switch action {
			case blackjack.Stand:
				break actionloop
			case blackjack.Hit:
				_, err := deck.Deal(&player.Hands[handIndex])
				if errors.Is(err, blackjack.ErrOutOfCards) {
					panic(err.Error())
				}
			case blackjack.Double:
				_, err := deck.Deal(&player.Hands[handIndex])
				if errors.Is(err, blackjack.ErrOutOfCards) {
					panic(err.Error())
				}
			case blackjack.Split:
				// TODO: implement split action
			}
		}

		for {
			printPlayers()
			if dealer.Hands[handIndex].Total() >= strategy.DealerStand {
				break
			}

			_, err := deck.Deal(&dealer.Hands[handIndex])
			if errors.Is(err, blackjack.ErrOutOfCards) {
				panic(err.Error())
			}
		}

		winner := score(&player, &dealer)
		if winner != nil {
			fmt.Printf("%s won this hand with %d\n", winner.Name, winner.Hands[handIndex].Total())
		}
	}
}
