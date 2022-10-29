package blackjack

type Card struct {
	Name  string
	Value int

	Suite
}

const (
	AceName   = "Ace"
	KingName  = "King"
	QueenName = "Queen"
	JackName  = "Jack"
	TenName   = "Ten"
	NineName  = "Nine"
	EightName = "Eight"
	SevenName = "Seven"
	SixName   = "Six"
	FiveName  = "Five"
	FourName  = "Four"
	ThreeName = "Three"
	TwoName   = "Two"
)

const (
	AceValue    = 11
	AceLowValue = 1
	FaceValue   = 10
	TenValue    = 10
	NineValue   = 9
	EightValue  = 8
	SevenValue  = 7
	SixValue    = 6
	FiveValue   = 5
	FourValue   = 4
	ThreeValue  = 3
	TwoValue    = 2
)

var (
	Ace   = Card{Name: AceName, Value: AceValue}
	King  = Card{Name: KingName, Value: FaceValue}
	Queen = Card{Name: QueenName, Value: FaceValue}
	Jack  = Card{Name: JackName, Value: FaceValue}
	Ten   = Card{Name: TenName, Value: TenValue}
	Nine  = Card{Name: NineName, Value: NineValue}
	Eight = Card{Name: EightName, Value: EightValue}
	Seven = Card{Name: SevenName, Value: SevenValue}
	Six   = Card{Name: SixName, Value: SixValue}
	Five  = Card{Name: FiveName, Value: FiveValue}
	Four  = Card{Name: FourName, Value: FourValue}
	Three = Card{Name: ThreeName, Value: ThreeValue}
	Two   = Card{Name: TwoName, Value: TwoValue}
)

var CardsSlice = []Card{
	Ace,
	King,
	Queen,
	Jack,
	Ten,
	Nine,
	Eight,
	Seven,
	Six,
	Five,
	Four,
	Three,
	Two,
}
