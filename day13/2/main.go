package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

// const inputFilePath = "test.txt"
const inputFilePath = "input.txt"

const (
	intersection rune = '+'
	horizontal   rune = '-'
	vertical     rune = '|'
	curve1       rune = '/'
	curve2       rune = '\\'

	right rune = '>'
	left  rune = '<'
	top   rune = '^'
	down  rune = 'v'
)

type coord struct {
	x int
	y int
}

type tracks map[coord]rune

var turnOrder = []rune{left, top, right}

type cart struct {
	id        int
	position  coord
	direction rune
	crashed   bool

	turnIdx int
}

func buildTracks(lines [][]byte) tracks {
	t := make(tracks)

	for y, line := range lines {
		for x, element := range string(line) {
			if element != ' ' {
				if element == top || element == down {
					t[coord{x, y}] = vertical
				} else if element == right || element == left {
					t[coord{x, y}] = horizontal
				} else {
					t[coord{x, y}] = element
				}
			}
		}
	}

	return t
}

func buildCarts(lines [][]byte) []cart {
	carts := []cart{}
	index := 0

	for y, line := range lines {
		for x, direction := range string(line) {
			if direction == down || direction == top || direction == left || direction == right {
				carts = append(carts, cart{
					id:        index,
					position:  coord{x, y},
					direction: direction,
					turnIdx:   0,
				})
				index++
			}
		}
	}

	return carts
}

func computeCrashes(carts []cart) []coord {
	crashes := []coord{}

	for idx1 := range carts {
		if carts[idx1].crashed {
			continue
		}
		for idx2 := range carts {
			if carts[idx2].crashed {
				continue
			}

			if idx1 != idx2 && carts[idx1].position == carts[idx2].position {
				crashPos := carts[idx1].position
				crashes = append(crashes, crashPos)
			}
		}
	}

	return crashes
}

const (
	sizeX = 150
	sizeY = 150
)

func printTracks(t tracks, c []cart) {

	grid := [sizeX][sizeY]rune{}

	for coord, element := range t {
		grid[coord.x][coord.y] = element
	}

	for _, cart := range c {
		if cart.crashed {
			continue
		}

		x := cart.position.x
		y := cart.position.y
		if grid[x][y] != vertical &&
			grid[x][y] != horizontal &&
			grid[x][y] != intersection &&
			grid[x][y] != curve1 &&
			grid[x][y] != curve2 {
			grid[x][y] = 'X'
		} else {
			// grid[x][y] = cart.direction
			grid[x][y] = rune(cart.id + 65)
		}
	}

	f := bufio.NewWriter(os.Stdout)

	lines := ""
	for y := 0; y < sizeY; y++ {
		line := ""
		for x := 0; x < sizeX; x++ {
			if grid[x][y] != rune(0) {
				// if grid[x][y] == top || grid[x][y] == down || grid[x][y] == right || grid[x][y] == left {
				if grid[x][y] >= 65 && grid[x][y] <= 77 {
					line += fmt.Sprintf("%s%s%c%s", "\033[32m", "\033[1m", grid[x][y], "\033[0m")
				} else if grid[x][y] == 'X' {
					line += fmt.Sprintf("%s%s%c%s", "\033[91m", "\033[1m", grid[x][y], "\033[0m")
				} else {
					line += string(grid[x][y])
				}
			} else {
				line += " "
			}
		}
		lines = fmt.Sprintf("%s%s\n", lines, line)
	}

	f.Write([]byte(lines))
	f.Flush()
}

func remainingCarts(c []cart) (bool, coord) {
	alive := 0
	alivePos := coord{}

	for _, cart := range c {
		if !cart.crashed {
			alive++
			alivePos = cart.position
		}
	}

	if alive == 1 {
		return true, alivePos
	}

	return false, coord{}
}

func computeCollision(t tracks, c []cart) coord {
	tick := 0

	for {
		last, lastCoords := remainingCarts(c)
		if last {
			return lastCoords
		}

		// Arrange order of carts based on their current position
		sort.Slice(c, func(a, b int) bool {
			if c[a].position.y == c[b].position.y {
				return c[a].position.x < c[b].position.x
			}
			return c[a].position.y < c[b].position.y
		})

		for idx := range c {
			var nextElement rune

			// Move
			if c[idx].direction == down {
				c[idx].position.y++
				nextElement = t[c[idx].position]
			} else if c[idx].direction == top {
				c[idx].position.y--
				nextElement = t[c[idx].position]
			} else if c[idx].direction == right {
				c[idx].position.x++
				nextElement = t[c[idx].position]
			} else if c[idx].direction == left {
				c[idx].position.x--
				nextElement = t[c[idx].position]
			}

			// Update direction if needed
			if nextElement == intersection { // +
				if turnOrder[c[idx].turnIdx] == left {
					if c[idx].direction == left {
						c[idx].direction = down
					} else if c[idx].direction == top {
						c[idx].direction = left
					} else if c[idx].direction == down {
						c[idx].direction = right
					} else if c[idx].direction == right {
						c[idx].direction = top
					}
				} else if turnOrder[c[idx].turnIdx] == right {
					if c[idx].direction == left {
						c[idx].direction = top
					} else if c[idx].direction == top {
						c[idx].direction = right
					} else if c[idx].direction == down {
						c[idx].direction = left
					} else if c[idx].direction == right {
						c[idx].direction = down
					}
				}

				if c[idx].turnIdx == 2 {
					c[idx].turnIdx = 0
				} else {
					c[idx].turnIdx++
				}
			} else if nextElement == curve1 && c[idx].direction == right { // --->/
				c[idx].direction = top
			} else if nextElement == curve1 && c[idx].direction == left { // /<----
				c[idx].direction = down
			} else if nextElement == curve2 && c[idx].direction == left { // \<----
				c[idx].direction = top
			} else if nextElement == curve2 && c[idx].direction == right { // ---->\
				c[idx].direction = down
			} else if nextElement == curve1 && c[idx].direction == top { // /---->
				c[idx].direction = right
			} else if nextElement == curve1 && c[idx].direction == down { // <----/
				c[idx].direction = left
			} else if nextElement == curve2 && c[idx].direction == down { // \---->
				c[idx].direction = right
			} else if nextElement == curve2 && c[idx].direction == top { // <----\
				c[idx].direction = left
			}

			crashes := computeCrashes(c)
			if len(crashes) > 0 {
				log.Printf("%d crashes on %d", len(crashes)/2, tick)

				for _, crash := range crashes {
					for idx := range c {
						if c[idx].crashed {
							continue
						}

						if c[idx].position == crash {
							c[idx].crashed = true
						}
					}
				}
			}
		}

		tick++
	}
}

func solveExercise(filePath string) coord {
	contents, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		log.Fatal(fmt.Sprint("Unable to read input file:", err))
	}

	lines := bytes.Split(contents, []byte("\n"))

	tracks := buildTracks(lines)
	carts := buildCarts(lines)

	return computeCollision(tracks, carts)
}

func main() {
	log.Println("Beginning day13ex02...")

	collision := solveExercise(inputFilePath)
	// solveExercise(inputFilePath)

	log.Println("Successfully computed upcoming cart collision on tracks")
	log.Printf("Position: %d, %d", collision.x, collision.y)
}
