package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const inputFilePath = "test.txt"

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
	// Add empty pots to the left of the current state
	for i := -4; i < 0; i++ {
		pots = push(pots, pot{
			index: pots[0].index + i,
			plant: false,
		})
	}

	// Add empty pots to the right of the current state
	for i := 1; i < 5; i++ {
		pots = append(pots, pot{
			index: pots[len(pots)-1].index + i,
			plant: false,
		})
	}

	return pots
}

func removeEmptyPots(pots []pot) []pot {
	// Remove empty pots to the left
	for pots[0].plant == false {
		pots = pots[1:]
	}

	// Remove empty pots to the right
	for i := len(pots) - 1; pots[i].plant == false; i-- {
		pots = pots[:i]
	}

	return pots
}

func computeNextGeneration(pots []pot, rules []rule) []pot {
	state := []bool{false, false, false, false}

	pots = addEmptyPots(pots)

	for idx, pot := range pots {
		state = append(state, pot.plant)

		for _, rule := range rules {
			if stateMatches(rule.state, state) {
				pots[idx].plant = rule.effect
			}
		}
		if len(state) == 5 {
			state = state[1:]
		}
	}

	pots = removeEmptyPots(pots)

	return pots
}

func printPots(generation int, pots []pot) {
	line := fmt.Sprintf("%d:\t", generation)

	for _, pot := range pots {
		if pot.plant {
			line += "#"
		} else {
			line += "."
		}
	}

	log.Println(line)
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
		printPots(i, pots)
		pots = computeNextGeneration(pots, rules)
	}

	return computeTotalSum(pots)
}

func main() {
	log.Println("Beginning day12ex01...")

	log.Printf("Sum of pot numbers after 20 generations: %d", solveExercise(inputFilePath))
}
