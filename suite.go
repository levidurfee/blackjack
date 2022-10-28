package main

type Suite struct {
	Name string
}

type Suites []Suite

var (
	Clubs    = Suite{Name: "Clubs"}
	Diamonds = Suite{Name: "Diamonds"}
	Hearts   = Suite{Name: "Hearts"}
	Spades   = Suite{Name: "Spades"}
)

var SuitesSlice = Suites{
	Clubs,
	Diamonds,
	Hearts,
	Spades,
}
