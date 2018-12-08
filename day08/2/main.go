package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

const inputFilePath = "input.txt"

var metadata = regexp.MustCompile("[0-9]+")

type node struct {
	nChildren int
	nMetadata int

	value int

	values []int
	meta   []int
}

func computeNode(meta []int) (int, []int) {
	node := node{
		nChildren: meta[0],
		nMetadata: meta[1],
		meta:      meta[2:],
	}

	// Iterate over each child and store their values in a slice
	for i := 0; i < node.nChildren; i++ {
		value, remaining := computeNode(node.meta)

		node.meta = remaining
		node.values = append(node.values, value)
	}

	// If no children, value is the node's metadata
	if node.nChildren == 0 {
		for _, value := range node.meta[:node.nMetadata] {
			node.value += value
		}

		return node.value, node.meta[node.nMetadata:]
	}

	// If children, add to value only from indexes in metadata
	for _, val := range node.meta[:node.nMetadata] {
		if val == 0 {
			continue
		}

		if val > len(node.values) {
			continue
		}

		// Add value from child pointed to by metadata
		node.value += node.values[val-1]
	}

	return node.value, node.meta[node.nMetadata:]
}

func computeMetadata(meta []int) int {
	value, _ := computeNode(meta)

	return value
}

func solveExercise(filePath string) int {
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(fmt.Sprint("Unable to read input file:", err))
	}

	lines := bytes.Split(contents, []byte("\n"))

	var meta []int
	for _, line := range lines {
		matches := metadata.FindAllSubmatch(line, -1)
		if len(matches) == 0 {
			continue
		}
		for _, match := range matches {
			data, err := strconv.ParseInt(string(match[0]), 10, 64)
			if err != nil {
				log.Printf("could not parse data: %v", err)
			}

			meta = append(meta, int(data))
		}
	}

	return computeMetadata(meta)
}

func main() {
	log.Println("Beginning day08ex02...")

	meta := solveExercise(inputFilePath)
	log.Printf("Metadata entries computed")
	log.Printf("Sum of metadata entries: %d", meta)
}
