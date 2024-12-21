package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Robot struct {
	X int
	Y int

	XM int
	YM int
}

func Parse(in string) []Robot {
	robots := []Robot{}
	for _, line := range strings.Split(in, "\n") {
		aLine := strings.Split(line, "p=")
		midLine := strings.Split(aLine[1], " v=")
		pParts := strings.Split(midLine[0], ",")
		pX, _ := strconv.Atoi(pParts[0])
		pY, _ := strconv.Atoi(pParts[1])
		vParts := strings.Split(midLine[1], ",")
		vX, _ := strconv.Atoi(vParts[0])
		vY, _ := strconv.Atoi(vParts[1])

		robots = append(robots, Robot{
			X:  pX,
			Y:  pY,
			XM: vX,
			YM: vY,
		})
	}
	return robots
}

func Day1(robots []Robot, height, width int) int {
	fmt.Println(robots)
	rs := map[int]int{}
	for _, robot := range robots {
		// fmt.Println("robot", "X", robot.X, "Y", robot.Y)
		secs := 100
		x := robot.X + robot.XM*secs
		y := robot.Y + robot.YM*secs

		// height := 11
		// width := 7
		dx := x % width
		// fmt.Println("dx mod", dx)
		if dx < 0 {
			dx = width + dx
		}
		dy := y % height
		// fmt.Println("dy mod", dy)
		if dy < 0 {
			dy = height + dy
		}
		fmt.Println(dx, dy)

		halfWidth := width / 2
		halfHeight := height / 2
		if dx < halfWidth && dy < halfHeight {
			rs[0]++
		}
		if dx > halfWidth && dy < halfHeight {
			rs[1]++
		}
		if dx < halfWidth && dy > halfHeight {
			rs[2]++
		}
		if dx > halfWidth && dy > halfHeight {
			rs[3]++
		}
	}

	fmt.Println(rs)
	sum := 1
	for _, r := range rs {
		sum *= r
	}

	return sum
}
