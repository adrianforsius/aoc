package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Parse(in []byte) [][]int {
	ranges := strings.Split(string(in), ",")
	out := [][]int{}
	for _, rs := range ranges {
		parts := strings.Split(rs, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		out = append(out, []int{start, end})
	}
	return out
}

func Puzzle(in [][]int) int {
	sum := 0
	ids := map[string]int{}
	for _, set := range in {
		for i := set[0]; i <= set[1]; i++ {
			val := fmt.Sprint(i)
			if len(val)%2 == 1 {
				continue
			}
			if _, ok := ids[val]; ok {
				continue
			}

			if repeat(val) {
				ids[val] = 1
				sum += i
				fmt.Println("invalid", i, "sum", sum)
			}
		}
	}
	return sum
}

func repeat(digits string) bool {
	for s, e := 0, len(digits)/2; s < len(digits)/2; s, e = s+1, e+1 {
		if digits[s] != digits[e] {
			return false
		}
	}
	return true
}
