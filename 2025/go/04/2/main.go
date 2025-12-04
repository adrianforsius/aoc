package main

import (
	"errors"
	"strings"
)

func Parse(in []byte) [][]string {
	str := strings.TrimSpace(string(in))
	lines := strings.Split(str, "\n")
	out := [][]string{}
	for _, l := range lines {
		out = append(out, strings.Split(l, ""))
	}
	return out
}

func Puzzle(in [][]string) int {
	sum := 0
	queue := [][]int{}
	for y, set := range in {
		for x, char := range set {
			if char == "@" {
				queue = append(queue, []int{y, x})
			}
		}
	}

	changed := true
	for changed {
		changed = false
		size := len(queue)
		for i := 0; i < size; i++ {
			roll, q := queue[0], queue[1:]
			queue = q

			y := roll[0]
			x := roll[1]
			coords := [][]int{
				{0, -1},
				{0, 1},
				{1, 0},
				{-1, 0},
				{1, 1},
				{-1, -1},
				{1, -1},
				{-1, 1},
			}

			count := 0
			for _, coord := range coords {
				adj, err := get(y+coord[0], x+coord[1], in)
				if err != nil {
					continue
				}
				if adj == "@" {
					count++
				}
			}
			if count < 4 {
				changed = true
				// fmt.Println("found")
				in[y][x] = "."
				sum++
			} else {
				// fmt.Println("give back")
				queue = append(queue, roll)
			}
			// time.Sleep(time.Millisecond * 50)
			// fmt.Println("size", size, "i", i, len(queue))
		}
	}
	return sum
}

func get(y, x int, grid [][]string) (string, error) {
	if y < 0 || y >= len(grid) {
		return "", errors.New("oob")
	}
	if x < 0 || x >= len(grid[0]) {
		return "", errors.New("oob")
	}
	return grid[y][x], nil
}
