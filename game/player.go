package main

type Player struct {
	name string
}

func (p *Player) Id() string {
	return p.name
}