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
		t.Fatalf("score() = %d, want = %d", score, want)
	}
}

func TestScoreAfterOneRollNoBonus(t *testing.T) {
	g := newGame("ScoreOfFreshGame")

	g.roll(1)
	want := uint(1)
	score := g.score()
	if score != want {
		t.Fatalf("score() = %d, want = %d", score, want)
	}
}

func TestScoreAfterTwoRollBonus(t *testing.T) {
	g := newGame("ScoreOfFreshGame")
	g.roll(1)
	g.roll(1)

	want2 := uint(2)
	score2 := g.score()
	if score2 != want2 {
		t.Fatalf("score() = %d, want = %d", score2, want2)
	}
}

func TestScoreAfterSpare(t *testing.T) {
	g := newGame("ScoreOfFreshGame")

	g.roll(1)
	g.roll(9)

	g.roll(1)
	g.roll(1)

	// 10 for pins in the first frame + 1 bonus + 2 of the second frame
	// frame0: 10 + bonus
	// frame1: 1+1 = 2. but only 1 is bonus
	want := uint(10 + 1 + 2)
	score := g.score()
	if score != want {
		t.Fatalf("score() = %d, want = %d", score, want)
	}
}

func TestScoreDoubleSpare(t *testing.T) {
	g := newGame("ScoreOfFreshGame")

	g.roll(1)
	g.roll(9)

	g.roll(1)
	g.roll(9)

	g.roll(1)
	g.roll(1)

	want := uint(10 + 1 + 10 + 1 + 2)
	score := g.score()
	if score != want {
		t.Fatalf("score = %d, want = %d", score, want)
	}
}

func TestScoreTwoNonSpareAfterSpare(t *testing.T) {
	g := newGame("ScoreOfFreshGame")

	g.roll(1)
	g.roll(9)

	g.roll(1)
	g.roll(1)

	g.roll(1)
	g.roll(1)

	want := uint(10 + 1 + 2 + 2)
	score := g.score()
	if score != want {
		t.Fatalf("score = %d, want = %d", score, want)
	}
}

func TestScoreAfterStrike(t *testing.T) {
	g := newGame("ScoreOfFreshGame")

	g.roll(10)

	g.roll(1)
	g.roll(1)

	want := uint(10 + 1 + 1 + 2)
	score := g.score()
	if score != want {
		t.Fatalf("score() = %d, want = %d", score, want)
	}
}

func TestScoreAfterDoubleStrike(t *testing.T) {
	g := newGame("ScoreOfFreshGame")

	g.roll(10)

	g.roll(10)

	g.roll(1)
	g.roll(1)

	// 10 + bonus (10) + 10 + bonus (2) + 2
	want := uint(34)
	score := g.score()
	if score != want {
		t.Fatalf("ScoreOfFreshGame = %d, want = %d", score, want)
	}
}

func TestScoreStrikeThenSpareThenRegular(t *testing.T) {
	g := newGame("ScoreOfFreshGame")

	g.roll(10)

	g.roll(9)
	g.roll(1)

	g.roll(1)
	g.roll(1)

	// 10 + b(10) + 10 + b(1) + 2
	want := uint(33)
	score := g.score()
	if score != want {
		t.Fatalf("score = %d, want = %d", score, want)
	}
}

func TestScoreTwoRegsAfterStrike(t *testing.T) {
	g := newGame("ScoreOfFreshGame")

	g.roll(10)

	g.roll(1)
	g.roll(1)

	g.roll(1)
	g.roll(1)

	// 10 + b(2) + 2 + 2
	want := uint(16)
	score := g.score()
	if score != want {
		t.Fatalf("score = %d, want = %d", score, want)
	}
}

// flip Strike->Spare. to same but to Spare->Strike
func TestScoreSpareStrikeReg(t *testing.T) {
	g := newGame("ScoreOfFreshGame")

	g.roll(1)
	g.roll(9)

	g.roll(10)

	g.roll(1)
	g.roll(1)

	// 10 + bonus (10) + 10 + bonus (2) + 2
	want := uint(34)
	score := g.score()
	if score != want {
		t.Fatalf("score() = %d, want = %d", score, want)
	}
}
