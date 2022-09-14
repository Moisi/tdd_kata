package main

import (
	"fmt"
)

type Game struct {
	GameName    string
	FrameNumber uint
	RollNumber  uint
	Score       uint
}

func (e Game) printGame() {
	fmt.Printf("%s %d %d %d\n", e.GameName, e.FrameNumber, e.RollNumber, e.Score)
}
