package poker

import (
	"fmt"
	"strings"
	"testing"
)

// Go には Rust の unwrap のような機能は無いので自前で作成する.
func unwrap[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func Test_カードのsuitとrankを表示することができる(t *testing.T) {

	tests := []struct {
		suit string
		rank string
		want string
	}{
		{suit: "♠", rank: "3", want: "3♠"},
		{suit: "♥", rank: "J", want: "J♥"},
	}

	for _, tt := range tests {
		tt := tt
		name := fmt.Sprintf("want:%v", tt.want)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			card := Card{suit: unwrap(ofSuit(tt.suit)), rank: unwrap(ofRank(tt.rank))}

			if card.Notation() != tt.want {
				t.Errorf(`Card(1) is %q`, card)
			}
		})
	}
}

func Test_存在しないランクを指定するとエラーになる(t *testing.T) {
	tests := []struct {
		rank string
		want string
	}{
		{rank: "B", want: "failed to parse rank string:"},
	}

	for _, tt := range tests {
		tt := tt
		name := fmt.Sprintf("want:%v", tt.want)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var err error
			_, err = ofRank(tt.rank)

			println(err.Error())

			if err == nil {
				t.Errorf(`Error did not occurred. rank_string: %q`, tt.rank)
			}

			if !strings.HasPrefix(err.Error(), tt.want) {
				t.Errorf(`Unexpected error: %q`, err.Error())
			}
		})
	}

}

func Test_存在しないスートを指定するとエラーになる(t *testing.T) {
	tests := []struct {
		suit string
		want string
	}{
		{suit: "hoge", want: "failed to parse suit string"},
	}

	for _, tt := range tests {
		tt := tt
		name := fmt.Sprintf("want:%v", tt.want)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var err error
			_, err = ofSuit(tt.suit)

			println(err.Error())

			if err == nil {
				t.Errorf(`Error did not occurred. rank_string: %q`, tt.suit)
			}

			if !strings.HasPrefix(err.Error(), tt.want) {
				t.Errorf(`Unexpected error: %q`, err.Error())
			}
		})
	}

}

func Test_カードが同じsuitを持つか判定できる(t *testing.T) {

	tests := []struct {
		suit1 string
		rank1 string
		suit2 string
		rank2 string
		want  bool
	}{
		{suit1: "♠", rank1: "3", suit2: "♠", rank2: "A", want: true},
		{suit1: "♠", rank1: "5", suit2: "♥", rank2: "2", want: false},
	}

	for _, tt := range tests {
		tt := tt
		name := fmt.Sprintf("want:%v", tt.want)

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			card1 := Card{suit: unwrap(ofSuit(tt.suit1)), rank: unwrap(ofRank(tt.rank1))}
			card2 := Card{suit: unwrap(ofSuit(tt.suit2)), rank: unwrap(ofRank(tt.rank2))}

			if card1.hasSameSuit(card2) != tt.want {
				t.Errorf(`Card(1) is %q, Card(2) is %q`, card1, card2)
			}
		})
	}
}

func Test_カードが同じrankを持つか判定できる(t *testing.T) {

	tests := []struct {
		suit1 string
		rank1 string
		suit2 string
		rank2 string
		want  bool
	}{
		{suit1: "♥", rank1: "A", suit2: "♠", rank2: "A", want: true},
		{suit1: "♥", rank1: "5", suit2: "♥", rank2: "2", want: false},
	}

	for _, tt := range tests {
		tt := tt
		name := fmt.Sprintf("want:%v", tt.want)

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			var rank1 Rank
			var rank2 Rank
			rank1, _ = ofRank(tt.rank1)
			rank2, _ = ofRank(tt.rank2)

			card1 := Card{suit: unwrap(ofSuit(tt.suit1)), rank: rank1}
			card2 := Card{suit: unwrap(ofSuit(tt.suit2)), rank: rank2}

			if card1.hasSameRank(card2) != tt.want {
				t.Errorf(`Card(1) is %q, Card(2) is %q`, card1, card2)
			}
		})
	}
}

func Test_ツーカードポーカーの役を判定できる(t *testing.T) {
	tests := []struct {
		cards []Card
		want  Hand
	}{
		{cards: []Card{Card{"♥", unwrap(ofRank("A"))}, Card{"♠", unwrap(ofRank("A"))}}, want: Pair},
		{cards: []Card{Card{"♥", unwrap(ofRank("A"))}, Card{"♥", unwrap(ofRank("J"))}}, want: Flush},
		{cards: []Card{Card{"♥", unwrap(ofRank("A"))}, Card{"♠", unwrap(ofRank("J"))}}, want: HighCard},
		{cards: []Card{Card{"♥", unwrap(ofRank("A"))}, Card{"♠", unwrap(ofRank("2"))}}, want: Straight},
		{cards: []Card{Card{"♥", unwrap(ofRank("K"))}, Card{"♠", unwrap(ofRank("A"))}}, want: Straight},
		{cards: []Card{Card{"♥", unwrap(ofRank("A"))}, Card{"♥", unwrap(ofRank("2"))}}, want: StraightFlush},
	}

	for _, tt := range tests {
		tt := tt
		name := fmt.Sprintf("want:%v", tt.want)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			cards := Cards{Cards: tt.cards}

			if cards.hand() != tt.want {
				t.Errorf(`Cards is %q, want: %q, actual: %q`, cards, tt.want, cards.hand())
			}
		})
	}
}

func Test_ツーカードポーカーの強さを比較できる(t *testing.T) {
	tests := []struct {
		testName    string
		cardsPlayer []Card
		cardsEnemy  []Card
		want        PokerResult
	}{
		{
			testName:    "自分: フラッシュ / 敵: ハイカード だった場合は自分が勝つ",
			cardsPlayer: []Card{Card{"♥", unwrap(ofRank("A"))}, Card{"♥", unwrap(ofRank("3"))}},
			cardsEnemy:  []Card{Card{"♥", unwrap(ofRank("A"))}, Card{"♠", unwrap(ofRank("3"))}},
			want:        WIN,
		},
		{
			testName:    "自分: ハイカード / 敵: フラッシュ だった場合は敵が勝つ",
			cardsPlayer: []Card{Card{"♥", unwrap(ofRank("A"))}, Card{"♠", unwrap(ofRank("3"))}},
			cardsEnemy:  []Card{Card{"♥", unwrap(ofRank("A"))}, Card{"♥", unwrap(ofRank("3"))}},
			want:        LOSE,
		},
		{
			testName:    "自分: ハイカード, max-A / 敵: ハイカード, max-K だった場合は自分が勝つ",
			cardsPlayer: []Card{Card{"♥", unwrap(ofRank("A"))}, Card{"♠", unwrap(ofRank("3"))}},
			cardsEnemy:  []Card{Card{"♥", unwrap(ofRank("K"))}, Card{"♦︎", unwrap(ofRank("3"))}},
			want:        WIN,
		},
		{
			testName:    "自分: ハイカード, max-K / 敵: ハイカード, max-A だった場合は敵が勝つ",
			cardsPlayer: []Card{Card{"♥", unwrap(ofRank("K"))}, Card{"♠", unwrap(ofRank("3"))}},
			cardsEnemy:  []Card{Card{"♥", unwrap(ofRank("A"))}, Card{"♦︎", unwrap(ofRank("3"))}},
			want:        LOSE,
		},
		{
			testName:    "自分: ハイカード, max-A, 2nd max-K / 敵: ハイカード, max-A, 2nd max-Q だった場合は自分が勝つ",
			cardsPlayer: []Card{Card{"♥", unwrap(ofRank("A"))}, Card{"♠", unwrap(ofRank("K"))}},
			cardsEnemy:  []Card{Card{"♥", unwrap(ofRank("A"))}, Card{"♦︎", unwrap(ofRank("Q"))}},
			want:        WIN,
		},
		{
			testName:    "自分: ハイカード, max-A, 2nd max-Q / 敵: ハイカード, max-A, 2nd max-K だった場合は相手が勝つ",
			cardsPlayer: []Card{Card{"♥", unwrap(ofRank("A"))}, Card{"♠", unwrap(ofRank("Q"))}},
			cardsEnemy:  []Card{Card{"♥", unwrap(ofRank("A"))}, Card{"♦︎", unwrap(ofRank("K"))}},
			want:        LOSE,
		},
		{
			testName:    "自分: フラッシュ, max-A / 敵: フラッシュ, max-K だった場合は自分が勝つ",
			cardsPlayer: []Card{Card{"♥", unwrap(ofRank("A"))}, Card{"♥", unwrap(ofRank("K"))}},
			cardsEnemy:  []Card{Card{"♦", unwrap(ofRank("K"))}, Card{"♦︎", unwrap(ofRank("Q"))}},
			want:        WIN,
		},
		{
			testName:    "自分: フラッシュ, max-A, 2nd max-Q / 敵: フラッシュ, max-A, 2nd max-K だった場合は相手が勝つ",
			cardsPlayer: []Card{Card{"♥", unwrap(ofRank("A"))}, Card{"♥", unwrap(ofRank("Q"))}},
			cardsEnemy:  []Card{Card{"♦", unwrap(ofRank("A"))}, Card{"♦︎", unwrap(ofRank("K"))}},
			want:        LOSE,
		},
		{
			testName:    "自分: ストレート, A-K / 敵: ストレート, K-Q だった場合は自分が勝つ",
			cardsPlayer: []Card{Card{"♥", unwrap(ofRank("K"))}, Card{"♠", unwrap(ofRank("A"))}},
			cardsEnemy:  []Card{Card{"♦", unwrap(ofRank("K"))}, Card{"♠", unwrap(ofRank("Q"))}},
			want:        WIN,
		},
		{
			testName:    "自分: ストレート, A-K / 敵: ストレート, 2-A だった場合は自分が勝つ",
			cardsPlayer: []Card{Card{"♥", unwrap(ofRank("A"))}, Card{"♠", unwrap(ofRank("K"))}},
			cardsEnemy:  []Card{Card{"♦", unwrap(ofRank("A"))}, Card{"♠", unwrap(ofRank("2"))}},
			want:        WIN,
		},
		{
			testName:    "自分: ペア, K-K / 敵: ペア, A-A だった場合は相手が勝つ",
			cardsPlayer: []Card{Card{"♥", unwrap(ofRank("K"))}, Card{"♠", unwrap(ofRank("K"))}},
			cardsEnemy:  []Card{Card{"♦", unwrap(ofRank("A"))}, Card{"♠", unwrap(ofRank("A"))}},
			want:        LOSE,
		},
		{
			testName:    "自分: ストレートフラッシュ, A-K / 敵: ストレートフラッシュ, A-2 だった場合は自分が勝つ",
			cardsPlayer: []Card{Card{"♥", unwrap(ofRank("A"))}, Card{"♥", unwrap(ofRank("K"))}},
			cardsEnemy:  []Card{Card{"♦", unwrap(ofRank("A"))}, Card{"♦", unwrap(ofRank("2"))}},
			want:        WIN,
		},
	}

	for _, tt := range tests {
		tt := tt
		name := fmt.Sprintf("%v", tt.testName)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			cardsPlayer := Cards{Cards: tt.cardsPlayer}
			cardsEnemy := Cards{Cards: tt.cardsEnemy}
			result := cardsPlayer.battle(cardsEnemy)

			if result != tt.want {
				t.Errorf(`cardsPlayer is %q, cardsEnemy is %q, want: %q, actual: %q`, cardsPlayer, cardsEnemy, tt.want, result)
			}
		})
	}
}
