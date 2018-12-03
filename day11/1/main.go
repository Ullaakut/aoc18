package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const inputFilePath = "input.txt"

func computeXXX(content []byte) int {
	return 0
}

func main() {
	log.Println("Beginning day11ex01...")

	contents, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		log.Fatal(fmt.Sprint("Unable to read input file:", err))
	}

	XXX := computeXXX(contents)

	log.Println("XXX successfully computed")
	log.Printf("XXX: %d\n", XXX)
}
