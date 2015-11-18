package deck

// Suit represents the suit of the card (spade, heart, diamon, club)
type Suit string

// Face represents the face of the card (ace, two...queen, king)
type Face string

// Contants for Suit ♠♥♦♣
const (
	CLUB    Suit = "♣"
	DIAMOND      = "♦"
	HEART        = "♥"
	SPADE        = "♠"
)

// Contants for Face
const (
	ACE   Face = "A"
	TWO        = "2"
	THREE      = "3"
	FOUR       = "4"
	FIVE       = "5"
	SIX        = "6"
	SEVEN      = "7"
	EIGHT      = "8"
	NINE       = "9"
	TEN        = "T"
	JACK       = "J"
	QUEEN      = "Q"
	KING       = "K"
)

// Global Variables representing the default suits and faces in a deck of cards
var (
	SUITS = []Suit{CLUB, DIAMOND, HEART, SPADE}
	FACES = []Face{ACE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING}
)
