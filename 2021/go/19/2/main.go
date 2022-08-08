package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type number struct {
	value *int

	parent *number

	leftNumber  *number
	rightNumber *number
}

func main() {
	lines := strings.Split(string(input), "\n")
	roots := parseLines(lines)
	fmt.Println("roots", len(roots))

	var root *number
	max := 0
	for i := 0; i < len(roots); i++ {
		for e := 0; e < len(roots); e++ {
			if i == e {
				continue
			}
			lines := strings.Split(string(input), "\n")
			roots := parseLines(lines)

			root = roots[i]
			nextRoot := roots[e]

			root = addition(root, nextRoot)
			rootRep := walk(roots[i], "")
			rootRepNext := walk(roots[e], "")

			for {
				explotionNode := findExplotion(root, 0)
				// fmt.Printf("node %+v\n", explotionNode)

				if explotionNode != nil {
					// fmt.Printf("none empty values %+v %p \n%+v %p\n\n", explotionNode.parent.leftNumber, explotionNode.parent.leftNumber, explotionNode.parent.rightNumber, explotionNode.parent.rightNumber)
					// if node.parent.leftNumber.value != nil && node.parent.rightNumber.value != nil {
					leftAdd(explotionNode.parent, *explotionNode.parent.leftNumber.value)
					rightAdd(explotionNode.parent, *explotionNode.parent.rightNumber.value)
					if explotionNode.parent == explotionNode.parent.parent.leftNumber {
						val := 0
						explotionNode.parent.parent.leftNumber.value = &val
					}
					if explotionNode.parent == explotionNode.parent.parent.rightNumber {
						val := 0
						explotionNode.parent.parent.rightNumber.value = &val
					}
					explotionNode.parent.leftNumber = nil
					explotionNode.parent.rightNumber = nil
					continue

					// strs := rootsToString(roots)
					// fmt.Println("explode", strs)
				}
				// fmt.Printf("node %+v, val %d\n", explotionNode, *explotionNode.value)

				splitNode := findSplit(root)
				if splitNode != nil {
					// fmt.Printf("node %+v, val %d\n", splitNode, *splitNode.value)
					val := float64(*splitNode.value) / 2
					floor := int(math.Floor(val))
					ceil := int(math.Ceil(val))
					left := &number{parent: splitNode, value: &floor}
					right := &number{parent: splitNode, value: &ceil}

					splitNode.leftNumber = left
					splitNode.rightNumber = right
					splitNode.value = nil
					continue

					// strs := rootsToString(roots)
					// fmt.Println("split", strs)
				}
				if splitNode == nil && explotionNode == nil {
					break
				}
			}

			magnitude(root)

			fmt.Printf("new root %d max %d node %s next %s\n", *root.value, max, rootRep, rootRepNext)
			if *root.value > max {
				max = *root.value
			}
		}
	}

	fmt.Println("max", max)
}

func magnitude(node *number) {
	if node.leftNumber != nil {
		magnitude(node.leftNumber)
	}
	if node.rightNumber != nil {
		magnitude(node.rightNumber)
	}

	if node.value != nil {
		if node.parent != nil && node.parent.leftNumber != nil && node.parent.rightNumber != nil {
			if node.parent.leftNumber.value != nil && node.parent.rightNumber.value != nil {
				result := *node.parent.leftNumber.value*3 + *node.parent.rightNumber.value*2
				node.parent.value = &result
				node.parent.leftNumber = nil
				node.parent.rightNumber = nil
				magnitude(node.parent)
				return
			}
		}
	}
}

func addition(left, right *number) *number {
	newRoot := &number{leftNumber: left, rightNumber: right}
	left.parent = newRoot
	right.parent = newRoot
	return newRoot
}

func leftAdd(node *number, value int) {
	if node.parent == nil {
		return
	}
	if node.parent.leftNumber.rightNumber != nil && node.parent.leftNumber != node {
		addToChildRight(node.parent.leftNumber, value)
		return
	}
	if node.parent.leftNumber.value != nil {
		*node.parent.leftNumber.value += value
		// fmt.Println("found add", *node.parent.leftNumber.value)
		return
	}
	leftAdd(node.parent, value)
}

func addToChildRight(node *number, value int) {
	if node.value != nil {
		*node.value += value
		return
	}
	addToChildRight(node.rightNumber, value)
}

func addToChildLeft(node *number, value int) {
	if node.value != nil {
		*node.value += value
		return
	}
	addToChildLeft(node.leftNumber, value)
}

func rightAdd(node *number, value int) {
	if node.parent == nil {
		return
	}
	if node.parent.rightNumber.rightNumber != nil && node.parent.rightNumber != node {
		addToChildLeft(node.parent.rightNumber, value)
		return
	}
	if node.parent.rightNumber.value != nil {
		*node.parent.rightNumber.value += value
		// fmt.Println("found add", *node.parent.rightNumber.value)
		return
	}
	rightAdd(node.parent, value)
}

func findExplotion(node *number, depth int) *number {
	if node.leftNumber != nil {
		depth++
		n := findExplotion(node.leftNumber, depth)
		depth--
		if n != nil {
			return n
		}
	}
	if node.rightNumber != nil {
		depth++
		n := findExplotion(node.rightNumber, depth)
		depth--
		if n != nil {
			return n
		}
	}
	if node.value != nil {
		if depth >= 5 {
			if node.parent.leftNumber.value != nil && node.parent.rightNumber.value != nil {
				return node
			}
		}
	}
	return nil
}

func findSplit(node *number) *number {
	// value := 0
	// if node.value != nil {
	// 	value = *node.value
	// }
	// fmt.Printf("entring node %+v, %p, %d\n", node, node, value)

	if node.leftNumber != nil {
		node := findSplit(node.leftNumber)
		if node != nil {
			return node
		}
	}
	if node.rightNumber != nil {
		node := findSplit(node.rightNumber)
		if node != nil {
			return node
		}
	}
	if node.value != nil {
		// fmt.Println("======================", *node.value)
		if *node.value >= 10 {
			return node
		}
	}
	return nil
}

func walk(node *number, str string) string {
	// value := 0
	// if node.value != nil {
	// 	value = *node.value
	// }
	// fmt.Printf("entring node %+v, %p, %d\n", node, node, value)

	if node.leftNumber != nil {
		str += "["
		str = walk(node.leftNumber, str)
		str += ","
	}
	if node.rightNumber != nil {
		str = walk(node.rightNumber, str)
		str += "]"
	}
	if node.value != nil {
		str += fmt.Sprint(*node.value)
	}
	return str
}

func rootsToString(roots []*number) []string {
	strs := []string{}
	for _, root := range roots {
		start := ""
		str := walk(root, start)
		strs = append(strs, str)
	}
	return strs
}

func parseLine(line string, val []string) []string {
	if len(line) == 0 {
		return val
	}
	char := string(line[0])
	if char == "0" || !strings.Contains("1234567890", char) {
		val = append(val, char)
		return parseLine(line[1:], val)
	}

	index := 1
	for {
		char = string(line[:index])
		if !strings.Contains("1234567890", string(char[len(char)-1])) {
			char = string(line[:index-1])
			break
		}
		index++
	}
	// fmt.Println("new char", char)

	val = append(val, char)
	// fmt.Println("val", val, "char", char)
	return parseLine(line[len(char):], val)
}

func parseLines(lines []string) []*number {
	roots := []*number{}
	for _, line := range lines {
		root := &number{}
		node := root
		for _, c := range parseLine(line, []string{}) {
			char := string(c)
			// fmt.Println("character", char)
			n := &number{parent: node}
			if char == "[" {
				node.leftNumber = n
				node = n
			}
			if char == "]" {
				node = node.parent
			}
			if char == "," {
				node.rightNumber = n
				node = n
			}
			if char != "," && char != "]" && char != "[" {
				val, err := strconv.Atoi(char)
				if err != nil {
					panic(err)
				}
				// fmt.Println("val", val)
				n.parent.value = &val
				node = node.parent
			}
		}
		roots = append(roots, root)
	}
	return roots
}

var input = `[[[0,[4,4]],6],[[[7,6],6],[[5,3],[3,2]]]]
[[[[4,6],[1,7]],[5,8]],[[9,7],[9,6]]]
[[[2,[7,1]],[[8,2],[9,3]]],3]
[[[[2,1],6],2],4]
[[[[0,3],0],6],[[9,[0,8]],[[2,1],[0,2]]]]
[[[5,1],[[0,5],1]],[[[9,9],[8,7]],7]]
[[[[0,2],8],8],[0,[7,[2,7]]]]
[[[[3,8],[6,4]],[[2,0],2]],3]
[[[[1,5],3],[[5,3],[5,4]]],[[0,1],[[1,2],8]]]
[[[1,1],[[9,3],9]],[[9,[6,5]],[2,6]]]
[[[9,3],[6,[1,5]]],[[3,8],[[4,6],[8,0]]]]
[[3,[6,7]],[[3,0],[5,[3,4]]]]
[1,[2,[[4,1],[2,3]]]]
[[6,[7,8]],[[0,[0,3]],[6,7]]]
[[8,[[0,0],[9,3]]],[[2,6],[[9,1],[4,9]]]]
[[3,0],[[8,[7,1]],4]]
[[[1,0],[[9,7],[7,8]]],[[[0,0],5],[[4,9],4]]]
[[[[4,2],7],[[4,0],0]],[[[5,4],[6,7]],[0,[1,2]]]]
[[[4,[4,3]],[[1,4],[1,1]]],6]
[[[0,[5,9]],[[7,4],2]],[[9,1],[4,7]]]
[[[[5,5],[7,0]],[8,[5,3]]],[[0,[0,2]],[[1,3],[5,8]]]]
[[9,[[9,9],2]],[[9,6],[[4,7],5]]]
[[[[8,7],[5,3]],9],[3,[6,9]]]
[[3,[0,3]],[2,6]]
[[[2,[7,0]],[6,6]],[[7,0],[[3,8],[8,5]]]]
[[[2,6],[2,7]],[[3,6],[0,[9,5]]]]
[[[5,4],1],[5,[[4,9],5]]]
[[[6,3],6],[[[6,0],0],[[4,0],7]]]
[[[[4,1],2],[3,[9,0]]],[0,8]]
[[[2,[3,9]],[[8,3],8]],[[1,[2,2]],[8,[6,4]]]]
[[[[4,3],[5,2]],0],[9,[5,[7,5]]]]
[[[3,2],5],[[[6,3],9],[[2,0],[6,7]]]]
[[[3,9],[[0,6],[0,7]]],[6,[3,2]]]
[0,0]
[[[[0,3],9],[8,[3,9]]],[[0,2],[[0,1],[3,7]]]]
[[0,[4,[3,0]]],[[7,9],[5,[8,7]]]]
[[2,9],[[0,[2,2]],1]]
[[[[5,4],[1,7]],6],[2,[[5,3],[7,7]]]]
[[[[0,4],4],[[6,6],[1,4]]],4]
[[[[4,8],5],[[6,4],[2,3]]],[9,[[8,6],[4,0]]]]
[[1,[6,[1,9]]],[3,[[4,2],[1,8]]]]
[[[[3,7],[5,9]],[[3,8],[3,3]]],[[[7,8],3],[7,3]]]
[[[[0,4],5],[4,[9,0]]],[3,[[4,1],6]]]
[[[7,[2,1]],[[1,9],1]],[[[3,4],[8,6]],6]]
[[[4,1],[5,[8,2]]],[[[1,6],9],[[4,4],2]]]
[[[7,[6,4]],[[0,1],4]],[[5,2],[[9,5],[9,3]]]]
[[[4,2],[1,8]],2]
[[[1,6],5],[8,[2,[2,3]]]]
[[[[0,2],[5,0]],[7,[0,0]]],[[6,[5,9]],5]]
[[[7,6],[9,[2,4]]],[[5,[2,6]],2]]
[[[6,2],4],[[2,9],[[3,0],[4,3]]]]
[[8,[[6,4],[0,2]]],1]
[[[4,1],[7,5]],[9,[[2,4],4]]]
[[[[4,8],[7,5]],[1,[8,5]]],[[3,5],[[9,9],[4,2]]]]
[[[7,[8,4]],[4,[5,8]]],5]
[[7,9],[2,[[9,1],[7,1]]]]
[3,[[[5,8],[4,8]],[5,4]]]
[[[0,[5,5]],[[5,4],[5,4]]],[[9,6],[[9,4],[6,5]]]]
[[7,2],[1,[8,[1,7]]]]
[2,[9,[2,[2,3]]]]
[[[3,[5,1]],[8,[6,4]]],[[2,8],[[2,2],8]]]
[[[7,3],[0,4]],[[4,0],[6,[3,4]]]]
[[4,[2,[2,8]]],[4,[[7,1],9]]]
[[8,[[6,1],2]],[1,[[1,5],9]]]
[[0,[2,[9,4]]],[[[7,4],7],8]]
[[[2,[7,0]],5],[3,[4,4]]]
[[7,[4,[6,0]]],[[4,7],[[3,7],5]]]
[[2,[[8,0],[6,1]]],[[6,[6,1]],[3,9]]]
[[9,0],[[[3,7],0],[[5,8],4]]]
[6,[[[5,8],8],[3,[4,1]]]]
[[8,[[9,3],[8,4]]],[4,[8,2]]]
[[[[8,0],8],[3,7]],[[7,[4,3]],0]]
[[[7,[2,6]],[8,0]],[4,[[1,3],[4,1]]]]
[1,[[[4,9],[4,9]],[[7,0],[6,6]]]]
[9,4]
[[6,7],4]
[[2,[5,2]],[[2,4],[[4,6],[5,5]]]]
[[[5,2],[[5,5],[8,1]]],[[9,[1,6]],3]]
[[[[4,3],1],[8,9]],6]
[[[[3,2],[4,5]],4],[[[4,3],[0,0]],[[3,0],1]]]
[[6,7],[[8,5],[[7,2],4]]]
[[[[8,1],[5,8]],7],[[[5,2],[4,3]],1]]
[[2,[[4,9],5]],[1,1]]
[[9,1],[[[0,8],[1,8]],7]]
[[9,3],[6,4]]
[[8,[4,2]],[[7,[7,4]],[[0,9],[6,1]]]]
[[[[0,5],7],[[7,7],2]],[[2,[5,8]],[9,6]]]
[[[2,1],[7,[1,3]]],[[2,[7,1]],0]]
[[[8,[8,4]],[2,[4,3]]],[[2,[5,6]],[[2,0],[7,3]]]]
[[4,[[4,3],[5,2]]],[1,3]]
[[5,[5,0]],9]
[[[2,[7,6]],[1,8]],[[[5,2],2],0]]
[[2,[2,3]],[[9,8],[[0,1],[3,5]]]]
[[7,[[3,7],3]],[[[7,6],[4,8]],[[1,7],[8,6]]]]
[[[0,0],[[6,1],5]],[5,[5,4]]]
[[2,3],[4,[3,5]]]
[[8,[7,7]],[8,[4,[8,1]]]]
[[[[4,0],3],[[0,0],[0,0]]],[[[6,0],4],[[1,7],0]]]
[[[[6,4],[3,1]],[[2,8],[1,2]]],[4,[[6,5],4]]]
[[[[5,3],7],[4,[2,6]]],[[6,[4,5]],[1,[9,0]]]]`

// var input = `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
// [[[5,[2,8]],4],[5,[[9,9],0]]]
// [6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
// [[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
// [[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
// [[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
// [[[[5,4],[7,7]],8],[[8,3],8]]
// [[9,3],[[9,9],[6,[4,9]]]]
// [[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
// [[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`
