package main

import (
	"testing"
)

func TestGameInit(t *testing.T) {
	g := Game{GameName: "tmpName",
		FrameNumber: 1,
		RollNumber:  2,
		Score:       3}
	g.printGame()
}
