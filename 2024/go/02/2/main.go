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

func isAsc(curr, next int) bool {
	return next > curr && next < curr+4
}

func isDesc(curr, next int) bool {
	return next < curr && next > curr-4
}

func asc(levels []int) int {
	points := 0
	for i, j := 0, 1; j < len(levels); i, j = i+1, j+1 {
		if isAsc(levels[i], levels[j]) {
			points++
		}
	}
	return points
}

func desc(levels []int) int {
	points := 0
	for i, j := 0, 1; j < len(levels); i, j = i+1, j+1 {
		if isDesc(levels[i], levels[j]) {
			points++
		}
	}
	return points
}

func Day2(reportCard [][]int) int {
	safe := 0
	for _, levels := range reportCard {
		for i := 0; i < len(levels); i++ {
			ss := make([]int, len(levels))
			copy(ss, levels)
			s := remove(ss, i)
			ascPoints := asc(s)
			descPoints := desc(s)
			if ascPoints >= len(s)-1 || descPoints >= len(s)-1 {
				safe++
				break
			}
			// fmt.Println("break", s, ascPoints, descPoints)
		}
	}

	return safe
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
