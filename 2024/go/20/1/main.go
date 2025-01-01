package main

import (
	"container/heap"
	"fmt"
	"strings"
)

type Point struct {
	Val       string
	Score     int
	Direction Coord
	Coord
	Path []CheatCheck
}

type CheatCheck struct {
	Score     int
	Coord     Coord
	Direction Coord
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

func (g Grid) PrintMap(c map[Coord]CheatCheck) {
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

func Day1(grid Grid) int {
	start := grid.start
	start.Path = []CheatCheck{{Coord: start.Coord, Direction: start.Direction, Score: 0}}

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
				next.Path = append(append([]CheatCheck{}, prev.Path...), CheatCheck{Coord: next.Coord, Direction: next.Direction, Score: newScore})
				next.Score = newScore
				next.Direction = turn
				heap.Push(queue, next)
			}
		}
	}

	fmt.Println("base", len(base.Path))
	// fmt.Println(base.Path)
	wallMap := make(map[Coord]CheatCheck, len(base.Path)*2)
	pathMap := make(map[Coord]CheatCheck, len(base.Path))
	for _, b := range base.Path {
		pathMap[b.Coord] = b
		for _, side := range dir {
			next := grid.Coord(Add(b.Coord, side))
			if _, ok := wallMap[next.Coord]; !ok && next.Val == "#" {
				wallMap[next.Coord] = CheatCheck{Direction: side, Score: b.Score + 1, Coord: next.Coord}
			}
		}
	}

	// fmt.Println("wallMap", len(wallMap))
	// fmt.Println("pathMap", len(pathMap))
	cheats := map[int]int{}

	sum := 0
	for _, check := range wallMap {
		queue = &Heap{}
		heap.Init(queue)
		queue.Push(Point{Direction: check.Direction, Coord: check.Coord, Score: check.Score})
		visited = map[Coord]int{}
		for queue.Len() > 0 {
			next := heap.Pop(queue).(Point)
			if _, ok := visited[next.Coord]; ok {
				continue
			}
			visited[next.Coord] = 1

			if v, ok := pathMap[next.Coord]; ok {
				if v.Score <= next.Score {
					continue
				}
				if v.Score-next.Score >= 100 {
					sum++
				}

				break
			}

			prev := next
			for _, turn := range dir {
				nextCoord := Add(prev.Coord, turn)
				if grid.Valid(nextCoord) {
					next := grid.Coord(nextCoord)
					if next.Val == "#" {
						continue
					}

					next.Score = prev.Score + 1
					heap.Push(queue, next)
				}
			}
		}
	}

	// grid.PrintMap(pathMap)
	// grid.PrintMap(wallMap)
	// fmt.Println(cheats)

	// out := 0
	// for sec := range cheatSheet {
	// 	if base-sec >= 100 {
	// 		out++
	// 	}
	// }
	// fmt.Println(sum)
	return sum
}
