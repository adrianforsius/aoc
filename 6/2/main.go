package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	fishNumbers := strings.Split(string(input), ",")
	var fishs [9]int
	for _, number := range fishNumbers {
		timer, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}

		fishs[timer] += 1
	}
	log.Println("fishs", fishs)

	previousGeneration := fishs
	for i := 1; i <= 256; i++ {
		var fishs [9]int
		respawn := previousGeneration[0]
		for e := 1; e <= 8; e++ {
			fishs[e-1] = previousGeneration[e]
		}
		fishs[6] += respawn
		fishs[8] += respawn

		previousGeneration = fishs
		log.Println("day", i, "fish", fishs, "prev", previousGeneration)
	}
	count := 0
	for time, section := range previousGeneration {
		log.Println("count", time, count, section)
		count += section
	}
	log.Println("fishes", count)
}

// var input = `3,4,3,1,2`

var input = `2,5,3,4,4,5,3,2,3,3,2,2,4,2,5,4,1,1,4,4,5,1,2,1,5,2,1,5,1,1,1,2,4,3,3,1,4,2,3,4,5,1,2,5,1,2,2,5,2,4,4,1,4,5,4,2,1,5,5,3,2,1,3,2,1,4,2,5,5,5,2,3,3,5,1,1,5,3,4,2,1,4,4,5,4,5,3,1,4,5,1,5,3,5,4,4,4,1,4,2,2,2,5,4,3,1,4,4,3,4,2,1,1,5,3,3,2,5,3,1,2,2,4,1,4,1,5,1,1,2,5,2,2,5,2,4,4,3,4,1,3,3,5,4,5,4,5,5,5,5,5,4,4,5,3,4,3,3,1,1,5,2,4,5,5,1,5,2,4,5,4,2,4,4,4,2,2,2,2,2,3,5,3,1,1,2,1,1,5,1,4,3,4,2,5,3,4,4,3,5,5,5,4,1,3,4,4,2,2,1,4,1,2,1,2,1,5,5,3,4,1,3,2,1,4,5,1,5,5,1,2,3,4,2,1,4,1,4,2,3,3,2,4,1,4,1,4,4,1,5,3,1,5,2,1,1,2,3,3,2,4,1,2,1,5,1,1,2,1,2,1,2,4,5,3,5,5,1,3,4,1,1,3,3,2,2,4,3,1,1,2,4,1,1,1,5,4,2,4,3`
