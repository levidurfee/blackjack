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
		// Loop for new hands.
		for {
			// Deal new hands

			// Show player hand and dealer face-up card

			// Player's turn
			for {
				// Ask player what they wanna do
				fmt.Scanln(&cmd)
				if cmd == "s" {
					break
				}
			}

			// Dealer's turn, this will be automated based on the dealer strategy
			// for {

			// }
		}
		// Ask if they want another hand or if they wanna quit.
		fmt.Scanln(&cmd)
		switch cmd {
		case "q":
			return
		}
		// If they wanna keep playing, empty the current hands
	}
}
