package main

type Hand struct {
	Cards []Card
}

const (
	AceCardName = "Ace"
	AceLowValue = 1
)

func (h Hand) Total() (int, int) {
	low := 0
	high := 0

	for _, card := range h.Cards {
		high += card.Value
		if card.Name == AceCardName {
			low += AceLowValue
			continue
		}
		low += card.Value
	}

	return low, high
}
