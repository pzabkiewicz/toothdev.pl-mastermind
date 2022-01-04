package main

import "testing"

func TestProperHints(t *testing.T) {
	secret := []int{3, 5, 7, 7}
	guesses := []int{7, 8, 8, 5}
	want := []int{2, 1, 0, 0}
	got := AnalyzeGuessesAndGetHints(secret, guesses)

	if isEqual(got, want) {
		t.Errorf("Failed analysis and generation of hints. Want: %#v, Got: %#v", want, got)
	}
}

func isEqual(got []int, want []int) bool {
	if len(got) != len(want) {
		return false
	}
	for i, item := range got {
		if item != want[i] {
			return false
		}
	}
	return true
}
