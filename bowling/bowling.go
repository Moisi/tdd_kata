package main

import (
	"fmt"
)

const NIL_PREVROLL = 100

type Frame struct {
	Rolls [3]int
}

func newFrame() Frame {
	f := Frame{[3]int{0, 0, 0}}
	return f
}

func (f Frame) IsSpare() bool {
	return f.Rolls[0]+f.Rolls[1] == 10
}

func (f Frame) IsStrike() bool {
	return f.Rolls[0] == 10
}

func (f Frame) printFrame() {
	fmt.Println("Rolls: ", f.Rolls)
	fmt.Println("IsSpare: ", f.IsSpare())
	fmt.Println("IsStrike: ", f.IsStrike())

}

type Game struct {
	GameName    string
	FrameNumber uint
	RollNumber  uint
	Score       uint
	IsSpare     bool
	IsStrike    bool
	PrevRoll    uint

	FrameArray [10]Frame
}

func (e Game) printGame() {
	fmt.Printf("%s %d %d %d\n", e.GameName, e.FrameNumber, e.RollNumber, e.Score)
}

func newGame(GameName string) Game {
	g := Game{GameName, 0, 0, 0, false, false, NIL_PREVROLL,
		[10]Frame{newFrame(), newFrame(), newFrame(),
			newFrame(), newFrame(), newFrame(),
			newFrame(), newFrame(), newFrame(),
			newFrame()}}
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
	} else if e.IsStrike {
		e.Score += NumOfPinsDown
	}

	// prepare state for next roll:
	if NumOfPinsDown+e.PrevRoll == 10 {
		// congrats we have a spare:
		e.IsSpare = true
		e.IsStrike = false
		e.PrevRoll = NIL_PREVROLL
	} else if e.PrevRoll == NIL_PREVROLL && NumOfPinsDown == 10 {
		// found a strike:
		e.IsStrike = true
	} else if e.PrevRoll == NIL_PREVROLL {
		// PrevRoll == NIL_PREVROLL == 100, meaning this is first roll in frame.
		// need to save it for second roll
		// might do this in a different way (via some state) but for now.
		e.PrevRoll = NumOfPinsDown
	} else {
		// this was second roll. make sure IsStrike = false;
		e.IsStrike = false
		e.PrevRoll = NIL_PREVROLL
	}
}

func (e *Game) score() uint {
	return e.Score
}
