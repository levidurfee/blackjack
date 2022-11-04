package strategy

import (
	"git.x6c.co/go/blackjack/pkg/blackjack"
)

// StrategyTable is a map of the suggested actions based on the player's hand
// and the dealer's up-card. We use the total of the player's hand as an int,
// and the string name of the dealer's card.
//
// There are different strategies based on what type of hand the player has.
type StrategyTable map[int]map[string]blackjack.Action

func (s StrategyTable) Get(playerTotal int, dealerCard string) blackjack.Action {
	return s[playerTotal][dealerCard]
}

func Evaluate(hand blackjack.Hand, dealerCard string) blackjack.Action {
	if hand.IsBlackjack() {
		return blackjack.Stand
	}

	switch hand.Type() {
	case blackjack.PairHand:
		return Pair.Get(hand.Total(), dealerCard)
	case blackjack.HardHand:
		return Hard.Get(hand.Total(), dealerCard)
	case blackjack.SoftHand:
		// TODO: implement split action
		// For now use the hard hand strategy
		return Hard.Get(hand.Total(), dealerCard)
	}

	return blackjack.Stand
}
