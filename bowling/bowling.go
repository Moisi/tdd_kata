package main

import (
	"fmt"
)

const NIL_PREVROLL = 100

type Frame struct {
	Rolls           [3]int
	FinalScoreCache int
}

func newFrame() Frame {
	f := Frame{[3]int{0, 0, 0}, -1}
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
	FrameIndex  int
	RollIndex   int
	PrevRoll    uint

	FramesArray [10]Frame
}

func (e Game) printGame() {
	fmt.Printf("%s %d %d %d\n", e.GameName, e.FrameNumber, e.RollNumber, e.Score)
}

func newGame(GameName string) Game {
	g := Game{GameName, 0, 0, 0, 0, 0, NIL_PREVROLL,
		[10]Frame{newFrame(), newFrame(), newFrame(),
			newFrame(), newFrame(), newFrame(),
			newFrame(), newFrame(), newFrame(),
			newFrame()}}
	return g
}

func (e *Game) roll(NumOfPinsDown uint) {
	if e.FrameIndex == 10 || (e.FrameIndex == 9 && e.RollIndex == 2) {
		return
	}

	currentFrame := e.FramesArray[e.FrameIndex]
	e.FramesArray[e.FrameIndex].Rolls[e.RollIndex] = int(NumOfPinsDown)

	if e.FrameIndex == 9 && currentFrame.Rolls[0]+currentFrame.Rolls[1] == 10 {
		// last frame + strike | spare
		e.RollIndex = 2
	} else if e.RollIndex == 0 && NumOfPinsDown == 10 {
		// strike, jump to next Frame
		e.FrameIndex++
		e.RollIndex = 0
	} else if e.RollIndex == 1 {
		e.FrameIndex++
		e.RollIndex = 0
	} else {
		e.RollIndex = 1
	}
}

func (e *Game) getBonus(i int, SorS bool) int {
	if i == 10 {
		// this is last frame, we store next roll in:
		// Rolls[2]:
		return e.FramesArray[i-1].Rolls[2]
	} else {
		f := e.FramesArray[i]
		if SorS {
			// strike - return sum
			// of next two rolls
			if f.Rolls[0] == 10 {
				// next two rolls are in two frames:
				return 10 + e.getBonus(i+1, false)
			} else {
				// two are in same frame:
				return f.Rolls[0] + f.Rolls[1]
			}
		} else {
			// spare - return next roll
			return f.Rolls[0]
		}
	}

}

func (e *Game) score() uint {
	score := 0
	for i, f := range e.FramesArray {
		score += f.Rolls[0]
		score += f.Rolls[1]
		if f.Rolls[0] == 10 {
			score += e.getBonus(i+1, true)
		} else if f.Rolls[0]+f.Rolls[1] == 10 {
			score += e.getBonus(i+1, false)
		}
	}
	return uint(score)
}
