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

	total int

	meta []int
}

func computeNode(meta []int) (int, []int) {
	node := node{
		nChildren: meta[0],
		nMetadata: meta[1],
		meta:      meta[2:],
	}

	// Iterate over children
	for i := 0; i < node.nChildren; i++ {
		total, remaining := computeNode(node.meta)

		node.total += total
		node.meta = remaining
	}

	// Iterate over node's metadata
	for _, val := range node.meta[:node.nMetadata] {
		node.total += val
	}

	return node.total, node.meta[node.nMetadata:]
}

func computeMetadata(meta []int) int {
	total, _ := computeNode(meta)

	return total
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
	log.Println("Beginning day08ex01...")

	meta := solveExercise(inputFilePath)
	log.Printf("Metadata entries computed")
	log.Printf("Sum of metadata entries: %d", meta)
}
