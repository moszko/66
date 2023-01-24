package main

import (
	"66_game/helper"
	"errors"
)

type maxPlayerCount int

const (
	Three maxPlayerCount = iota + 3
	Four
)

// type PlayerCollection interface {
// 	addPlayer(name string, player Player) error
// 	isPlayerCountMaxed() bool
// }

type GamePlayerCollection struct {
	maxCount  maxPlayerCount
	players   map[string]Player
	observers map[helper.Observer]bool
}

func (p *GamePlayerCollection) addPlayer(name string, player Player) error {
	if p.isPlayerCountMaxed() {
		return errors.New("the number of players is already at maximum level")
	}

	if p.exist(name) {
		return errors.New("that name has already been taken")
	}

	p.players[name] = player
	if p.isPlayerCountMaxed() {
		p.Notify()
	}

	return nil
}

func (p *GamePlayerCollection) exist(name string) bool {
	_, ok := p.players[name]

	return ok
}

func (p *GamePlayerCollection) isPlayerCountMaxed() bool {
	return len(p.players) == int(p.maxCount)
}

func (p *GamePlayerCollection) getPlayerNames() []string {
	playerNames := []string{}
	for k := range p.players {
		playerNames = append(playerNames, k)
	}

	return playerNames
}

func (p *GamePlayerCollection) getPlayers() []Player {
	players := []Player{}
	for _, player := range p.players {
		// fmt.Println(player)
		players = append(players, player)
	}

	return players
}

func (p *GamePlayerCollection) AddObserver(o helper.Observer) {
	p.observers[o] = true
}

func (p *GamePlayerCollection) RemoveObserver(o helper.Observer) {
	delete(p.observers, o)
}

func (p *GamePlayerCollection) Notify() {
	for o := range p.observers {
		o.Update()
	}
}
