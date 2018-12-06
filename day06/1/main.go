package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"regexp"
	"strconv"
)

type areas map[coord][]coord

const inputFilePath = "input.txt"

var coordRegex = regexp.MustCompile("^([0-9]+), ([0-9]+)$")

type coord struct {
	x int
	y int
}

func calcDistance(a, b coord) int {
	distance := math.Abs(float64(b.x)-float64(a.x)) + math.Abs(float64(b.y)-float64(a.y))

	return int(distance)
}

func computeAreasFromCoordinates(coordinates []coord) areas {
	areas1 := make(areas)
	areas2 := make(areas)

	// Run with low boundaries
	for x := -1000; x <= 1000; x++ {
		for y := -1000; y <= 1000; y++ {
			point := coord{x: x, y: y}

			var closestCoord coord
			closestDistance := -1
			previousClosestDistance := -1

			for _, coordinate := range coordinates {
				distance := calcDistance(coordinate, point)

				if distance <= closestDistance || closestDistance == -1 {
					closestCoord = coordinate
					previousClosestDistance = closestDistance
					closestDistance = distance
				}
			}

			// If equidistant, ignore point
			if closestDistance == previousClosestDistance {
				continue
			}

			// Add point to this area
			areas1[closestCoord] = append(areas1[closestCoord], point)
		}
	}

	// Run with high boundaries
	for x := -1100; x <= 1100; x++ {
		for y := -1100; y <= 1100; y++ {
			point := coord{x: x, y: y}

			var closestCoord coord
			closestDistance := -1
			previousClosestDistance := -1

			for _, coordinate := range coordinates {
				distance := calcDistance(coordinate, point)

				if distance <= closestDistance || closestDistance == -1 {
					closestCoord = coordinate
					previousClosestDistance = closestDistance
					closestDistance = distance

					// log.Printf("Found new closest coordinate at %d,%d close to %d,%d with distance %d", point.x, point.y, coordinate.x, coordinate.y, distance)
				}
			}

			// If equidistant, ignore point
			if closestDistance == previousClosestDistance {
				continue
			}

			// Add point to this area
			areas2[closestCoord] = append(areas2[closestCoord], point)
		}
	}

	areas3 := make(areas)

	// Check difference in size between different computations
	// If area size changed by changing the size of the board
	// it means that the area is infinite.
	for coord, area := range areas1 {
		// If it changed after changing the bounds, then it's infinite and we ignore it
		if len(areas2[coord]) != len(area) {
			log.Printf("Ignoring infinite area for %v: was %d became %d", coord, len(area), len(areas2[coord]))
		} else {
			areas3[coord] = area
			log.Printf("Found finite area for %v of size %d", coord, len(area))
		}
	}

	return areas3
}

func computeAreas(coordinateLines [][]byte) areas {
	coordinates := []coord{}

	for _, coordinateLine := range coordinateLines {
		r := coordRegex.FindAllStringSubmatch(string(coordinateLine), -1)
		if len(r) == 0 || len(r[0]) != 3 {
			continue
		}

		x, err := strconv.ParseInt(r[0][1], 10, 64)
		if err != nil {
			log.Printf("Could not parse coordinate %q: %v", coordinateLine, err)
			continue
		}
		y, err := strconv.ParseInt(r[0][2], 10, 64)
		if err != nil {
			log.Printf("Could not parse coordinate %q: %v", coordinateLine, err)
			continue
		}

		coordinates = append(coordinates, coord{x: int(x), y: int(y)})
	}

	areas := computeAreasFromCoordinates(coordinates)

	return areas
}

func solveExercise(inputPath string) int {
	contents, err := ioutil.ReadFile(inputPath)
	if err != nil {
		log.Fatal(fmt.Sprint("Unable to read input file:", err))
	}

	lines := bytes.Split(contents, []byte("\n"))

	areas := computeAreas(lines)

	biggestArea := 0
	for closest, area := range areas {
		if len(area) > biggestArea {
			biggestArea = len(area)
			log.Printf("Found new biggest area, close to %d,%d of size %d", closest.x, closest.y, biggestArea)
		}
	}

	// Remove 1 from the length, for the point with distance = 0
	return biggestArea
}

func main() {
	log.Println("Beginning day06ex01...")

	biggestArea := solveExercise(inputFilePath)

	log.Println("Biggest area successfully computed")
	log.Printf("Biggest area size: %d\n", biggestArea)
}
