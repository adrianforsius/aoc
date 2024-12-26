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
	bytes []Point
	size  int
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

func newGrid(size int) [][]Point {
	grid := [][]Point{}

	for y := 0; y <= size; y++ {
		row := []Point{}
		for x := 0; x <= size; x++ {
			row = append(row, Point{Coord: Coord{X: x, Y: y}, Val: "."})
		}
		grid = append(grid, row)
	}
	return grid
}

func Parse(in string, size int) Grid {
	grid := newGrid(size)

	start := Point{Coord: Coord{X: 0, Y: 0}, Val: "S", Score: 1}
	end := Point{Coord: Coord{X: len(grid[0]) - 1, Y: len(grid) - 1}, Val: "E"}
	bytes := []Point{}
	for _, p := range strings.Split(in, "\n") {
		parts := strings.Split(p, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		bytes = append(bytes, Point{Coord: Coord{X: x, Y: y}, Val: "#"})
	}
	// fmt.Println(bytes)
	g := Grid{grid, start, end, bytes, size}
	g.Set(end)
	g.Set(start)
	return g
}

func Day2(grid Grid) string {
	curr := Point{}
	// fmt.Println(grid.bytes)
	prevs := []Point{}
	for _, b := range grid.bytes {
		// fmt.Println("byte", b)
		prevs = append(prevs, b)
		curr = b
		queue := &Heap{}
		heap.Init(queue)
		grid.grid = newGrid(grid.size)
		for _, p := range prevs {
			grid.Set(p)
		}
		score := math.MaxInt
		start := grid.start
		queue.Push(start)
		visited := map[Coord]int{}
		for queue.Len() > 0 {
			next := heap.Pop(queue).(Point)
			// grid.Print(next.Coord)

			if scr, ok := visited[next.Coord]; ok && scr < next.Score {
				continue
			}

			if next.X == grid.size && next.Y == grid.size {
				score = 10
				break
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
		if score != 10 {
			break
		}
	}

	return fmt.Sprintf("%d,%d", curr.X, curr.Y)
}
