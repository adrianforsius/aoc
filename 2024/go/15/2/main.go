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

var box = map[string]Coord{
	"[": {1, 0},
	"]": {-1, 0},
}

var opp = map[string]string{
	"[": "]",
	"]": "[",
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

func (g Grid) Set(p Coord, v string) {
	g.grid[p.Y][p.X].Val = v
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

	parts := strings.Split(in, "\n\n")
	gridIn := parts[0]
	movesIn := parts[1]
	moves := strings.Split(strings.Join(strings.Split(movesIn, "\n"), ""), "")
	start := Point{}
	for y, line := range strings.Split(gridIn, "\n") {
		row := []Point{}
		for x, char := range strings.Split(line, "") {
			if char == "@" {
				start = Point{Coord: Coord{X: x * 2, Y: y}, Val: char}
				row = append(row, Point{Coord: Coord{X: x * 2, Y: y}, Val: "."})
				row = append(row, Point{Coord: Coord{X: x*2 + 1, Y: y + 1}, Val: "."})
				continue
			}
			if char == "O" {
				row = append(row, Point{Coord: Coord{X: x * 2, Y: y}, Val: "["})
				row = append(row, Point{Coord: Coord{X: x*2 + 1, Y: y + 1}, Val: "]"})
				continue
			}
			row = append(row, Point{Coord: Coord{X: x * 2, Y: y}, Val: char})
			row = append(row, Point{Coord: Coord{X: x*2 + 1, Y: y + 1}, Val: char})

		}
		grid = append(grid, row)
	}

	g := Grid{grid, start, moves}
	// g.Print(start.Coord)
	return g
}

func Day2(grid Grid) int {
	step := grid.starts.Coord
	for i, moveChar := range grid.moves {
		d := dir[moveChar]
		next := Add(step, d)
		// grid.Print(step)
		fmt.Println("next", next, "d", d, "moveChar", moveChar, "step", step, "move", i)
		if grid.Coord(next) == "#" {
			continue
		}

		if (grid.Coord(next) == "[" || grid.Coord(next) == "]") && (moveChar == "<" || moveChar == ">") {
			checkNext := next
			nexts := []Coord{}
			move := true
			for {
				nexts = append(nexts, checkNext)
				if grid.grid[checkNext.Y][checkNext.X].Val == "#" {
					move = false
					break
				}
				if grid.grid[checkNext.Y][checkNext.X].Val == "." {
					break
				}
				checkNext = Add(checkNext, d)
			}
			if move {
				grid.grid[nexts[0].Y][nexts[0].X].Val = "."
				penUlti := grid.grid[nexts[len(nexts)-2].Y][nexts[len(nexts)-2].X].Val
				grid.grid[nexts[len(nexts)-1].Y][nexts[len(nexts)-1].X].Val = penUlti
				for i := 1; i < len(nexts)-1; i++ {
					grid.grid[nexts[i].Y][nexts[i].X].Val = opp[grid.grid[nexts[i].Y][nexts[i].X].Val]
				}
				step = next
			}
			continue

		}

		if (grid.Coord(next) == "[" || grid.Coord(next) == "]") && (moveChar == "v" || moveChar == "^") {
			list := map[Coord]string{}
			d := dir[moveChar]
			nexts := []Coord{next}
			move := true
			visited := map[Coord]int{}
			for len(nexts) > 0 {
				next := Coord{}
				if len(nexts) == 1 {
					next = nexts[0]
					nexts = []Coord{}
				} else {
					next, nexts = nexts[0], nexts[1:]
				}

				if grid.Coord(next) == "#" {
					move = false
					break
				}

				if _, ok := visited[next]; ok {
					continue
				}
				visited[next] = 1

				if grid.Coord(next) == "." {
					delete(list, next)
					continue
				}

				pairDir := box[grid.Coord(next)]
				pairNext := Add(next, pairDir)
				nextNext := Add(next, d)
				nexts = append(nexts, pairNext)
				nexts = append(nexts, nextNext)
				list[pairNext] = grid.Coord(pairNext)
				list[nextNext] = grid.Coord(nextNext)
			}

			if move {
				for l := range list {
					grid.Set(l, ".")
				}
				for l, val := range list {
					grid.Set(Add(l, d), val)
				}
				step = next
			}
			continue

		}
		step = next
	}

	sum := 0
	for y, row := range grid.grid {
		for x, c := range row {
			if c.Val == "[" {
				sum += (y * 100) + x
			}
		}
	}

	return sum
}
