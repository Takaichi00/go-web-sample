package poker

import (
	"fmt"
	"testing"
)

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
			card := Card{suit: tt.suit, rank: ofRank(tt.rank)}
			if card.Notation() != tt.want {
				t.Errorf(`Card(1) is %q`, card)
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
			card1 := Card{suit: tt.suit1, rank: ofRank(tt.rank1)}
			card2 := Card{suit: tt.suit2, rank: ofRank(tt.rank2)}

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
			card1 := Card{suit: tt.suit1, rank: ofRank(tt.rank1)}
			card2 := Card{suit: tt.suit2, rank: ofRank(tt.rank2)}

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
		{cards: []Card{Card{"♥", ofRank("A")}, Card{"♠", ofRank("A")}}, want: Pair},
		{cards: []Card{Card{"♥", ofRank("A")}, Card{"♥", ofRank("J")}}, want: Flush},
		{cards: []Card{Card{"♥", ofRank("A")}, Card{"♠", ofRank("J")}}, want: HighCard},
		{cards: []Card{Card{"♥", ofRank("A")}, Card{"♠", ofRank("2")}}, want: Straight},
		{cards: []Card{Card{"♥", ofRank("K")}, Card{"♠", ofRank("A")}}, want: Straight},
		{cards: []Card{Card{"♥", ofRank("A")}, Card{"♥", ofRank("2")}}, want: StraightFlush},
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
			cardsPlayer: []Card{Card{"♥", ofRank("A")}, Card{"♥", ofRank("3")}},
			cardsEnemy:  []Card{Card{"♥", ofRank("A")}, Card{"♠", ofRank("3")}},
			want:        WIN,
		},
		{
			testName:    "自分: ハイカード / 敵: フラッシュ だった場合は敵が勝つ",
			cardsPlayer: []Card{Card{"♥", ofRank("A")}, Card{"♠", ofRank("3")}},
			cardsEnemy:  []Card{Card{"♥", ofRank("A")}, Card{"♥", ofRank("3")}},
			want:        LOSE,
		},
		{
			testName:    "自分: ハイカード, max-A / 敵: ハイカード, max-K だった場合は自分が勝つ",
			cardsPlayer: []Card{Card{"♥", ofRank("A")}, Card{"♠", ofRank("3")}},
			cardsEnemy:  []Card{Card{"♥", ofRank("K")}, Card{"♦︎", ofRank("3")}},
			want:        WIN,
		},
		{
			testName:    "自分: ハイカード, max-K / 敵: ハイカード, max-A だった場合は敵が勝つ",
			cardsPlayer: []Card{Card{"♥", ofRank("K")}, Card{"♠", ofRank("3")}},
			cardsEnemy:  []Card{Card{"♥", ofRank("A")}, Card{"♦︎", ofRank("3")}},
			want:        LOSE,
		},
		{
			testName:    "自分: ハイカード, max-A, 2nd max-K / 敵: ハイカード, max-A, 2nd max-Q だった場合は自分が勝つ",
			cardsPlayer: []Card{Card{"♥", ofRank("A")}, Card{"♠", ofRank("K")}},
			cardsEnemy:  []Card{Card{"♥", ofRank("A")}, Card{"♦︎", ofRank("Q")}},
			want:        WIN,
		},
		{
			testName:    "自分: ハイカード, max-A, 2nd max-Q / 敵: ハイカード, max-A, 2nd max-K だった場合は相手が勝つ",
			cardsPlayer: []Card{Card{"♥", ofRank("A")}, Card{"♠", ofRank("Q")}},
			cardsEnemy:  []Card{Card{"♥", ofRank("A")}, Card{"♦︎", ofRank("K")}},
			want:        LOSE,
		},
		{
			testName:    "自分: フラッシュ, max-A / 敵: フラッシュ, max-K だった場合は自分が勝つ",
			cardsPlayer: []Card{Card{"♥", ofRank("A")}, Card{"♥", ofRank("K")}},
			cardsEnemy:  []Card{Card{"♦", ofRank("K")}, Card{"♦︎", ofRank("Q")}},
			want:        WIN,
		},
		{
			testName:    "自分: フラッシュ, max-A, 2nd max-Q / 敵: フラッシュ, max-A, 2nd max-K だった場合は相手が勝つ",
			cardsPlayer: []Card{Card{"♥", ofRank("A")}, Card{"♥", ofRank("Q")}},
			cardsEnemy:  []Card{Card{"♦", ofRank("A")}, Card{"♦︎", ofRank("K")}},
			want:        LOSE,
		},
		{
			testName:    "自分: ストレート, A-K / 敵: ストレート, K-Q だった場合は自分が勝つ",
			cardsPlayer: []Card{Card{"♥", ofRank("K")}, Card{"♠", ofRank("A")}},
			cardsEnemy:  []Card{Card{"♦", ofRank("K")}, Card{"♠", ofRank("Q")}},
			want:        WIN,
		},
		{
			testName:    "自分: ストレート, A-K / 敵: ストレート, 2-A だった場合は自分が勝つ",
			cardsPlayer: []Card{Card{"♥", ofRank("A")}, Card{"♠", ofRank("K")}},
			cardsEnemy:  []Card{Card{"♦", ofRank("A")}, Card{"♠", ofRank("2")}},
			want:        WIN,
		},
		{
			testName:    "自分: ペア, K-K / 敵: ペア, A-A だった場合は相手が勝つ",
			cardsPlayer: []Card{Card{"♥", ofRank("K")}, Card{"♠", ofRank("K")}},
			cardsEnemy:  []Card{Card{"♦", ofRank("A")}, Card{"♠", ofRank("A")}},
			want:        LOSE,
		},
		{
			testName:    "自分: ストレートフラッシュ, A-K / 敵: ペア, A-2 だった場合は自分が勝つ",
			cardsPlayer: []Card{Card{"♥", ofRank("A")}, Card{"♥", ofRank("K")}},
			cardsEnemy:  []Card{Card{"♦", ofRank("A")}, Card{"♦", ofRank("2")}},
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
