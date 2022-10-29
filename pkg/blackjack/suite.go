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

type Suite struct {
	Name   string
	Symbol string
}

type Suites []Suite

var (
	Clubs    = Suite{Name: ClubsName, Symbol: ClubsSymbol}
	Diamonds = Suite{Name: DiamondsName, Symbol: DiamondsSymbol}
	Hearts   = Suite{Name: HeartsName, Symbol: HeartsSymbol}
	Spades   = Suite{Name: SpadesName, Symbol: SpadesSymbol}
)

var SuitesSlice = Suites{
	Clubs,
	Diamonds,
	Hearts,
	Spades,
}
