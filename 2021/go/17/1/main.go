package main

import (
	"fmt"
)

func main() {
	p, ok := try(input)
	fmt.Println("out", p, ok)
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

func try(target coord) (point, bool) {
	maxStepDistance := target.yEnd*-1 - 1

	fmt.Println("step max", maxStepDistance)
	for x := 0; ; x++ {
		maxHeight := 0
		startPos := point{0, 0}
		startVelocity := point{x, maxStepDistance}
		currentPos := startPos
		currVelocity := startVelocity
		for {
			currVelocity, currentPos = pos(currVelocity, currentPos)

			fmt.Println("velocity", currVelocity, "pos", currentPos)
			if currentPos.y > maxHeight {
				maxHeight = currentPos.y
			}

			if onTarget(currentPos, target) {
				fmt.Println("max height", maxHeight)
				return startVelocity, true
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
