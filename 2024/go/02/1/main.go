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
	for i, j := 0, 1; j < len(levels); i, j = i+1, j+1 {
		curr, next := levels[i], levels[j]
		if next <= curr || next >= curr+4 {
			return false
		}
	}

	return true
}

func desc(levels []int) bool {
	for i, j := 0, 1; j < len(levels); i, j = i+1, j+1 {
		curr, next := levels[i], levels[j]
		if next >= curr || next <= curr-4 {
			return false
		}
	}
	return true
}

func Day1(reportCard [][]int) int {
	safe := 0
	for _, levels := range reportCard {
		if asc(levels) || desc(levels) {
			safe++
		}
	}

	return safe
}
