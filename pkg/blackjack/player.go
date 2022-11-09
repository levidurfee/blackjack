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
		Hands:    NewHands(),
		Score:    0,
	}
}

func NewDealer() Player {
	return Player{
		Name:     "Dealer",
		IsDealer: true,
		Hands:    NewHands(),
		Score:    0,
	}
}

func (p Player) String() string {
	var str = ""
	for _, hand := range p.Hands {
		str += "==================================================\n"
		str += fmt.Sprintf("Name:\t%s\n", p.Name)
		str += fmt.Sprintf("Score:\t%d\n", p.Score)
		str += fmt.Sprintf("Total:\t%d\n", hand.Total())
		str += "Cards:\n"
		for _, card := range hand.Cards {
			str += "\t" + fmt.Sprint(card) + "\n"
		}
		str += "==================================================\n"
	}

	return str
}
