package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Parse(in string) []int {
	line := []int{}
	for _, char := range strings.Split(in, "") {
		num, _ := strconv.Atoi(char)
		line = append(line, num)
	}
	return line
}

func toString(line []int) []string {
	count := 0
	display := []string{}
	file := true
	for _, num := range line {
		for i := 0; i < num; i++ {
			if file {
				display = append(display, fmt.Sprint(count))
			} else {
				display = append(display, ".")
			}
		}
		if file {
			count++
		}
		file = !file
	}
	return display
}

func Day1(line []int) int {
	blocks := toString(line)
	// fmt.Println(blocks)

	first := slices.Index(blocks, ".")
	for i := len(blocks) - 1; i > first; i-- {
		if blocks[i] != "." {
			blocks[first] = blocks[i]
			blocks[i] = "."
		}
		first = slices.Index(blocks, ".")
		// fmt.Println(blocks)
	}

	sum := 0
	for i, v := range blocks {
		val, _ := strconv.Atoi(v)
		sum += i * val
	}

	return sum
}
