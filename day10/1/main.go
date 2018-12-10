package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

const inputFilePath = "input.txt"

type star struct {
	x int
	y int

	vx int
	vy int
}

type coord struct {
	x int
	y int
}

func printSentence(stars []star, min, max coord) string {
	galaxy := make(map[coord]bool)

	for _, star := range stars {
		galaxy[coord{star.x, star.y}] = true
	}

	sentence := ""

	for y := min.y; y <= max.y; y++ {
		for x := min.x; x <= max.x; x++ {
			if galaxy[coord{x, y}] {
				sentence = fmt.Sprint(sentence, "#")
			} else {
				sentence = fmt.Sprint(sentence, ".")
			}
		}
		sentence = fmt.Sprint(sentence, "\n")
	}

	return sentence
}

func calcCloseStars(stars []star) (coord, coord) {
	xmin, ymin := stars[0].x, stars[0].y
	xmax, ymax := stars[0].x, stars[0].y

	for _, star := range stars {
		if star.x > xmax {
			xmax = star.x
		}
		if star.y > ymax {
			ymax = star.y
		}
		if star.x < xmin {
			xmin = star.x
		}
		if star.y < ymin {
			ymin = star.y
		}
	}

	return coord{xmin, ymin}, coord{xmax, ymax}
}

func aligned(min, max coord) bool {
	// we assume that they must be aligned if they are close enough to each other
	return math.Abs(float64(max.x)-float64(min.x)) <= 65 && math.Abs(float64(max.y)-float64(min.y)) <= 65
}

func findSentence(stars []star) string {
	for i := 0; i < 100000; i++ {
		min, max := calcCloseStars(stars)

		if aligned(min, max) {
			return printSentence(stars, min, max)
		}

		for idx := range stars {
			stars[idx].x += stars[idx].vx
			stars[idx].y += stars[idx].vy
		}
	}

	return "Not found"
}

func solveExercise(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Unable to open input file: %v", err)
		os.Exit(1)
	}

	stars := []star{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		star := star{}
		_, err := fmt.Sscanf(scanner.Text(), "position=<%d, %d> velocity=<%d,  %d>", &star.x, &star.y, &star.vx, &star.vy)
		if err != nil {
			log.Printf("Unable to parse line %q: %v", scanner.Text(), err)
			os.Exit(1)
		}

		stars = append(stars, star)
	}

	return findSentence(stars)
}

func main() {
	log.Println("Beginning day10ex01...")

	sentence := solveExercise(inputFilePath)

	log.Printf("Star velocities computed")
	log.Printf("Sentence found: \n%s", sentence)
}
