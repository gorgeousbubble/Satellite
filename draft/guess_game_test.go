package draft

import "testing"

func TestGuessGame1(t *testing.T) {
	guess := []int{1, 2, 3}
	answer := []int{1, 2, 3}
	count := guess_game(guess, answer)
	if count != 3 {
		t.Fatal("Error execute guess_game.")
	}
}

func TestGuessGame2(t *testing.T) {
	guess := []int{2, 2, 3}
	answer := []int{3, 2, 1}
	count := guess_game(guess, answer)
	if count != 1 {
		t.Fatal("Error execute guess_game.")
	}
}