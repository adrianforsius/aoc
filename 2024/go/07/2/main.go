package main

import (
	"strconv"
	"strings"
)

type Equation struct {
	Values []int
	Result int
}

func Parse(in string) []Equation {
	eqs := []Equation{}
	for _, line := range strings.Split(in, "\n") {
		parts := strings.Split(line, ": ")

		iVals := []int{}
		for _, v := range strings.Split(parts[1], " ") {
			i, _ := strconv.Atoi(v)
			iVals = append(iVals, i)
		}
		res, _ := strconv.Atoi(parts[0])
		eqs = append(eqs, Equation{
			Result: res,
			Values: iVals,
		})

	}
	return eqs
}

var ops = []func(int, int) int{
	func(a, b int) int {
		return a + b
	},
	func(a, b int) int {
		return a * b
	},
	func(a, b int) int {
		res, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
		return res
	},
}

func calc(vals []int, res int) bool {
	val1, rest := vals[0], vals[1:]
	val2 := rest[0]

	if len(vals) <= 2 {
		if ops[0](val1, val2) == res {
			return true
		}
		if ops[1](val1, val2) == res {
			return true
		}
		if ops[2](val1, val2) == res {
			return true
		}
		return false
	}
	ret1 := calc(append([]int{ops[0](val1, val2)}, rest[1:]...), res)
	ret2 := calc(append([]int{ops[1](val1, val2)}, rest[1:]...), res)
	ret3 := calc(append([]int{ops[2](val1, val2)}, rest[1:]...), res)
	return ret1 || ret2 || ret3
}

func Day2(eqs []Equation) int {
	sum := 0
	for _, eq := range eqs {
		if calc(eq.Values, eq.Result) {
			sum += eq.Result
		}
	}

	return sum
}
