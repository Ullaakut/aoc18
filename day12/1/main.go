package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const inputFilePath = "input.txt"

type pot struct {
	index int
	plant bool
}

type rule struct {
	state  []bool
	effect bool
}

func computeTotalSum(pots []pot) int {
	total := 0

	for _, pot := range pots {
		if pot.plant {
			total += pot.index
		}
	}

	return total
}

func stateMatches(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}

	for idx := range a {
		if a[idx] != b[idx] {
			return false
		}
	}

	return true
}

func push(pots []pot, p pot) []pot {
	return append([]pot{p}, pots...)
}

func addEmptyPots(pots []pot) []pot {
	firstIndex := pots[0].index

	// Add empty pots to the left
	for i := -4; i < 0; i++ {
		pots = push(pots, pot{
			index: firstIndex + i,
			plant: false,
		})
	}

	// Add empty pots to the right
	for i := 1; i < 5; i++ {
		pots = append(pots, pot{
			index: pots[len(pots)-1].index + i,
			plant: false,
		})
	}

	return pots
}

func arrangePots(pots []pot) []pot {
	for fIdx := range pots {
		if pots[fIdx].index == 0 {
			for idx := range pots {
				pots[idx].index = idx - fIdx
			}
			break
		}
	}

	return pots
}

func computeNextGeneration(pots []pot, rules []rule, debug bool) []pot {
	state := []bool{}

	pots = addEmptyPots(pots)

	for idx, pot := range pots {
		state = append(state, pot.plant)

		match := false
		for _, rule := range rules {
			if stateMatches(rule.state, state) {
				pots[idx-2].plant = rule.effect
				match = true
				break
			}
		}

		if match == false && idx > 2 {
			pots[idx-2].plant = false
		}

		if len(state) == 5 {
			state = state[1:]
		}
	}

	pots = arrangePots(pots)

	return pots
}

func solveExercise(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Unable to open input file: %v", err)
		os.Exit(1)
	}

	pots := []pot{}
	rules := []rule{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		potState := ""
		fmt.Sscanf(scanner.Text(), "initial state: %s", &potState)

		if potState != "" {
			for index, char := range potState {
				pots = append(pots, pot{
					index: index,
					plant: char == '#',
				})
			}
		}

		state, effect := "", ""
		fmt.Sscanf(scanner.Text(), "%s => %s", &state, &effect)

		if state != "" && effect != "" {
			rule := rule{
				effect: effect == "#",
			}

			for _, s := range state {
				rule.state = append(rule.state, s == '#')
			}

			rules = append(rules, rule)
		}
	}

	for i := 1; i <= 20; i++ {
		if i > 2 && i < 5 {
			pots = computeNextGeneration(pots, rules, true)
			printPots(i, pots)
			// printPotsIdx(i, pots)
		} else {
			pots = computeNextGeneration(pots, rules, false)
			printPots(i, pots)
		}
	}

	return computeTotalSum(pots)
}

func main() {
	log.Println("Beginning day12ex01...")
	log.Printf("Sum of pot numbers after 20 generations: %d", solveExercise(inputFilePath))
}

func printState(state []bool) {
	line := ""

	for _, s := range state {
		if s {
			line += "#"
		} else {
			line += "."
		}
	}

	log.Println(line)
}

func printPots(generation int, pots []pot) {
	// line := fmt.Sprintf("%2d: ", generation)

	line := ""
	// for _, pot := range pots {
	// 	if pot.plant {
	// 		line += "#"
	// 	} else {
	// 		line += "."
	// 	}
	// }

	for i := -6; i < 120; i++ {
		plant := false
		for _, pot := range pots {
			if pot.index == i && pot.plant {
				plant = true
				break
			}
		}

		if plant {
			line += "#"
		} else {
			line += "."
		}
	}

	fmt.Println(line)
}

func printPotsIdx(generation int, pots []pot) {
	line := fmt.Sprintf("%2d: ", generation)

	for _, pot := range pots {
		if pot.plant {
			line += fmt.Sprintf("%d, ", pot.index)
		}
	}

	fmt.Println(line)
}
