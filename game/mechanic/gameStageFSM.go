package mechanic

import "errors"

type gameStageFSM struct {
	currentGameStage gameStage
	possibleShifts   map[gameStage]map[gameStage]gameStage
}

func newGameStageFSM() *gameStageFSM {
	gameStageFSM := new(gameStageFSM)
	gameStageFSM.currentGameStage = AUCTION
	gameStageFSM.possibleShifts = map[gameStage]map[gameStage]gameStage{
		AUCTION: {
			ROUND: ROUND,
		},
		ROUND: {
			AUCTION: AUCTION,
			FINISH:  FINISH,
		},
		FINISH: {},
	}

	return gameStageFSM
}

func (p *gameStageFSM) change(from, to gameStage) error {
	if _, ok := p.possibleShifts[from][to]; !ok {
		return errors.New("impossible change of state in gameStageFSM")
	}
	p.currentGameStage = to

	return nil
}







type gameStage int

const (
	AUCTION gameStage = iota
	ROUND
	FINISH
)
