package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const inputFilePath = "input.txt"

// similarBoxIDs takes two box IDs and returns whether or not they have only one
// letter that differs, as well as the common characters in the ID.
func similarBoxIDs(boxID1, boxID2 []byte) (bool, []byte) {
	if len(boxID1) != len(boxID2) {
		return false, nil
	}

	var (
		differences   int
		commonLetters []byte
	)
	for index, character := range boxID1 {
		if boxID2[index] != character {
			differences++
		} else {
			commonLetters = append(commonLetters, character)
		}
	}

	// BoxIDs are similar only if they have less than 2 differences.
	return differences < 2, commonLetters
}

func solveExercise(inputPath string) []byte {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(fmt.Sprint("Unable to open input file:", err))
	}

	content, err := ioutil.ReadAll(inputFile)
	if err != nil {
		log.Fatal(fmt.Sprint("Unable to read input file:", err))
	}

	boxIDs := bytes.Split(content, []byte("\n"))
	for position1, boxID1 := range boxIDs {
		for position2, boxID2 := range boxIDs {
			// Skip this element, as it's the same box ID
			if position1 == position2 {
				continue
			}

			similarBoxIDs, commonLetters := similarBoxIDs(boxID1, boxID2)
			if similarBoxIDs {
				log.Println("Similar BoxIDs found")
				log.Printf("BoxIDs: %s - %s\n", boxID1, boxID2)
				return commonLetters
			}
		}
	}

	return nil
}

func main() {
	log.Println("Beginning day02ex02...")

	log.Printf("Common letters: %s\n", solveExercise(inputFilePath))
}
