package main

import (
	"testing"
)

const (
	testInputFilePath = "test.txt"
)

// The instructions specify a series of steps and requirements about
// which steps must be finished before others can begin (your puzzle input).
// Each step is designated by a single letter. For example, suppose you
// have the following instructions:

// Step C must be finished before step A can begin.
// Step C must be finished before step F can begin.
// Step A must be finished before step B can begin.
// Step A must be finished before step D can begin.
// Step B must be finished before step E can begin.
// Step D must be finished before step E can begin.
// Step F must be finished before step E can begin.
// Visually, these requirements look like this:

//   -->A--->B--
//  /    \      \
// C      -->D----->E
//  \           /
//   ---->F-----
// Your first goal is to determine the order in which the steps should
//  be completed. If more than one step is ready, choose the step which
//  is first alphabetically. In this example, the steps would be completed
//  as follows:

// Only C is available, and so it is done first.
// Next, both A and F are available. A is first alphabetically, so it is done next.
// Then, even though F was available earlier, steps B and D are now also
//  available, and B is the first alphabetically of the three.
// After that, only D and F are available. E is not available because
//  only some of its prerequisites are complete. Therefore, D is completed next.
// F is the only choice, so it is done next.
// Finally, E is completed.
// So, in this example, the correct order is CABDFE.

func TestSolveExercise(t *testing.T) {
	t.Run("known output", func(t *testing.T) {
		expectedInstructionsOrder := "CABDFE"
		instructionsOrder := solveExercise(testInputFilePath)

		if instructionsOrder != expectedInstructionsOrder {
			t.Errorf("expected safe region size to be equal to %q, got %q", expectedInstructionsOrder, instructionsOrder)
		}
	})
}
