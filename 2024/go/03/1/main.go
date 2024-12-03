package main

import (
	"regexp"
	"strconv"
	"strings"
)

func Parse(in string) string {
	return in
}

func Day1(in string) int {
	r := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	muls := r.FindAllString(in, -1)

	score := 0
	for _, m := range muls {
		ss := strings.Split(m, "mul(")
		ss = strings.Split(ss[1], ")")
		nums := strings.Split(ss[0], ",")
		mul1, _ := strconv.Atoi(nums[0])
		mul2, _ := strconv.Atoi(nums[1])
		score += mul1 * mul2
	}
	return score
}
