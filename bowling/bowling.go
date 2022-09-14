package main

import (
	"fmt"
)

const NIL_PREVROLL = 100

type Game struct {
	GameName    string
	FrameNumber uint
	RollNumber  uint
	Score       uint
	IsSpare     bool
	PrevRoll    uint
}

func (e Game) printGame() {
	fmt.Printf("%s %d %d %d\n", e.GameName, e.FrameNumber, e.RollNumber, e.Score)
}

func newGame(GameName string) Game {
	g := Game{GameName, 0, 0, 0, false, NIL_PREVROLL}
	g.IsSpare = false
	return g
}

func (e *Game) roll(NumOfPinsDown uint) {
	// calc current run
	e.Score += NumOfPinsDown

	// add bonus if prev was spare
	if e.IsSpare {
		// add the run after spare
		e.Score += NumOfPinsDown
		e.IsSpare = false
	}

	// prepare state for next roll:
	if NumOfPinsDown+e.PrevRoll == 10 {
		// congrats we have a spare:
		e.IsSpare = true
		e.PrevRoll = NIL_PREVROLL
	} else if e.PrevRoll == NIL_PREVROLL {
		// PrevRoll == NIL_PREVROLL == 100, meaning this is first roll in frame.
		// need to save it for second roll
		// might do this in a different way (via some state) but for now.
		e.PrevRoll = NumOfPinsDown
	} else {
		e.PrevRoll = NIL_PREVROLL
	}
}

func (e *Game) score() uint {
	return e.Score
}
