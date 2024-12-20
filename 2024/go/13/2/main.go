package main

import (
	"strconv"
	"strings"
)

type Coord struct {
	X int
	Y int
}

type Machine struct {
	A     Coord
	B     Coord
	Prize Coord
}

func Parse(in string) []Machine {
	machines := []Machine{}
	for _, machine := range strings.Split(in, "\n\n") {
		lines := strings.Split(machine, "\n")
		aLine := strings.Split(lines[0], "Button A: X+")
		aStrs := strings.Split(aLine[1], ", Y+")
		aX, _ := strconv.Atoi(aStrs[0])
		aY, _ := strconv.Atoi(aStrs[1])

		bLine := strings.Split(lines[1], "Button B: X+")
		bStrs := strings.Split(bLine[1], ", Y+")
		bX, _ := strconv.Atoi(bStrs[0])
		bY, _ := strconv.Atoi(bStrs[1])

		pLine := strings.Split(lines[2], "Prize: X=")
		pStrs := strings.Split(pLine[1], ", Y=")
		pX, _ := strconv.Atoi(pStrs[0])
		pY, _ := strconv.Atoi(pStrs[1])

		machines = append(machines, Machine{
			A:     Coord{X: aX, Y: aY},
			B:     Coord{X: bX, Y: bY},
			Prize: Coord{X: pX + 1_000_000_000_0000, Y: pY + 1_000_000_000_0000},
		})
	}
	return machines
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func Day2(machines []Machine) int {
	sum := 0
	for _, m := range machines {
		d := m.A.X*m.B.Y - m.A.Y*m.B.X

		d1 := m.Prize.X*m.B.Y - m.Prize.Y*m.B.X
		d2 := m.Prize.Y*m.A.X - m.Prize.X*m.A.Y

		a, b := 0, 0
		if d1%d == 0 && d2%d == 0 {
			a = d1 / d
			b = d2 / d
		}
		if a <= 0 || b <= 0 {
			continue
		}
		sum += (3 * a) + b
	}

	return sum
}
