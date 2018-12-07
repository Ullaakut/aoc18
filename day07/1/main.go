package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
)

type instructions map[byte][]byte

var stepRegex = regexp.MustCompile("^Step ([A-Z]+) must be finished before step ([A-Z]+) can begin.$")

const inputFilePath = "input.txt"

func computeInstructionsOrder(instructions instructions) string {
	done := []byte{}

	for {
		notNext := []byte{}

		if len(instructions) == 1 {
			for id, prev := range instructions {
				log.Printf("Next instruction should be %q", id)
				sort.Slice(prev, func(i, j int) bool {
					return prev[i] < prev[j]
				})
				res := append(done, id)

				for _, p := range prev {
					log.Printf("Next instruction should be %q", p)
				}
				res = append(res, prev...)

				return string(res)
			}
		}

		for _, previous := range instructions {
			for _, prevID := range previous {
				notNext = append(notNext, prevID)
			}
		}

		potentialNext := []byte{}
		for id := range instructions {
			if !bytes.Contains(notNext, []byte{id}) {
				potentialNext = append(potentialNext, id)
			}
		}

		sort.Slice(potentialNext, func(i, j int) bool {
			return potentialNext[i] < potentialNext[j]
		})

		log.Printf("Next instruction should be %q", potentialNext[0])
		done = append(done, potentialNext[0])
		delete(instructions, potentialNext[0])
	}
}

func solveExercise(filePath string) string {
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(fmt.Sprint("Unable to read input file:", err))
	}

	lines := bytes.Split(contents, []byte("\n"))

	instructions := make(instructions)
	for _, line := range lines {
		matches := stepRegex.FindAllSubmatch(line, -1)
		if len(matches) == 0 || len(matches[0]) != 3 {
			continue
		}

		ID := matches[0][1][0]
		prev := matches[0][2][0]

		instructions[ID] = append(instructions[ID], prev)
	}

	return computeInstructionsOrder(instructions)
}

func main() {
	log.Println("Beginning day07ex01...")

	instructionsOrder := solveExercise(inputFilePath)

	log.Println("Instructions order successfully computed")
	log.Printf("Instructions order: %s\n", instructionsOrder)
}
