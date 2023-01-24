package mechanic

import "testing"

type goodMoveTest struct {
	movesList         []Mover
	resultingGameType *GameType
	gameStage         gameStage
}

type badMoveTest struct {
	movesList []Mover
}

var makeGoodMoveTest = []goodMoveTest{
	{[]Mover{}, &WARSAW, AUCTION},
	{[]Mover{PASS}, &WARSAW, AUCTION},
	{[]Mover{PASS, PASS}, &WARSAW, AUCTION},
	{[]Mover{PASS, PASS, PASS, START_ROUND}, &WARSAW, ROUND},
	{[]Mover{CHOOSE_HEARTS_LESS_UNASKED}, &HEARTS_LESS_UNASKED, ROUND},
	{[]Mover{CHOOSE_DIAMONDS_LESS_UNASKED}, &DIAMONDS_LESS_UNASKED, ROUND},
	{[]Mover{CHOOSE_CLUBS_LESS_UNASKED}, &CLUBS_LESS_UNASKED, ROUND},
	{[]Mover{CHOOSE_SPADES_LESS_UNASKED}, &SPADES_LESS_UNASKED, ROUND},
	{[]Mover{CHOOSE_HEARTS_LESS_ASKED, PASS, PASS, START_ROUND}, &HEARTS_LESS_ASKED, ROUND},
	{[]Mover{CHOOSE_DIAMONDS_LESS_ASKED, PASS, PASS, START_ROUND}, &DIAMONDS_LESS_ASKED, ROUND},
	{[]Mover{CHOOSE_CLUBS_LESS_ASKED, PASS, PASS, START_ROUND}, &CLUBS_LESS_ASKED, ROUND},
	{[]Mover{CHOOSE_SPADES_LESS_ASKED, PASS, PASS, START_ROUND}, &SPADES_LESS_ASKED, ROUND},
	{[]Mover{CHOOSE_HEARTS_LESS_ASKED, PASS, PASS, LOOK_AT_ALL_CARDS}, &HEARTS_LESS_ASKED, AUCTION},
}

var makeBadMoveTest = []badMoveTest{
	{[]Mover{START_ROUND}},
	{[]Mover{DOUBLING}},
	{[]Mover{LOOK_AT_ALL_CARDS}},
	{[]Mover{CHOOSE_WORSE}},
	{[]Mover{CHOOSE_BETTER}},
	{[]Mover{CHOOSE_HEARTS_ASKED}},
	{[]Mover{CHOOSE_HEARTS_UNASKED}},
	{[]Mover{CHOOSE_DIAMONDS_ASKED}},
	{[]Mover{CHOOSE_DIAMONDS_UNASKED}},
	{[]Mover{CHOOSE_CLUBS_ASKED}},
	{[]Mover{CHOOSE_CLUBS_UNASKED}},
	{[]Mover{CHOOSE_SPADES_ASKED}},
	{[]Mover{CHOOSE_SPADES_UNASKED}},
	{[]Mover{PASS, DOUBLING}},
	{[]Mover{PASS, START_ROUND}},
	{[]Mover{PASS, DOUBLING}},
	{[]Mover{PASS, LOOK_AT_ALL_CARDS}},
	{[]Mover{PASS, CHOOSE_WORSE}},
	{[]Mover{PASS, CHOOSE_BETTER}},
	{[]Mover{PASS, CHOOSE_HEARTS_ASKED}},
	{[]Mover{PASS, CHOOSE_HEARTS_UNASKED}},
	{[]Mover{PASS, CHOOSE_DIAMONDS_ASKED}},
	{[]Mover{PASS, CHOOSE_DIAMONDS_UNASKED}},
	{[]Mover{PASS, CHOOSE_CLUBS_ASKED}},
	{[]Mover{PASS, CHOOSE_CLUBS_UNASKED}},
	{[]Mover{PASS, CHOOSE_SPADES_ASKED}},
	{[]Mover{PASS, CHOOSE_SPADES_UNASKED}},
	{[]Mover{PASS, PASS, DOUBLING}},
	{[]Mover{PASS, PASS, START_ROUND}},
	{[]Mover{PASS, PASS, DOUBLING}},
	{[]Mover{PASS, PASS, LOOK_AT_ALL_CARDS}},
	{[]Mover{PASS, PASS, CHOOSE_WORSE}},
	{[]Mover{PASS, PASS, CHOOSE_BETTER}},
	{[]Mover{PASS, PASS, CHOOSE_HEARTS_ASKED}},
	{[]Mover{PASS, PASS, CHOOSE_HEARTS_UNASKED}},
	{[]Mover{PASS, PASS, CHOOSE_DIAMONDS_ASKED}},
	{[]Mover{PASS, PASS, CHOOSE_DIAMONDS_UNASKED}},
	{[]Mover{PASS, PASS, CHOOSE_CLUBS_ASKED}},
	{[]Mover{PASS, PASS, CHOOSE_CLUBS_UNASKED}},
	{[]Mover{PASS, PASS, CHOOSE_SPADES_ASKED}},
	{[]Mover{PASS, PASS, CHOOSE_SPADES_UNASKED}},
	{[]Mover{PASS, PASS, PASS, DOUBLING}},
	{[]Mover{PASS, PASS, PASS, DOUBLING}},
	{[]Mover{PASS, PASS, PASS, LOOK_AT_ALL_CARDS}},
	{[]Mover{PASS, PASS, PASS, CHOOSE_WORSE}},
	{[]Mover{PASS, PASS, PASS, CHOOSE_BETTER}},
	{[]Mover{PASS, PASS, PASS, CHOOSE_HEARTS_ASKED}},
	{[]Mover{PASS, PASS, PASS, CHOOSE_HEARTS_UNASKED}},
	{[]Mover{PASS, PASS, PASS, CHOOSE_DIAMONDS_ASKED}},
	{[]Mover{PASS, PASS, PASS, CHOOSE_DIAMONDS_UNASKED}},
	{[]Mover{PASS, PASS, PASS, CHOOSE_CLUBS_ASKED}},
	{[]Mover{PASS, PASS, PASS, CHOOSE_CLUBS_UNASKED}},
	{[]Mover{PASS, PASS, PASS, CHOOSE_SPADES_ASKED}},
	{[]Mover{PASS, PASS, PASS, CHOOSE_SPADES_UNASKED}},
	{[]Mover{PASS, PASS, PASS, PASS, DOUBLING}},
	{[]Mover{CHOOSE_HEARTS_LESS_ASKED, PASS, PASS, PASS}},
	{[]Mover{CHOOSE_DIAMONDS_LESS_ASKED, PASS, PASS, PASS}},
	{[]Mover{CHOOSE_CLUBS_LESS_ASKED, PASS, PASS, PASS}},
	{[]Mover{CHOOSE_SPADES_LESS_ASKED, PASS, PASS, PASS}},
	{[]Mover{CHOOSE_HEARTS_LESS_UNASKED, PASS}},
	{[]Mover{CHOOSE_BETTER_LESS, CHOOSE_HEARTS_LESS_ASKED}},
	{[]Mover{CHOOSE_BETTER_LESS, CHOOSE_DIAMONDS_LESS_UNASKED}},
	{[]Mover{CHOOSE_BETTER_LESS, CHOOSE_CLUBS_ASKED}},
	{[]Mover{CHOOSE_BETTER_LESS, CHOOSE_SPADES_UNASKED}},
	{[]Mover{CHOOSE_WORSE_LESS, CHOOSE_BETTER_LESS}},
	{[]Mover{CHOOSE_WORSE_LESS, CHOOSE_HEARTS_LESS_ASKED}},
	{[]Mover{CHOOSE_WORSE_LESS, CHOOSE_DIAMONDS_LESS_UNASKED}},
	{[]Mover{CHOOSE_WORSE_LESS, CHOOSE_CLUBS_ASKED}},
	{[]Mover{CHOOSE_WORSE_LESS, CHOOSE_SPADES_UNASKED}},
	{[]Mover{CHOOSE_HEARTS_LESS_ASKED, CHOOSE_HEARTS_ASKED}},
}

func TestMakeMove(t *testing.T) {
	//TODO: write tests for 4 players
	p1 := testPlayer{"1"}
	p2 := testPlayer{"2"}
	p3 := testPlayer{"3"}
	for _, test := range makeGoodMoveTest {
		gm := NewGame(p1, p2, p3)
		for _, move := range test.movesList {
			if err := gm.MakeMove(move, gm.players[gm.smallTurn]); err != nil {
				t.Errorf("couldn't make move: %q", move)
			}
		}
		if test.resultingGameType != gm.gameTypeShiftsFSM.getCurrentGameType() {
			t.Errorf("got %q, wanted %q", test.resultingGameType, gm.gameTypeShiftsFSM.getCurrentGameType())
		}
		if test.gameStage != gm.gameStageFSM.currentGameStage {
			t.Errorf("wanted %q, got %q", test.gameStage, gm.gameStageFSM.currentGameStage)
		}
	}
	for _, test := range makeBadMoveTest {
		gm := NewGame(p1, p2, p3)
		hasErrorOccured := false
		for _, move := range test.movesList {
			error := gm.MakeMove(move, gm.players[gm.smallTurn])
			if error != nil {
				hasErrorOccured = true
			}
		}
		if !hasErrorOccured {
			t.Error("no error happend during bad moves from ", test.movesList)
		}
	}
}

type testPlayer struct {
	name string
}

func (p testPlayer) Id() string {
	return p.name
}