package deck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecrement(t *testing.T) {
	face, err := Decrement(TWO)
	assert.EqualValues(t, ACE, face)
	assert.Nil(t, err)

	face, err = Decrement(JACK)
	assert.EqualValues(t, TEN, face)
	assert.Nil(t, err)

	face, err = Decrement(KING)
	assert.EqualValues(t, QUEEN, face)
	assert.Nil(t, err)

	_, err = Decrement(ACE)
	assert.Error(t, err)
}

func TestIncrement(t *testing.T) {
	face, err := Increment(ACE)
	assert.EqualValues(t, TWO, face)
	assert.Nil(t, err)

	face, err = Increment(TEN)
	assert.EqualValues(t, JACK, face)
	assert.Nil(t, err)

	face, err = Increment(QUEEN)
	assert.EqualValues(t, KING, face)
	assert.Nil(t, err)

	_, err = Increment(KING)
	assert.Error(t, err)
}
