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
	number   int
	strength int
	display  string
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
	strength := i - 1
	if rankString == "A" {
		strength = 13
	}
	return Rank{number: i, strength: strength, display: rankString}
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

func (p *Card) isStraight(card Card) bool {
	return math.Abs(float64(p.rank.number-card.rank.number)) == 1 || math.Abs(float64(p.rank.number-card.rank.number)) == 12
}

type Cards struct {
	Cards []Card
}

// refer to: https://zenn.dev/kyoh86/articles/qiita-18b8bfc6ffe045aaf380

type Hand int

const (
	StraightFlush = Hand(5)
	Pair          = Hand(4)
	Straight      = Hand(3)
	Flush         = Hand(2)
	HighCard      = Hand(1)
)

func (p *Cards) hand() Hand {
	if p.Cards[0].hasSameRank(p.Cards[1]) {
		return Pair
	}
	if p.Cards[0].hasSameSuit(p.Cards[1]) && p.Cards[0].isStraight(p.Cards[1]) {
		return StraightFlush
	}
	if p.Cards[0].hasSameSuit(p.Cards[1]) {
		return Flush
	}
	if p.Cards[0].isStraight(p.Cards[1]) {
		return Straight
	}

	return HighCard
}

type PokerResult string

const (
	WIN  = PokerResult("WIN")
	DRAW = PokerResult("DRAW")
	LOSE = PokerResult("WIN")
)

func (p *Cards) battle(enemy Cards) PokerResult {
	if p.hand() > enemy.hand() {
		return WIN
	}
	if p.hand() < enemy.hand() {
		return LOSE
	}
	// 一番強いランクを取得して比較
	playlerMaxRank := p.Cards[0].rank
	for _, card := range p.Cards {
		if card.rank.strength > playlerMaxRank.strength {
			playlerMaxRank = card.rank
		}
	}

	enemyMaxRank := enemy.Cards[0].rank
	for _, card := range enemy.Cards {
		if card.rank.strength > enemyMaxRank.strength {
			enemyMaxRank = card.rank
		}
	}

	if playlerMaxRank.strength > enemyMaxRank.strength {
		return WIN
	}

	if playlerMaxRank.number < enemyMaxRank.number {
		return LOSE
	}
	//
	//// 二番目に強いランクを取得して比較
	return DRAW
}
