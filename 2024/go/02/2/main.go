package main

import (
	"strconv"
	"strings"
)

func Parse(input string) [][]int {
	lines := strings.Split(input, "\n")
	reportCard := [][]int{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		nextLine := []int{}
		for _, p := range parts {
			num, _ := strconv.Atoi(p)
			nextLine = append(nextLine, num)
		}
		reportCard = append(reportCard, nextLine)
	}
	return reportCard
}

func asc(levels []int) bool {
	lifeline := 1
	for i, j := 0, 1; j < len(levels); i, j = i+1, j+1 {
		curr, next := levels[i], levels[j]
		if next <= curr || next >= curr+4 {
			if j+1 >= len(levels) {
				return false
			}
			curr, next := levels[i], levels[j+1]
			if next <= curr || next >= curr+4 {
				return false
			}
			i = j
			j = j + 1
			lifeline--
		}
		if lifeline < 0 {
			return false
		}
	}

	return true
}

func desc(levels []int) bool {
	lifeline := 1
	for i, j := 0, 1; j < len(levels); i, j = i+1, j+1 {
		curr, next := levels[i], levels[j]
		if next >= curr || next <= curr-4 {
			if j+1 >= len(levels) {
				return false
			}
			curr, next := levels[i], levels[j+1]
			if next >= curr || next <= curr-4 {
				return false
			}
			i = j
			j = j + 1
			lifeline--
			i--
		}
		if lifeline < 0 {
			return false
		}
	}
	return true
}

func Day2(reportCard [][]int) int {
	safe := 0
	for _, levels := range reportCard {
		if asc(levels) || desc(levels) {
			safe++
		}
	}

	return safe
}
