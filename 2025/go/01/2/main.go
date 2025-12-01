package main

import (
	"math"
	"strconv"
	"strings"
)

func Parse(in []byte) [][]string {
	lines := strings.Split(string(in), "\n")
	lines = lines[:len(lines)-1]
	out := [][]string{}
	for _, line := range lines {
		out = append(out, []string{string(line[0]), string(line[1:])})
	}
	return out
}

func Puzzle(in [][]string) int {
	sum := 0
	dial := 50
	for _, i := range in {
		val, _ := strconv.Atoi(i[1])
		before := dial

		if i[0] == "L" {
			dial -= val
			dial = dial % 100
			if dial < 0 {
				dial += 100
			}
			if dial > before && before != 0 || dial == 0 {
				sum++
			}
		} else {
			dial += val
			dial = dial % 100
			if dial < before {
				sum++
			}
		}
		sum += int(math.Trunc(float64(val) / 100))
	}
	return sum
}
