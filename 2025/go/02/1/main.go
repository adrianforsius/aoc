package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Parse(in []byte) [][]int {

	str := strings.TrimSpace(string(in))
	ranges := strings.Split(str, ",")
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
	// ids := map[string]int{}
	for _, set := range in {
		// fmt.Println("new section", set[0], set[1])
		for i := set[0]; i <= set[1]; i++ {
			val := fmt.Sprint(i)
			if len(val)%2 == 1 {
				continue
			}
			// if _, ok := ids[val]; ok {
			// 	continue
			// }

			if val[:len(val)/2] == val[len(val)/2:] {
				// ids[val] = 1
				sum += i
				// fmt.Println("invalid", i, "sum", sum)
			}
		}
	}
	return sum
}
