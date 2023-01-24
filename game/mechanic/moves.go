package mechanic

import (
	"errors"
	"time"
)

var (
	START_ROUND       Mover = &StartRound{bid: true}
	DOUBLING          Mover = &Doubling{bid: true}
	PASS              Mover = &Pass{bid: true}
	LOOK_AT_ALL_CARDS Mover = &LookAtAllCards{bid: true}

	CHOOSE_HEARTS_LESS_ASKED     Mover = &ChooseGameType{bid: true, GameType: &HEARTS_LESS_ASKED}
	CHOOSE_DIAMONDS_LESS_ASKED   Mover = &ChooseGameType{bid: true, GameType: &DIAMONDS_LESS_ASKED}
	CHOOSE_CLUBS_LESS_ASKED      Mover = &ChooseGameType{bid: true, GameType: &CLUBS_LESS_ASKED}
	CHOOSE_SPADES_LESS_ASKED     Mover = &ChooseGameType{bid: true, GameType: &SPADES_LESS_ASKED}
	CHOOSE_HEARTS_LESS_UNASKED   Mover = &ChooseGameType{bid: true, GameType: &HEARTS_LESS_UNASKED}
	CHOOSE_DIAMONDS_LESS_UNASKED Mover = &ChooseGameType{bid: true, GameType: &DIAMONDS_LESS_UNASKED}
	CHOOSE_CLUBS_LESS_UNASKED    Mover = &ChooseGameType{bid: true, GameType: &CLUBS_LESS_UNASKED}
	CHOOSE_SPADES_LESS_UNASKED   Mover = &ChooseGameType{bid: true, GameType: &SPADES_LESS_UNASKED}

	CHOOSE_HEARTS_ASKED     Mover = &ChooseGameType{bid: true, GameType: &HEARTS_ASKED}
	CHOOSE_DIAMONDS_ASKED   Mover = &ChooseGameType{bid: true, GameType: &DIAMONDS_ASKED}
	CHOOSE_CLUBS_ASKED      Mover = &ChooseGameType{bid: true, GameType: &CLUBS_ASKED}
	CHOOSE_SPADES_ASKED     Mover = &ChooseGameType{bid: true, GameType: &SPADES_ASKED}
	CHOOSE_HEARTS_UNASKED   Mover = &ChooseGameType{bid: true, GameType: &HEARTS_UNASKED}
	CHOOSE_DIAMONDS_UNASKED Mover = &ChooseGameType{bid: true, GameType: &DIAMONDS_UNASKED}
	CHOOSE_CLUBS_UNASKED    Mover = &ChooseGameType{bid: true, GameType: &CLUBS_UNASKED}
	CHOOSE_SPADES_UNASKED   Mover = &ChooseGameType{bid: true, GameType: &SPADES_UNASKED}

	CHOOSE_WORSE_LESS  Mover = &ChooseGameType{bid: true, GameType: &WORSE_LESS}
	CHOOSE_BETTER_LESS Mover = &ChooseGameType{bid: true, GameType: &BETTER_LESS}
	CHOOSE_WORSE       Mover = &ChooseGameType{bid: true, GameType: &WORSE}
	CHOOSE_BETTER      Mover = &ChooseGameType{bid: true, GameType: &BETTER}

	PLAY_ACE_OF_HEARTS   Mover = &PlayCard{card: ACE_OF_HEARTS}
	PLAY_TEN_OF_HEARTS   Mover = &PlayCard{card: TEN_OF_HEARTS}
	PLAY_KING_OF_HEARTS  Mover = &PlayCard{card: KING_OF_HEARTS}
	PLAY_QUEEN_OF_HEARTS Mover = &PlayCard{card: QUEEN_OF_HEARTS}
	PLAY_JACK_OF_HEARTS  Mover = &PlayCard{card: JACK_OF_HEARTS}
	PLAY_NINE_OF_HEARTS  Mover = &PlayCard{card: NINE_OF_HEARTS}

	PLAY_ACE_OF_DIAMONDS   Mover = &PlayCard{card: ACE_OF_DIAMONDS}
	PLAY_TEN_OF_DIAMONDS   Mover = &PlayCard{card: TEN_OF_DIAMONDS}
	PLAY_KING_OF_DIAMONDS  Mover = &PlayCard{card: KING_OF_DIAMONDS}
	PLAY_QUEEN_OF_DIAMONDS Mover = &PlayCard{card: QUEEN_OF_DIAMONDS}
	PLAY_JACK_OF_DIAMONDS  Mover = &PlayCard{card: JACK_OF_DIAMONDS}
	PLAY_NINE_OF_DIAMONDS  Mover = &PlayCard{card: NINE_OF_DIAMONDS}

	PLAY_ACE_OF_CLUBS   Mover = &PlayCard{card: ACE_OF_CLUBS}
	PLAY_TEN_OF_CLUBS   Mover = &PlayCard{card: TEN_OF_CLUBS}
	PLAY_KING_OF_CLUBS  Mover = &PlayCard{card: KING_OF_CLUBS}
	PLAY_QUEEN_OF_CLUBS Mover = &PlayCard{card: QUEEN_OF_CLUBS}
	PLAY_JACK_OF_CLUBS  Mover = &PlayCard{card: JACK_OF_CLUBS}
	PLAY_NINE_OF_CLUBS  Mover = &PlayCard{card: NINE_OF_CLUBS}

	PLAY_ACE_OF_SPADES   Mover = &PlayCard{card: ACE_OF_SPADES}
	PLAY_TEN_OF_SPADES   Mover = &PlayCard{card: TEN_OF_SPADES}
	PLAY_KING_OF_SPADES  Mover = &PlayCard{card: KING_OF_SPADES}
	PLAY_QUEEN_OF_SPADES Mover = &PlayCard{card: QUEEN_OF_SPADES}
	PLAY_JACK_OF_SPADES  Mover = &PlayCard{card: JACK_OF_SPADES}
	PLAY_NINE_OF_SPADES  Mover = &PlayCard{card: NINE_OF_SPADES}

	ALL_MOVES_NAMES map[Mover]string = map[Mover]string{
		START_ROUND:       "startRound",
		DOUBLING:          "doubling",
		PASS:              "pass",
		LOOK_AT_ALL_CARDS: "lookAtAllCards",

		CHOOSE_HEARTS_LESS_ASKED:     "chooseHeartsLessAsked",
		CHOOSE_DIAMONDS_LESS_ASKED:   "chooseDiamondsLessAsked",
		CHOOSE_CLUBS_LESS_ASKED:      "chooseClubsLessAsked",
		CHOOSE_SPADES_LESS_ASKED:     "chooseSpadesLessAsked",
		CHOOSE_HEARTS_LESS_UNASKED:   "chooseHeartsLessUnasked",
		CHOOSE_DIAMONDS_LESS_UNASKED: "chooseDiamondsLessUnasked",
		CHOOSE_CLUBS_LESS_UNASKED:    "chooseClubsLessUnasked",
		CHOOSE_SPADES_LESS_UNASKED:   "chooseSpadesLessUnasked",

		CHOOSE_HEARTS_ASKED:     "chooseHeartsAsked",
		CHOOSE_DIAMONDS_ASKED:   "chooseDiamondsAsked",
		CHOOSE_CLUBS_ASKED:      "chooseClubsAsked",
		CHOOSE_SPADES_ASKED:     "chooseSpadesAsked",
		CHOOSE_HEARTS_UNASKED:   "chooseHeartsUnasked",
		CHOOSE_DIAMONDS_UNASKED: "chooseDiamondsUnasked",
		CHOOSE_CLUBS_UNASKED:    "chooseClubsUnasked",
		CHOOSE_SPADES_UNASKED:   "chooseSpadesUnasked",

		CHOOSE_WORSE_LESS:  "chooseWorseLess",
		CHOOSE_BETTER_LESS: "chooseBetterLess",
		CHOOSE_WORSE:       "chooseWorse",
		CHOOSE_BETTER:      "chooseBetter",

		PLAY_ACE_OF_HEARTS:   "playAceOfHearts",
		PLAY_TEN_OF_HEARTS:   "playTenOfHearts",
		PLAY_KING_OF_HEARTS:  "playKingOfHearts",
		PLAY_QUEEN_OF_HEARTS: "playQueenOfHearts",
		PLAY_JACK_OF_HEARTS:  "playJackOfHearts",
		PLAY_NINE_OF_HEARTS:  "playNineOfHearts",

		PLAY_ACE_OF_DIAMONDS:   "playAceOfDiamonds",
		PLAY_TEN_OF_DIAMONDS:   "playTenOfDiamonds",
		PLAY_KING_OF_DIAMONDS:  "playKingOfDiamonds",
		PLAY_QUEEN_OF_DIAMONDS: "playQueenOfDiamonds",
		PLAY_JACK_OF_DIAMONDS:  "playJackOfDiamonds",
		PLAY_NINE_OF_DIAMONDS:  "playNineOfDiamonds",

		PLAY_ACE_OF_CLUBS:   "playAceOfClubs",
		PLAY_TEN_OF_CLUBS:   "playTenOfClubs",
		PLAY_KING_OF_CLUBS:  "playKingOfClubs",
		PLAY_QUEEN_OF_CLUBS: "playQueenOfClubs",
		PLAY_JACK_OF_CLUBS:  "playJackOfClubs",
		PLAY_NINE_OF_CLUBS:  "playNineOfClubs",

		PLAY_ACE_OF_SPADES:   "playAceOfSpades",
		PLAY_TEN_OF_SPADES:   "playTenOfSpades",
		PLAY_KING_OF_SPADES:  "playKingOfSpades",
		PLAY_QUEEN_OF_SPADES: "playQueenOfSpades",
		PLAY_JACK_OF_SPADES:  "playJackOfSpades",
		PLAY_NINE_OF_SPADES:  "playNineOfSpades",
	}
	ALL_NAMES_MOVES map[string]Mover = flipMap(ALL_MOVES_NAMES)
)

type Mover interface {
	IsPossible(gameMechanic *Game) bool
	Move(gameMechanic *Game) error
	IsBid() bool //TODO: sprawdzić na koniec, czy przenieść definicję ruchu licytacyjnego do slice ALL_BIDS i usunąć metodę oraz pole "bid"
}

type ChooseGameType struct {
	bid      bool
	GameType *GameType
}

func (p *ChooseGameType) IsPossible(gm *Game) bool {
	hasEverybodyBid, error := gm.hasEverybodyBid()
	if error != nil || hasEverybodyBid || gm.doubleCount != 0 {
		return false
	}
	_, ok := gm.gameTypeShiftsFSM.getPossibleShifts().nodes[p.GameType]
	_, canLookAtAllCardsGameType := gameTypesCanLookAtAllCards[gm.gameTypeShiftsFSM.getCurrentGameType()]
	if canLookAtAllCardsGameType && !gm.allCardsSeen {
		return ok && (p.GameType == &WORSE_LESS || p.GameType == &BETTER_LESS)
	}
	if canLookAtAllCardsGameType && gm.allCardsSeen {
		return ok && p.GameType != &WORSE_LESS && p.GameType != &BETTER_LESS
	}

	return ok
}

func (p *ChooseGameType) Move(gameMechanic *Game) error {
	if !p.IsPossible(gameMechanic) {
		return errors.New("now this move ain't possible, bro")
	}
	gameMechanic.gameTypeShiftsFSM.changeTo(p.GameType)
	gameMechanic.passCount = 0
	gameMechanic.currentAuctionLeader = gameMechanic.smallTurn
	_, gameStopsAuction := gamesThatStopAuction[p.GameType]
	if gameStopsAuction {
		gameMechanic.gameStageFSM.change(AUCTION, ROUND)

		return nil
	}
	gameMechanic.nextPlayerSmallTurn()

	return nil
}

func (p *ChooseGameType) IsBid() bool {
	return p.bid
}

type Doubling struct {
	bid bool
}

func (p *Doubling) IsPossible(gm *Game) bool {
	if gm.gameTypeShiftsFSM.getCurrentGameType() == &WARSAW {
		return false
	}
	if gm.currentAuctionLeader == gm.smallTurn {
		return gm.doubleCount != 0 && gm.currentAuctionLeader != gm.lastDoublingPlayer
	}

	return true
}

func (p *Doubling) Move(gameMechanic *Game) error {
	if !p.IsPossible(gameMechanic) {
		return errors.New("now this move ain't possible, bro")
	}
	if gameMechanic.playerCount() == 4 {
		gameMechanic.nextPlayerSmallTurn()
	} else if gameMechanic.currentAuctionLeader != gameMechanic.smallTurn {
		gameMechanic.lastDoublingPlayer = gameMechanic.smallTurn
		gameMechanic.giveSmallTurnTo(gameMechanic.currentAuctionLeader)
	} else {
		gameMechanic.giveSmallTurnTo(gameMechanic.lastDoublingPlayer)
		gameMechanic.lastDoublingPlayer = gameMechanic.currentAuctionLeader
	}
	gameMechanic.passCount = 0
	gameMechanic.doubleCount++

	return nil
}

func (p *Doubling) IsBid() bool {
	return p.bid
}

type LookAtAllCards struct {
	bid bool
}

func (p *LookAtAllCards) IsPossible(gm *Game) bool {
	_, canLookAtAllCardsGameType := gameTypesCanLookAtAllCards[gm.gameTypeShiftsFSM.getCurrentGameType()]
	hasEverybodyBid, error := gm.hasEverybodyBid()
	if error != nil {
		return false
	}

	return hasEverybodyBid && canLookAtAllCardsGameType && gm.doubleCount == 0
}

func (p *LookAtAllCards) Move(gameMechanic *Game) error {
	if !p.IsPossible(gameMechanic) {
		return errors.New("now this move ain't possible, bro")
	}
	gameMechanic.passCount = 0
	gameMechanic.allCardsSeen = true

	return nil
}

func (p *LookAtAllCards) IsBid() bool {
	return p.bid
}

type Pass struct {
	bid bool
}

func (p *Pass) IsPossible(gm *Game) bool {
	hasEverybodyBid, error := gm.hasEverybodyBid()
	if error != nil {
		return false
	}
	if gm.playerCount() == 3 && gm.currentAuctionLeader == gm.smallTurn && gm.doubleCount != 0 {
		return false
	}

	return !hasEverybodyBid
}

func (p *Pass) Move(gameMechanic *Game) error {
	if !p.IsPossible(gameMechanic) {
		return errors.New("now this move ain't possible, bro")
	}
	gameMechanic.passCount++
	if gameMechanic.playerCount() == 3 && gameMechanic.doubleCount != 0 {
		gameMechanic.giveSmallTurnTo(gameMechanic.currentAuctionLeader)
		return nil
	}
	gameMechanic.nextPlayerSmallTurn()

	return nil
}

func (p *Pass) IsBid() bool {
	return p.bid
}

type StartRound struct {
	bid bool
}

func (p *StartRound) IsPossible(gm *Game) bool {
	hasEverybodyBid, error := gm.hasEverybodyBid()
	if error != nil {
		return false
	}
	if gm.playerCount() == 3 && gm.currentAuctionLeader == gm.smallTurn && gm.doubleCount != 0 {
		return true
	}

	return hasEverybodyBid
}

func (p *StartRound) Move(gameMechanic *Game) error {
	if !p.IsPossible(gameMechanic) {
		return errors.New("now this move ain't possible, bro")
	}
	err := gameMechanic.gameStageFSM.change(AUCTION, ROUND)
	if err != nil {
		return err
	}
	gameMechanic.passCount = 0

	return nil
}

func (p *StartRound) IsBid() bool {
	return p.bid
}

type PlayCard struct {
	bid  bool
	card *card
}

func (p *PlayCard) IsPossible(gm *Game) bool {
	if len(gm.trick) == 0 {
		return true
	}
	if len(gm.trick) == int(gm.playerCount()) {
		return false
	}

	hasOnHand := false
	for _, v := range gm.playersCards[gm.smallTurn] {
		if v == p.card {
			hasOnHand = true
			break
		}
	}
	if !hasOnHand {
		return false
	}

	//jeśli masz w kolorze inicjalnym, to musisz w nim położyć
	if p.card.suit != gm.trick[0].suit {
		for _, v := range gm.playersCards[gm.smallTurn] {
			if v.suit == gm.trick[0].suit {
				return false
			}
		}
		//*albo nie mam karty w kolorze inicjalnym, albo właśnie chcę ją zagrać*
	} else {
		//*karta jest w kolorze inicjalnym*
		//jeśli jest w kolorze inicjalnym, ale jest niższa od najwyższej na stole,
		//a mam wyższą nieatutową(tutaj jest w kolorze inicjalnym), to nie mogę tej niskiej grać
		if !gm.isCardHigherThanTrick(p.card) {
			for _, v := range gm.playersCards[gm.smallTurn] {
				if v.suit != gm.trick[0].suit {
					continue
				}
				if gm.isCardHigherThanTrick(v) {
					return false
				}
			}
		}

		return true
	}

	//*na tym etapie nie masz koloru inicjalnego*
	//jeśli masz atu wyższe od kart na stole, to musisz atu
	trump := gm.gameTypeShiftsFSM.getCurrentGameType().trump
	if trump != NONE_SUIT && p.card.suit != trump {
		for _, v := range gm.playersCards[gm.smallTurn] {
			if v.suit == trump {
				return false
			}
		}
	}
	//atu musi być wyższe od tego na stole, jeśli takie masz
	if trump != NONE_SUIT && !gm.isCardHigherThanTrick(p.card) {
		for _, v := range gm.playersCards[gm.smallTurn] {
			if gm.isCardHigherThanTrick(v) {
				return false
			}
		}
	}

	return true
}

func (p *PlayCard) Move(gm *Game) error {
	if !p.IsPossible(gm) {
		return errors.New("now this move ain't possible, bro")
	}
	//TODO: meldunki
	newHand := []*card{}
	for _, v := range gm.playersCards[gm.smallTurn] {
		if v == p.card {
			continue
		}
		newHand = append(newHand, v)
	}
	gm.playersCards[gm.smallTurn] = newHand
	if gm.isCardHigherThanTrick(p.card) {
		gm.trickLeader = gm.smallTurn
	}
	gm.trick = append(gm.trick, p.card)
	if len(gm.trick) == int(gm.playerCount()) {
		go func() {
			time.Sleep(2 * time.Second)
			gm.cleanTrick()
			gm.giveSmallTurnTo(gm.trickLeader)
			gm.trickLeader = -1
			gm.Notify()
		}()

		return nil
	}
	gm.nextPlayerSmallTurn()

	return nil
}

func (p *PlayCard) IsBid() bool {
	return p.bid
}

func flipMap(input map[Mover]string) map[string]Mover {
	output := map[string]Mover{}
	for k, v := range input {
		output[v] = k
	}
	return output
}
