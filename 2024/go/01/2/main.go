package main

import (
	"strconv"
	"strings"
)

func Parse(input string) (map[int]int, map[int]int) {
	lines := strings.Split(input, "\n")
	left, right := make(map[int]int, 0), make(map[int]int, 0)
	for _, line := range lines {
		parts := strings.Split(line, "   ")

		leftNum, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		if _, ok := left[leftNum]; !ok {
			left[leftNum] = 1
		} else {
			left[leftNum] += 1
		}

		rightNum, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		if _, ok := right[rightNum]; !ok {
			right[rightNum] = 1
		} else {
			right[rightNum] += 1
		}
	}
	return left, right
}

func Day2(left, right map[int]int) int {
	sum := 0
	for key, val := range left {
		if rightVal, ok := right[key]; ok {
			sum += key * val * rightVal
		}
	}

	return sum
}
