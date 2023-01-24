package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"github.com/gammazero/nexus/v3/client"
	"github.com/gammazero/nexus/v3/wamp"
)

func main() {
	a := app.New()
	w := a.NewWindow("")
	closeConnection := make(chan int)
	connectionErr := make(chan error)
	gClient := new(client.Client)
	gameState := newGameState()
	screens := newScreens(w, closeConnection, connectionErr, gClient, gameState)
	go func() {
		for i := range connectionErr {
			w.SetContent(screens.connectScreen)
			fmt.Println("connection error: ", i)
		}
	}()
	w.SetContent(screens.mainMenuScreen)
	w.Show()
	a.Run()
}

type gameStage int

const (
	AUCTION gameStage = iota
	ROUND
	FINISH
)

type gameState struct {
	gameStage                         gameStage
	clientPlayerName                  string
	playerCards                       []*card
	playersCardCount                  map[string]int
	playersCyclicOrder                []string
	playerListBinding, bidListBinding binding.StringList
	bidListBindingData                []string
	cardPlayList                      map[string]bool
	turn                              binding.String
	trick                             []*card
	observers                         map[Observer]bool
}

func newGameState() *gameState {
	state := &gameState{}
	state.playerListBinding = binding.NewStringList()
	state.bidListBinding = binding.NewStringList()
	state.bidListBindingData = []string{}
	state.playerCards = []*card{}
	state.turn = binding.NewString()
	state.trick = []*card{}
	state.observers = map[Observer]bool{}

	return state
}

func (p *gameState) updateGameState(dict wamp.Dict) {
	p.gameStage = gameStage(dict["gameStage"].(uint64))
	p.turn.Set(dict["turn"].(string))

	p.bidListBinding.Set([]string{})
	p.bidListBindingData = []string{}
	bidList := dict["possibleBids"].([]interface{})
	for _, v := range bidList {
		if !p.isPlayerTurn() { // TODO: może przenieść tę logikę do apki gry?
			break
		}
		bid := v.(string)
		p.bidListBinding.Append(ALL_MOVES_NAMES[bid])
		p.bidListBindingData = append(p.bidListBindingData, bid)
	}

	p.cardPlayList = map[string]bool{}
	cardPlayList := dict["possibleCardPlays"].([]interface{})
	for _, v := range cardPlayList {
		if !p.isPlayerTurn() { // TODO: może przenieść tę logikę do apki gry?
			break
		}
		cardPlay := v.(string)
		p.cardPlayList[cardPlay] = true
	}

	p.playerCards = []*card{}
	cards := dict["playerCards"].([]interface{})
	fmt.Println(cards)
	for _, v := range cards {
		p.playerCards = append(p.playerCards, ALL_CARDS[v.(string)])
	}

	p.trick = []*card{}
	trick := dict["trick"].([]interface{})
	fmt.Println(trick)
	for _, v := range trick {
		p.trick = append(p.trick, ALL_CARDS[v.(string)])
	}

	p.playersCardCount = map[string]int{}
	cardCounts := dict["playerCardCounts"].(map[string]interface{})
	for k, v := range cardCounts {
		p.playersCardCount[k] = int(v.(uint64))
	}

	p.playersCyclicOrder = []string{}
	cyclicOrder := dict["playersCyclicOrder"].([]interface{})
	for _, v := range cyclicOrder {
		p.playersCyclicOrder = append(p.playersCyclicOrder, v.(string))
	}

	p.notify()
}

func (p *gameState) hiddenCardsFor(playerName string) []*card {
	if p.playerListBinding.Length() == 0 {
		return []*card{}
	}
	count := p.playersCardCount[playerName]
	if p.clientPlayerName == playerName {
		count -= len(p.playerCards)
	}
	cards := make([]*card, count)
	for i := range cards {
		cards[i] = NO_CARD
	}

	return cards
}

func (p *gameState) AddObserver(o Observer) {
	p.observers[o] = true
}

func (p *gameState) RemoveObserver(o Observer) {
	delete(p.observers, o)
}

func (p *gameState) notify() {
	for o := range p.observers {
		o.Update()
	}
}

func (p *gameState) isPlayerTurn() bool {
	turn, err := p.turn.Get()
	if err != nil {
		panic(err) //TODO log error
	}
	return p.clientPlayerName == turn
}

type Observer interface {
	Update()
}

type Observable interface {
	AddObserver(o Observer)
	RemoveObserver(o Observer)
	notify()
}
