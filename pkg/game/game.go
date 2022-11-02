package game

import (
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
	deck   = blackjack.NewDecks(4)
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
	var cmd string
	for {
		deal()
		printPlayers()
		if !deck.HasEnoughCards() {
			return
		}

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
				deck.Deal(&player.Hands[handIndex])
			case blackjack.Double:
				deck.Deal(&player.Hands[handIndex])
			}
		}

		// TODO: implement dealer strategy

		winner := score(&player, &dealer)
		if winner != nil {
			fmt.Printf("%s won this hand\n", winner.Name)
		}

		fmt.Println("Another hand? y/n")
		fmt.Scanln(&cmd)
		if cmd == "n" {
			break
		}
	}
}
