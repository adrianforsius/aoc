package main

import (
	"log"
	"strconv"
	"strings"
)

type Fish struct {
	Timer           int
	FirstGeneration bool
}

func main() {
	fishNumbers := strings.Split(string(input), ",")
	var fishs []Fish
	for _, number := range fishNumbers {
		timer, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}
		fishs = append(fishs, Fish{Timer: timer})
	}

	previousGeneration := fishs
	for i := 0; i < 80; i++ {
		var fishs []Fish
		for _, fish := range previousGeneration {
			time := fish.Timer - 1
			if fish.Timer == 0 {
				fishs = append(fishs, Fish{
					Timer:           8,
					FirstGeneration: true,
				})
				time = 6
			}
			fishs = append(fishs, Fish{
				Timer:           time,
				FirstGeneration: false,
			})
		}
		previousGeneration = fishs
	}
	log.Println("fishes")
}

var input = `2,5,3,4,4,5,3,2,3,3,2,2,4,2,5,4,1,1,4,4,5,1,2,1,5,2,1,5,1,1,1,2,4,3,3,1,4,2,3,4,5,1,2,5,1,2,2,5,2,4,4,1,4,5,4,2,1,5,5,3,2,1,3,2,1,4,2,5,5,5,2,3,3,5,1,1,5,3,4,2,1,4,4,5,4,5,3,1,4,5,1,5,3,5,4,4,4,1,4,2,2,2,5,4,3,1,4,4,3,4,2,1,1,5,3,3,2,5,3,1,2,2,4,1,4,1,5,1,1,2,5,2,2,5,2,4,4,3,4,1,3,3,5,4,5,4,5,5,5,5,5,4,4,5,3,4,3,3,1,1,5,2,4,5,5,1,5,2,4,5,4,2,4,4,4,2,2,2,2,2,3,5,3,1,1,2,1,1,5,1,4,3,4,2,5,3,4,4,3,5,5,5,4,1,3,4,4,2,2,1,4,1,2,1,2,1,5,5,3,4,1,3,2,1,4,5,1,5,5,1,2,3,4,2,1,4,1,4,2,3,3,2,4,1,4,1,4,4,1,5,3,1,5,2,1,1,2,3,3,2,4,1,2,1,5,1,1,2,1,2,1,2,4,5,3,5,5,1,3,4,1,1,3,3,2,2,4,3,1,1,2,4,1,1,1,5,4,2,4,3`
