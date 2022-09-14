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

func TestScoreAfterOneRollNoBonus(t *testing.T) {
	g := newGame("ScoreOfFreshGame")

	g.roll(1)
	want := uint(1)
	score := g.score()
	if score != want {
		t.Fatalf("ScoreOfFreshGame = %d, want = %d", score, want)
	}
}

func TestScoreAfterTwoRollBonus(t *testing.T) {
	g := newGame("ScoreOfFreshGame")
	g.roll(1)
	g.roll(1)

	want2 := uint(2)
	score2 := g.score()
	if score2 != want2 {
		t.Fatalf("ScoreOfFreshGame = %d, want = %d", score2, want2)
	}
}
