package mechanic

import "errors"

type gtShiftsNode struct {
	nodes map[*GameType]*gtShiftsNode
}

type gameTypeShiftsFSM struct {
	changeLog []*GameType
	finished  bool
}

func newGameTypeShiftsFSM() *gameTypeShiftsFSM {
	gameTypeShiftsFSM := new(gameTypeShiftsFSM)
	gameTypeShiftsFSM.changeLog = []*GameType{}
	gameTypeShiftsFSM.finished = false

	return gameTypeShiftsFSM
}

func (p *gameTypeShiftsFSM) getPossibleShifts() *gtShiftsNode {
	var possibleShiftsLeft *gtShiftsNode = &allShifts
	for _, entry := range p.changeLog {
		possibleShiftsLeft = possibleShiftsLeft.nodes[entry]
	}
	return possibleShiftsLeft
}

func (p *gameTypeShiftsFSM) changeTo(g *GameType) error {
	err := errors.New("impossible change of state in gameTypesShiftsFSM")
	if p.finished {
		return err
	}

	var possibleShiftsLeft *gtShiftsNode = p.getPossibleShifts()
	_, ok := possibleShiftsLeft.nodes[g]
	if !ok {
		return err
	}

	p.changeLog = append(p.changeLog, g)
	if len(possibleShiftsLeft.nodes) == 0 {
		p.finished = true
	}

	return nil
}

func (p *gameTypeShiftsFSM) getCurrentGameType() *GameType {
	changeLogLength := len(p.changeLog)
	if changeLogLength == 0 {
		return &WARSAW
	}

	return p.changeLog[changeLogLength-1]
}

var worse gtShiftsNode = gtShiftsNode{nodes: map[*GameType]*gtShiftsNode{}}
var worseLess gtShiftsNode = gtShiftsNode{nodes: map[*GameType]*gtShiftsNode{}}
var better gtShiftsNode = gtShiftsNode{nodes: map[*GameType]*gtShiftsNode{
	&WORSE: &worse,
}}
var betterLess gtShiftsNode = gtShiftsNode{nodes: map[*GameType]*gtShiftsNode{
	&WORSE_LESS: &worseLess,
}}

var heartsAsked gtShiftsNode = gtShiftsNode{nodes: map[*GameType]*gtShiftsNode{
	&BETTER: &better,
	&WORSE:  &worse,
}}
var heartsUnasked gtShiftsNode = gtShiftsNode{nodes: map[*GameType]*gtShiftsNode{}}
var heartsLessAsked gtShiftsNode = gtShiftsNode{nodes: map[*GameType]*gtShiftsNode{
	&BETTER_LESS:    &betterLess,
	&WORSE_LESS:     &worseLess,
	&HEARTS_ASKED:   &heartsAsked,
	&HEARTS_UNASKED: &heartsUnasked,
}}
var heartsLessUnasked gtShiftsNode = gtShiftsNode{nodes: map[*GameType]*gtShiftsNode{}}

var diamondsAsked gtShiftsNode = gtShiftsNode{nodes: map[*GameType]*gtShiftsNode{
	&BETTER: &better,
	&WORSE:  &worse,
}}
var diamondsUnasked gtShiftsNode = gtShiftsNode{nodes: map[*GameType]*gtShiftsNode{}}
var diamondsLessAsked gtShiftsNode = gtShiftsNode{nodes: map[*GameType]*gtShiftsNode{
	&BETTER_LESS:      &betterLess,
	&WORSE_LESS:       &worseLess,
	&DIAMONDS_ASKED:   &diamondsAsked,
	&DIAMONDS_UNASKED: &diamondsUnasked,
}}
var diamondsLessUnasked gtShiftsNode = gtShiftsNode{nodes: map[*GameType]*gtShiftsNode{}}

var clubsAsked gtShiftsNode = gtShiftsNode{nodes: map[*GameType]*gtShiftsNode{
	&BETTER: &better,
	&WORSE:  &worse,
}}
var clubsUnasked gtShiftsNode = gtShiftsNode{nodes: map[*GameType]*gtShiftsNode{}}
var clubsLessAsked gtShiftsNode = gtShiftsNode{nodes: map[*GameType]*gtShiftsNode{
	&BETTER_LESS:   &betterLess,
	&WORSE_LESS:    &worseLess,
	&CLUBS_ASKED:   &clubsAsked,
	&CLUBS_UNASKED: &clubsUnasked,
}}
var clubsLessUnasked gtShiftsNode = gtShiftsNode{nodes: map[*GameType]*gtShiftsNode{}}

var spadesAsked gtShiftsNode = gtShiftsNode{nodes: map[*GameType]*gtShiftsNode{
	&BETTER: &better,
	&WORSE:  &worse,
}}
var spadesUnasked gtShiftsNode = gtShiftsNode{nodes: map[*GameType]*gtShiftsNode{}}
var spadesLessAsked gtShiftsNode = gtShiftsNode{nodes: map[*GameType]*gtShiftsNode{
	&BETTER_LESS:    &betterLess,
	&WORSE_LESS:     &worseLess,
	&SPADES_ASKED:   &spadesAsked,
	&SPADES_UNASKED: &spadesUnasked,
}}
var spadesLessUnasked gtShiftsNode = gtShiftsNode{nodes: map[*GameType]*gtShiftsNode{}}

var allShifts gtShiftsNode = gtShiftsNode{nodes: map[*GameType]*gtShiftsNode{
	&HEARTS_LESS_ASKED:     &heartsLessAsked,
	&HEARTS_LESS_UNASKED:   &heartsLessUnasked,
	&DIAMONDS_LESS_ASKED:   &diamondsLessAsked,
	&DIAMONDS_LESS_UNASKED: &diamondsLessUnasked,
	&CLUBS_LESS_ASKED:      &clubsLessAsked,
	&CLUBS_LESS_UNASKED:    &clubsLessUnasked,
	&SPADES_LESS_ASKED:     &spadesLessAsked,
	&SPADES_LESS_UNASKED:   &spadesLessUnasked,
	&BETTER_LESS:           &betterLess,
	&WORSE_LESS:            &worseLess,
}}

var gameTypesCanLookAtAllCards map[*GameType]*GameType = map[*GameType]*GameType{
	&HEARTS_LESS_ASKED:   &HEARTS_LESS_ASKED,
	&DIAMONDS_LESS_ASKED: &DIAMONDS_LESS_ASKED,
	&CLUBS_LESS_ASKED:    &CLUBS_LESS_ASKED,
	&SPADES_LESS_ASKED:   &SPADES_LESS_ASKED,
}

type GameType struct {
	maxPoints uint8
	trump     suit
	name      string
}

var WARSAW GameType = GameType{3, NONE_SUIT, "warsaw"}

var HEARTS_LESS_UNASKED GameType = GameType{12, HEARTS_SUIT, "heartsLessUnasked"}
var HEARTS_LESS_ASKED GameType = GameType{6, HEARTS_SUIT, "heartsLessAsked"}
var HEARTS_ASKED GameType = GameType{3, HEARTS_SUIT, "heartsAsked"}
var HEARTS_UNASKED GameType = GameType{6, HEARTS_SUIT, "heartsUnasked"}

var DIAMONDS_LESS_UNASKED GameType = GameType{12, DIAMONDS_SUIT, "diamondsLessUnasked"}
var DIAMONDS_LESS_ASKED GameType = GameType{6, DIAMONDS_SUIT, "diamondsLessAsked"}
var DIAMONDS_ASKED GameType = GameType{3, DIAMONDS_SUIT, "diamondsAsked"}
var DIAMONDS_UNASKED GameType = GameType{6, DIAMONDS_SUIT, "diamondsUnasked"}

var CLUBS_LESS_UNASKED GameType = GameType{12, CLUBS_SUIT, "clubsLessUnasked"}
var CLUBS_LESS_ASKED GameType = GameType{6, CLUBS_SUIT, "clubsLessAsked"}
var CLUBS_ASKED GameType = GameType{3, CLUBS_SUIT, "clubsAsked"}
var CLUBS_UNASKED GameType = GameType{6, CLUBS_SUIT, "clubsUnasked"}

var SPADES_LESS_UNASKED GameType = GameType{12, SPADES_SUIT, "spadesLessUnasked"}
var SPADES_LESS_ASKED GameType = GameType{6, SPADES_SUIT, "spadesLessAsked"}
var SPADES_ASKED GameType = GameType{3, SPADES_SUIT, "spadesAsked"}
var SPADES_UNASKED GameType = GameType{6, SPADES_SUIT, "spadesUnasked"}

var WORSE_LESS GameType = GameType{12, NONE_SUIT, "worseLess"}
var WORSE GameType = GameType{6, NONE_SUIT, "worse"}
var BETTER_LESS GameType = GameType{12, NONE_SUIT, "betterLess"}
var BETTER GameType = GameType{6, NONE_SUIT, "better"}

var gamesThatCanBeStoppedOnDemand [16]*GameType = [16]*GameType{
	&HEARTS_LESS_UNASKED,
	&HEARTS_LESS_ASKED,
	&HEARTS_ASKED,
	&HEARTS_UNASKED,
	&DIAMONDS_LESS_UNASKED,
	&DIAMONDS_LESS_ASKED,
	&DIAMONDS_ASKED,
	&DIAMONDS_UNASKED,
	&CLUBS_LESS_UNASKED,
	&CLUBS_LESS_ASKED,
	&CLUBS_ASKED,
	&CLUBS_UNASKED,
	&SPADES_LESS_UNASKED,
	&SPADES_LESS_ASKED,
	&SPADES_ASKED,
	&SPADES_UNASKED,
}

var gamesThatCanBeDoubled []*GameType = []*GameType{
	&HEARTS_ASKED,
	&HEARTS_LESS_ASKED,
	&DIAMONDS_ASKED,
	&DIAMONDS_LESS_ASKED,
	&CLUBS_ASKED,
	&CLUBS_LESS_ASKED,
	&SPADES_ASKED,
	&SPADES_LESS_ASKED,
	&BETTER_LESS,
	&BETTER,
	&WORSE_LESS,
	&WORSE,
}

var gamesThatStopAuction map[*GameType]*GameType = map[*GameType]*GameType{
	&HEARTS_LESS_UNASKED:   &HEARTS_LESS_UNASKED,
	&HEARTS_UNASKED:        &HEARTS_UNASKED,
	&DIAMONDS_LESS_UNASKED: &DIAMONDS_LESS_UNASKED,
	&DIAMONDS_UNASKED:      &DIAMONDS_UNASKED,
	&CLUBS_LESS_UNASKED:    &CLUBS_LESS_UNASKED,
	&CLUBS_UNASKED:         &CLUBS_UNASKED,
	&SPADES_LESS_UNASKED:   &SPADES_LESS_UNASKED,
	&SPADES_UNASKED:        &SPADES_UNASKED,
}

// var warsaw gtShiftsNode = gtShiftsNode{nodes: map[*gameType]*gtShiftsNode{}}
// var possibleGameTypeShifts map[*gameType]map[*gameType]map[*gameType]map[*gameType]*gameType = map[*gameType]map[*gameType]map[*gameType]map[*gameType]*gameType{
// 	&WARSAW: {},

// 	&HEARTS_LESS_ASKED: {
// 		&BETTER_LESS: {
// 			&WORSE_LESS: {},
// 		},
// 		&WORSE_LESS: {},
// 		&HEARTS_ASKED: {
// 			&BETTER: {
// 				&WORSE: {},
// 			},
// 			&WORSE: {},
// 		},
// 		&HEARTS_UNASKED: {},
// 	},
// 	&HEARTS_LESS_UNASKED: {},

// 	&DIAMONDS_LESS_ASKED: {
// 		&BETTER_LESS: {
// 			&WORSE_LESS: {},
// 		},
// 		&WORSE_LESS: {},
// 		&DIAMONDS_ASKED: {
// 			&BETTER: {
// 				&WORSE: {},
// 			},
// 			&WORSE: {},
// 		},
// 		&DIAMONDS_UNASKED: {},
// 	},
// 	&DIAMONDS_LESS_UNASKED: {},

// 	&CLUBS_LESS_ASKED: {
// 		&BETTER_LESS: {
// 			&WORSE_LESS: {},
// 		},
// 		&WORSE_LESS: {},
// 		&CLUBS_ASKED: {
// 			&BETTER: {
// 				&WORSE: {},
// 			},
// 			&WORSE: {},
// 		},
// 		&CLUBS_UNASKED: {},
// 	},
// 	&CLUBS_LESS_UNASKED: {},

// 	&SPADES_LESS_ASKED: {
// 		&BETTER_LESS: {
// 			&WORSE_LESS: {},
// 		},
// 		&WORSE_LESS: {},
// 		&SPADES_ASKED: {
// 			&BETTER: {
// 				&WORSE: {},
// 			},
// 			&WORSE: {},
// 		},
// 		&SPADES_UNASKED: {},
// 	},
// 	&SPADES_LESS_UNASKED: {},

// 	&BETTER_LESS: {
// 		&WORSE_LESS: {},
// 	},
// 	&WORSE_LESS: {},
// }
