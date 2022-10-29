package blackjack

type Hand struct {
	Cards []Card
}

const (
	AceLowValue = 1
	MaxTotal    = 21
)

func (h Hand) Total() int {
	total := h.SoftTotal()

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
	if h.NumberOfAces() == 0 {
		return true
	}

	return MaxTotal-h.SoftTotal() < 10
}

// IsPair will return true if both of the first two cards are the same.
func (h Hand) IsPair() bool {
	if len(h.Cards) > 2 {
		return false
	}

	return h.Cards[0].Name == h.Cards[1].Name
}

// SoftTotal uses the low value of aces (1) instead of the high value (11) to
// calculate the total of the hand.
func (h Hand) SoftTotal() int {
	total := 0

	for _, card := range h.Cards {
		value := card.Value
		if card.Name == AceName {
			value = AceLowValue
		}

		total += value
	}

	return total
}
