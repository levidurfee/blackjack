package blackjack

type Suite struct {
	Name   string
	Symbol string
}

type Suites []Suite

var (
	Clubs    = Suite{Name: "Clubs", Symbol: "♣"}
	Diamonds = Suite{Name: "Diamonds", Symbol: "♦"}
	Hearts   = Suite{Name: "Hearts", Symbol: "♥"}
	Spades   = Suite{Name: "Spades", Symbol: "♠"}
)

var SuitesSlice = Suites{
	Clubs,
	Diamonds,
	Hearts,
	Spades,
}
