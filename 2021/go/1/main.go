package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	log.Println("day 1")

	// resp, err := http.Get("https://adventofcode.com/2021/day/1/input")
	// if err != nil {
	// 	log.Fatalf("get input %s", err)
	// }
	// defer resp.Body.Close()

	// content, err := ioutil.ReadAll(resp.Body)
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("read input %s", err)
	}
	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]
	var digits []int
	for _, line := range lines {
		digit, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalln("invalid input", err)
		}
		digits = append(digits, digit)
	}

	count := 0
	for i := 1; i < len(digits); i++ {
		if digits[i-1] < digits[i] {
			count += 1
		}
		if i < 10 {
			log.Println("calculating", digits[i-1], digits[i], digits[i-1] < digits[i])
		}
		if i > len(digits)-10 {
			log.Println("calculating", digits[i-1], digits[i], digits[i-1] < digits[i])
		}
	}
	log.Println("count", count)
}
