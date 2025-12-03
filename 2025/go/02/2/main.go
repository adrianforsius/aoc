package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Parse(in []byte) [][]int {

	str := strings.TrimSpace(string(in))
	ranges := strings.Split(str, ",")
	out := [][]int{}
	for _, rs := range ranges {
		parts := strings.Split(rs, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		out = append(out, []int{start, end})
	}
	return out
}

func Puzzle(in [][]int) int {
	sum := 0
	for _, set := range in {
		for i := set[0]; i <= set[1]; i++ {
			val := fmt.Sprint(i)

			for j := 1; j <= int(math.Trunc(float64(len(val)/2))); j++ {
				if len(val)%j != 0 {
					continue
				}

				parts := len(val) / j
				hit := true
				first := val[0:j]
				for e := 1; e < parts; e++ {
					if val[j*e:j+e*j] != first {
						hit = false
						break
					}
				}
				if hit {
					sum += i
					break
				}

			}
		}
	}
	return sum
}
