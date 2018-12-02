package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const inputFilePath = "input.txt"

func main() {
	log.Println("Beginning day01ex01...")

	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(fmt.Sprint("Unable to open input file:", err))
	}

	totalFrequencyChange := int64(0)
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		frequencyChange, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprint("Unable to parse frequency change:", err))
		}

		totalFrequencyChange = totalFrequencyChange + frequencyChange
	}

	log.Println("All frequency changes successfully computed")
	log.Printf("Frequency drift: %d\n", totalFrequencyChange)
}
