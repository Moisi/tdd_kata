package main

import (
	"testing"
)

func TestGameInit(t *testing.T) {
	g := newGame("tmpName")
	g.printGame()
}
