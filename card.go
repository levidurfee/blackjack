package main

type Card struct {
	Name  string
	Value int

	Suite
}

type Cards []Card

func (c Cards) Total() (int, int) {
	total := 0
	// TODO: Add counter for alternate value of ace

	for _, card := range c {
		total += card.Value
	}

	return total, 0
}

var (
	Ace   = Card{Name: "Ace", Value: 11}
	King  = Card{Name: "King", Value: 10}
	Queen = Card{Name: "Queen", Value: 10}
	Jack  = Card{Name: "Jack", Value: 10}
	Ten   = Card{Name: "Ten", Value: 10}
	Nine  = Card{Name: "Nine", Value: 9}
	Eight = Card{Name: "Eight", Value: 8}
	Seven = Card{Name: "Seven", Value: 7}
	Six   = Card{Name: "Six", Value: 6}
	Five  = Card{Name: "Five", Value: 5}
	Four  = Card{Name: "Four", Value: 4}
	Three = Card{Name: "Three", Value: 3}
	Two   = Card{Name: "Two", Value: 2}
)

var CardsSlice = []Card{
	Ace,
	King,
	Queen,
	Jack,
	Ten,
	Nine,
	Eight,
	Seven,
	Six,
	Five,
	Four,
	Three,
	Two,
}
