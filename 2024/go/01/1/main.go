package main

import (
	"sort"
	"strconv"
	"strings"
)

func Parse(input string) ([]int, []int) {
	lines := strings.Split(input, "\n")
	left, right := []int{}, []int{}
	for _, line := range lines {
		parts := strings.Split(line, "   ")

		leftNum, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		left = append(left, leftNum)

		rightNum, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		right = append(right, rightNum)
	}
	return left, right
}

func Day1(left, right []int) int {
	sort.Slice(left, func(i, j int) bool {
		return left[i] > left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] > right[j]
	})

	sum := 0
	for i := range left {
		sum += max(left[i], right[i]) - min(left[i], right[i])
	}
	return sum
}
