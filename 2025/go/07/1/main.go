package main

import (
	"errors"
	"strings"
)

func Parse(in []byte) [][]string {
	str := strings.TrimSpace(string(in))
	lines := strings.Split(str, "\n")
	out := [][]string{}
	for _, line := range lines {
		out = append(out, strings.Split(line, ""))
	}

	return out
}

func get(c coord, grid [][]string) (string, error) {
	y := c.y
	x := c.x
	if y < 0 || y >= len(grid) {
		return "", errors.New("oob")
	}
	if x < 0 || x >= len(grid[0]) {
		return "", errors.New("oob")
	}
	return grid[y][x], nil
}

type coord struct {
	y int
	x int
}

func Puzzle(in [][]string) int {
	stream := map[coord]int{}
	queue := []coord{}
	for x, char := range in[0] {
		if char == "S" {
			stream[coord{0, x}] = 1
			queue = append(queue, coord{0, x})
		}
	}

	sum := 0
	for len(queue) > 0 {
		val, temp := queue[0], queue[1:]
		queue = temp

		next := coord{val.y + 1, val.x}
		char, err := get(next, in)
		if err != nil {
			// fmt.Println(next)
			continue
		}

		if _, ok := stream[next]; ok {
			continue
		}

		if char == "." {
			stream[next] = 1
			queue = append(queue, next)
		} else {
			sum++
			left := coord{next.y, next.x - 1}
			_, err := get(left, in)
			if err != nil {
				continue
			}
			stream[left] = 1
			queue = append(queue, left)

			right := coord{next.y, next.x + 1}
			_, err = get(right, in)
			if err != nil {
				continue
			}
			stream[right] = 1
			queue = append(queue, right)
		}
	}
	return sum
}
