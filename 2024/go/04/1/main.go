package main

import (
	"strings"
)

func Parse(in string) [][]string {
	rows := strings.Split(in, "\n")
	out := [][]string{}
	for _, row := range rows {
		out = append(out, strings.Split(row, ""))
	}
	return out
}

type Coord struct {
	X int
	Y int
}

var dirs = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func Day1(in [][]string) int {
	word := "XMAS"

	score := 0
	width := len(in[0])
	height := len(in)
	for yI, row := range in {
		for xI := range row {
			for _, d := range dirs {
				count := 0
				for i := 0; i < len(word); i++ {
					c := Coord{d[0] * i, d[1] * i}
					x := c.X + xI
					if x < 0 || x >= width {
						break
					}
					y := c.Y + yI
					if y < 0 || y >= height {
						break
					}
					if in[y][x] == string(word[i]) {
						count++
					}
				}
				if count == len(word) {
					score++
				}
			}
		}
	}
	return score
}
