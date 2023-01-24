package main

import (
	"context"
	"errors"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/gammazero/nexus/v3/client"
	"github.com/gammazero/nexus/v3/wamp"
)

type screens struct {
	connectScreen, mainMenuScreen, gameContainer *fyne.Container
}

func newScreens(w fyne.Window, closeConnection chan int, connectionErr chan error, gClient *client.Client, gameState *gameState) *screens {
	screens := &screens{}

	nameFormEntry, addressFormEntry := widget.NewEntry(), widget.NewEntry()
	nameFormEntry.Validator, addressFormEntry.Validator = zeroLengthStringValidator, zeroLengthStringValidator

	connect := newConnectElements(w, screens, nameFormEntry, addressFormEntry, gameState, closeConnection, connectionErr, gClient)
	mainMenu := newMainMenuElements(w, screens, nameFormEntry)
	game := newGameElements(w, gClient, closeConnection, connectionErr, screens, gameState)
	gameState.AddObserver(game)

	screens.connectScreen = container.NewGridWithRows(3, connect.label, connect.addressForm, connect.backToMainMenuButton)
	screens.mainMenuScreen = container.NewGridWithRows(2, mainMenu.newGameButton, mainMenu.exitButton)
	screens.gameContainer = container.NewGridWithRows(1, container.NewGridWithColumns(1, game.turn, game.bidList, game.playerList, game.quitGameButton), container.NewBorder(game.hands[2], game.hands[0], game.hands[1], game.hands[3], game.trick))

	return screens
}

type mainMenuElements struct {
	newGameButton, exitButton *widget.Button
}

func newMainMenuElements(w fyne.Window, le *screens, nameFormEntry *widget.Entry) *mainMenuElements {
	mainMenu := &mainMenuElements{}
	mainMenu.newGameButton = widget.NewButton("NOWA GRA", func() {
		w.SetContent(le.connectScreen)
		w.Canvas().Focus(nameFormEntry)
	})

	mainMenu.exitButton = widget.NewButton("WYJDŹ", func() {
		w.Close()
	})
	return mainMenu
}

type connectElements struct {
	addressForm          *widget.Form
	backToMainMenuButton *widget.Button
	label                *widget.Label
}

func newConnectElements(
	w fyne.Window,
	le *screens,
	nameFormEntry *widget.Entry,
	addressFormEntry *widget.Entry,
	gameState *gameState,
	closeConnection chan int,
	connectionErr chan error,
	gClient *client.Client,
) *connectElements {
	connect := &connectElements{}
	connect.label = widget.NewLabel("Wpisz dane")
	connect.backToMainMenuButton = widget.NewButton("MENU GŁÓWNE", func() {
		w.SetContent(le.mainMenuScreen)
	})
	connect.addressForm = &widget.Form{
		Items: []*widget.FormItem{
			{Text: "nick", Widget: nameFormEntry},
			{Text: "adres", Widget: addressFormEntry},
		},
		OnSubmit: func() {
			gameState.clientPlayerName = nameFormEntry.Text
			go func() {
				go play(addressFormEntry.Text, nameFormEntry.Text, closeConnection, connectionErr, gameState, gClient)
				w.SetContent(le.gameContainer)
			}()
		},
		SubmitText: "OK",
	}

	return connect
}

type gameElements struct {
	bidList, playerList *widget.List
	turn                *widget.Label
	hands               [4]*handWidget
	trick               *trickWidget
	quitGameButton      *widget.Button
	gameState           *gameState
}

func newGameElements(
	w fyne.Window,
	gClient *client.Client,
	closeConnection chan int,
	connectionErr chan error,
	screens *screens,
	gameState *gameState,
) *gameElements {
	game := &gameElements{}

	doNothing := func(card *card) {}
	playCard := func(card *card) {
		fmt.Println(gameState.cardPlayList)
		if gameState.gameStage != ROUND || !gameState.isPlayerTurn() || len(gameState.playersCyclicOrder) == len(gameState.trick) || !gameState.cardPlayList[CARDS_MOVES[card]] {
			return
		}
		_, err := gClient.Call(context.Background(), makeMoveProcedure, nil, wamp.List{gameState.clientPlayerName, CARDS_MOVES[card]}, nil, nil)
		if err != nil {
			connectionErr <- err
		}
	}
	game.hands[0] = NewHandWidget([]*card{}, []*card{}, true, playCard)
	game.hands[1] = NewHandWidget([]*card{}, []*card{}, false, doNothing)
	game.hands[2] = NewHandWidget([]*card{}, []*card{}, true, doNothing)
	game.hands[3] = NewHandWidget([]*card{}, []*card{}, false, doNothing)

	game.trick = NewTrickWidget(0, [4]*card{ /*NINE_OF_CLUBS, TEN_OF_CLUBS, nil, QUEEN_OF_CLUBS*/ })

	game.playerList = widget.NewListWithData(
		gameState.playerListBinding,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		},
	)

	game.bidList = widget.NewListWithData(
		gameState.bidListBinding,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		},
	)
	game.bidList.OnSelected = func(id widget.ListItemID) {
		_, err := gClient.Call(context.Background(), makeMoveProcedure, nil, wamp.List{gameState.clientPlayerName, gameState.bidListBindingData[id]}, nil, nil)
		if err != nil {
			connectionErr <- err
		}
		game.bidList.UnselectAll()
	}

	game.turn = widget.NewLabelWithData(gameState.turn) //TODO: może coś innego niż ten obsrany binding? to jest bardzo niewygodne
	game.quitGameButton = widget.NewButton("WYJDŹ", func() {
		closeConnection <- 0
		w.SetContent(screens.mainMenuScreen)
	})
	game.gameState = gameState

	return game
}

func (p *gameElements) Update() {
	for i := 0; i < 4; i++ {
		name, err := p.getNameForHandWidgetsOrder(i)
		if err != nil {
			continue
		}
		if i == 0 {
			p.hands[i].SetHandwidget(p.gameState.playerCards, p.gameState.hiddenCardsFor(name))
			continue
		}
		p.hands[i].SetHandwidget([]*card{}, p.gameState.hiddenCardsFor(name))
	}
	var turnIndex, clientPlayerIndex, leadIndex int
	playersCount := len(p.gameState.playersCyclicOrder)
	turnName, err := p.gameState.turn.Get()
	if err != nil {
		panic(err) //TODO log error albo w ogóle wywal to, gdy pozbędziesz się bindingu
	}
	for i, v := range p.gameState.playersCyclicOrder {
		if turnName == v {
			turnIndex = i
		}
		if v == p.gameState.clientPlayerName {
			clientPlayerIndex = i
		}
	}
	leadIndex = (turnIndex - len(p.gameState.trick) + playersCount) % playersCount

	trick := [4]*card{}
	leadWidgetIndex := 0
	for i := range trick {
		gameStateTrickCardIndex := (clientPlayerIndex - leadIndex + i + playersCount) % playersCount
		if playersCount == 3 && i > 2 {
			gameStateTrickCardIndex--
			gameStateTrickCardIndex += playersCount
			gameStateTrickCardIndex %= playersCount
		}
		if gameStateTrickCardIndex >= len(p.gameState.trick) || (playersCount == 3 && i == 2) {
			continue
		}
		if gameStateTrickCardIndex == 0 {
			leadWidgetIndex = i
		}
		trick[i] = p.gameState.trick[gameStateTrickCardIndex]
	}

	p.trick.SetTrickWidget(uint8(leadWidgetIndex), trick)
}

func (p *gameElements) getNameForHandWidgetsOrder(index int) (string, error) {
	if index > 3 || index < 0 {
		return "", errors.New("out of bounds error - ")
	}
	if len(p.gameState.playersCardCount) == 3 && index == 2 {
		return "", errors.New("top hand should remain empty while there are only three players")
	}
	if len(p.gameState.playersCardCount) == 3 && index == 3 {
		index = 2
	}

	var clientPlayerIndex int
	playersCount := len(p.gameState.playersCyclicOrder)
	for i, v := range p.gameState.playersCyclicOrder {
		if v == p.gameState.clientPlayerName {
			clientPlayerIndex = i
		}
	}

	return p.gameState.playersCyclicOrder[(index+clientPlayerIndex)%playersCount], nil
}

func zeroLengthStringValidator(s string) error {
	if len(s) == 0 {
		return errors.New("pole musi zawierać co najmniej jeden znak")
	}

	return nil
}
