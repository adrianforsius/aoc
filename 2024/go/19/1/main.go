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

type Queue [][]int

func (q *Queue) Peek() []int {
	qq := *q
	if len(*q) > 0 {
		return qq[len(qq)-1]
	}
	return []int{}
}

func (q *Queue) Append(pair []int) {
	*q = append(*q, pair)
}

func (q *Queue) Next() ([]int, bool) {
	qq := *q
	if len(*q) > 1 {
		val := qq[len(qq)-1]
		*q = qq[0 : len(qq)-1]
		return val, true
	}
	if len(*q) == 1 {
		val := qq[len(qq)-1]
		*q = [][]int{}
		return val, true
	}
	return nil, false
}

func Day1(d Design) int {
	sum := 0
	for _, design := range d.display {
		queue := &Queue{}
		queue.Append([]int{0, min(d.maxLength, len(design))})
		for {
			next, ok := queue.Next()
			if !ok {
				break
			}

			for {
				changed := false
				for i := next[1]; i > next[0]; i-- {
					fmt.Println("next", next, "i", i, "part", design[next[0]:i])
					if _, ok := d.patterns[design[next[0]:i]]; ok {
						queue.Append([]int{i, min(next[0]+d.maxLength, len(design))})
						changed = true
						break
					}
				}
				if !changed {
					next[1]--
				}
				if next[1] == next[0] {
					break
				}
			}

			q := *queue
			if len(q) == 0 {
				fmt.Println("empty quueu")
				break
			}

			if queue.Peek()[0] == len(design) {
				fmt.Println("found")
				sum++
				break
			}
		}
	}

	return sum
}
