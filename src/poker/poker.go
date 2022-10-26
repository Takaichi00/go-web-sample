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

// refer to: https://zenn.dev/kyoh86/articles/qiita-18b8bfc6ffe045aaf380

type Hand string

const (
	Pair     = Hand("Pair")
	Flush    = Hand("Flush")
	HighCard = Hand("High Card")
)

func (p *Cards) hand() Hand {

	for i := 0; i < len(p.Cards)-1; i++ {
		for j := i + 1; j < len(p.Cards); j++ {
			if p.Cards[i].hasSameRank(p.Cards[j]) {
				return Pair
			}
			if p.Cards[i].hasSameSuit(p.Cards[j]) {
				return Flush
			}
		}
	}
	return HighCard
}
