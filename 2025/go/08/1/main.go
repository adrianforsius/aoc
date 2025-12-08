package main

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Parse(in []byte) []node {
	str := strings.TrimSpace(string(in))
	lines := strings.Split(str, "\n")
	out := []node{}
	for _, line := range lines {
		parts := strings.Split(line, ",")
		o := []int{}
		for _, part := range parts {
			p, _ := strconv.Atoi(part)
			o = append(o, p)
		}
		out = append(out, node{line, o})
	}

	return out
}

func CalculateDistance3D(p1, p2 []int) (float64, error) {
	// Ensure both points have exactly 3 coordinates
	if len(p1) < 3 || len(p2) < 3 {
		return 0, fmt.Errorf("both points must have at least 3 coordinates")
	}

	// Calculate the difference for each axis
	dx := float64(p2[0] - p1[0])
	dy := float64(p2[1] - p1[1])
	dz := float64(p2[2] - p1[2])

	// Sum the squares of the differences
	sumOfSquares := math.Pow(dx, 2) + math.Pow(dy, 2) + math.Pow(dz, 2)

	// Return the square root
	return math.Sqrt(sumOfSquares), nil
}

func AddSorted(slice []vertex, item vertex, size int) []vertex {
	// 1. Find the position where 'item' should be inserted
	// If the item exists, 'i' is its index. If not, 'i' is where it should go.
	i, _ := slices.BinarySearchFunc(slice, item, func(a, b vertex) int {
		return cmp.Compare(a.distance, b.distance)
	})

	// Optional: If you want to allow duplicates, ignore 'found'.
	// If you want unique items only, check: if found { return slice }

	// 2. Insert the item at index 'i'
	s := slices.Insert(slice, i, item)
	return s[:min(len(s), size)]
}

type node struct {
	name string
	val  []int
}

type edge struct {
	node1 node
	node2 node
}

type vertex struct {
	distance float64
	edge     edge
}

func Puzzle(in []node, size int) int {
	distances := []vertex{}
	for i, point := range in {
		for c, other := range in[1+i:] {
			distance, _ := CalculateDistance3D(point.val, other.val)
			distances = AddSorted(distances, vertex{distance: distance, edge: edge{point, other}}, size)
		}
	}

	count := map[string][]node{}
	for _, n := range distances[:size] {
		if _, ok := count[n.edge.node1.name]; ok {
			count[n.edge.node1.name] = append(count[n.edge.node1.name], n.edge.node2)
		} else {
			count[n.edge.node1.name] = []node{n.edge.node2}
		}
		if _, ok := count[n.edge.node2.name]; ok {
			count[n.edge.node2.name] = append(count[n.edge.node2.name], n.edge.node1)
		} else {
			count[n.edge.node2.name] = []node{n.edge.node1}
		}
	}

	var v func(node, string)

	out := map[string]int{}
	visited := map[string]int{}
	proccesed := map[string]int{}
	v = func(n node, name string) {
		if _, ok := visited[n.name]; ok {
			return
		}
		// if _, ok := visited[n.name]; ok {
		out[name]++
		proccesed[n.name] = 1
		visited[n.name] = 1

		for _, nn := range count[n.name] {
			v(nn, name)
		}
	}

	top := make([]int, 0, 3)
	for name, n := range count {
		if _, ok := proccesed[name]; ok {
			continue
		}
		for _, no := range n {
			v(no, name)
		}
		visited = map[string]int{}

		idx, _ := slices.BinarySearch(top, out[name])
		top = slices.Insert(top, idx, out[name])

		if len(top) > 3 {
			top = top[1:]
		}
	}

	sum := 1
	for _, s := range top {
		sum *= s
	}
	return sum
}
