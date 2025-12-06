package main

import (
	"strconv"
	"strings"
)

type Index struct {
	index []int
	sign  string
}

func Parse(in []byte) []string {
	str := strings.TrimSpace(string(in))
	return strings.Split(str, "\n")
}

func calc(vals []string, sign string) int {
	sum := 0
	for _, val := range vals {
		v, _ := strconv.Atoi(val)
		if sign == "+" {
			sum += v
		}
		if sign == "*" {
			if sum == 0 {
				sum += v
			} else {
				sum *= v
			}
		}
	}
	return sum
}

func Puzzle(lines []string) int {
	ind := []Index{}

	start := 0
	sign := string(lines[len(lines)-1][0])
	for count, char := range lines[len(lines)-1][1:] {
		if string(char) != " " {
			ind = append(ind, Index{[]int{start, count}, sign})
			start = count
			sign = string(char)
		}
	}
	ind = append(ind, Index{[]int{ind[len(ind)-1].index[1], -1}, sign})

	sum := 0
	for _, ind := range ind {
		groups := []string{}
		for e := 0; e < len(lines)-1; e++ {
			start, end := ind.index[0], ind.index[1]
			var group string
			if end == -1 {
				group = lines[e][start:]
			} else {
				group = lines[e][start:end]
			}
			groups = append(groups, group)
		}
		nums := []string{}
		for i := 0; i < len(groups[0]); i++ {
			num := ""
			for e := 0; e < len(groups); e++ {
				val := string(groups[e][i])
				if val != " " {
					num += val
				}
			}
			nums = append(nums, num)
		}
		sum += calc(nums, ind.sign)

	}
	return sum
}
