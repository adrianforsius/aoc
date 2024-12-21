package main

import (
	"fmt"
	"strings"
)

type Point struct {
	Val string
	Coord
}

type Coord struct {
	X int
	Y int
}

type Grid struct {
	grid   [][]Point
	starts Point
	moves  []string
}

var dir = map[string]Coord{
	">": {1, 0},
	"<": {-1, 0},
	"v": {0, 1},
	"^": {0, -1},
}

func (g Grid) valid(c Coord) bool {
	return c.X < 0 || c.X >= len(g.grid[0]) || c.Y < 0 || c.Y >= len(g.grid)
}

func (g Grid) Coord(p Coord) string {
	return g.grid[p.Y][p.X].Val
}

func Add(a, b Coord) Coord {
	return Coord{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (g Grid) Walk(x, y int) int {
	sum := 0
	return sum
}

func (g Grid) Print(inX, inY int) {
	for y, row := range g.grid {
		for x, v := range row {
			if y == inY && x == inX {
				fmt.Print("(" + fmt.Sprint(v.Val) + ")")
			} else {
				fmt.Print(" " + fmt.Sprint(v.Val) + " ")
			}
		}
		fmt.Println()
	}
	fmt.Println("=================")
}

func Parse(in string) Grid {
	grid := [][]Point{}

	parts := strings.Split(in, "\n\n")
	gridIn := parts[0]
	movesIn := parts[1]
	moves := strings.Split(strings.Join(strings.Split(movesIn, "\n"), ""), "")
	start := Point{}
	for y, line := range strings.Split(gridIn, "\n") {
		row := []Point{}
		for x, char := range strings.Split(line, "") {
			if char == "@" {
				start = Point{Coord: Coord{X: x, Y: y}, Val: char}
				char = "."
			}
			row = append(row, Point{Coord: Coord{X: x, Y: y}, Val: char})
		}
		grid = append(grid, row)
	}
	return Grid{grid, start, moves}
}

func Day1(grid Grid) int {
	step := grid.starts.Coord
	for _, moveChar := range grid.moves {
		d := dir[moveChar]
		next := Add(step, d)
		// grid.Print(step.X, step.Y)
		// fmt.Println("next", next, "d", d, "moveChar", moveChar, "step", step)
		if grid.Coord(next) == "#" {
			continue
		}

		if grid.Coord(next) == "O" {
			checkNext := Add(next, d)
			for {
				if grid.grid[checkNext.Y][checkNext.X].Val == "." {
					grid.grid[next.Y][next.X].Val = "."
					grid.grid[checkNext.Y][checkNext.X].Val = "O"
					step = next
					break
				}
				if grid.grid[checkNext.Y][checkNext.X].Val == "#" {
					break
				}
				checkNext = Add(checkNext, d)
			}
			continue
		}
		step = next
	}

	sum := 0
	for y, row := range grid.grid {
		for x, c := range row {
			if c.Val == "O" {
				sum += (y * 100) + x
			}
		}
	}

	return sum
}
