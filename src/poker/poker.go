package poker

type Card struct {
	Suit string
	Rank string
}

// TODO Rank を内部的には数字で持ち、不正な値だった場合はエラーを返す

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
	Straight = Hand("Straight")
)

func (p *Cards) hand() Hand {
	if p.Cards[0].hasSameRank(p.Cards[1]) {
		return Pair
	}
	if p.Cards[0].hasSameSuit(p.Cards[1]) {
		return Flush
	}
	return HighCard
}
