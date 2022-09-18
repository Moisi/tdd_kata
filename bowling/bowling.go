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

	if e.RollIndex < 2 && NumOfPinsDown == 10 && e.FrameIndex == 9 {
		// case of strike in roll1 or roll2. we'll use the third roll
		// in the same frame:
		e.RollIndex++
	} else if e.FrameIndex == 9 && currentFrame.Rolls[0]+currentFrame.Rolls[1] == 10 {
		// case we got a spare in last frame
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

func (e *Game) getBonusForLastFrame(frame int, SorS bool) int {
	/*
		hacky workaround for edge case.
		we have 3 optins for bonus in 10th frame:
		1. spare <- get bonus From Rolls[2]
		2. Rolls[0] && Rolls[1] are strike <- bonus on Rolls[2]
		3. Strike on Roll[0] only <- bonus on Rolls[1] + [2]
	*/
	if e.FramesArray[frame].Rolls[1]+e.FramesArray[frame].Rolls[0] == 10 {
		// case 1:
		return e.FramesArray[frame].Rolls[2]
	} else if e.FramesArray[frame].Rolls[1] == 10 && e.FramesArray[frame].Rolls[0] == 10 {
		// case 2:
		return e.FramesArray[frame].Rolls[1]
	} else /* if e.FramesArray[frame].Rolls[0] == 10*/ {
		// case 3:
		return e.FramesArray[frame].Rolls[1] +
			e.FramesArray[frame].Rolls[2]
	}
}

func (e *Game) getBonusForFrame(frame int, SorS bool) int {
	if frame == 9 {
		return e.getBonusForLastFrame(frame, SorS)
	} else {
		nextFrame := e.FramesArray[frame+1]
		if SorS {
			// strike - return sum
			// of next two rolls
			if nextFrame.Rolls[0] == 10 {
				// found another strike, thus return 10 plus
				// next rolls which will the in next frame:
				return 10 + e.getBonusForFrame(frame+1, false)
			} else {
				// two are in same frame:
				return nextFrame.Rolls[0] + nextFrame.Rolls[1]
			}
		} else {
			// spare - return next roll
			return nextFrame.Rolls[0]
		}
	}

}

func (e *Game) score() uint {
	score := 0
	for i, f := range e.FramesArray {
		score += f.Rolls[0]
		score += f.Rolls[1]
		if f.Rolls[0] == 10 {
			score += e.getBonusForFrame(i, true)
		} else if f.Rolls[0]+f.Rolls[1] == 10 {
			score += e.getBonusForFrame(i, false)
		}
	}
	return uint(score)
}
