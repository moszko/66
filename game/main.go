package main

import (
	"66_game/helper"
	"66_game/mechanic"
	"context"
	"fmt"

	"log"
	"os"
	"os/signal"

	"github.com/gammazero/nexus/v3/client"
	"github.com/gammazero/nexus/v3/wamp"
)

const (
	addr  = "ws://localhost:8080/"
	realm = "realm1"

	playersTopic           = "game.players"
	connectToGameProcedure = "game.connect"
	gameStateChangeTopic   = "game.change"
	getGameStateProcedure  = "game.state"
	makeMoveProcedure      = "game.move"
)

func main() {
	logger := log.New(os.Stdout, "", 0)
	cfg := client.Config{
		Realm:  realm,
		Logger: logger,
	}

	gClient, err := client.ConnectNet(context.Background(), addr, cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer gClient.Close()

	var game mechanic.Game
	players := &GamePlayerCollection{maxCount: Three, players: map[string]Player{}, observers: map[helper.Observer]bool{}}
	starter := gameStarter{game: &game, clientObserver: &clientObserver{client: gClient}, players: players}
	players.AddObserver(&starter)

	if err = gClient.Register(connectToGameProcedure, connectToGame(logger, gClient, starter), nil); err != nil {
		logger.Fatal("Failed to register procedure:", err)
	}
	logger.Println("Registered procedure", connectToGameProcedure, "with router")

	if err = gClient.Register(getGameStateProcedure, getGameState(&game, logger, players), nil); err != nil {
		logger.Fatal("Failed to register procedure:", err)
	}
	logger.Println("Registered procedure", getGameStateProcedure, "with router")

	if err = gClient.Register(makeMoveProcedure, makeMove(&game, players), nil); err != nil {
		logger.Fatal("Failed to register procedure:", err)
	}
	logger.Println("Registered procedure", makeMoveProcedure, "with router")

	// Wait for CTRL-c or client close while handling events.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	select {
	case <-sigChan:
	case <-gClient.Done():
		logger.Print("Router gone, exiting")
		return
	}
}

type clientObserver struct {
	client *client.Client
}

func (p *clientObserver) Update() {
	p.client.Publish(gameStateChangeTopic, nil, nil, nil)
}

type gameStarter struct {
	game           *mechanic.Game
	clientObserver *clientObserver
	players        *GamePlayerCollection
}

func (p *gameStarter) Update() {
	players := p.players.getPlayers()
	*p.game = *mechanic.NewGame(&players[0], &players[1], &players[2])
	p.game.AddObserver(p.clientObserver)
	fmt.Println("start the game")
	p.clientObserver.client.Publish(gameStateChangeTopic, nil, nil, nil)
}

func connectToGame(logger *log.Logger, gClient *client.Client, starter gameStarter) func(ctx context.Context, i *wamp.Invocation) client.InvokeResult {
	return func(ctx context.Context, i *wamp.Invocation) client.InvokeResult {
		newPlayerName, ok := wamp.AsString(i.Arguments[0])
		if !ok {
			return client.InvokeResult{Err: wamp.URI("err.bad.player.name")}
		}
		newPlayer := Player{name: newPlayerName}

		error := starter.players.addPlayer(newPlayerName, newPlayer)
		if error != nil {
			return client.InvokeResult{Err: wamp.URI("err.bad.player.name")}
		}
		logger.Println("connected player:", newPlayerName)

		playerNames, ok := wamp.AsList(starter.players.getPlayerNames())
		if !ok {
			logger.Fatal(ok)
		}

		err := gClient.Publish(playersTopic, nil, playerNames, nil)
		if err != nil {
			logger.Fatal("publish error:", err)
		}

		return client.InvokeResult{Args: playerNames}
	}
}

func getGameState(game *mechanic.Game, logger *log.Logger, players *GamePlayerCollection) func(ctx context.Context, i *wamp.Invocation) client.InvokeResult {
	return func(ctx context.Context, i *wamp.Invocation) client.InvokeResult {
		playerName, ok := wamp.AsString(i.Arguments[0])
		// fmt.Println(playerName)
		if !ok || !players.exist(playerName) {
			return client.InvokeResult{Err: wamp.URI("err.bad.player.name")}
		}

		player := players.players[playerName]
		gameState, ok := wamp.AsDict(game.StateFor(&player))
		// defer fmt.Println("game", gameState, ok)
		if !ok {
			logger.Fatal(ok)
		}

		return client.InvokeResult{Kwargs: gameState}
	}
}

func makeMove(game *mechanic.Game, players *GamePlayerCollection) func(ctx context.Context, i *wamp.Invocation) client.InvokeResult {
	return func(ctx context.Context, i *wamp.Invocation) client.InvokeResult {
		playerName, ok := wamp.AsString(i.Arguments[0])
		fmt.Println(i.Arguments...)
		if !ok || !players.exist(playerName) {
			return client.InvokeResult{Err: wamp.URI("err.bad.player.absent")}
		}
		moveName, ok := wamp.AsString(i.Arguments[1])
		if !ok {
			return client.InvokeResult{Err: wamp.URI("err.bad.move")}
		}

		err := game.MakeMove(mechanic.ALL_NAMES_MOVES[moveName], &Player{playerName})
		if err != nil {
			return client.InvokeResult{Err: wamp.URI("err.bad.player.or.move")}
		}

		return client.InvokeResult{}
	}
}
