package deck

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallDeckToString(t *testing.T) {
	deck := NewEmptyDeck()
	deck.Cards = append(deck.Cards, Card{ACE, HEART}, Card{KING, HEART})
	result := fmt.Sprintf("%s", deck)
	assert.Equal(t, "A♥ K♥ ", result, "These should be equal")
}

func TestDeckToString(t *testing.T) {
	deck := NewDeck(false)
	result := fmt.Sprintf("%s", deck)
	assert.Equal(t, "A♣ 2♣ 3♣ 4♣ 5♣ 6♣ 7♣ 8♣ 9♣ T♣ J♣ Q♣ K♣ "+
		"A♦ 2♦ 3♦ 4♦ 5♦ 6♦ 7♦ 8♦ 9♦ T♦ J♦ Q♦ K♦ "+
		"A♥ 2♥ 3♥ 4♥ 5♥ 6♥ 7♥ 8♥ 9♥ T♥ J♥ Q♥ K♥ "+
		"A♠ 2♠ 3♠ 4♠ 5♠ 6♠ 7♠ 8♠ 9♠ T♠ J♠ Q♠ K♠ ", result, "These should be equal")
}

func TestNewTinyDeck(t *testing.T) {
	deck := NewSpecificDeck(true, FACES, []Suit{SPADE})
	assert.Equal(t, 13, deck.NumberOfCards())
}

func TestNewSmallDeck(t *testing.T) {
	deck := NewSpecificDeck(false, FACES, []Suit{SPADE, HEART})
	assert.Equal(t, 26, deck.NumberOfCards())
}

func BenchmarkTinyDeckShuffle(b *testing.B) {
	deck := NewSpecificDeck(false, FACES, []Suit{SPADE})
	for n := 0; n < b.N; n++ {
		deck.Shuffle()
	}
}

func BenchmarkSmallDeckShuffle(b *testing.B) {
	deck := NewSpecificDeck(false, FACES, []Suit{SPADE, HEART})
	for n := 0; n < b.N; n++ {
		deck.Shuffle()
	}
}

func BenchmarkMediumDeckShuffle(b *testing.B) {
	deck := NewSpecificDeck(false, FACES, []Suit{SPADE, HEART, DIAMOND})
	for n := 0; n < b.N; n++ {
		deck.Shuffle()
	}
}

func BenchmarkDeckShuffle(b *testing.B) {
	deck := NewSpecificDeck(false, FACES, SUITS)
	for n := 0; n < b.N; n++ {
		deck.Shuffle()
	}
}
