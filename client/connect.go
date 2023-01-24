package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gammazero/nexus/v3/client"
	"github.com/gammazero/nexus/v3/wamp"
)

const (
	ws    = "ws://"
	port  = ":8080/"
	realm = "realm1"

	allPlayersTopic        = "game.players"
	gameStateChangeTopic   = "game.change"
	connectToGameProcedure = "game.connect"
	getGameStateProcedure  = "game.state"
	makeMoveProcedure      = "game.move"
)

func play(
	address string,
	nick string,
	close chan int,
	connectionErr chan<- error,
	gameState *gameState,
	gClient *client.Client,
) {
	logger := log.New(os.Stdout, "", 0)
	cfg := client.Config{
		Realm:  realm,
		Logger: logger,
	}

	gameClient, err := client.ConnectNet(context.Background(), ws+address+port, cfg)
	if err != nil {
		connectionErr <- err
		return
	}
	*gClient = *gameClient
	defer gClient.Close()
	defer fmt.Println("Client closed")

	callArgs := wamp.List{nick}

	//funkcja initAllPlayersHandling nadąża powiedzieć routerowi, że się zapisuje do topicu, ale jeszcze nie ogarnęła handlera u siebie,
	//przez ten czas do ogarnięcia handlera ten klient właśnie się podłączył do rozgrywki - miejsce A poniżej,
	//i router rozporowadza informację od gry o nowym stanie graczy, choć nie ma handlera, który mógłby to obsłużyć
	//rozwiązanie: waitGroup albo może mutex? doczytać i się zastanowić
	go initAllPlayersHandling(logger, gameState, gClient, close, connectionErr)
	go initGameStateChangeHandling(logger, gClient, close, connectionErr, callArgs, gameState)

	//miejsce A
	_, err = gClient.Call(context.Background(), connectToGameProcedure, nil, callArgs, nil, nil)
	if err != nil {
		connectionErr <- err
		return
	}

	<-close
}

func initAllPlayersHandling(logger *log.Logger, gameState *gameState, gClient *client.Client, close chan int, connectionErr chan<- error) {
	eventHandler := func(event *wamp.Event) {
		logger.Println("Received", allPlayersTopic, "event")
		updatePlayerList(event.Arguments, gameState)
	}
	if err := gClient.Subscribe(allPlayersTopic, eventHandler, nil); err != nil {
		connectionErr <- err
		close <- 0
	}
}

func initGameStateChangeHandling(logger *log.Logger, gClient *client.Client, close chan int, connectionErr chan<- error, callArgs wamp.List, gameState *gameState) {
	eventHandler := func(event *wamp.Event) {
		go func() {
			gameStateResult, err := gClient.Call(context.Background(), getGameStateProcedure, nil, callArgs, nil, nil)
			if err != nil {
				logger.Println("call error:", err)
				connectionErr <- err
				close <- 0
			}
			gameState.updateGameState(gameStateResult.ArgumentsKw)
		}()
	}
	if err := gClient.Subscribe(gameStateChangeTopic, eventHandler, nil); err != nil {
		connectionErr <- err
		close <- 0
	}
}

func updatePlayerList(args wamp.List, gameState *gameState) {
	gameState.playerListBinding.Set([]string{})
	for _, p := range args {
		if playerName, ok := wamp.AsString(p); ok {
			gameState.playerListBinding.Append(playerName)
		}
	}
}
