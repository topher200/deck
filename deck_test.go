package deck

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewDeck(t *testing.T) {
	deck := NewDeck(true)
	assert.Equal(t, len(deck.Cards), 52)
	deckShuffled := NewDeck(true)
	assert.Equal(t, len(deckShuffled.Cards), 52)
}

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

func TestCopyDeck(t *testing.T) {
	deck := NewDeck(false)
	newDeck := deck.Copy()
	assert.Equal(t, deck.Cards, newDeck.Cards)
	assert.True(t, reflect.DeepEqual(deck.Cards, newDeck.Cards))
	assert.Equal(t, deck, newDeck)
	assert.True(t, reflect.DeepEqual(deck, newDeck))
	newDeck.Shuffle()
	assert.NotEqual(t, deck.Cards, newDeck.Cards)
	assert.NotEqual(t, deck, newDeck)
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
