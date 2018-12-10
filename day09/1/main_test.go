package main

import (
	"fmt"
	"testing"
)

const (
	testInputFilePath = "test.txt"
)

// The Elves play this game by taking turns arranging the marbles in a circle according to very particular rules.
// The marbles are numbered starting with 0 and increasing by 1 until every marble has a number.

// First, the marble numbered 0 is placed in the circle. At this point, while it contains only a single marble,
// it is still a circle: the marble is both clockwise from itself and counter-clockwise from itself. This marble
// is designated the current marble.

// Then, each Elf takes a turn placing the lowest-numbered remaining marble into the circle between the marbles that
// are 1 and 2 marbles clockwise of the current marble. (When the circle is large enough, this means that there is one
// marble between the marble that was just placed and the current marble.) The marble that was just placed then becomes
// the current marble.

// However, if the marble that is about to be placed has a number which is a multiple of 23, something entirely
// different happens. First, the current player keeps the marble they would have placed, adding it to their score.
// In addition, the marble 7 marbles counter-clockwise from the current marble is removed from the circle and also
// added to the current player's score. The marble located immediately clockwise of the marble that was removed
// becomes the new current marble.

// For example, suppose there are 9 players. After the marble with value 0 is placed in the middle, each player
// (shown in square brackets) takes a turn. The result of each of those turns would produce circles of marbles
// like this, where clockwise is to the right and the resulting current marble is in parentheses:

// [1]  0 (1)
// [2]  0 (2) 1
// [3]  0  2  1 (3)
// [4]  0 (4) 2  1  3
// [5]  0  4  2 (5) 1  3
// [6]  0  4  2  5  1 (6) 3
// [7]  0  4  2  5  1  6  3 (7)
// [8]  0 (8) 4  2  5  1  6  3  7
// [9]  0  8  4 (9) 2  5  1  6  3  7
// [1]  0  8  4  9  2(10) 5  1  6  3  7
// [2]  0  8  4  9  2 10  5(11) 1  6  3  7
// [3]  0  8  4  9  2 10  5 11  1(12) 6  3  7
// [4]  0  8  4  9  2 10  5 11  1 12  6(13) 3  7
// [5]  0  8  4  9  2 10  5 11  1 12  6 13  3(14) 7
// [6]  0  8  4  9  2 10  5 11  1 12  6 13  3 14  7(15)
// [7]  0(16) 8  4  9  2 10  5 11  1 12  6 13  3 14  7 15
// [8]  0 16  8(17) 4  9  2 10  5 11  1 12  6 13  3 14  7 15
// [9]  0 16  8 17  4(18) 9  2 10  5 11  1 12  6 13  3 14  7 15
// [1]  0 16  8 17  4 18  9(19) 2 10  5 11  1 12  6 13  3 14  7 15
// [2]  0 16  8 17  4 18  9 19  2(20)10  5 11  1 12  6 13  3 14  7 15
// [3]  0 16  8 17  4 18  9 19  2 20 10(21) 5 11  1 12  6 13  3 14  7 15
// [4]  0 16  8 17  4 18  9 19  2 20 10 21  5(22)11  1 12  6 13  3 14  7 15
// [5]  0 16  8 17  4 18(19) 2 20 10 21  5 22 11  1 12  6 13  3 14  7 15
// [6]  0 16  8 17  4 18 19  2(24)20 10 21  5 22 11  1 12  6 13  3 14  7 15
// [7]  0 16  8 17  4 18 19  2 24 20(25)10 21  5 22 11  1 12  6 13  3 14  7 15

// The goal is to be the player with the highest score after the last marble is used up.
// Assuming the example above ends after the marble numbered 25, the winning score is 23+9=32
// (because player 5 kept marble 23 and removed marble 9, while no other player got any points
// in this very short example game).

func TestSolveExercise(t *testing.T) {
	testCases := []struct {
		players         int
		lastMarbleScore int

		expectedHighScore int
	}{
		{players: 9, lastMarbleScore: 25, expectedHighScore: 32},
		{players: 10, lastMarbleScore: 1618, expectedHighScore: 8317},
		{players: 13, lastMarbleScore: 7999, expectedHighScore: 146373},
		{players: 17, lastMarbleScore: 1104, expectedHighScore: 2764},
		{players: 21, lastMarbleScore: 6111, expectedHighScore: 54718},
		{players: 30, lastMarbleScore: 5807, expectedHighScore: 37305},
		{players: 479, lastMarbleScore: 71035, expectedHighScore: 367634},
		{players: 479, lastMarbleScore: 7103500, expectedHighScore: 3020072891},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("%d/%d/%d", test.players, test.lastMarbleScore, test.expectedHighScore), func(t *testing.T) {
			highScore := solveExercise(test.players, test.lastMarbleScore)

			if highScore != test.expectedHighScore {
				t.Errorf("expected high score to be %d, got %d", test.expectedHighScore, highScore)
			}
		})
	}
}
