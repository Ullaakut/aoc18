package main

import (
	"bytes"
	"testing"
)

const (
	testInputFilePath = "test.txt"
)

// The boxes will have IDs which differ by exactly one character
// at the same position in both strings. For example, given the
// following box IDs:

// abcde
// fghij
// klmno
// pqrst
// fguij
// axcye
// wvxyz

// The IDs abcde and axcye are close, but they differ by two
// characters (the second and fourth). However, the IDs fghij and
// fguij differ by exactly one character, the third (h and u).
// Those must be the correct boxes.

// In the example above, this is found by removing the differing
// character from either ID, producing fgij.

func TestSolveExercise(t *testing.T) {
	t.Run("known output", func(t *testing.T) {
		expectedLetters := []byte("fgij")
		letters := solveExercise(testInputFilePath)

		if bytes.Compare(letters, expectedLetters) != 0 {
			t.Errorf("expected common letters to be equal to %s, got %s", expectedLetters, letters)
		}
	})
}
