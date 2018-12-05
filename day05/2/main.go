package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

const inputFilePath = "input.txt"

func improvePolymer(polymer []byte) int {
	var (
		bestImprovedUnit    byte
		bestImprovedPolymer int
	)

	for i := byte('a'); i <= byte('z'); i++ {
		improvedPolymer := bytes.Replace(polymer, []byte{i}, []byte{}, -1)
		improvedPolymer = bytes.Replace(improvedPolymer, []byte{i - 32}, []byte{}, -1)

		reducedPolymer := computePolymerReduction(improvedPolymer)
		if len(reducedPolymer) < bestImprovedPolymer || bestImprovedPolymer == 0 {
			bestImprovedPolymer = len(reducedPolymer)
			bestImprovedUnit = i

			log.Printf("Discovered better polymer improvement by removing unit %q, resulting in a %d unit long polymer", bestImprovedUnit, bestImprovedPolymer)
		}
	}

	return bestImprovedPolymer
}

func computePolymerReduction(polymer []byte) []byte {
	stable := false
	for {
		if stable {
			break
		}

		for idx, unit := range polymer {
			if idx+1 < len(polymer) {
				neighbor := polymer[idx+1]
				// If unit is lowercase and its neighbor is its
				// uppercase equivalent, they react
				if 65 <= unit && unit <= 90 && neighbor-unit == 32 {
					polymer = append(polymer[:idx], polymer[idx+2:]...)
					break
				}

				// If unit is uppercase and its neighbor is its
				// lowercase equivalent, they react
				if 97 <= unit && unit <= 122 && unit-neighbor == 32 {
					polymer = append(polymer[:idx], polymer[idx+2:]...)
					break
				}
			} else {
				stable = true
			}
		}
	}

	return polymer
}

func solveExercise(inputPath string) int {
	contents, err := ioutil.ReadFile(inputPath)
	if err != nil {
		log.Fatal(fmt.Sprint("Unable to read input file:", err))
	}

	contents = bytes.Replace(contents, []byte("\n"), []byte(""), -1)

	return improvePolymer(contents)
}

func main() {
	log.Println("Beginning day05ex02...")

	polymerLength := solveExercise(inputFilePath)

	log.Println("Polymer improvements successfully computed")
	log.Printf("Improved polymer length: %d\n", polymerLength)
}
