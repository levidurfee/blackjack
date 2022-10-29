package blackjack

const (
	ClubsName    = "Clubs"
	DiamondsName = "Diamonds"
	HeartsName   = "Hearts"
	SpadesName   = "Spades"
)

const (
	ClubsSymbol    = "♣"
	DiamondsSymbol = "♦"
	HeartsSymbol   = "♥"
	SpadesSymbol   = "♠"
)

type Suit struct {
	Name   string
	Symbol string
}

type Suits []Suit

var (
	Clubs    = Suit{Name: ClubsName, Symbol: ClubsSymbol}
	Diamonds = Suit{Name: DiamondsName, Symbol: DiamondsSymbol}
	Hearts   = Suit{Name: HeartsName, Symbol: HeartsSymbol}
	Spades   = Suit{Name: SpadesName, Symbol: SpadesSymbol}
)

var SuitsSlice = Suits{
	Clubs,
	Diamonds,
	Hearts,
	Spades,
}
