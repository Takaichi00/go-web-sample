package poker

type Card struct {
	Suit string
	Rank string
}

func (p *Card) Notation() string {
	return p.Rank + p.Suit
}
