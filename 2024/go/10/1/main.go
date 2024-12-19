package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	Val     int
	Visited bool
	Coord
}

type Coord struct {
	X int
	Y int
}

type Grid struct {
	grid   [][]Point
	starts [][]int
	taken  map[Coord]int
}

var dir = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func (g Grid) valid(x, y int, currX, currY int) bool {
	if x < 0 || x >= len(g.grid[0]) || y < 0 || y >= len(g.grid) {
		return false
	}
	if g.grid[y][x].Visited {
		return false
	}
	if g.grid[y][x].Val > g.grid[currY][currX].Val+1 || g.grid[y][x].Val < g.grid[currY][currX].Val-1 {
		return false
	}
	if g.grid[y][x].Val == 0 {
		return false
	}
	return true
}

func (g Grid) Walk(x, y int) int {
	// g.Print(x, y)
	g.grid[y][x].Visited = true
	if x < 0 || x >= len(g.grid[0]) || y < 0 || y >= len(g.grid) {
		return 0
	}

	if g.grid[y][x].Val == 9 {
		if _, ok := g.taken[Coord{x, y}]; ok {
			return 0
		}
		// found++
		g.taken[Coord{x, y}] = 1
		return 1
	}

	sum := 0
	if g.valid(x+dir[0][0], y+dir[0][1], x, y) {
		sum += g.Walk(x+dir[0][0], y+dir[0][1])
	}
	if g.valid(x+dir[1][0], y+dir[1][1], x, y) {
		sum += g.Walk(x+dir[1][0], y+dir[1][1])
	}
	if g.valid(x+dir[2][0], y+dir[2][1], x, y) {
		sum += g.Walk(x+dir[2][0], y+dir[2][1])
	}
	if g.valid(x+dir[3][0], y+dir[3][1], x, y) {
		sum += g.Walk(x+dir[3][0], y+dir[3][1])
	}
	// fmt.Println(sum)
	return sum
}

func (g Grid) Print(inX, inY int) {
	for y, row := range g.grid {
		for x, v := range row {
			if y == inY && x == inX {
				fmt.Print("(" + fmt.Sprint(v.Val) + ")")
			} else {
				if v.Visited {
					fmt.Print(" # ")
				} else {
					fmt.Print(" " + fmt.Sprint(v.Val) + " ")
				}
			}
		}
		fmt.Println()
	}
	fmt.Println("=================")
}

func Parse(in string) Grid {
	grid := [][]Point{}
	starts := [][]int{}
	for y, line := range strings.Split(in, "\n") {
		chars := strings.Split(line, "")
		nums := []Point{}
		for x, char := range chars {
			num, _ := strconv.Atoi(char)
			if num == 0 {
				starts = append(starts, []int{x, y})
			}
			nums = append(nums, Point{Val: num, Visited: false})

		}
		grid = append(grid, nums)
	}
	return Grid{grid, starts, make(map[Coord]int, 0)}
}

func New(g Grid) Grid {
	gr := make([][]Point, 0, len(g.grid))
	for _, row := range g.grid {
		nextRow := make([]Point, len(row))
		copy(nextRow, row)
		gr = append(gr, nextRow)
	}
	return Grid{grid: gr, taken: make(map[Coord]int, 0)}
}

func Day1(grid Grid) int {
	sum := 0
	for _, start := range grid.starts {
		g := New(grid)
		grid.taken = make(map[Coord]int, 0)

		v := g.Walk(start[0], start[1])
		sum += v
	}

	return sum
}

// too high: 28399
