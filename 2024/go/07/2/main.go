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
	return GuardMap{board, guard, guard.Pos, map[Coord]map[Coord]int{
		{-1, 0}: make(map[Coord]int),
		{0, -1}: make(map[Coord]int),
		{1, 0}:  make(map[Coord]int),
		{0, 1}:  make(map[Coord]int),
	}}
}

type Coord struct {
	X int
	Y int
}

func (g GuardMap) Next() string {
	return g.board[g.guard.Pos.Y+g.guard.Dir.Y][g.guard.Pos.X+g.guard.Dir.X]
}

type GuardMap struct {
	board      [][]string
	guard      Guard
	startPos   Coord
	directions map[Coord]map[Coord]int
}

func (g GuardMap) Print(curr Coord) {
	b := g.board
	saved := b[curr.Y][curr.X]
	b[curr.Y][curr.X] = "O"
	for _, row := range b {
		for _, char := range row {
			fmt.Print(char)
		}
		fmt.Println()
	}
	fmt.Println()
	b[curr.Y][curr.X] = saved
}

func (g GuardMap) Loop() int {
	nextDir := g.guard.Turn()
	next := Coord{g.guard.Pos.X + nextDir.X, g.guard.Pos.Y + nextDir.Y}
	for _, d := range g.directions {
		if _, ok := d[Coord{g.guard.Pos.X + g.guard.Dir.X, g.guard.Pos.Y + g.guard.Dir.Y}]; ok {
			return 0
		}
	}
	for {
		if next.X >= len(g.board[0]) || next.X < 0 || next.Y >= len(g.board) || next.Y < 0 {
			return 0
		}
		if _, ok := g.directions[nextDir][next]; ok {
			// g.Print(Coord{next.X - nextDir.X + g.guard.Dir.X, next.Y - nextDir.Y + g.guard.Dir.Y})
			// g.Print(Coord{next.X, next.Y})
			return 1
		}

		nextCheck := next
		for g.board[nextCheck.Y][nextCheck.X] == "#" {
			nextDir = g.guard.Turn()
			nextCheck = Coord{nextCheck.X + nextDir.X, nextCheck.Y + nextDir.Y}
			if nextCheck.X >= len(g.board[0]) || nextCheck.X < 0 || nextCheck.Y >= len(g.board) || nextCheck.Y < 0 {
				break
			}
			fmt.Println(nextCheck)
		}
		next = Coord{next.X + nextDir.X, next.Y + nextDir.Y}
	}
}

func (g GuardMap) Move() int {
	loops := 0
	// i := 0
	for {
		if g.guard.Pos.Y+g.guard.Dir.Y >= len(g.board) || g.guard.Pos.Y+g.guard.Dir.Y < 0 {
			return loops
		}
		if g.guard.Pos.X+g.guard.Dir.X >= len(g.board[0]) || g.guard.Pos.X+g.guard.Dir.X < 0 {
			return loops
		}

		g.directions[g.guard.Dir][g.guard.Pos] = 1
		for g.Next() == "#" {
			g.guard.Dir = g.guard.Turn()
		}
		loops += g.Loop()

		g.guard.Pos = Coord{g.guard.Pos.X + g.guard.Dir.X, g.guard.Pos.Y + g.guard.Dir.Y}
		g.board[g.guard.Pos.Y][g.guard.Pos.X] = "X"
	}
}

// too low: 567

func Day2(gm GuardMap) int {
	return gm.Move()
}
