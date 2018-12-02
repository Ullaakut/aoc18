package main

import (
	"bufio"
	"fmt"
	"index/suffixarray"
	"log"
	"os"
)

const inputFilePath = "input.txt"

// inspectBoxID takes a box ID and returns whether or not it contains doubles and triples
func inspectBoxID(boxID []byte) (double bool, triple bool) {
	index := suffixarray.New(boxID)

	for _, letter := range boxID {
		results := index.Lookup([]byte{letter}, -1)
		if len(results) == 2 {
			double = true
		} else if len(results) == 3 {
			triple = true
		}
	}

	return
}

func main() {
	log.Println("Beginning day02ex01...")

	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(fmt.Sprint("Unable to open input file:", err))
	}

	var triples, doubles int

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		double, triple := inspectBoxID(scanner.Bytes())

		if double {
			doubles++
		}

		if triple {
			triples++
		}
	}

	log.Println("Box IDs checksum successfully computed")
	log.Printf("Checksum: %d\n", triples*doubles)
}
