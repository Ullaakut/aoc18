package main

import (
	"testing"
)

const (
	testInputFilePath = "test.txt"
)

// Suppose you want the sum of the Manhattan distance to all of the
// coordinates to be less than 32. For each location, add up the distances
// to all of the given coordinates; if the total of those distances is
// less than 32, that location is within the desired region. Using the
// same coordinates as above, the resulting region looks like this:

// ..........
// .A........
// ..........
// ...###..C.
// ..#D###...
// ..###E#...
// .B.###....
// ..........
// ..........
// ........F.

// In particular, consider the highlighted location 4,3 located at the top middle of
// the region. Its calculation is as follows, where abs() is the absolute value function:

// Distance to coordinate A: abs(4-1) + abs(3-1) =  5
// Distance to coordinate B: abs(4-1) + abs(3-6) =  6
// Distance to coordinate C: abs(4-8) + abs(3-3) =  4
// Distance to coordinate D: abs(4-3) + abs(3-4) =  2
// Distance to coordinate E: abs(4-5) + abs(3-5) =  3
// Distance to coordinate F: abs(4-8) + abs(3-9) = 10
// Total distance: 5 + 6 + 4 + 2 + 3 + 10 = 30
// Because the total distance to all coordinates (30) is less than 32, the location
// is within the region.

// This region, which also includes coordinates D and E, has a total size of 16.

func TestSolveExercise(t *testing.T) {
	t.Run("known output", func(t *testing.T) {
		expectedSafeRegionSize := 16
		safeRegionSize := solveExercise(testInputFilePath, 32)

		if safeRegionSize != expectedSafeRegionSize {
			t.Errorf("expected safe region size to be equal to %d, got %d", expectedSafeRegionSize, safeRegionSize)
		}
	})
}
