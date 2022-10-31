package blackjack

type Hand struct {
	Cards []Card
}

const (
	MaxTotal           = 21
	AceValueDifference = AceValue - AceLowValue
)

// Total calculates the total score of the hand. If an ace can be worth 11, then
// it will calculate the total with at most one ace as an 11. So, if the other
// cards only add up to 10 or less, then it'll allow an ace to be worth 11.
func (h Hand) Total() int {
	total := h.SoftTotal()

	if total+AceValueDifference <= MaxTotal && h.HasAces() {
		// Only one ace can have the value of 11 in one hand, if two aces were
		// 11, then you'd bust with at least 22.
		total += AceValueDifference
	}

	return total
}

// NumberOfAces counts the number of aces in the hand.
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

// HasAces uses NumberOfAces to determine if there are any aces in the hand.
//
// We could iterate over the cards and return early if an ace is found, but
// since we already have the `NumberOfAces` method, we might as well use it.
// Otherwise we'd duplicate a lot of the code with not much benefit. When
// iterating over every card in the hand, it'll never be more than 22 cards.
// If someone had 22 aces, then they'd bust and they couldn't get anymore
// cards in their hand.
func (h Hand) HasAces() bool {
	return h.NumberOfAces() > 0
}

// DealerTotal returns the total for the dealer. We only return the value for
// the first card, since the other cards aren't revealed until it is no longer
// the player's turn.
func (h Hand) DealerTotal() int {
	return h.Cards[0].Value
}

// Bust checks if the hand has exceeded the max total allowed, `MaxTotal`, 21.
//
// It only checks the low value, since the low value will always be equal to or
// less than the high value.
func (h Hand) Bust() bool {
	return h.Total() > MaxTotal
}

// IsHard will return `true` if the hand is hard and `soft` if the hand is soft.
//
// A hand is considered soft if it has an ace that can be counted as an 11. If a
// hand does not have any aces, then it is hard. Otherwise, we get the total of
// the hand using the 1 for the value of the ace(s); if 21 minus the soft total
// is less than 10, then we know one of the aces can be counted as 11.
//
//	21 - SoftTotal < 10
//
//	21 - 10 = 11 // This would be a soft hand.
//	21 - 12 = 9  // This would be a hard hand.
func (h Hand) IsHard() bool {
	if h.NumberOfAces() == 0 {
		return true
	}

	return MaxTotal-h.SoftTotal() < AceValueDifference
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
