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
	velocity := point{0, 0}
	for {
		currentPos := point{0, 0}
		currVelocity := velocity
		for {
			currVelocity, currentPos = pos(currVelocity, currentPos)
			if onTarget(currentPos, target) {
				return currentPos, true
			}

			if currentPos.x > target.xEnd {
				velocity.y++
				break
			}

			if currentPos.y < target.yEnd {
				velocity.x++
				break
			}
			fmt.Println("velocity", velocity, "pos", currentPos)
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
var input = coord{281, 311, 74, -54}
