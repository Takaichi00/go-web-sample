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
			card := Card{Suit: tt.suit, Rank: tt.rank}
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
			card1 := Card{Suit: tt.suit1, Rank: tt.rank1}
			card2 := Card{Suit: tt.suit2, Rank: tt.rank2}

			if card1.hasSameSuit(card2) != tt.want {
				t.Errorf(`Card(1) is %q, Card(2) is %q`, card1, card2)
			}
		})
	}
}
