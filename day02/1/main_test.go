package main

import (
	"testing"
)

const (
	testInputFilePath = "test.txt"
)

// abcdef contains no letters that appear exactly two or three times.
// bababc contains two a and three b, so it counts for both.
// abbcde contains two b, but no letter appears exactly three times.
// abcccd contains three c, but no letter appears exactly two times.
// aabcdd contains two a and two d, but it only counts once.
// abcdee contains two e.
// ababab contains three a and three b, but it only counts once.

// Of these box IDs, four of them contain a letter which appears exactly
// twice, and three of them contain a letter which appears exactly three
// times. Multiplying these together produces a checksum of 4 * 3 = 12.

func TestSolveExercise(t *testing.T) {
	t.Run("known output", func(t *testing.T) {
		expectedTriples := 3
		expectedDoubles := 4
		doubles, triples := solveExercise(testInputFilePath)

		if doubles != expectedDoubles {
			t.Errorf("expected doubles to be equal to %d, got %d", expectedDoubles, doubles)
		}

		if triples != expectedTriples {
			t.Errorf("expected triples to be equal to %d, got %d", expectedTriples, triples)
		}
	})
}
