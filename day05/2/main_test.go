package main

import (
	"testing"
)

const (
	testInputFilePath = "test.txt"
)

// One of the unit types is causing problems; it's preventing the polymer from collapsing as much as it should.
// Your goal is to figure out which unit type is causing the most problems, remove all instances of it (regardless of polarity),
// fully react the remaining polymer, and measure its length.

// For example, again using the polymer dabAcCaCBAcCcaDA from above:

// Removing all A/a units produces dbcCCBcCcD. Fully reacting this polymer produces dbCBcD, which has length 6.
// Removing all B/b units produces daAcCaCAcCcaDA. Fully reacting this polymer produces daCAcaDA, which has length 8.
// Removing all C/c units produces dabAaBAaDA. Fully reacting this polymer produces daDA, which has length 4.
// Removing all D/d units produces abAcCaCBAcCcaA. Fully reacting this polymer produces abCBAc, which has length 6.
// In this example, removing all C/c units was best, producing the answer 4.

func TestSolveExercise(t *testing.T) {
	t.Run("known output", func(t *testing.T) {
		expectedPolymerLength := 4
		polymerLength := solveExercise(testInputFilePath)

		if polymerLength != expectedPolymerLength {
			t.Errorf("expected improved polymer length to be equal to %q, got %q", expectedPolymerLength, polymerLength)
		}
	})
}
