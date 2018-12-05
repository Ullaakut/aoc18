package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

const inputFilePath = "test.txt"

func computePolymerReduction(content []byte) []byte {
	stable := false
	for {
		if stable {
			break
		}

		for idx, unit := range content {
			if idx+1 < len(content) {
				neighbor := content[idx+1]
				// If unit is lowercase and its neighbor is its
				// uppercase equivalent, they react
				if 65 <= unit && unit <= 90 && neighbor-unit == 32 {
					content = append(content[:idx], content[idx+2:]...)
					log.Printf("Reaction between %s and %s: Polymer reduced to: %s", []byte{unit}, []byte{neighbor}, content)
					break
				}

				// If unit is uppercase and its neighbor is its
				// lowercase equivalent, they react
				if 97 <= unit && unit <= 122 && unit-neighbor == 32 {
					content = append(content[:idx], content[idx+2:]...)
					log.Printf("Reaction between %s and %s: Polymer reduced to: %s", []byte{unit}, []byte{neighbor}, content)
					break
				}
			} else {
				stable = true
			}
		}
	}

	return content
}

func solveExercise(inputPath string) []byte {
	contents, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		log.Fatal(fmt.Sprint("Unable to read input file:", err))
	}

	contents = bytes.Replace(contents, []byte("\n"), []byte(""), -1)

	return computePolymerReduction(contents)
}

func main() {
	log.Println("Beginning day05ex01...")

	polymer := solveExercise(inputFilePath)

	log.Println("Alchemical reduction successfully computed")
	log.Printf("Reduced polymer: %s\n", polymer)
}
