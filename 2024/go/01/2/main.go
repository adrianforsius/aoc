package main

import (
	"fmt"
	"strconv"
	"strings"
)

const input = `
`

func main() {
	lines := strings.Split("\n", input)

	left, right := []int{}, []int{}
	for _, line := range lines {
		parts := strings.Split(" ", line)
		leftNum, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		left = append(left, leftNum)

		rightNum, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		right = append(right, rightNum)
	}

	out := Day2(left, right)
	fmt.Println("out", out)
}

func Day2(left, right []int) int {
	return 0
}
