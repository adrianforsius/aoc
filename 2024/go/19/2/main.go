package main

import (
	"fmt"
	"strings"
)

type Design struct {
	patterns  map[string]string
	display   []string
	maxLength int
}

func Parse(in string) Design {
	parts := strings.Split(in, "\n\n")
	patterns := map[string]string{}
	for _, s := range strings.Split(parts[0], ", ") {
		patterns[s] = s
	}

	display := strings.Split(parts[1], "\n")
	maxLength := len(display[0])
	for _, d := range display {
		maxLength = max(maxLength, len(d))
	}

	return Design{patterns, display, maxLength}
}

func Day2(d Design) int {
	sum := 0
	cache := map[string]int{}

	var search func(string) int
	search = func(s string) int {
		if len(s) == 0 {
			return 1
		}
		if v, ok := cache[s]; ok {
			return v
		}

		out := 0
		for _, pattern := range d.patterns {
			if strings.HasPrefix(s, pattern) {
				fmt.Println(s, len(s), len(pattern), pattern, "next", len(s[len(pattern):]))
				temp := search(s[len(pattern):])
				cache[s[len(pattern):]] = temp
				out += temp
			}
		}
		return out
	}

	for _, design := range d.display {
		sum += search(design)
	}

	return sum
}
