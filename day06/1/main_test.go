package main

import (
	"testing"
)

const (
	testInputFilePath = "test.txt"
)

// Your goal is to find the size of the largest area that isn't infinite.
// For example, consider the following list of coordinates:

// 1, 1
// 1, 6
// 8, 3
// 3, 4
// 5, 5
// 8, 9
// If we name these coordinates A through F, we can draw them on
// a grid, putting 0,0 at the top left:

// ..........
// .A........
// ..........
// ........C.
// ...D......
// .....E....
// .B........
// ..........
// ..........
// ........F.
// This view is partial - the actual grid extends infinitely in all directions.
// Using the Manhattan distance, each location's closest coordinate can be
// determined, shown here in lowercase:

// aaaaa.cccc
// aAaaa.cccc
// aaaddecccc
// aadddeccCc
// ..dDdeeccc
// bb.deEeecc
// bBb.eeee..
// bbb.eeefff
// bbb.eeffff
// bbb.ffffFf
// Locations shown as . are equally far from two or more coordinates,
// and so they don't count as being closest to any.

// In this example, the areas of coordinates A, B, C, and F are
// infinite - while not shown here, their areas extend forever outside
// the visible grid. However, the areas of coordinates D and E are finite:
// D is closest to 9 locations, and E is closest to 17 (both including the
// coordinate's location itself). Therefore, in this example, the
// size of the largest area is 17.

func TestSolveExercise(t *testing.T) {
	t.Run("known output", func(t *testing.T) {
		expectedBiggestArea := 17
		biggestArea := solveExercise(testInputFilePath)

		if biggestArea != expectedBiggestArea {
			t.Errorf("expected biggest area size to be equal to %d, got %d", expectedBiggestArea, biggestArea)
		}
	})
}
