package main

import (
	"fmt"
	"math"
	"strings"
)

type Point struct {
	Val       string
	Direction Coord
	Score     int
	Coord
}

type Coord struct {
	X int
	Y int
}

type Grid struct {
	grid  [][]Point
	start Point
	end   Point
}

var dir = []Coord{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func (g Grid) Valid(c Coord) bool {
	if c.X < 0 || c.X >= len(g.grid[0]) || c.Y < 0 || c.Y >= len(g.grid) {
		return false
	}
	if g.Coord(c).Val == "#" {
		return false
	}
	return true
}

func (g Grid) Coord(p Coord) Point {
	return g.grid[p.Y][p.X]
}

func (g Grid) Score(p Coord, score int) {
	g.grid[p.Y][p.X].Score = score
}

func (g Grid) Direction(p Coord, dir Coord) {
	g.grid[p.Y][p.X].Direction = dir
}

func Add(a, b Coord) Coord {
	return Coord{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (g Grid) Print(c Coord) {
	for y, row := range g.grid {
		for x, v := range row {
			if y == c.Y && x == c.X {
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

	start := Point{}
	end := Point{}
	for y, line := range strings.Split(in, "\n") {
		row := []Point{}
		for x, char := range strings.Split(line, "") {
			if char == "S" {
				start = Point{Coord: Coord{X: x, Y: y}, Val: char}
			}
			if char == "E" {
				end = Point{Coord: Coord{X: x, Y: y}, Val: char}
			}
			row = append(row, Point{Coord: Coord{X: x, Y: y}, Val: char})
		}
		grid = append(grid, row)
	}
	g := Grid{grid, start, end}
	// g.Print(start.Coord)
	return g
}

func Day1(grid Grid) int {
	start := grid.start
	score := math.MaxInt
	grid.Score(start.Coord, 1)
	p := grid.Coord(start.Coord)
	grid.Score(start.Coord, 1)
	grid.Direction(start.Coord, Coord{0, 1})
	queue := []Point{p}
	for len(queue) > 0 {
		var next Point
		if len(queue) == 1 {
			next = queue[0]
			queue = []Point{}
		} else {
			next = queue[0]
			queue = queue[1:]
		}

		if next.Val == "E" {
			end := grid.Coord(next.Coord)
			if end.Score != 0 {
				score = min(score, end.Score)
			} else {
				score = end.Score
			}
			continue
		}

		prev := next
		for _, turn := range dir {
			next := grid.Coord(Add(next.Coord, turn))
			if grid.Valid(next.Coord) {
				grid.Direction(next.Coord, turn)
				// fmt.Println("valid directions")
				// fmt.Println("prev", prev.Score, "next", next.Score)
				newScore := prev.Score + 1
				if prev.Direction != turn {
					newScore += 1000
				}

				// fmt.Println("new score", newScore, "prev score", prev.Score, "new score", grid.Coord(next.Coord).Score, next.Score)
				// grid.Print(next.Coord)

				if next.Score == 0 || newScore < grid.Coord(next.Coord).Score {
					grid.Score(next.Coord, newScore)
					queue = append(queue, grid.Coord(next.Coord))
				}
			}
		}

	}

	return score - 1
}
