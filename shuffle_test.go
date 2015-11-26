package deck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShuffle(t *testing.T) {
	original := NewDeck(false)
	shuffled := NewDeck(false)
	shuffled.Shuffle()
	decksAreDifferent := false
	for i, _ := range original.Cards {
		if original.Cards[i] != shuffled.Cards[i] {
			decksAreDifferent = true
		}
	}
	assert.True(t, decksAreDifferent)
}

func TestMultiShuffle(t *testing.T) {
	original := NewDeck(false)
	shuffled := NewDeck(false)
	shuffled.MultiShuffle(10)
	decksAreDifferent := false
	for i, _ := range original.Cards {
		if original.Cards[i] != shuffled.Cards[i] {
			decksAreDifferent = true
		}
	}
	assert.True(t, decksAreDifferent)
}
