package main

import (
	"testing"
)

func TestGameInit(t *testing.T) {
	g := newGame("tmpName")
	g.printGame()
}

func TestScoreOfFreshGame(t *testing.T) {
	// this is bad practice of TDD. because this test never failed first
	g := newGame("ScoreOfFreshGame")
	want := uint(0)
	score := g.score()
	if score != want {
		t.Fatalf("ScoreOfFreshGame = %d, want = %d", score, want)
	}
}

