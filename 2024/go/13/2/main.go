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

func Day2(line []int) int {
	blocks := toString(line)

	for i := len(blocks) - 1; i > 0; i-- {
		if blocks[i] == "." {
			continue
		}

		char := blocks[i]
		fileFrom := i
		fileTo := i
		for fileTo > 0 && blocks[fileTo] == char {
			fileTo--
		}
		i = fileTo + 1
		fileSize := fileFrom - fileTo

		blockFrom := slices.IndexFunc(blocks, func(s string) bool {
			return s != "."
		})
		blockTo := blockFrom
		for k := blockFrom; k < len(blocks)-1; k++ {
			if blocks[k] == "." {
				blockTo++
			} else {
				blockFrom = k
				blockTo = k
			}
			if blockTo-blockFrom == fileSize {
				break
			}
		}
		if fileTo < blockTo {
			continue
		}

		if blockTo-blockFrom == fileSize {
			for f := 0; f < fileSize; f++ {
				blocks[blockFrom+1] = blocks[fileTo+1]
				blocks[fileTo+1] = "."
				blockFrom++
				fileTo++
			}
		}
	}

	sum := 0
	for i, v := range blocks {
		val, _ := strconv.Atoi(v)
		sum += i * val
	}

	return sum
}
