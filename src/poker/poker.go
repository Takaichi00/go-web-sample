package poker

import (
	"math"
	"strconv"
)

type Card struct {
	suit string
	rank Rank
}

type Rank struct {
	number  int
	display string
}

func ofRank(rankString string) Rank {
	var i int
	if rankString == "A" {
		i = 1
	} else if rankString == "J" {
		i = 11
	} else if rankString == "Q" {
		i = 12
	} else if rankString == "K" {
		i = 13
	} else {
		i, _ = strconv.Atoi(rankString)
	}
	return Rank{number: i, display: rankString}
}

// TODO rank を内部的には数字で持ち、不正な値だった場合はエラーを返す

func (p *Card) Notation() string {
	return p.rank.display + p.suit
}

func (p *Card) hasSameSuit(card Card) bool {
	return p.suit == card.suit
}

func (p *Card) hasSameRank(card Card) bool {
	return p.rank == card.rank
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
	if math.Abs(float64(p.Cards[0].rank.number-p.Cards[1].rank.number)) == 1 {
		return Straight
	}
	return HighCard
}
