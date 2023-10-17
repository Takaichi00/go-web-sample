package poker

import (
	"errors"
	"fmt"
	"github.com/samber/lo"
	"math"
	"strconv"
)

type Card struct {
	suit Suit
	rank Rank
}

type Suit string

const (
	SPADE   = Suit("♠")
	HEART   = Suit("♥")
	DIAMOND = Suit("♦")
	CLUB    = Suit("♣")
)

func ofSuit(suitString string) (Suit, error) {
	if suitString == "♠" {
		return SPADE, nil
	}
	if suitString == "♥" {
		return HEART, nil
	}
	if suitString == "♦" {
		return DIAMOND, nil
	}
	if suitString == "♣" {
		return CLUB, nil
	}
	return "", errors.New("failed to parse suit string")
}

type Rank struct {
	number   int
	strength int
	display  string
}

func ofRank(rankString string) (Rank, error) {
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
		var err error
		i, err = strconv.Atoi(rankString)
		if err != nil {
			return Rank{}, fmt.Errorf("failed to parse rank string: %w", err)
		}
		if i < 2 || i > 10 {
			return Rank{}, errors.New("rank number out of range")
		}
	}
	strength := i - 1
	if rankString == "A" {
		strength = 13
	}
	return Rank{number: i, strength: strength, display: rankString}, nil
}

func (p *Card) Notation() string {
	return p.rank.display + string(p.suit)
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
	LOSE = PokerResult("LOSE")
)

func (p *Cards) battle(enemy Cards) PokerResult {
	if p.hand() > enemy.hand() {
		return WIN
	}
	if p.hand() < enemy.hand() {
		return LOSE
	}
	// 一番強いランクを取得して比較
	playerMaxRankCard := lo.MaxBy(p.Cards, func(card Card, max Card) bool {
		return card.rank.strength > max.rank.strength
	})

	enemyMaxRankCard := lo.MaxBy(enemy.Cards, func(card Card, max Card) bool {
		return card.rank.strength > max.rank.strength
	})

	if playerMaxRankCard.rank.strength > enemyMaxRankCard.rank.strength {
		return WIN
	}

	if playerMaxRankCard.rank.strength < enemyMaxRankCard.rank.strength {
		return LOSE
	}

	// 二番目に強いランクを取得して比較
	playerSecondMaxRankCard := lo.MaxBy(lo.Filter(p.Cards, func(card Card, index int) bool {
		return card != playerMaxRankCard
	}), func(card Card, max Card) bool {
		return card.rank.strength > max.rank.strength
	})

	enemySecondMaxRankCard := lo.MaxBy(lo.Filter(enemy.Cards, func(card Card, index int) bool {
		return card != enemyMaxRankCard
	}), func(card Card, max Card) bool {
		return card.rank.strength > max.rank.strength
	})

	if playerSecondMaxRankCard.rank.strength > enemySecondMaxRankCard.rank.strength {
		return WIN
	}

	if playerSecondMaxRankCard.rank.strength < enemySecondMaxRankCard.rank.strength {
		return LOSE
	}

	return DRAW
}
