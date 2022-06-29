package main

import (
	"fmt"
)

func main() {
	velocitys := try(input)
	fmt.Println("out", len(velocitys))
}

func pos(velocity, pos point) (point, point) {
	newPoint := point{
		x: pos.x + velocity.x,
		y: pos.y + velocity.y,
	}
	if velocity.x > 0 {
		velocity.x--
	}
	velocity.y--
	return velocity, newPoint
}

func onTarget(pos point, target coord) bool {
	if pos.x <= target.xEnd && pos.x >= target.xStart && pos.y >= target.yEnd && pos.y <= target.yStart {
		return true
	}
	return false
}

func try(target coord) map[string]point {
	ps := make(map[string]point)
	for x := 1; x <= target.xEnd; x++ {
		for y := target.yEnd; y < target.yEnd*-1; y++ {
			startPos := point{0, 0}
			startVelocity := point{x, y}
			currentPos := startPos
			currVelocity := startVelocity
			for {
				currVelocity, currentPos = pos(currVelocity, currentPos)

				fmt.Println("velocity", currVelocity, "pos", currentPos)
				if onTarget(currentPos, target) {
					ps[fmt.Sprintf("%d,%d", startVelocity.x, startVelocity.y)] = startVelocity
				}

				if currentPos.x > target.xEnd {
					break
				}
				if currentPos.y < target.yEnd {
					break
				}
			}

		}
	}
	return ps
}

type point struct {
	x int
	y int
}

type coord struct {
	xStart int
	xEnd   int
	yStart int
	yEnd   int
}

// puzzle input
// var input = coord{20, 30, -5, -10}

var input = coord{281, 311, -54, -74}
