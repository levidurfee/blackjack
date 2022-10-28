package main

import "fmt"

type Player struct {
	Name  string
	Cards []Card

	IsDealer bool
}

func NewPlayer(name string) Player {
	return Player{
		Name:     name,
		IsDealer: false,
	}
}

func NewDealer() Player {
	return Player{
		Name:     "Dealer",
		IsDealer: true,
	}
}

func (p Player) Total() (int, int) {
	low := 0
	high := 0

	for _, card := range p.Cards {
		high += card.Value
		if card.Name == "Ace" {
			low += 1
			continue
		}
		low += card.Value
	}

	return low, high
}

func (p Player) String() string {
	low, high := p.Total()
	str := "==================================================\n"
	str += fmt.Sprintf("Name:\t%s\n", p.Name)
	str += fmt.Sprintf("Total:\t%d/%d\n", low, high)
	str += "Cards:\n"
	for _, card := range p.Cards {
		str += fmt.Sprintf("\t%s %d %s\n", card.Name, card.Value, card.Suite.Name)
	}
	str += "==================================================\n"

	return str
}
