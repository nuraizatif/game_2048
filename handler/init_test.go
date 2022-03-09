package handler

import "testing"

var testGame = [4][4]int{
	{0, 0, 0, 0},
	{0, 0, 0, 0},
	{0, 0, 0, 1},
	{0, 0, 0, 0},
}

func TestInitGame(t *testing.T) {
	game := InitGame()
	if game != testGame {
		t.Errorf("got %q, wanted %q", game, testGame)
	}
}
