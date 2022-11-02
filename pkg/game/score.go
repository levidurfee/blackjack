package game

import (
	"git.x6c.co/go/blackjack/pkg/blackjack"
)

func score(player *blackjack.Player, dealer *blackjack.Player) *blackjack.Player {
	if player.Hands[handIndex].Bust() && !dealer.Hands[handIndex].Bust() {
		dealer.Score++
		return dealer
	}

	if !player.Hands[handIndex].Bust() && dealer.Hands[handIndex].Bust() {
		player.Score++
		return player
	}

	if player.Hands[handIndex].IsBlackjack() && !dealer.Hands[handIndex].IsBlackjack() {
		player.Score++
		return player
	}

	if !player.Hands[handIndex].IsBlackjack() && dealer.Hands[handIndex].IsBlackjack() {
		dealer.Score++
		return dealer
	}

	if player.Hands[handIndex].Total() > dealer.Hands[handIndex].Total() {
		player.Score++
		return player
	}

	if player.Hands[handIndex].Total() < dealer.Hands[handIndex].Total() {
		dealer.Score++
		return dealer
	}

	return nil
}
