package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	lines := strings.Split(input, "\n")
	var grid [][]int
	for _, line := range lines {
		pos := strings.Split(line, "")
		l := []int{}
		for _, p := range pos {
			number, err := strconv.Atoi(p)
			if err != nil {
				log.Fatal(err)
			}
			l = append(l, number)
		}
		grid = append(grid, l)
	}

	sum := 0
	for i := 0; i < 100; i++ {
		flashes := [][]int{}
		log.Println("grid", grid[0])
		for lineIdx, line := range grid {
			for numberIdx, number := range line {
				if number == 9 {
					sum += 1
					flashes = append(flashes, []int{lineIdx, numberIdx})
					// grid[lineIdx][numberIdx] = 0
				} else {
					grid[lineIdx][numberIdx] += 1
				}
			}

		}
		// log.Println("grid ---")
		// for _, g := range grid {
		// 	log.Println(g)
		// }
		for 0 < len(flashes) {
			var flash []int
			flash, flashes = flashes[len(flashes)-1], flashes[:len(flashes)-1]
			x := flash[0]
			y := flash[1]

			if grid[x][y] == 0 {
				sum -= 1
				continue
			}

			//top
			if x != 0 {
				if grid[x-1][y] != 0 {
					if grid[x-1][y] == 9 {
						sum += 1
						flashes = append(flashes, []int{x - 1, y})
					} else {
						grid[x-1][y] += 1
					}
				}
			}

			// top-left
			if x != 0 && y != 0 {
				if grid[x-1][y-1] != 0 {
					if grid[x-1][y-1] == 9 {
						sum += 1
						flashes = append(flashes, []int{x - 1, y - 1})
					} else {
						grid[x-1][y-1] += 1
					}
				}
			}

			// top-right
			if x != 0 && y+1 < len(grid[0]) {
				if grid[x-1][y+1] != 0 {
					if grid[x-1][y+1] == 9 {
						sum += 1
						flashes = append(flashes, []int{x - 1, y + 1})
					} else {
						grid[x-1][y+1] += 1
					}
				}
			}

			//left
			if y != 0 {
				if grid[x][y-1] != 0 {
					if grid[x][y-1] == 9 {
						sum += 1
						flashes = append(flashes, []int{x, y - 1})
					} else {
						grid[x][y-1] += 1
					}
				}
			}

			//right
			if y+1 < len(grid[0]) {
				if grid[x][y+1] != 0 {
					if grid[x][y+1] == 9 {
						sum += 1
						flashes = append(flashes, []int{x, y + 1})
					} else {
						grid[x][y+1] += 1
					}
				}
			}

			// bottom-right
			if x+1 < len(grid) && y+1 < len(grid[0]) {
				if grid[x+1][y+1] != 0 {
					if grid[x+1][y+1] == 9 {
						sum += 1
						flashes = append(flashes, []int{x + 1, y + 1})
					} else {
						grid[x+1][y+1] += 1
					}
				}
			}

			// bottom-left
			if x+1 < len(grid) && y != 0 {
				if grid[x+1][y-1] != 0 {
					if grid[x+1][y-1] == 9 {
						sum += 1
						flashes = append(flashes, []int{x + 1, y - 1})
					} else {
						grid[x+1][y-1] += 1
					}
				}
			}

			//bottom
			if x+1 < len(grid) {
				if grid[x+1][y] != 0 {
					if grid[x+1][y] == 9 {
						sum += 1
						flashes = append(flashes, []int{x + 1, y})
					} else {
						grid[x+1][y] += 1
					}
				}
			}
			grid[x][y] = 0
			// log.Println("after flash ---")
			// for _, g := range grid {
			// 	log.Println(g)
			// }
		}
	}
	log.Println(sum)
}

// var input = `54831
// 27458
// 52645
// 61413
// 63573`

// var input = `5483143223
// 2745854711
// 5264556173
// 6141336146
// 6357385478
// 4167524645
// 2176841721
// 6882881134
// 4846848554
// 5283751526`

var input = `8577245547
1654333653
5365633785
1333243226
4272385165
5688328432
3175634254
6775142227
6152721415
2678227325`
