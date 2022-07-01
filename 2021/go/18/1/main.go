package main

import (
	"fmt"
	"strconv"
	"strings"
)

type number struct {
	value *int

	parent *number

	leftNumber  *number
	rightNumber *number
}

func (n *number) IsValid() bool {
	return n.value != nil && *n.leftNumber == number{} && *n.rightNumber == number{}
}

func main() {
	lines := strings.Split(string(input), "\n")
	roots := parseLines(lines[0:1])
	// for _, node := range roots {
	// 	fmt.Printf("entring node %+v\n", node)
	// }
	strs := rootsToString(roots[0:1])

	fmt.Println("out", lines[0:1], "tree", strs)
}

func walk(node *number, str *string) {
	fmt.Printf("entring node %+v, %p str %s, %d\n", node, node, *str, node.IsValid())

	if node.value != nil {
		// if node.parent.leftNumber == node {
		// 	fmt.Println("left number", *node.value)
		// 	str += fmt.Sprint(*node.value)
		// } else if node.parent.rightNumber == node {
		// 	fmt.Println("right number", *node.value)
		// 	str += fmt.Sprint(*node.value)
		// }
		*str += fmt.Sprint(*node.value)
		if node.parent != nil && node.parent.leftNumber == node {
			*str += ","
		}
	}
	// if node.parent != nil && node.parent.leftNumber != nil && node.parent.rightNumber != nil {
	// 	*str += ","
	// }

	// if node.parent != nil && node.parent.leftNumber != nil && node.parent.rightNumber != nil && node.parent.rightNumber == node {
	// 	*str += "]"
	// }

	if node.leftNumber != nil {
		*str += "["
		walk(node.leftNumber, str)
	}

	if node.rightNumber != nil {
		walk(node.rightNumber, str)
		*str += "]"
	}
}

func rootsToString(roots []*number) []string {
	strs := []string{}
	for _, root := range roots {
		start := ""
		walk(root, &start)
		strs = append(strs, start)
	}
	return strs
}

func parseLines(lines []string) []*number {
	roots := []*number{}
	for _, line := range lines {
		root := &number{}
		node := root
		for _, c := range line {
			char := string(c)
			n := &number{parent: node}
			if char == "[" {
				node.leftNumber = n
				node = n
			}
			if char == "]" {
				node = n.parent
			}
			if char == "," {
				node.rightNumber = n
				node = n
			}
			if char != "," && char != "]" && char != "[" {
				fmt.Println("char", char)
				val, err := strconv.Atoi(char)
				if err != nil {
					panic(err)
				}
				n.parent.value = &val
				node = n.parent
			}
		}
		roots = append(roots, root)
	}
	return roots
}

var input = `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`
