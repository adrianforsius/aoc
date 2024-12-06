package main

import (
	"fmt"
	"strings"
)

type Guard struct {
	Dir Coord
	Pos Coord
}

func (gg Guard) Turn() Coord {
	chars := map[Coord]Coord{
		{0, 1}:  {-1, 0},
		{-1, 0}: {0, -1},
		{0, -1}: {1, 0},
		{1, 0}:  {0, 1},
	}
	return chars[gg.Dir]
}

func NewGuard(char string) (Guard, bool) {
	chars := map[string]Guard{
		"<": {Dir: Coord{-1, 0}},
		">": {Dir: Coord{1, 0}},
		"^": {Dir: Coord{0, -1}},
		"v": {Dir: Coord{0, 1}},
	}
	if guard, ok := chars[char]; ok {
		return guard, ok
	}
	return Guard{}, false
}

func Parse(in string) GuardMap {
	board := [][]string{}
	for _, row := range strings.Split(in, "\n") {
		char := strings.Split(row, "")
		board = append(board, char)
	}

	guard := Guard{}
	for y, row := range board {
		for x, char := range row {
			if g, ok := NewGuard(char); ok {
				g.Pos = Coord{x, y}
				guard = g
			}
		}
	}
	board[guard.Pos.Y][guard.Pos.X] = "X"
	return GuardMap{board, guard}
}

type Coord struct {
	X int
	Y int
}

func (g GuardMap) Next() string {
	return g.board[g.guard.Pos.Y+g.guard.Dir.Y][g.guard.Pos.X+g.guard.Dir.X]
}

type GuardMap struct {
	board [][]string
	guard Guard
}

func (g GuardMap) Print() {
	// b := g.board
	// b[g.guard.Pos.Y][g.guard.Pos.X] = "@"
	for _, row := range g.board {
		for _, char := range row {
			fmt.Print(char)
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g GuardMap) Move() int {
	moves := 1
	for {
		if g.guard.Pos.Y+g.guard.Dir.Y >= len(g.board) || g.guard.Pos.Y+g.guard.Dir.Y < 0 {
			return moves
		}
		if g.guard.Pos.X+g.guard.Dir.X >= len(g.board[0]) || g.guard.Pos.X+g.guard.Dir.X < 0 {
			return moves
		}

		for g.Next() == "#" {
			g.guard.Dir = g.guard.Turn()
		}

		if g.Next() != "X" {
			moves++
		}
		g.guard.Pos = Coord{g.guard.Pos.X + g.guard.Dir.X, g.guard.Pos.Y + g.guard.Dir.Y}
		g.board[g.guard.Pos.Y][g.guard.Pos.X] = "X"

	}
}

func Day1(gm GuardMap) int {
	return gm.Move()
	// gm.Move()
	// xs := 0
	// gm.Print()
	// for _, row := range gm.board {
	// 	for _, char := range row {
	// 		if char == "X" {
	// 			xs++
	// 		}
	// 	}
	// }
	// return xs
}

// too high; 4698, 4699, 5700
