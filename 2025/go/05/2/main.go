package main

import (
	"fmt"
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
	fmt.Println(len(out))

	return Ingridients{out, available}
}

func Puzzle(in Ingridients) int {
	nums := in.Ranges

	out := [][]int{}
	for len(nums) > 1 {
		num, temp := nums[0], nums[1:]
		nums = temp

		hit := false
		startP := num[0]
		endP := num[1]
		for i, next := range nums {

			startN := next[0]
			endN := next[1]

			if startP <= startN && startN <= endP || startP <= endN && endN <= endP || startN <= startP && startP <= endN || startN <= endP && endP <= endN {
				nums = append(nums[:i], nums[i+1:]...)
				nums = append(nums, []int{min(startN, startP), max(endN, endP)})
				hit = true
				break
			}
		}

		if hit == false {
			out = append(out, num)
		}
	}
	out = append(out, nums[0])

	sum := 0
	for _, num := range out {
		sum += num[1] + 1 - num[0]
	}
	return sum
}
