package main

import (
	"fmt"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Grid struct {
	grid    [][]string
	letters map[string][]Point
}

func (g Grid) Print(m map[Point]int) {
	for y, row := range g.grid {
		for x, v := range row {
			if _, ok := m[Point{x, y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(v)
			}
		}
		fmt.Println()
	}
}

func Parse(in string) Grid {
	grid := [][]string{}
	letters := map[string][]Point{}
	for y, line := range strings.Split(in, "\n") {
		chars := strings.Split(line, "")
		grid = append(grid, chars)
		for x, char := range chars {
			if char == "." {
				continue
			}
			if _, ok := letters[char]; ok {
				letters[char] = append(letters[char], Point{x, y})
			} else {
				letters[char] = []Point{{x, y}}
			}
		}
	}
	return Grid{grid, letters}
}

func Day2(grid Grid) int {
	sum := map[Point]int{}
	for i := range grid.letters {
		for _, curr := range grid.letters[i] {
			for _, next := range grid.letters[i] {
				if curr == next {
					continue
				}
				dx := curr.X - next.X
				dy := curr.Y - next.Y
				sum[Point{next.X, next.Y}] = 1

				opp := Point{next.X + dx*2, next.Y + dy*2}
				for {
					if opp.X < len(grid.grid[0]) && opp.X >= 0 && opp.Y < len(grid.grid) && opp.Y >= 0 {
						sum[opp] = 1
					} else {
						break
					}
					opp = Point{opp.X + dx, opp.Y + dy}
				}
			}
		}
	}
	grid.Print(sum)
	return len(sum)
}
