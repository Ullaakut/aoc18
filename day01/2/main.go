package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const inputFilePath = "input.txt"

func main() {
	log.Println("Beginning day01ex02...")

	frequencyDrift := int64(0)
	previousFrequencies := make(map[int64]struct{})
	previousFrequencies[0] = struct{}{}

	content, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		log.Fatal(fmt.Sprint("Unable to read input file:", err))
	}

	shifts := strings.Split(string(content), "\n")

	for {
		for _, shift := range shifts {
			if shift == "" {
				continue
			}

			frequencyChange, err := strconv.ParseInt(shift, 10, 64)
			if err != nil {
				log.Fatal(fmt.Sprint("Unable to parse frequency change:", err))
			}

			frequencyDrift = frequencyDrift + frequencyChange

			_, exists := previousFrequencies[frequencyDrift]
			if exists {
				log.Println("Duplicate frequency discovered")
				log.Printf("Frequency: %d\n", frequencyDrift)
				return
			}

			previousFrequencies[frequencyDrift] = struct{}{}
		}
	}
}
