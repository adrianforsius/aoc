package main

import (
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

func Day1(stones []int, blinks int) int {
	// for i := 0; i < 6; i++ {
	for i := 0; i < 25; i++ {
		stones = NewStones(stones)
		// fmt.Println(stones)
	}

	return len(stones)
}

func NewStones(stones []int) []int {
	out := []int{}
	for _, s := range stones {
		if s == 0 {
			out = append(out, 1)
			continue
		}
		if len(strconv.Itoa(s))%2 == 0 {
			sStr := strconv.Itoa(s)
			first, last := sStr[:len(sStr)/2], sStr[len(sStr)/2:]
			firstNum, lastNum := 0, 0
			for i, f := range first {
				if string(f) != "0" {
					firstNum, _ = strconv.Atoi(first[i:])
					// fmt.Println("first", firstNum)
					break
				}
			}
			for i, f := range last {
				if string(f) != "0" {
					lastNum, _ = strconv.Atoi(last[i:])
					// fmt.Println("last", lastNum)
					break
				}
			}
			out = append(out, firstNum)
			out = append(out, lastNum)
			continue
		}
		out = append(out, s*2024)
	}
	return out
}
