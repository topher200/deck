package deck

import "fmt"

// Card represents a Card with a Face and a Suit
type Card struct {
	Face Face
	Suit Suit
}

func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.Face, c.Suit)
}

// GetSignature is the hex representation of the Face and Suit of the card
func (c *Card) GetSignature() string {
	return fmt.Sprintf("%x%x", c.Face, c.Suit)
}

// Compare compares 2 cards 1 if the passed in card is greater -1 if its lesser 0 of equal.
func (c *Card) Compare(k Card) int {
	if k.Face > c.Face {
		return 1
	}

	if k.Face < c.Face {
		return -1
	}

	if k.Suit > c.Suit {
		return 1
	}

	if k.Suit < c.Suit {
		return -1
	}

	return 0
}

//IsLessThan returns bool if card passed in is less then
func (c *Card) IsLessThan(k Card) bool {
	return c.Compare(k) == 1
}

//IsGreaterThan return bool if card passed in is greater then
func (c *Card) IsGreaterThan(k Card) bool {
	return c.Compare(k) == -1
}

//IsEqualTo returns true if the card is equal in face and
func (c *Card) IsEqualTo(k Card) bool {
	return c.Compare(k) == 0
}
