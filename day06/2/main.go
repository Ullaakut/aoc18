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

	safeness int
}

func calcDistance(a, b coord) int {
	distance := math.Abs(float64(b.x)-float64(a.x)) + math.Abs(float64(b.y)-float64(a.y))

	return int(distance)
}

func computeSafeRegionFromCoordinates(coordinates []coord, safeLimit int) []coord {
	safeZone := []coord{}
	points := []coord{}

	for x := -1000; x <= 1000; x++ {
		for y := -1000; y <= 1000; y++ {
			point := coord{x: x, y: y}

			for _, coordinate := range coordinates {
				distance := calcDistance(point, coordinate)
				point.safeness += distance
			}

			points = append(points, point)
		}
	}

	for _, point := range points {
		if point.safeness < safeLimit {
			safeZone = append(safeZone, point)
		}
	}

	return safeZone
}

func computeSafeZone(coordinateLines [][]byte, safeLimit int) []coord {
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

	safeZone := computeSafeRegionFromCoordinates(coordinates, safeLimit)

	return safeZone
}

func solveExercise(inputPath string, safeLimit int) int {
	contents, err := ioutil.ReadFile(inputPath)
	if err != nil {
		log.Fatal(fmt.Sprint("Unable to read input file:", err))
	}

	lines := bytes.Split(contents, []byte("\n"))

	safeZone := computeSafeZone(lines, safeLimit)

	return len(safeZone)
}

func main() {
	log.Println("Beginning day06ex02...")

	safeZone := solveExercise(inputFilePath, 10000)

	log.Println("SafeZone successfully computed")
	log.Printf("SafeZone size: %d\n", safeZone)
}
