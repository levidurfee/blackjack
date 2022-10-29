package blackjack

type Hand struct {
	Cards []Card
}

const (
	AceCardName = "Ace"
	AceLowValue = 1
	MaxTotal    = 21
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

func (h Hand) DealerTotal() int {
	return h.Cards[0].Value
}

// Bust checks if the hand has exceeded the max total allowed.
//
// It only checks the low value, since the low value will always be equal to or
// less than the high value.
func (h Hand) Bust() bool {
	low, _ := h.Total()

	return low > MaxTotal
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
