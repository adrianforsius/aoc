package main

import (
	"container/heap"
	"fmt"
	"math"
	"strings"
)

type Point struct {
	Val   string
	Score int
	Coord
	Cheat int
	Path  []Step
}

type Step struct {
	Coord
	Score int
}

type Coord struct {
	X int
	Y int
}

type Heap []Point

func (h Heap) Len() int { return len(h) }

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

func sides(c Coord) []Coord {
	out := []Coord{}
	for _, d := range dir {
		if d.X+c.X == 0 && c.Y+d.Y == 0 {
			continue
		}
		if diff(d.X, c.X) == 2 || diff(c.Y, d.Y) == 2 {
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
	return true
}

func (g Grid) Coord(p Coord) Point {
	return g.grid[p.Y][p.X]
}

func (g Grid) Set(p Point) {
	g.grid[p.Y][p.X] = p
}

func Add(a, b Coord) Coord {
	return Coord{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (g Grid) PrintMap(c map[Coord]Point) {
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
				start = Point{
					Coord: Coord{X: x, Y: y},
					Val:   char,
				}
			}
			if char == "E" {
				end = Point{Coord: Coord{X: x, Y: y}, Val: char}
			}

			row = append(row, Point{Coord: Coord{X: x, Y: y}, Val: char})
		}
		grid = append(grid, row)
	}
	g := Grid{grid, start, end}

	return g
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func Day2(grid Grid, limit int) int {
	start := grid.start
	// start.Cheat = 20
	start.Path = []Step{{Coord: start.Coord, Score: 0}}

	queue := &Heap{}
	heap.Init(queue)
	queue.Push(start)
	visited := map[Coord]int{}
	base := Point{}
	for queue.Len() > 0 {
		next := heap.Pop(queue).(Point)
		if next.Val == "E" {
			base = next
			break
		}

		if _, ok := visited[next.Coord]; ok {
			continue
		}
		visited[next.Coord] = 1

		prev := next
		for _, turn := range dir {
			nextCoord := Add(prev.Coord, turn)
			if grid.Valid(nextCoord) {
				next := grid.Coord(nextCoord)
				if next.Val == "#" {
					continue
				}
				newScore := prev.Score + 1
				next.Path = append(append([]Step{}, prev.Path...), Step{Coord: next.Coord, Score: newScore})
				next.Score = newScore
				heap.Push(queue, next)
			}
		}
	}

	fmt.Println("base", len(base.Path))
	cheat := map[int]int{}
	sum := 0
	for i, a := range base.Path[:len(base.Path)-limit] {
		for j, b := range base.Path[i+limit:] {
			d := int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
			if d <= 20 && d <= j {
				sum++
			}
		}
	}
	fmt.Println(cheat)

	return sum
}
