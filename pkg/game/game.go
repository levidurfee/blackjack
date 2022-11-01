package game

import "fmt"

const (
	cls = "\033[H\033[2J"
)

func Run() {
	// Setup stuff, welcome message, create players
	fmt.Print(cls)

	var cmd string
	// Main loop for the game.
	for {
		fmt.Print(cls)
		// Player's turn
		for {
			// Check if player bust
			// Check if player has blackjack

			// Show player their options
			fmt.Println("hit, (s)tand, double")
			fmt.Scanln(&cmd)
			if cmd == "s" {
				break
			}
		}

		// Dealer's turn
		// for {
		// Check if dealer bust
		// Check if dealer has blackjack
		// }

		// Determine winner

		fmt.Println("Another hand? y/n")
		fmt.Scanln(&cmd)
		if cmd == "n" {
			return
		}
		// hand.Empty()
	}
}
