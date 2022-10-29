package blackjack

type Hand struct {
	Cards []Card
}

const (
	AceLowValue = 1
	MaxTotal    = 21
)

func (h Hand) Total() int {
	total := 0

	for _, card := range h.Cards {
		value := card.Value
		if card.Name == AceName {
			value = AceLowValue
		}

		total += value
	}

	if total+10 <= MaxTotal && h.NumberOfAces() > 0 {
		// Only one ace can have the value of 11 if one hand, if two aces were
		// 11, then you'd bust with at least 22.
		total += 10
	}

	return total
}

func (h Hand) NumberOfAces() int {
	ctr := 0

	for _, card := range h.Cards {
		if card.Name != AceName {
			continue
		}
		ctr++
	}

	return ctr
}

func (h Hand) DealerTotal() int {
	return h.Cards[0].Value
}

// Bust checks if the hand has exceeded the max total allowed.
//
// It only checks the low value, since the low value will always be equal to or
// less than the high value.
func (h Hand) Bust() bool {
	return h.Total() > MaxTotal
}

// IsHard will return true if the hand does not have an Ace.
func (h Hand) IsHard() bool {
	for _, card := range h.Cards {
		if card.Name == AceCardName {
			return false
		}
	}

	return true
}

// IsPair will return true if both of the first two cards are the same.
func (h Hand) IsPair() bool {
	if len(h.Cards) > 2 {
		return false
	}

	return h.Cards[0].Name == h.Cards[1].Name
}
