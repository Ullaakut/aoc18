package main

import (
	"fmt"
	"strconv"
)

const input = 768071

func main() {
	scores := []byte{'3', '7'}
	firstElf, secondElf := 0, 1

	for len(scores) < input+10 {
		score := []byte(strconv.Itoa(int(scores[firstElf] - '0' + scores[secondElf] - '0')))
		scores = append(scores, score...)

		firstElf = (firstElf + 1 + int(scores[firstElf]-'0')) % len(scores)
		secondElf = (secondElf + 1 + int(scores[secondElf]-'0')) % len(scores)
	}

	fmt.Println(string(scores[input : input+10]))
}
