package blackjack

import (
	"fmt"
)

type Player struct {
	Name  string
	Hands []Hand
	Score int

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
		Score: 0,
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
		Score: 0,
	}
}

func (p Player) String() string {
	var str = ""
	for _, hand := range p.Hands {
		str += "==================================================\n"
		str += fmt.Sprintf("Name:\t%s\n", p.Name)
		str += fmt.Sprintf("Score:\t%d\n", p.Score)
		if p.IsDealer {
			str += fmt.Sprintf("Total:\t%d\n", hand.DealerTotal())
			str += fmt.Sprintf("Up:\t%s\n", hand.DealerUpCard())
			continue
		}
		str += fmt.Sprintf("Total:\t%d\n", hand.Total())
		str += "Cards:\n"
		for _, card := range hand.Cards {
			str += "\t" + fmt.Sprint(card) + "\n"
		}
		str += "==================================================\n"
	}

	return str
}
