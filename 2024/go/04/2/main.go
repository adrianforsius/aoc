package main

import (
	"fmt"
	"slices"
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
	{-1, 1},
	{1, -1},
	{1, 1},
}

type Grid struct {
	board   [][]string
	corners [][]int
	width   int
	height  int
}

func NewGrid(b [][]string, c [][]int) Grid {
	return Grid{
		board:   b,
		corners: c,
		width:   len(b[0]),
		height:  len(b),
	}
}

func (g Grid) Inside(x, y int) bool {
	if x < 0 || x >= g.width {
		return false
	}
	if y < 0 || y >= g.height {
		return false
	}
	return true
}

func (g Grid) InsideAndChar(x, y int) bool {
	return g.Inside(x, y) && g.IsChar(x, y)
}

func (g Grid) CharDistribution(cs []Coord) bool {
	mCount := 0
	sCount := 0
	for _, c := range cs {
		if g.board[c.Y][c.X] == "M" {
			mCount++
		}
		if g.board[c.Y][c.X] == "S" {
			sCount++
		}
	}
	return sCount == 2 && mCount == 2
}

func (g Grid) IsChar(x, y int) bool {
	letters := []string{"S", "M"}
	return slices.Contains(letters, g.board[y][x])
}

func (g Grid) Print(x, y int) {
	fmt.Println("X", x, "Y", y)
	for j := y - 1; j < y+2; j++ {
		for i := x - 1; i < x+2; i++ {
			fmt.Print(g.board[j][i])
		}
		fmt.Println()
	}
	fmt.Println("---")
}

func Day2(in [][]string) int {
	score := 0
	board := NewGrid(in, dirs)
	for yI, row := range in {
		for xI, val := range row {
			if val != "A" {
				continue
			}

			c1 := Coord{dirs[0][0], dirs[0][1]}
			x1 := c1.X + xI
			y1 := c1.Y + yI
			if !board.InsideAndChar(x1, y1) {
				continue
			}

			c2 := Coord{dirs[1][0], dirs[1][1]}
			x2 := c2.X + xI
			y2 := c2.Y + yI
			if !board.InsideAndChar(x2, y2) {
				continue
			}

			c3 := Coord{dirs[2][0], dirs[2][1]}
			x3 := c3.X + xI
			y3 := c3.Y + yI
			if !board.InsideAndChar(x3, y3) {
				continue
			}

			c4 := Coord{dirs[3][0], dirs[3][1]}
			x4 := c4.X + xI
			y4 := c4.Y + yI
			if !board.InsideAndChar(x4, y4) {
				continue
			}
			if !board.CharDistribution([]Coord{{x1, y1}, {x2, y2}, {x3, y3}, {x4, y4}}) {
				continue
			}

			if in[y1][x1] == in[y2][x2] || in[y1][x1] == in[y3][x3] {
				// board.Print(xI, yI)
				score++
			}
		}
	}
	return score
}

// 2572: too high
