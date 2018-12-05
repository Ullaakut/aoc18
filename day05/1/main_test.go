package main

import (
	"bytes"
	"testing"
)

const (
	testInputFilePath = "test.txt"
)

// The polymer is formed by smaller units which, when triggered, react with each other
// such that two adjacent units of the same type and opposite polarity are destroyed. Units'
// types are represented by letters; units' polarity is represented by capitalization. For
// instance, r and R are units with the same type but opposite polarity, whereas r and s are
// entirely different types and do not react.

// For example:

// In aA, a and A react, leaving nothing behind.
// In abBA, bB destroys itself, leaving aA. As above, this then destroys itself, leaving nothing.
// In abAB, no two adjacent units are of the same type, and so nothing happens.
// In aabAAB, even though aa and AA are of the same type, their polarities match, and so nothing happens.
// Now, consider a larger example, dabAcCaCBAcCcaDA:

// dabAcCaCBAcCcaDA  The first 'cC' is removed.
// dabAaCBAcCcaDA    This creates 'Aa', which is removed.
// dabCBAcCcaDA      Either 'cC' or 'Cc' are removed (the result is the same).
// dabCBAcaDA        No further actions can be taken.
// After all possible reactions, the resulting polymer contains 10 units.

func TestSolveExercise(t *testing.T) {
	t.Run("known output", func(t *testing.T) {
		expectedPolymer := []byte("dabCBAcaDA")
		polymer := solveExercise(testInputFilePath)

		if bytes.Compare(polymer, expectedPolymer) != 0 {
			t.Errorf("expected polymer to be equal to %q, got %q", expectedPolymer, polymer)
		}
	})
}
