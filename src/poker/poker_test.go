package poker

import (
	"fmt"
	"testing"
)

func TestCard(t *testing.T) {

	tests := []struct {
		suit string
		rank string
		want string
	}{
		{suit: "♠", rank: "3", want: "3♠"},
	}

	for _, tt := range tests {
		tt := tt
		name := fmt.Sprintf("want:%v", tt.want)

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			card := Card{Suit: tt.suit, Rank: tt.rank}
			if card.Notation() != "3♠" {
				t.Errorf(`Card(1) is %q`, card)
			}
		})
	}

}
