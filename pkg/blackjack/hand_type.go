package blackjack

type HandType int

const (
	HardHand HandType = iota
	SoftHand
	PairHand
)

func (h HandType) String() string {
	switch h {
	case HardHand:
		return "hard"
	case SoftHand:
		return "soft"
	case PairHand:
		return "pair"
	}

	return "undefined"
}
