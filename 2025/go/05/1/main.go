package main

import (
	"strconv"
	"strings"
)

type Ingridients struct {
	Ranges    [][]int
	Available []string
}

func Parse(in []byte) Ingridients {
	str := strings.TrimSpace(string(in))
	parts := strings.Split(str, "\n\n")

	out := [][]int{}
	for _, rs := range strings.Split(parts[0], "\n") {
		r := strings.Split(rs, "-")
		start, _ := strconv.Atoi(r[0])
		end, _ := strconv.Atoi(r[1])
		out = append(out, []int{start, end})
	}
	available := strings.Split(parts[1], "\n")

	return Ingridients{out, available}
}

func Puzzle(in Ingridients) int {
	sum := 0
	nums := map[int]int{}
	for _, av := range in.Available {
		num, _ := strconv.Atoi(av)
		for _, r := range in.Ranges {
			if num >= r[0] && num <= r[1] {
				if _, ok := nums[num]; !ok {
					nums[num] = 1
					sum++
				}
			}
		}

	}
	return sum
}
