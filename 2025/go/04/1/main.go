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
	for y, set := range in {
		for x, char := range set {
			if char == "@" {
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
					sum++
				}
			}
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
