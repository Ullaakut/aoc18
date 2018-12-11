package main

import (
	"log"
)

type coord struct {
	x int
	y int
}

type cells map[coord]int

func bestSubgrid(grid cells) (int, int) {
	bestX, bestY, bestTotalPower := 0, 0, 0

	for x := 0; x < 298; x++ {
		for y := 0; y < 298; y++ {
			totalPower := grid[coord{x, y}] + grid[coord{x + 1, y}] + grid[coord{x + 2, y}]
			totalPower += grid[coord{x, y + 1}] + grid[coord{x + 1, y + 1}] + grid[coord{x + 2, y + 1}]
			totalPower += grid[coord{x, y + 2}] + grid[coord{x + 1, y + 2}] + grid[coord{x + 2, y + 2}]

			if totalPower > bestTotalPower {
				bestTotalPower = totalPower
				bestX = x
				bestY = y
			}
		}
	}

	log.Printf("Found best X,Y at %d,%d with a total of %d power", bestX, bestY, bestTotalPower)
	return bestX, bestY
}

func solveExercise(serialNumber int) (int, int) {
	grid := make(cells)

	for x := 0; x < 300; x++ {
		for y := 0; y < 300; y++ {
			grid[coord{x, y}] = (((((x+10)*y)+serialNumber)*(x+10)/100)%10 - 5)
		}
	}

	return bestSubgrid(grid)
}

func main() {
	log.Println("Beginning day11ex01...")

	x, y := solveExercise(9306)

	log.Printf("Cell grid computed")
	log.Printf("Best 3x3 cell grid: %d,%d", x, y)
}
