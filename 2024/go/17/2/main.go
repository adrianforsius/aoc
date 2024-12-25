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

type Coord struct {
	X int
	Y int
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

func Day2(r []Robot, height, width int) int {
	coords := map[Coord][]Robot{}
	for _, v := range r {
		coords[Coord{v.X, v.Y}] = []Robot{v}
	}

	for i := 1; i < 10_000; i++ {
		rs := map[Coord][]Robot{}
		unique := true
		for _, coord := range coords {
			for _, robot := range coord {
				x := robot.X + robot.XM
				y := robot.Y + robot.YM

				dx := x % width
				if dx < 0 {
					dx = width + dx
				}
				dy := y % height
				if dy < 0 {
					dy = height + dy
				}

				// halfWidth := width / 2
				// halfHeight := height / 2
				robot.X = dx
				robot.Y = dy
				if _, ok := rs[Coord{robot.X, robot.Y}]; ok {
					rs[Coord{robot.X, robot.Y}] = append(rs[Coord{robot.X, robot.Y}], robot)
					unique = false
				} else {
					rs[Coord{robot.X, robot.Y}] = []Robot{robot}
				}
			}
			coords = rs
		}
		if unique {
			fmt.Println("count", i)
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					if _, ok := rs[Coord{x, y}]; ok {
						fmt.Print("#")
					} else {
						fmt.Print(".")
					}
				}
				fmt.Println()
			}
			return 0
		}
	}
	return 0
}
