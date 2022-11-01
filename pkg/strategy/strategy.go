package strategy

import "git.x6c.co/go/blackjack/pkg/blackjack"

// StrategyTable is a map of the suggested actions based on the player's hand
// and the dealer's up-card. We use the total of the player's hand as an int,
// and the string name of the dealer's card.
//
// There are different strategies based on what type of hand the player has.
type StrategyTable map[int]map[string]blackjack.Action

func (s StrategyTable) Get(playerTotal int, dealerCard string) blackjack.Action {
	return s[playerTotal][dealerCard]
}
