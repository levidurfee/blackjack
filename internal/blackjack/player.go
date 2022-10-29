package blackjack

import (
	"fmt"
)

type Player struct {
	Name  string
	Hands []Hand

	IsDealer bool
}

func NewPlayer(name string) Player {
	return Player{
		Name:     name,
		IsDealer: false,
		Hands: []Hand{
			{
				Cards: []Card{},
			},
		},
	}
}

func NewDealer() Player {
	return Player{
		Name:     "Dealer",
		IsDealer: true,
		Hands: []Hand{
			{
				Cards: []Card{},
			},
		},
	}
}

func (p Player) String() string {
	var str = ""
	for _, hand := range p.Hands {
		low, high := hand.Total()
		str += "==================================================\n"
		str += fmt.Sprintf("Name:\t%s\n", p.Name)
		if p.IsDealer {
			str += fmt.Sprintf("Total:\t%d\n", hand.DealerTotal())
			continue
		}
		str += fmt.Sprintf("Total:\t%d/%d\n", low, high)
		str += "Cards:\n"
		for _, card := range hand.Cards {
			str += fmt.Sprintf("\t%s %s\n", card.Suite.Symbol, card.Name)
		}
		str += "==================================================\n"
	}

	return str
}
