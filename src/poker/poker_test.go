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
		cardsPlayer []Card
		cardsEnemy  []Card
		want        PokerResult
	}{
		{
			cardsPlayer: []Card{Card{"♥", ofRank("A")}, Card{"♥", ofRank("3")}},
			cardsEnemy:  []Card{Card{"♥", ofRank("A")}, Card{"♠", ofRank("3")}},
			want:        WIN,
		},
		{
			cardsPlayer: []Card{Card{"♥", ofRank("A")}, Card{"♠", ofRank("3")}},
			cardsEnemy:  []Card{Card{"♥", ofRank("A")}, Card{"♥", ofRank("3")}},
			want:        LOSE,
		},
		{
			cardsPlayer: []Card{Card{"♥", ofRank("A")}, Card{"♠", ofRank("3")}},
			cardsEnemy:  []Card{Card{"♥", ofRank("K")}, Card{"♦︎", ofRank("3")}},
			want:        WIN,
		},

		//{
		//	cardsPlayer: []Card{Card{"♥", ofRank("A")}, Card{"♦", ofRank("3")}},
		//	cardsEnemy:  []Card{Card{"♠", ofRank("A")}, Card{"♥", ofRank("4")}},
		//	want:        LOSE,
		//},
		//
		//{
		//	cardsPlayer: []Card{Card{"♥", ofRank("A")}, Card{"♦", ofRank("3")}},
		//	cardsEnemy:  []Card{Card{"♠", ofRank("A")}, Card{"♥", ofRank("3")}},
		//	want:        DRAW,
		//},
	}

	for _, tt := range tests {
		tt := tt
		name := fmt.Sprintf("want:%v", tt.want)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			cardsPlayer := Cards{Cards: tt.cardsPlayer}
			cardsEnemy := Cards{Cards: tt.cardsEnemy}

			if cardsPlayer.battle(cardsEnemy) != tt.want {
				t.Errorf(`cardsPlayer is %q, cardsPlayer is %q, want: %q, actual: %q`, cardsPlayer, cardsEnemy, tt.want, cardsPlayer.battle(cardsEnemy))
			}
		})
	}
}
