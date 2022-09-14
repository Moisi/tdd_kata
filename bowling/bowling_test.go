package main

import (
	"testing"
)

func TestGameInit(t *testing.T) {
	g := newGame("tmpName", 1, 2, 3)
	g.printGame()
}
