package mechanic

type card struct {
	suit suit
	rank rank
}

func (p *card) String() string {
	return p.rank.String() + p.suit.String()
}

var NO_CARD = &card{NONE_SUIT, NONERANK}

var ACE_OF_SPADES = &card{SPADES_SUIT, ACE}
var TEN_OF_SPADES = &card{SPADES_SUIT, TEN}
var KING_OF_SPADES = &card{SPADES_SUIT, KING}
var QUEEN_OF_SPADES = &card{SPADES_SUIT, QUEEN}
var JACK_OF_SPADES = &card{SPADES_SUIT, JACK}
var NINE_OF_SPADES = &card{SPADES_SUIT, NINE}

var ACE_OF_CLUBS = &card{CLUBS_SUIT, ACE}
var TEN_OF_CLUBS = &card{CLUBS_SUIT, TEN}
var KING_OF_CLUBS = &card{CLUBS_SUIT, KING}
var QUEEN_OF_CLUBS = &card{CLUBS_SUIT, QUEEN}
var JACK_OF_CLUBS = &card{CLUBS_SUIT, JACK}
var NINE_OF_CLUBS = &card{CLUBS_SUIT, NINE}

var ACE_OF_DIAMONDS = &card{DIAMONDS_SUIT, ACE}
var TEN_OF_DIAMONDS = &card{DIAMONDS_SUIT, TEN}
var KING_OF_DIAMONDS = &card{DIAMONDS_SUIT, KING}
var QUEEN_OF_DIAMONDS = &card{DIAMONDS_SUIT, QUEEN}
var JACK_OF_DIAMONDS = &card{DIAMONDS_SUIT, JACK}
var NINE_OF_DIAMONDS = &card{DIAMONDS_SUIT, NINE}

var ACE_OF_HEARTS = &card{HEARTS_SUIT, ACE}
var TEN_OF_HEARTS = &card{HEARTS_SUIT, TEN}
var KING_OF_HEARTS = &card{HEARTS_SUIT, KING}
var QUEEN_OF_HEARTS = &card{HEARTS_SUIT, QUEEN}
var JACK_OF_HEARTS = &card{HEARTS_SUIT, JACK}
var NINE_OF_HEARTS = &card{HEARTS_SUIT, NINE}

var ALL_CARDS = []*card{
	ACE_OF_SPADES,
	TEN_OF_SPADES,
	KING_OF_SPADES,
	QUEEN_OF_SPADES,
	JACK_OF_SPADES,
	NINE_OF_SPADES,

	ACE_OF_CLUBS,
	TEN_OF_CLUBS,
	KING_OF_CLUBS,
	QUEEN_OF_CLUBS,
	JACK_OF_CLUBS,
	NINE_OF_CLUBS,

	ACE_OF_DIAMONDS,
	TEN_OF_DIAMONDS,
	KING_OF_DIAMONDS,
	QUEEN_OF_DIAMONDS,
	JACK_OF_DIAMONDS,
	NINE_OF_DIAMONDS,

	ACE_OF_HEARTS,
	TEN_OF_HEARTS,
	KING_OF_HEARTS,
	QUEEN_OF_HEARTS,
	JACK_OF_HEARTS,
	NINE_OF_HEARTS,
}

type suit int

func (v suit) String() string {
	switch v {
	case HEARTS_SUIT:
		return "♥"
	case DIAMONDS_SUIT:
		return "♦"
	case CLUBS_SUIT:
		return "♣"
	case SPADES_SUIT:
		return "♠"
	}
	return ""
}

const (
	NONE_SUIT suit = iota
	HEARTS_SUIT
	DIAMONDS_SUIT
	CLUBS_SUIT
	SPADES_SUIT
)

type rank int

func (v rank) String() string {
	switch v {
	case NINE:
		return "9"
	case JACK:
		return "J"
	case QUEEN:
		return "Q"
	case KING:
		return "K"
	case TEN:
		return "10"
	default:
		return "A"
	}
}

const (
	NONERANK rank = -1
	NINE     rank = 0
	JACK     rank = 2
	QUEEN    rank = 3
	KING     rank = 4
	TEN      rank = 10
	ACE      rank = 11
)
