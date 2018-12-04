package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	inputFilePath = "input.txt"
	timeLayout    = "2006-01-02 15:04"
)

var (
	beginShift  = regexp.MustCompile("^\\[(.*)\\] Guard #([0-9]+) begins shift$")
	wakeUp      = regexp.MustCompile("^\\[(.*)\\] wakes up$")
	fallsAsleep = regexp.MustCompile("^\\[(.*)\\] falls asleep$")
)

type sleepyTime struct {
	dateBegin time.Time
	dateEnd   time.Time
	duration  time.Duration
}

type guardLogs map[string][]*sleepyTime

type guard struct {
	ID string

	sleepyTimes []*sleepyTime
	totalSleep  time.Duration
}

func sortLogs(content []byte) []string {
	lines := strings.Split(string(content), "\n")

	sort.Strings(lines)

	return lines
}

func computeLogs(logs []string) guardLogs {
	var currentGuard string
	var currentSleepyTime *sleepyTime
	guardLogs := make(map[string][]*sleepyTime)

	for _, logLine := range logs {
		// If a guard begins its shift, set currentGuard to its ID and reset current sleepy time.
		// If the previous guard was sleeping, wake him up.
		beginShiftSubmatches := beginShift.FindAllStringSubmatch(logLine, -1)
		if len(beginShiftSubmatches) > 0 && len(beginShiftSubmatches[0]) == 3 {
			if currentSleepyTime != nil {
				endSleep, err := time.Parse(timeLayout, beginShiftSubmatches[0][1])
				if err != nil {
					log.Fatalf("Unable to parse time %q from guard logs: %v", beginShiftSubmatches[0][1], err)
				}

				currentSleepyTime.dateEnd = endSleep
				currentSleepyTime.duration = endSleep.Sub(currentSleepyTime.dateBegin)
				guardLogs[currentGuard] = append(guardLogs[currentGuard], currentSleepyTime)
				currentSleepyTime = nil
			}

			currentGuard = beginShiftSubmatches[0][2]
			currentSleepyTime = nil
		}

		// If a guard falls asleep, begin monitoring its sleep (unless it's already sleeping)
		fallsAsleepSubmatches := fallsAsleep.FindAllStringSubmatch(logLine, -1)
		if len(fallsAsleepSubmatches) > 0 && len(fallsAsleepSubmatches[0]) == 2 {
			if currentSleepyTime == nil {
				beginSleep, err := time.Parse(timeLayout, fallsAsleepSubmatches[0][1])
				if err != nil {
					log.Fatalf("Unable to parse time %q from guard logs: %v", fallsAsleepSubmatches[0][1], err)
				}

				currentSleepyTime = &sleepyTime{
					dateBegin: beginSleep,
				}
			}
		}

		// If a guard wakes up, end its current sleepy time (if it had one) and append it to its logs
		wakeUpSubmatches := wakeUp.FindAllStringSubmatch(logLine, -1)
		if len(wakeUpSubmatches) > 0 && len(wakeUpSubmatches[0]) == 2 {
			if currentSleepyTime != nil {
				endSleep, err := time.Parse(timeLayout, wakeUpSubmatches[0][1])
				if err != nil {
					log.Fatalf("Unable to parse time %q from guard logs: %v", wakeUpSubmatches[0][1], err)
				}

				currentSleepyTime.dateEnd = endSleep
				currentSleepyTime.duration = endSleep.Sub(currentSleepyTime.dateBegin)
				guardLogs[currentGuard] = append(guardLogs[currentGuard], currentSleepyTime)
				currentSleepyTime = nil
			}
		}
	}

	return guardLogs
}

func computeSleepyGuards(logs guardLogs) []guard {
	var guards []guard

	for ID, sleepyTimes := range logs {
		guard := guard{
			ID:          ID,
			sleepyTimes: sleepyTimes,
		}

		for _, sleepyTime := range sleepyTimes {
			guard.totalSleep += sleepyTime.duration
		}

		guards = append(guards, guard)
	}

	return guards
}

func computeSleepiestMinute(guards []guard) int {
	minutes := make(map[int]int)
	for _, guard := range guards {
		for _, sleepyTime := range guard.sleepyTimes {
			for min := sleepyTime.dateBegin.Minute(); min < sleepyTime.dateEnd.Minute(); min++ {
				minutes[min]++
			}
		}
	}

	var (
		sleepiestMinute int
		sleepiestAmount int
	)
	for minute, amount := range minutes {
		if amount > sleepiestAmount {
			sleepiestMinute = minute
			sleepiestAmount = amount
		}
	}

	return sleepiestMinute
}

func computeSleepiestGuardAtMinute(guards []guard, minute int) int {
	guardSleepsAtMinute := make(map[string]int)
	for _, guard := range guards {
		for _, sleepyTime := range guard.sleepyTimes {
			for min := sleepyTime.dateBegin.Minute(); min < sleepyTime.dateEnd.Minute(); min++ {
				if min == minute {
					guardSleepsAtMinute[guard.ID]++
				}
			}
		}
	}

	var (
		sleepiestGuard  string
		sleepiestAmount int
	)
	for guard, amount := range guardSleepsAtMinute {
		if amount > sleepiestAmount {
			sleepiestGuard = guard
			sleepiestAmount = amount
		}
	}

	log.Println("Sleepiest guard found: Elf", sleepiestGuard, "slept at minute", minute, sleepiestAmount, "times!")

	guardID, err := strconv.ParseInt(sleepiestGuard, 10, 64)
	if err != nil {
		log.Fatal(fmt.Sprint("Unable to read guard badge number:", err))
	}

	return int(guardID)
}

func solveExercise(filePath string) int {
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(fmt.Sprint("Unable to read input file:", err))
	}

	guardLogs := computeLogs(sortLogs(contents))
	log.Println("Guard logs successfully computed")

	guards := computeSleepyGuards(guardLogs)
	sleepiestMinute := computeSleepiestMinute(guards)
	log.Println("Sleepiest minute found:", sleepiestMinute)

	sleepiestGuard := computeSleepiestGuardAtMinute(guards, sleepiestMinute)

	return sleepiestGuard * sleepiestMinute
}

func main() {
	log.Println("Beginning day04ex02...")

	checksum := solveExercise(inputFilePath)

	log.Printf("Checksum: %d\n", checksum)
}
