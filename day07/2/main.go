package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"text/tabwriter"
)

type instructions map[byte][]byte

var stepRegex = regexp.MustCompile("^Step ([A-Z]+) must be finished before step ([A-Z]+) can begin.$")

const inputFilePath = "input.txt"

func computeInstructionsOrder(originalInstructions instructions) []byte {
	// Copy the original instructions in a new map to avoid modifying them
	instr := make(instructions)
	for k, v := range originalInstructions {
		instr[k] = v
	}

	done := []byte{}

	for {
		notNext := []byte{}

		if len(instr) == 1 {
			for id, prev := range instr {
				sort.Slice(prev, func(i, j int) bool {
					return prev[i] < prev[j]
				})
				res := append(done, id)
				res = append(res, prev...)

				return res
			}
		}

		for _, previous := range instr {
			for _, prevID := range previous {
				notNext = append(notNext, prevID)
			}
		}

		potentialNext := []byte{}
		for id := range instr {
			if !bytes.Contains(notNext, []byte{id}) {
				potentialNext = append(potentialNext, id)
			}
		}

		sort.Slice(potentialNext, func(i, j int) bool {
			return potentialNext[i] < potentialNext[j]
		})

		done = append(done, potentialNext[0])
		delete(instr, potentialNext[0])
	}
}

type worker struct {
	task     byte
	working  bool
	timeLeft int
}

func (w *worker) getTask() byte {
	if w.working {
		return w.task
	}

	// the ascii equivalent of a dot
	return 46
}

// Checks if a task is available for the taking
func available(task byte, instructions instructions, done, inProgress []byte) bool {
	if bytes.Contains(inProgress, []byte{task}) {
		// log.Printf("Ignoring task %q since it's in progress", task)
		return false
	}

	if bytes.Contains(done, []byte{task}) {
		// log.Printf("Ignoring task %q since it's been done", task)
		return false
	}

	for instruction, prev := range instructions {
		if task == instruction {
			continue
		}

		// log.Printf("Checking if task %q is blocked by instruction %q", task, instruction)
		// If the instruction is in progress, we must wait for it to be complete
		// before we can pick up its subsequent tasks
		if bytes.Contains(inProgress, []byte{instruction}) && bytes.Contains(prev, []byte{task}) {
			// log.Printf("Blocking task %q since it's a subtask of %q which is in progress", task, instruction)
			return false
		}

		if !bytes.Contains(done, []byte{instruction}) && bytes.Contains(prev, []byte{task}) {
			// log.Printf("Blocking task %q since it's a subtask of %q which is not done yet", task, instruction)
			return false
		}
	}

	// log.Printf("Task %q is available!", task)
	return true
}

func computeWorkTime(instructions instructions, order []byte, stepTime, numWorkers int) int {
	totalTime := 0
	done := []byte{}
	inProgress := []byte{}
	workers := []worker{}
	header := "Second\t"
	formatString := "%d\t"

	for i := 0; i < numWorkers; i++ {
		workers = append(workers, worker{})
		formatString += "%c\t"
		header += fmt.Sprintf("Worker %d\t", i+1)
	}

	formatString += "%s\n"
	header += "Done\n"

	w := tabwriter.NewWriter(os.Stdout, 10, 0, 3, ' ', 0)
	fmt.Fprintln(w, header)

	for {
		for id := range workers {
			if workers[id].working {
				workers[id].timeLeft--

				if workers[id].timeLeft == 0 {
					inProgress = bytes.Replace(inProgress, []byte{workers[id].getTask()}, []byte{}, -1)
					done = append(done, workers[id].getTask())
					workers[id].working = false
				}
			} else {
				// log.Printf("Worker %d looks for a task!", id+1)
			}

			// This worker is still working and does not need a new task
			if workers[id].working {
				// log.Printf("Worker %d works on %q", id+1, workers[id].getTask())
				continue
			}

			// This worker needs a new task! He will cycle through the remaining tasks,
			// looking for one that isn't blocked nor in progress
			for _, task := range order {
				if available(task, instructions, done, inProgress) {
					workers[id].task = task
					workers[id].timeLeft = stepTime + int(task) - 64
					workers[id].working = true
					inProgress = append(inProgress, task)
					// log.Printf("Worker %d takes task %q", id+1, task)
					break
				}
			}

			if !workers[id].working {
				// log.Printf("Worker %d waits for a task to be available", id+1)
			}
		}

		table := []interface{}{totalTime}
		for _, worker := range workers {
			table = append(table, worker.getTask())
		}
		table = append(table, done)

		fmt.Fprintf(w, formatString, table...)
		w.Flush()

		if len(done) == len(order) {
			return totalTime
		}

		totalTime++
	}
}

func solveExercise(filePath string, stepTime, workers int) int {
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(fmt.Sprint("Unable to read input file:", err))
	}

	lines := bytes.Split(contents, []byte("\n"))

	instructions := make(instructions)
	for _, line := range lines {
		matches := stepRegex.FindAllSubmatch(line, -1)
		if len(matches) == 0 || len(matches[0]) != 3 {
			continue
		}

		ID := matches[0][1][0]
		prev := matches[0][2][0]

		instructions[ID] = append(instructions[ID], prev)
	}

	order := computeInstructionsOrder(instructions)

	return computeWorkTime(instructions, order, stepTime, workers)
}

func main() {
	log.Println("Beginning day07ex02...")

	solveExercise(inputFilePath, 60, 5)
}
