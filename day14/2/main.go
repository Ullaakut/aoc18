package main

import (
	"fmt"
	"strconv"
	"strings"
)

const input = 768071

func main() {
	scores := []byte{'3', '7'}
	firstElf, secondElf := 0, 1

	for len(scores) < 50000000 {
		score := []byte(strconv.Itoa(int(scores[firstElf] - '0' + scores[secondElf] - '0')))
		scores = append(scores, score...)

		firstElf = (firstElf + 1 + int(scores[firstElf]-'0')) % len(scores)
		secondElf = (secondElf + 1 + int(scores[secondElf]-'0')) % len(scores)
	}

	fmt.Println(strings.Index(string(scores), strconv.Itoa(input)))
}
