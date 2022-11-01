package blackjack

type Action int

const (
	Hit Action = iota
	Stand
	Double
	Split
)

func (a Action) String() string {
	switch a {
	case Hit:
		return "hit"
	case Stand:
		return "stand"
	case Double:
		return "double"
	case Split:
		return "split"
	}

	return "undefined"
}
