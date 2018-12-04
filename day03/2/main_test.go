package main

import (
	"testing"
)

const (
	testInputFilePath = "test.txt"
)

// A claim like #123 @ 3,2: 5x4 means that claim ID 123 specifies
// a rectangle 3 inches from the left edge, 2 inches from the top
// edge, 5 inches wide, and 4 inches tall. Visually, it claims the
// square inches of fabric represented by # (and ignores the square
// inches of fabric represented by .) in the diagram below:

// ...........
// ...........
// ...#####...
// ...#####...
// ...#####...
// ...#####...
// ...........
// ...........
// ...........
// The problem is that many of the claims overlap, causing two or
// more claims to cover part of the same areas. For example,
// consider the following claims:

// #1 @ 1,3: 4x4
// #2 @ 3,1: 4x4
// #3 @ 5,5: 2x2
// Visually, these claim the following areas:

// ........
// ...2222.
// ...2222.
// .11XX22.
// .11XX22.
// .111133.
// .111133.
// ........

// In the example above, the only non-overlapping area is the one from the elf
// with the ID 3.

func TestSolveExercise(t *testing.T) {
	t.Run("known output", func(t *testing.T) {
		expectedClaimID := "3"
		claimID := solveExercise(testInputFilePath)

		if claimID != expectedClaimID {
			t.Errorf("expected unique claim ID to be equal to %s, got %s", expectedClaimID, claimID)
		}
	})
}
