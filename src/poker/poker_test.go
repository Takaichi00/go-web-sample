package poker

import (
	"testing"
)

func TestCard(t *testing.T) {
	card := Card{Suit: "♠", Rank: "3"}
	if card.Notation() != "3♠" {
		t.Errorf(`Card(1) is %q`, card)
	}
}
