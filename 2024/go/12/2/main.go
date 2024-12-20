package main

import (
	"fmt"
	"strings"
)

type Coord struct {
	X   int
	Y   int
	Val string
}

type Grid struct {
	grid    [][]Coord
	visited map[Coord]int
}

var dir = []Coord{
	{X: 1, Y: 0},
	{X: -1, Y: 0},
	{X: 0, Y: 1},
	{X: 0, Y: -1},
}

func (g Grid) valid(coord Coord) bool {
	if coord.X < 0 || coord.X >= len(g.grid[0]) || coord.Y < 0 || coord.Y >= len(g.grid) {
		return false
	}
	if coord.Val != g.grid[coord.Y][coord.X].Val {
		return false
	}
	return true
}

type Lot struct {
	Sides map[Coord]int
	Size  int
}

func (g Grid) Walk(coord Coord) Lot {
	if !g.valid(coord) {
		return Lot{Parameter: 1}
	}
	if _, ok := g.visited[coord]; ok {
		return Lot{}
	}
	g.visited[coord] = 1
	// fmt.Println("coord", coord)

	// sum := 0
	// sum += g.Walk(Coord{X: coord.X + dir[0][0], Y: coord.Y + dir[0][1], Val: coord.Val})
	// sum += g.Walk(Coord{X: coord.X + dir[1][0], Y: coord.Y + dir[1][1], Val: coord.Val})
	// sum += g.Walk(Coord{X: coord.X + dir[2][0], Y: coord.Y + dir[2][1], Val: coord.Val})
	// sum += g.Walk(Coord{X: coord.X + dir[3][0], Y: coord.Y + dir[3][1], Val: coord.Val})
	up := g.Walk(Coord{X: coord.X + dir[0].X, Y: coord.Y + dir[0].Y, Val: coord.Val})
	right := g.Walk(Coord{X: coord.X + dir[1].X, Y: coord.Y + dir[1].Y, Val: coord.Val})
	down := g.Walk(Coord{X: coord.X + dir[2].X, Y: coord.Y + dir[2].Y, Val: coord.Val})
	left := g.Walk(Coord{X: coord.X + dir[3].X, Y: coord.Y + dir[3].Y, Val: coord.Val})
	return Lot{
		Parameter: up.Parameter + right.Parameter + down.Parameter + left.Parameter,
		Size:      up.Size + right.Size + down.Size + left.Size + 1,
	}
}

func Parse(in string) Grid {
	grid := [][]Coord{}
	for y, line := range strings.Split(in, "\n") {
		chars := strings.Split(line, "")
		cs := []Coord{}
		for x, char := range chars {
			cs = append(cs, Coord{x, y, char})
		}
		grid = append(grid, cs)
	}
	return Grid{grid, make(map[Coord]int, 0)}
}

func Day2(grid Grid) int {
	sum := 0
	for _, row := range grid.grid {
		for _, coord := range row {
			if _, ok := grid.visited[coord]; !ok {
				lot := grid.Walk(coord)
				fmt.Println("lot", lot)
				sum += lot.Parameter * lot.Size
			}
		}
	}

	return sum
}
