package main

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	Val   string
	Score int
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

func Parse(in string, size int, bytes int) Grid {
	grid := [][]Point{}

	for y := 0; y <= size; y++ {
		row := []Point{}
		for x := 0; x <= size; x++ {
			row = append(row, Point{Coord: Coord{X: x, Y: y}, Val: "."})
		}
		grid = append(grid, row)
	}

	start := Point{Coord: Coord{X: 0, Y: 0}, Val: "S", Score: 1}
	end := Point{Coord: Coord{X: len(grid[0]) - 1, Y: len(grid) - 1}, Val: "E"}
	g := Grid{grid, start, end}
	for i, p := range strings.Split(in, "\n") {
		if i == bytes {
			break
		}
		parts := strings.Split(p, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		g.Set(Point{Coord: Coord{X: x, Y: y}, Val: "#"})
	}
	g.Set(end)
	g.Set(start)
	return g
}

func Day1(grid Grid) int {
	start := grid.start

	visited := map[Coord]int{}

	queue := &Heap{}
	score := math.MaxInt
	heap.Init(queue)
	queue.Push(start)
	for queue.Len() > 0 {
		next := heap.Pop(queue).(Point)
		// grid.Print(next.Coord)

		if scr, ok := visited[next.Coord]; ok && scr < next.Score {
			continue
		}

		if next.Val == "E" {
			if next.Score < score {
				score = next.Score
				continue
			}
		}

		prev := next
		for _, turn := range dir {
			nextCoord := Add(prev.Coord, turn)
			if grid.Valid(nextCoord) {
				next := grid.Coord(nextCoord)
				newScore := prev.Score + 1
				if next.Score == 0 || newScore < grid.Coord(next.Coord).Score {
					next.Score = newScore
					grid.Set(next)
					heap.Push(queue, next)
				}
			}
		}
	}

	return score - 1
}
