package main

import (
	"container/heap"
	"fmt"
	"math"
	"strings"
)

type Point struct {
	Val       string
	Direction Coord
	Score     int
	Path      []Coord
	Coord
}

type Coord struct {
	X int
	Y int
}

type Heap []Point

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].Score < h[j].Score }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(Point)) }
func (h *Heap) Pop() interface{} {
	heap := *h
	size := len(heap)
	next := heap[size-1]
	*h = heap[0 : size-1]
	return next
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

var (
	NORTH = Coord{0, -1}
	SOUTH = Coord{0, 1}
	EAST  = Coord{1, 0}
	WEST  = Coord{-1, 0}
)

func forward(c Coord) []Coord {
	out := []Coord{}
	for _, d := range dir {
		if d.X+c.X == 0 && c.Y+d.Y == 0 {
			continue
		}
		out = append(out, d)

	}
	return out
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

func (g Grid) Set(p Point) {
	g.grid[p.Y][p.X] = p
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

func (g Grid) PrintCoords(c map[Coord]int) {
	for y, row := range g.grid {
		for x, v := range row {
			if _, ok := c[Coord{x, y}]; ok {
				fmt.Print("(" + fmt.Sprint(v.Val) + ")")
			} else {
				fmt.Print(" " + fmt.Sprint(v.Val) + " ")
			}
		}
		fmt.Println()
	}
	fmt.Println("=================")
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

func Day2(grid Grid) int {
	start := grid.start
	score := math.MaxInt
	start.Direction = Coord{1, 0}
	start.Path = []Coord{start.Coord}
	start.Score = 1
	grid.Set(start)

	type combo struct {
		c   Coord
		dir Coord
	}
	visited := map[combo]int{}
	coords := map[Coord]int{}

	queue := &Heap{}
	heap.Init(queue)
	queue.Push(start)
	for queue.Len() > 0 {
		next := heap.Pop(queue).(Point)
		if next.Score > score {
			break
		}

		if scr, ok := visited[combo{next.Coord, next.Direction}]; ok && scr < next.Score {
			continue
		}
		visited[combo{next.Coord, next.Direction}] = next.Score

		if next.Val == "E" {
			if next.Score < score {
				score = next.Score
			}
			for _, p := range next.Path {
				coords[p] = 1
			}
		}

		prev := next
		prevPath := prev.Path
		for _, turn := range forward(prev.Direction) {
			next := grid.Coord(Add(prev.Coord, turn))
			// grid.Print(next.Coord)
			if grid.Valid(next.Coord) {
				newScore := prev.Score + 1
				if prev.Direction != turn {
					newScore += 1000
				}

				if next.Score == 0 || newScore < grid.Coord(next.Coord).Score {
					// fmt.Println("add", newScore, score)
					point := next
					point.Score = newScore
					point.Direction = turn
					point.Path = append(append([]Coord{}, prevPath...), next.Coord)
					// grid.Set(point)
					// fmt.Println("point", point.Coord, prev.Coord)
					// m := map[Coord]int{}
					// for _, p := range point.Path {
					// 	m[p] = 1
					// }
					// grid.PrintCoords(m)
					heap.Push(queue, point)
				}
			}
		}
	}
	// grid.PrintCoords(coords)

	return len(coords)
}
