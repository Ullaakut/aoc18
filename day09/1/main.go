package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"log"
	"os"
)

type playerMap map[int]int

const inputFilePath = "input.txt"

func solveExercise(nPlayers, lastMarbleScore int) int {
	playerIdx := 1

	players := make(playerMap)
	for i := 1; i < nPlayers; i++ {
		players[i] = 0
	}

	circle := ring.New(1)
	circle.Value = 0

	for i := 1; i <= lastMarbleScore; i++ {
		if i%23 == 0 {
			// Rotate 8 times counter-clockwise
			circle = circle.Move(-8)
			// Pop the next marble
			popped := circle.Unlink(1)
			players[playerIdx] += i + popped.Value.(int)
			// Rotate counter-clockwise 1 time
			circle = circle.Move(1)
		} else {
			// Rotate counter-clockwise 1 time
			circle = circle.Move(1)

			// Add new element to ring
			s := ring.New(1)
			s.Value = i
			circle.Link(s)

			// Rotate counter-clockwise 1 time
			circle = circle.Move(1)
		}

		if playerIdx < nPlayers {
			playerIdx++
		} else {
			playerIdx = 1
		}
	}

	highScore := 0
	for _, score := range players {
		if score > highScore {
			highScore = score
		}
	}

	return highScore
}

func main() {
	log.Println("Beginning day09ex01...")

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Printf("Unable to open input file: %v", err)
		os.Exit(1)
	}

	players, lms := 0, 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		_, err := fmt.Sscanf(scanner.Text(), "%d players; last marble is worth %d points", &players, &lms)
		if err != nil {
			log.Printf("Unable to parse line: %v", err)
			os.Exit(1)
		}
	}

	highScore := solveExercise(players, lms)
	log.Printf("Game computed")
	log.Printf("High score: %d", highScore)
}
