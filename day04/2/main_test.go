package main

import (
	"testing"
)

const (
	testInputFilePath = "test.txt"
)

// Date   ID   Minute
//             000000000011111111112222222222333333333344444444445555555555
//             012345678901234567890123456789012345678901234567890123456789
// 11-01  #10  .....####################.....#########################.....
// 11-02  #99  ........................................##########..........
// 11-03  #10  ........................#####...............................
// 11-04  #99  ....................................##########..............
// 11-05  #99  .............................................##########.....

// In the example above, Guard #99 spent minute 45 asleep more than any other guard or minute - three times in total.
// Of all guards, which guard is most frequently asleep on the same minute?
// In the above example, the answer would be 99 * 45 = 4455.)

func TestSolveExercise(t *testing.T) {
	t.Run("known output", func(t *testing.T) {
		expectedChecksum := 4455
		checksum := solveExercise(testInputFilePath)
		if checksum != expectedChecksum {
			t.Errorf("expected cheksum to be equal to %d, got %d", expectedChecksum, checksum)
		}
	})
}
