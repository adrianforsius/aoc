package main

import (
	"regexp"
	"strconv"
	"strings"
)

func Parse(in string) string {
	return in
}

func Day2(in string) int {
	doReg := regexp.MustCompile(`(do\(\)|don't\(\))`)
	mulReg := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)

	pos := doReg.FindAllIndex([]byte(in), -1)
	nextDoPos, pos := pos[0][0], pos[1:]
	dos := doReg.FindAllString(in, -1)
	nextDo, dos := dos[0], dos[1:]

	mulPos := mulReg.FindAllIndex([]byte(in), -1)
	muls := mulReg.FindAllString(in, -1)

	enable := true
	score := 0
	for i := 0; i < len(muls); i++ {
		for mulPos[i][0] > nextDoPos {
			if nextDo == "do()" {
				enable = true
			} else {
				enable = false
			}
			if len(pos) > 0 {
				nextDo, dos = dos[0], dos[1:]
				nextDoPos, pos = pos[0][0], pos[1:]
			} else {
				break
			}
		}
		if enable {
			ss := strings.Split(muls[i], "mul(")
			ss = strings.Split(ss[1], ")")
			nums := strings.Split(ss[0], ",")
			mul1, _ := strconv.Atoi(nums[0])
			mul2, _ := strconv.Atoi(nums[1])
			score += mul1 * mul2
		}
	}
	return score
}
