package main

import (
	"strconv"
	"strings"
)

func Parse(in []byte) [][]string {
	str := strings.TrimSpace(string(in))
	lines := strings.Split(str, "\n")

	divided := [][]string{}
	out := []string{}
	for _, line := range lines {
		word := []string{}
		for _, char := range line {
			if string(char) == " " {
				if len(word) > 0 {
					out = append(out, strings.Join(word, ""))
				}
				word = []string{}
			} else {
				word = append(word, string(char))
			}
		}
		out = append(out, strings.Join(word, ""))
		word = []string{}
		divided = append(divided, out)
		out = []string{}
	}

	return divided
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

func Puzzle(in [][]string) int {
	sum := 0
	for i := 0; i < len(in[0]); i++ {
		vals := []string{}
		for e := 0; e < len(in)-1; e++ {
			vals = append(vals, in[e][i])
		}
		sum += calc(vals, in[len(in)-1][i])
	}
	return sum
}
