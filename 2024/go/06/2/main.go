package main

import (
	"slices"
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

type RuleSet struct {
	rules []Rule
}

func (r RuleSet) Correct(page []int) []int {
	out := []int{page[0]}
	for _, p := range page[1:] {
		afterIndex := 0
		for i, v := range out {
			for _, rule := range r.rules {
				if rule.After == p && rule.Before == v {
					afterIndex = i + 1
				}
			}
		}
		out = slices.Insert(out, afterIndex, p)
	}
	return out
}

func (r RuleSet) Valid(num int, after map[int]int) bool {
	for _, rule := range r.rules {
		if rule.Before == num {
			if _, ok := after[rule.After]; ok {
				after[rule.After] = 1
			}
		}
		if _, ok := after[rule.Before]; ok {
			if rule.After == num {
				return false
			}
		}
	}
	return mapSet(after)
}

func mapSet(m map[int]int) bool {
	for _, v := range m {
		if v == 0 {
			return false
		}
	}
	return true
}

func Day2(rules []Rule, pages [][]int) int {
	score := 0
	set := RuleSet{rules}
	for _, page := range pages {
		valid := true
		for i := 0; i < len(page); i++ {
			after := map[int]int{}
			if i+1 < len(page) {
				for _, a := range page[i+1:] {
					after[a] = 0
				}
			}

			if valid {
				valid = set.Valid(page[i], after)
			}
		}
		if !valid {
			page = set.Correct(page)
			score += page[len(page)/2]
		}
	}

	return score
}
