package mechanic

import (
	"66_game/helper"
	"errors"
	"math/rand"
	"time"
)

type Game struct {
	players              []Player
	playersCards         [][]*card
	gameStageFSM         gameStageFSM
	gameTypeShiftsFSM    gameTypeShiftsFSM
	smallTurn            int
	bigTurn              int
	passCount            uint8
	currentAuctionLeader int
	allCardsSeen         bool
	doubleCount          uint8
	lastDoublingPlayer   int //used only in 3-player game
	trick                []*card
	trickLeader          int

	observers map[helper.Observer]bool
}

func NewGame(firstPlayer, secondPlayer, thirdPlayer Player) *Game {
	gameMechanic := new(Game)
	gameMechanic.players = append(gameMechanic.players, firstPlayer)
	gameMechanic.players = append(gameMechanic.players, secondPlayer)
	gameMechanic.players = append(gameMechanic.players, thirdPlayer)
	gameMechanic.doubleCount = 0
	gameMechanic.randomizeFirstPlayer()
	gameMechanic.gameStageFSM = *newGameStageFSM()
	gameMechanic.observers = make(map[helper.Observer]bool)

	error := gameMechanic.startNewAuction()
	if error != nil {
		panic(error)
	}
	gameMechanic.shuffleAndDeal()
	gameMechanic.cleanTrick()

	return gameMechanic
}

func (gameMechanic *Game) MakeMove(move Mover, player Player) error {
	if gameMechanic.players[gameMechanic.smallTurn].Id() != player.Id() {
		return errors.New("not your turn, buster")
	}
	if !move.IsPossible(gameMechanic) {
		return errors.New("now that move ain't right, dawg")
	}
	err := move.Move(gameMechanic)
	if err != nil {
		return err
	}
	gameMechanic.Notify()

	return nil
}

func (p *Game) SmallTurn() Player {
	return p.players[p.smallTurn]
}

func (p *Game) StateFor(player Player) map[string]interface{} {
	playerGameState := PlayerGameState{
		gameStage:            p.gameStageFSM.currentGameStage,
		playerCount:          int(p.playerCount()),
		players:              p.players,
		currentAuctionLeader: p.players[p.currentAuctionLeader].Id(),
		turn:                 p.SmallTurn().Id(),
		possibleBids:         p.PossibleBids(),
		possibleCardPlays:    p.PossibleCardPlays(),
		playersCards:         p.playersCardsByNames(),
		allCardsSeen:         p.allCardsSeen,
		trick:                p.trick,
	}
	return playerGameState.parse(player.Id())
}

func (p *Game) PossibleBids() []Mover {
	result := []Mover{}
	if p.gameStageFSM.currentGameStage != AUCTION { //TODO: sprawdzenie tury dopisać
		return result
	}
	for move := range ALL_MOVES_NAMES {
		if move.IsBid() && move.IsPossible(p) {
			result = append(result, move)
		}
	}

	return result
}

func (p *Game) PossibleCardPlays() []Mover {
	result := []Mover{}
	if p.gameStageFSM.currentGameStage != ROUND { //TODO: sprawdzenie tury dopisać
		return result
	}
	for move := range ALL_MOVES_NAMES {
		if !move.IsBid() && move.IsPossible(p) {
			result = append(result, move)
		}
	}

	return result
}

func (p *Game) playerCardsByName(playerName string) ([]*card, error) {
	cards := []*card{}
	_, index, err := p.playerByName(playerName)
	if err != nil {
		return cards, err
	}

	return p.playersCards[index], nil
}

func (p *Game) playersCardsByNames() map[string][]*card {
	cards := map[string][]*card{}
	for i, v := range p.playersCards {
		cards[p.players[i].Id()] = v
	}

	return cards
}

func (p *Game) playerByName(playerName string) (Player, int, error) {
	for i, v := range p.players {
		if v.Id() == playerName {
			return v, i, nil
		}
	}

	return nil, -1, errors.New("no such player")
}

func (p *Game) cleanTrick() {
	p.trick = []*card{}
}

func (p *Game) randomizeFirstPlayer() {
	rand.Seed(time.Now().UnixNano())
	p.bigTurn = rand.Intn(int(p.playerCount()))
	p.smallTurn = p.bigTurn
}

func (p *Game) startNewAuction() error {
	if p.gameStageFSM.currentGameStage == FINISH {
		return errors.New("the game is over")
	}
	if p.gameStageFSM.currentGameStage != AUCTION {
		p.gameStageFSM.change(ROUND, AUCTION)
	}
	p.passCount = 0
	p.currentAuctionLeader = p.bigTurn
	p.allCardsSeen = false
	p.gameTypeShiftsFSM = *newGameTypeShiftsFSM()

	return nil
}

func (p *Game) shuffleAndDeal() {
	deck := make([]*card, len(ALL_CARDS))
	copy(deck, ALL_CARDS)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
	p.playersCards = make([][]*card, p.playerCount())
	for i, card := range deck {
		p.playersCards[i%int(p.playerCount())] = append(p.playersCards[i%int(p.playerCount())], card)
	}
}

func (p *Game) isCardHigherThanTrick(card *card) bool {
	if len(p.trick) == 0 {
		return true
	}
	trump := p.gameTypeShiftsFSM.getCurrentGameType().trump
	lead := p.trick[0]
	if card.suit != lead.suit && card.suit != trump {
		return false
	}
	for _, trickCard := range p.trick {
		if card.suit != trump && trickCard.suit == trump {
			return false
		}
		if card.suit != lead.suit && card.suit != trump {
			return false
		}
		if card.suit == trump && trickCard.suit == trump && card.rank < trickCard.rank {
			return false
		}
		if card.suit == lead.suit && trickCard.suit == lead.suit && card.rank < trickCard.rank {
			return false
		}
	}

	return true
}

func (p *Game) nextPlayerSmallTurn() {
	p.smallTurn++
	p.smallTurn = p.smallTurn % int(p.playerCount())
}

func (p *Game) nextPlayerBigTurn() {
	p.bigTurn++
	p.bigTurn = p.bigTurn % int(p.playerCount())
	p.smallTurn = p.bigTurn
}

func (p *Game) hasEverybodyBid() (bool, error) {
	if p.gameStageFSM.currentGameStage != AUCTION {
		return false, errors.New("wrong stage of the game, fool")
	}
	if p.currentAuctionLeader != p.smallTurn {
		return false, nil
	}
	if p.gameTypeShiftsFSM.getCurrentGameType() == &WARSAW {
		return p.passCount > 0, nil
	}
	if p.passCount == 0 {
		return false, nil
	}
	return true, nil
}

func (p *Game) playerCount() uint8 {
	return uint8(len(p.players))
}

// func (p *Game) hasEverybodyPassed() (bool, error) {
// 	if p.gameStageFSM.currentGameStage != AUCTION {
// 		return false, errors.New("wrong stage of the game, fool")
// 	}
// 	return p.playerCount == p.passCount, nil
// }

func (p *Game) giveSmallTurnTo(playerId int) {
	p.smallTurn = playerId
}

func (p *Game) AddObserver(o helper.Observer) {
	p.observers[o] = true
}

func (p *Game) RemoveObserver(o helper.Observer) {
	delete(p.observers, o)
}

func (p *Game) Notify() {
	for o := range p.observers {
		o.Update()
	}
}

type Player interface {
	Id() string
}

type PlayerGameState struct {
	gameStage   gameStage
	playerCount int

	//empty when a round is being played
	currentAuctionLeader string
	currentGameType      GameType

	//available in a round
	trick []*card

	//clockwise order
	players []Player
	//cards the player can see
	playersCards      map[string][]*card
	possibleCardPlays []Mover
	possibleBids      []Mover
	turn              string
	allCardsSeen      bool
}

func (p *PlayerGameState) parse(playerName string) map[string]interface{} {
	result := map[string]interface{}{}
	result["gameStage"] = p.gameStage
	result["currentAuctionLeader"] = p.currentAuctionLeader
	result["currentGameType"] = map[string]interface{}{"name": p.currentGameType.name, "maxPoints": p.currentGameType.maxPoints, "triuph": p.currentGameType.trump}
	result["possibleBids"] = p.parsePossibleBids()
	result["possibleCardPlays"] = p.parsePossibleCardPlays()
	result["turn"] = p.turn
	result["playerCards"] = p.parsePlayerCards(playerName)
	result["playerCardCounts"] = p.parsePlayerCardCounts()
	result["playersCyclicOrder"] = p.parsePlayersCyclicOrder()
	result["trick"] = p.parseTrick()

	return result
}

func (p *PlayerGameState) parsePossibleBids() []string {
	result := []string{}
	for _, v := range p.possibleBids {
		result = append(result, ALL_MOVES_NAMES[v])
	}

	return result
}

func (p *PlayerGameState) parsePossibleCardPlays() []string {
	result := []string{}
	for _, v := range p.possibleCardPlays {
		result = append(result, ALL_MOVES_NAMES[v])
	}

	return result
}

func (p *PlayerGameState) parsePlayerCards(playerName string) []string {
	cards := []string{}
	for _, v := range p.playersCards[playerName] {
		cards = append(cards, v.String())
	}
	if !p.allCardsSeen {
		count := 24 / (p.playerCount * 2)
		return cards[count:]
	}

	return cards
}

func (p *PlayerGameState) parseTrick() []string {
	cards := []string{}
	for _, v := range p.trick {
		cards = append(cards, v.String())
	}

	return cards
}

func (p *PlayerGameState) parsePlayerCardCounts() map[string]int {
	counts := map[string]int{}
	for i, v := range p.playersCards {
		counts[i] = len(v)
	}

	return counts
}

func (p *PlayerGameState) parsePlayersCyclicOrder() []string {
	order := []string{}
	for _, v := range p.players {
		order = append(order, v.Id())
	}

	return order
}
