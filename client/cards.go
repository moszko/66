package main

type card struct {
	suit suit
	rank rank
}

func (p *card) String() string {
	return p.rank.String() + p.suit.String()
}

var (
	NO_CARD         = &card{NONE_SUIT, NONERANK}
	ACE_OF_SPADES   = &card{SPADES_SUIT, ACE}
	TEN_OF_SPADES   = &card{SPADES_SUIT, TEN}
	KING_OF_SPADES  = &card{SPADES_SUIT, KING}
	QUEEN_OF_SPADES = &card{SPADES_SUIT, QUEEN}
	JACK_OF_SPADES  = &card{SPADES_SUIT, JACK}
	NINE_OF_SPADES  = &card{SPADES_SUIT, NINE}

	ACE_OF_CLUBS   = &card{CLUBS_SUIT, ACE}
	TEN_OF_CLUBS   = &card{CLUBS_SUIT, TEN}
	KING_OF_CLUBS  = &card{CLUBS_SUIT, KING}
	QUEEN_OF_CLUBS = &card{CLUBS_SUIT, QUEEN}
	JACK_OF_CLUBS  = &card{CLUBS_SUIT, JACK}
	NINE_OF_CLUBS  = &card{CLUBS_SUIT, NINE}

	ACE_OF_DIAMONDS   = &card{DIAMONDS_SUIT, ACE}
	TEN_OF_DIAMONDS   = &card{DIAMONDS_SUIT, TEN}
	KING_OF_DIAMONDS  = &card{DIAMONDS_SUIT, KING}
	QUEEN_OF_DIAMONDS = &card{DIAMONDS_SUIT, QUEEN}
	JACK_OF_DIAMONDS  = &card{DIAMONDS_SUIT, JACK}
	NINE_OF_DIAMONDS  = &card{DIAMONDS_SUIT, NINE}

	ACE_OF_HEARTS   = &card{HEARTS_SUIT, ACE}
	TEN_OF_HEARTS   = &card{HEARTS_SUIT, TEN}
	KING_OF_HEARTS  = &card{HEARTS_SUIT, KING}
	QUEEN_OF_HEARTS = &card{HEARTS_SUIT, QUEEN}
	JACK_OF_HEARTS  = &card{HEARTS_SUIT, JACK}
	NINE_OF_HEARTS  = &card{HEARTS_SUIT, NINE}

	ALL_CARDS map[string]*card = map[string]*card{
		"A♠":  ACE_OF_SPADES,
		"10♠": TEN_OF_SPADES,
		"K♠":  KING_OF_SPADES,
		"Q♠":  QUEEN_OF_SPADES,
		"J♠":  JACK_OF_SPADES,
		"9♠":  NINE_OF_SPADES,

		"A♣":  ACE_OF_CLUBS,
		"10♣": TEN_OF_CLUBS,
		"K♣":  KING_OF_CLUBS,
		"Q♣":  QUEEN_OF_CLUBS,
		"J♣":  JACK_OF_CLUBS,
		"9♣":  NINE_OF_CLUBS,

		"A♦":  ACE_OF_DIAMONDS,
		"10♦": TEN_OF_DIAMONDS,
		"K♦":  KING_OF_DIAMONDS,
		"Q♦":  QUEEN_OF_DIAMONDS,
		"J♦":  JACK_OF_DIAMONDS,
		"9♦":  NINE_OF_DIAMONDS,

		"A♥":  ACE_OF_HEARTS,
		"10♥": TEN_OF_HEARTS,
		"K♥":  KING_OF_HEARTS,
		"Q♥":  QUEEN_OF_HEARTS,
		"J♥":  JACK_OF_HEARTS,
		"9♥":  NINE_OF_HEARTS,
	}
	CARDS_MOVES map[*card]string = map[*card]string{
		ACE_OF_HEARTS: "playAceOfHearts",
		TEN_OF_HEARTS: "playTenOfHearts",
		KING_OF_HEARTS: "playKingOfHearts",
		QUEEN_OF_HEARTS: "playQueenOfHearts",
		JACK_OF_HEARTS: "playJackOfHearts",
		NINE_OF_HEARTS: "playNineOfHearts",

		ACE_OF_DIAMONDS: "playAceOfDiamonds",
		TEN_OF_DIAMONDS: "playTenOfDiamonds",
		KING_OF_DIAMONDS: "playKingOfDiamonds",
		QUEEN_OF_DIAMONDS: "playQueenOfDiamonds",
		JACK_OF_DIAMONDS: "playJackOfDiamonds",
		NINE_OF_DIAMONDS: "playNineOfDiamonds",

		ACE_OF_CLUBS: "playAceOfClubs",
		TEN_OF_CLUBS: "playTenOfClubs",
		KING_OF_CLUBS: "playKingOfClubs",
		QUEEN_OF_CLUBS: "playQueenOfClubs",
		JACK_OF_CLUBS: "playJackOfClubs",
		NINE_OF_CLUBS: "playNineOfClubs",

		ACE_OF_SPADES: "playAceOfSpades",
		TEN_OF_SPADES: "playTenOfSpades",
		KING_OF_SPADES: "playKingOfSpades",
		QUEEN_OF_SPADES: "playQueenOfSpades",
		JACK_OF_SPADES: "playJackOfSpades",
		NINE_OF_SPADES: "playNineOfSpades",
	}
)

type suit int

const (
	NONE_SUIT suit = iota
	HEARTS_SUIT
	DIAMONDS_SUIT
	CLUBS_SUIT
	SPADES_SUIT
)

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

type rank int

const (
	NONERANK rank = -1
	NINE     rank = 0
	JACK     rank = 2
	QUEEN    rank = 3
	KING     rank = 4
	TEN      rank = 10
	ACE      rank = 11
)

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
	case ACE:
		return "A"
	default:
		return "NN"
	}
}
