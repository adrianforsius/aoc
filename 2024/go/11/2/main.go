package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Parse(in string) []int {
	line := []int{}
	for _, char := range strings.Split(in, " ") {
		num, _ := strconv.Atoi(char)
		line = append(line, num)
	}
	return line
}

func Day2(stones []int, blinks int) int {
	vals := map[int]int{}
	for _, s := range stones {
		vals[s]++
	}

	for i := 0; i < blinks; i++ {
		vals = NewStones(vals)
	}

	out := 0
	for _, count := range vals {
		out += count
	}

	return out
}

var lookup = map[int][]int{
	0: {1},
	1: {2024},
}

func NewStones(c map[int]int) map[int]int {
	cache := map[int]int{}
	for i, v := range c {
		if next, ok := lookup[i]; ok {
			for _, n := range next {
				cache[n] += v
			}
			continue
		}

		if len(strconv.Itoa(i))%2 == 0 {
			sStr := strconv.Itoa(i)
			first, last := sStr[:len(sStr)/2], sStr[len(sStr)/2:]
			firstNum, lastNum := 0, 0
			for i, f := range first {
				if string(f) != "0" {
					firstNum, _ = strconv.Atoi(first[i:])
					break
				}
			}
			for i, f := range last {
				if string(f) != "0" {
					lastNum, _ = strconv.Atoi(last[i:])
					break
				}
			}
			cache[firstNum] += v
			cache[lastNum] += v
			lookup[i] = []int{firstNum, lastNum}
			continue
		}

		val := i * 2024
		cache[val] += v
		lookup[i] = []int{val}
	}
	fmt.Println("cache", cache)

	return cache
}
