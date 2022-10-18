package poker

type Card struct {
	Suit string
	Rank string
}

func (p *Card) Notation() string {
	return p.Rank + p.Suit
}

func (p *Card) hasSameSuit(card Card) bool {
	return p.Suit == card.Suit
}

func (p *Card) hasSameRank(card Card) bool {
	return p.Rank == card.Rank
}

type Cards struct {
	Cards []Card
}

func (p *Cards) isPair() bool {
	for i := 0; i < len(p.Cards)-1; i++ {
		for j := i + 1; j < len(p.Cards); j++ {
			if p.Cards[i].hasSameRank(p.Cards[j]) {
				return true
			}
		}
	}
	return false
}
