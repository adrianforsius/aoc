package main

import (
	"strconv"
	"strings"
)

func Parse(in string) ([]Rule, [][]int) {
	parts := strings.Split(in, "\n\n")

	rules := []Rule{}
	for _, row := range strings.Split(parts[0], "\n") {
		rule := strings.Split(row, "|")
		before, _ := strconv.Atoi(rule[0])
		after, _ := strconv.Atoi(rule[1])
		rules = append(rules, Rule{before, after})
	}

	pages := [][]int{}
	for _, row := range strings.Split(parts[1], "\n") {
		pageRow := []int{}
		pageStrs := strings.Split(row, ",")
		for _, pageStr := range pageStrs {
			pageNum, _ := strconv.Atoi(pageStr)
			pageRow = append(pageRow, pageNum)

		}
		pages = append(pages, pageRow)
	}
	return rules, pages
}

type Rule struct {
	Before int
	After  int
}

func Day1(rules []Rule, pages [][]int) int {
	score := 0
	for _, page := range pages {
		valid := true
		for i := 0; i < len(page); i++ {
			after := map[int]int{}
			if i+i < len(page) {
				for _, a := range page[i+1:] {
					after[a] = 0
				}
			}

			num := page[i]

			for _, rule := range rules {
				if rule.Before == num {
					if _, ok := after[rule.After]; ok {
						after[rule.After] = 1
					}
				}
				if _, ok := after[rule.Before]; ok {
					if rule.After == num {
						valid = false
					}
				}
			}
			if valid {
				for _, a := range after {
					if a == 0 {
						valid = false
					}
				}
			}
		}
		if valid {
			score += page[len(page)/2]
		}
	}

	return score
}
