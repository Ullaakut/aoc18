package main

import (
	"log"
)

type coord struct {
	x int
	y int
}

type cells [301][301]int

func bestSubgrid(grid cells) (int, int, int) {
	bestX, bestY, bestSize, bestTotalPower := 0, 0, 0, 0

	for s := 1; s <= 300; s++ {
		for x := s; x <= 300; x++ {
			for y := s; y <= 300; y++ {
				total := grid[x][y] - grid[x-s][y] - grid[x][y-s] + grid[x-s][y-s]

				if total > bestTotalPower {
					bestTotalPower = total
					bestX = x
					bestY = y
					bestSize = s
				}
			}
		}
	}

	return bestX - bestSize + 1, bestY - bestSize + 1, bestSize
}

func solveExercise(serialNumber int) (int, int, int) {
	grid := cells{}

	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			power := (((((x+10)*y)+serialNumber)*(x+10)/100)%10 - 5)
			grid[x][y] = power + grid[x-1][y] + grid[x][y-1] - grid[x-1][y-1]
		}
	}

	return bestSubgrid(grid)
}

func main() {
	log.Println("Beginning day11ex02...")

	x, y, size := solveExercise(9306)

	log.Printf("Cell grid computed")
	log.Printf("Best cell grid: %d,%d,%d", x, y, size)
}
