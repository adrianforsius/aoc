package main

import (
	"log"
	"strings"
	"unicode"
)

type Edge struct {
	Name  string
	Large bool
	Edges []*Edge
}

func main() {
	lines := strings.Split(input, "\n")
	edges := make(map[string]*Edge)
	for _, line := range lines {
		parts := strings.Split(line, "-")
		edge0, ok := edges[parts[0]]
		if !ok {
			edge0 = &Edge{
				Large: IsUpper(parts[0]),
				Name:  parts[0],
			}
		}

		edge1, ok := edges[parts[1]]
		if !ok {
			edge1 = &Edge{
				Large: IsUpper(parts[1]),
				Name:  parts[1],
			}
		}
		edge1.Edges = append(edge1.Edges, edge0)
		edges[parts[1]] = edge1

		edge0.Edges = append(edge0.Edges, edge1)
		edges[parts[0]] = edge0
		// for _, e := range edges {
		// 	log.Println("line", i, e.Name, len(e.Edges))
		// }
	}

	path := []string{}
	paths := [][]string{}
	log.Println(edges["start"], edges["start"].Name)
	log.Println(edges["start"].Edges[0], edges["start"].Edges[0].Name)
	DFS(edges["start"], path, &paths)
	// log.Println("paths", len(paths), paths)
	count := 0
	for _, ps := range paths {
		log.Println("ps", ps)
	}
	log.Println("paths", len(paths), "count", count)
}

func DFS(edge *Edge, path []string, paths *[][]string) {
	// time.Sleep(time.Second * 1)
	// log.Println("edge", edge.Name, "path", path, "paths", paths)
	p := path
	if invalidOccurences(p) {
		// log.Println("small cave return", path, edge.Name)
		return
	}

	p = append(p, edge.Name)
	if edge.Name == "end" {
		*paths = append(*paths, p)
		// log.Println("adding path", p, edge.Name)
		return
	}

	for _, e := range edge.Edges {
		DFS(e, p, paths)
	}
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func invalidOccurences(s []string) bool {
	occ := make(map[string]int)
	for _, a := range s {
		if IsUpper(a) {
			continue
		}
		if _, ok := occ[a]; !ok {
			occ[a] = 1
		} else {
			occ[a] += 1
		}
	}

	if occ["start"] > 1 {
		return true
	}
	if occ["end"] > 1 {
		return true
	}
	// log.Println("occ", occ)
	count := 0
	for _, o := range occ {
		if o > 2 {
			return true
		}
		if o > 1 {
			count += 1
		}
		if count == 2 {
			return true
		}
	}
	return false
}

// var input = `start-A
// start-b
// A-c
// A-b
// b-d
// A-end
// b-end`

// slightly larger
// var input = `dc-end
// HN-start
// start-kj
// dc-start
// dc-HN
// LN-dc
// HN-end
// kj-sa
// kj-HN
// kj-dc`

// puzzle input
var input = `yb-start
de-vd
rj-yb
rj-VP
OC-de
MU-de
end-DN
vd-end
WK-vd
rj-de
DN-vd
start-VP
DN-yb
vd-MU
DN-rj
de-VP
yb-OC
start-rj
oa-MU
yb-de
oa-VP
jv-MU
yb-MU
end-OC`
