package blackjack

import (
	"errors"
	"math/rand"
	"sort"
	"time"
)

var (
	ErrOutOfCards = errors.New("out of cards")
)

const (
	MinCardsToPlay = 4
)

type Deck struct {
	Cards []Card
}

var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixMicro()))

func (d *Deck) Shuffle() {
	sort.Sort(d)
}

func NewDeck() Deck {
	deck := Deck{}

	for _, card := range CardsSlice {
		for _, suit := range SuitsSlice {
			card.Suit = suit
			deck.Cards = append(deck.Cards, card)
		}
	}

	return deck
}

func NewDecks(t int) Deck {
	var deck Deck

	for i := 0; i < t; i++ {
		d := NewDeck()
		d.Shuffle()

		deck.Cards = append(deck.Cards, d.Cards...)
	}

	return deck
}

func (d Deck) Len() int      { return len(d.Cards) }
func (d Deck) Swap(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] }
func (d Deck) Less(i, j int) bool {
	k := r.Int()
	l := r.Int()

	return k > l
}

func (d *Deck) Pop() Card {
	card, cards := d.Cards[0], d.Cards[1:]
	d.Cards = cards

	return card
}

func (d *Deck) Deal(hand *Hand) (int, error) {
	if len(d.Cards) <= 0 {
		return 0, ErrOutOfCards
	}

	hand.Cards = append(hand.Cards, d.Pop())

	return len(d.Cards), nil
}

func (d Deck) HasEnoughCards() bool {
	return len(d.Cards) > MinCardsToPlay
}
