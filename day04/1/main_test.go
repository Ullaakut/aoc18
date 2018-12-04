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

// In the example above, Guard #10 spent the most minutes asleep, a total of 50 minutes (20+25+5),
// while Guard #99 only slept for a total of 30 minutes (10+10+10). Guard #10 was asleep most during
// minute 24 (on two days, whereas any other minute the guard was asleep was only seen on one day).
// What is the ID of the guard you chose multiplied by the minute you chose?
// (In the above example, the answer would be 10 * 24 = 240.)

func TestSolveExercise(t *testing.T) {
	t.Run("known output", func(t *testing.T) {
		expectedChecksum := 240
		checksum := solveExercise(testInputFilePath)
		if checksum != expectedChecksum {
			t.Errorf("expected cheksum to be equal to %d, got %d", expectedChecksum, checksum)
		}
	})
}
