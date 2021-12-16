package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	lines := strings.Split(string(input), "\n")

	oxygen := lines
	oxyCount := 0
	for len(oxygen) > 1 {
		ones, zeros := popular(oxygen, oxyCount)

		oxygen = ones
		if len(zeros) > len(ones) {
			oxygen = zeros
		}
		log.Println("oxygen", oxygen, "count", oxyCount)
		oxyCount += 1
	}

	scrubber := lines
	scrubberCount := 0
	for len(scrubber) > 1 {
		ones, zeros := popular(scrubber, scrubberCount)

		scrubber = zeros
		if len(ones) < len(zeros) {
			scrubber = ones
		}
		scrubberCount += 1
	}
	log.Println("oxy", oxygen)
	log.Println("scrubber", scrubber)

	oxygenVal, err := strconv.ParseInt(oxygen[0], 2, 64)
	if err != nil {
		log.Fatalf("invalid oxy %s", err)
	}
	scrubberVal, err := strconv.ParseInt(scrubber[0], 2, 64)
	if err != nil {
		log.Fatalf("invalid scrubber %s", err)
	}

	log.Println("scrubber", scrubberVal)
	log.Println("oxygen", oxygenVal)
	log.Println("count", oxygenVal*scrubberVal)
}

func popular(report []string, pos int) ([]string, []string) {
	var ones []string
	var zeros []string
	for _, line := range report {
		if string(line[pos]) == "1" {
			ones = append(ones, line)
		} else {
			zeros = append(zeros, line)
		}
	}
	return ones, zeros
}

var input = `000001110001
000001111101
010011000000
000000000011
010000100111
101000100001
101000001101
010011001111
001000111011
110010110011
010110110000
110000111100
111110101000
100011001101
100100000001
101100011011
110111111100
010110101011
000010010001
100000000111
100000111100
111011100010
101010001101
111011110000
111010001101
010010110110
000101011011
001101101110
101100100000
101010111101
100010010111
000100100101
001000100011
000001011111
000100001101
011111100101
101001110011
100010000001
010111010100
100101101001
100001001000
110011010111
010111010001
111000011010
110101100100
011100001000
001010000000
011111001010
010111110111
100000000001
111011101011
100000100010
111110100110
001000000001
100001111110
000110010110
011001101010
101111001011
010010010100
000110110111
001000011111
100111001110
111110111111
100000010010
110111111001
000111011100
100100000010
111000110001
000011101010
010111111011
101010111000
010101001101
000010101101
001000000111
011111101110
010011100010
100010100000
001101000011
010011101010
111100000011
001100010010
011111100001
000001101101
011011100010
010011010000
111010001011
100111011101
011110010110
001110010100
010001110101
001011011100
001111011101
111111000001
100101001111
110111011111
001110100010
111111111110
001101111010
001110001101
001001110100
100000110000
001100000001
100001101100
110100110011
111101100000
000000000001
000010101100
101000000011
110011110100
000001010011
001101010111
100001101011
111000101011
000100010001
011111101010
010110011011
110111110000
100110111111
110111000011
110100101011
110011011001
101101001010
101110111100
101010110010
011111110000
100000011010
010100011111
100001111011
100000000000
100101000111
001111010111
110111111010
011011011111
111101000101
101111000111
101111000001
010100101110
001001101110
100001011000
100100010000
111111111011
111000000100
000010111000
010111111100
111110001100
011001111011
101011100111
111101001001
001111100011
101000101110
010001101001
001101001110
011011000110
111000000000
010010001100
111011000110
011011001011
111000000011
110110001110
001001101001
100011010101
100101111011
101111110001
011001011011
111101000100
110000100001
010001000110
011000101101
000101001001
000011111100
010001110011
111111100010
000110010000
101101001101
110000010111
000010010010
001110111001
110011001111
000010110010
010010000110
101100101100
111111011100
000111011111
011011111001
011100101011
100011010110
011101011111
010000010111
000110111001
010010111010
010011101000
011010101110
011100111101
010110110011
011110100100
100000100011
100001001011
001000111110
101101001011
101111101000
001011000111
101111101110
101001100011
110010110010
101111010001
110100010000
100011000010
001010110010
110111101101
000010001000
001101100000
101110101001
000111000000
011010101000
100111110011
111110000011
111101111011
100111110110
110100000111
101100111001
011111000001
110110101111
111011100101
000001001111
001111001010
010110010100
000001010001
111101111110
101010000110
001111011100
110001101010
101011110101
001001001011
101100010101
101011101011
100010110011
101001011111
101010010100
100001110011
001110000101
011110111110
111010111111
100110011001
010101111100
000101010100
111101101001
011100011010
001000111010
011111010011
100010001010
000111010011
010011100000
011011001000
001011101110
001001111100
111000011101
101011111101
110000110110
100010111001
010111101000
100101101101
011011011000
010101010011
011010000001
000001101011
101001110000
101111111100
000111100001
010111010110
010110000000
101010100100
010011110000
111011000100
110010100111
101101011111
011011101100
000011101110
011001100000
010110111011
100110011000
111101000011
000111011001
001101011110
010000100110
111011100001
001100100101
110110001001
111101001111
000111011000
001110011010
101000000000
111110011100
000100110000
100001110010
101101011000
101001100110
010000101110
010110111100
111000100100
110101101000
001111010100
110111010000
110010011100
110111000110
010100010011
001000010111
100111101001
100110101011
101010000100
111000010001
001001010011
100100111001
010000010011
010110010110
100101100010
011010011100
101101011110
100110101111
011100011100
011000110110
001000001110
101011111001
011001001011
000100111101
011111110100
101000110001
110111011000
100100010111
000111111100
110001110101
111001001000
101000011110
110111100101
010010001010
100110100010
111100011010
110111100110
011100110110
110100111000
011110010010
110100010001
010010100000
010100000000
010110110010
110111111101
101000000110
100010101000
111101000010
001011010001
010011111110
000001010000
100101001011
011000110001
001111101011
000001110010
101010101010
011011010000
011011001101
111001000100
110111001010
000111010100
000111111101
010111000110
010100010111
010001001010
001001010110
101111001110
000111100000
110111000100
001001000001
011111110111
101111000100
110100011010
100000111010
000000011101
001010010100
001000001100
001101101010
011110110001
100101011101
101000101001
101010000011
100111100001
001011001000
000111111001
000011110101
111111110000
100010000101
001100111000
001001011011
001111111110
011000110111
110111001110
111001111101
000111010000
001110100011
100111111000
111011110010
001000110101
110101111111
101110001000
100100000000
101001110100
000100010011
001110110010
011011000111
011110011100
111010110011
100011010000
101110100001
101011001110
010101011000
101100011101
011100000111
101010010011
000101000000
111000000110
110010001001
001111111010
010001101110
111011100110
010111101011
101111010100
011111110101
010110110001
000100101011
101010101001
111110110010
011110000010
111001001101
101100110110
011000001110
001110000111
100001000001
100111110010
101101101111
000110011001
000000110011
111110001101
000000111100
101111100010
100110100001
011110110110
101011110100
101101110100
011111110011
110110011110
111010001110
100101110000
100000111101
001011010110
101001111100
000010011001
000111010001
100000101010
100010011000
100100001000
001010100001
111100001011
110000000110
111011000111
000000111010
101100000111
010111010010
011001011100
000011100010
110111101010
111111010001
100011111000
111111100101
110011110101
100110001010
101011101111
111111100110
101111110010
110000010001
110001100100
101100110001
011010111001
001000111111
110110001000
011110001011
000100010101
010101100100
000001001000
000000010101
100000101101
010000000100
010110000110
111100110010
100010001101
011001111010
100000111011
110111101110
111000100010
101011010010
111000011011
001101000111
110011111100
110101010101
010100101010
110110010110
101011100000
111101111010
101000011000
000101011000
111010000001
000100000010
100110000100
010111100001
110001010101
110011111101
110000101010
011100111110
011010100100
000111100011
101000000111
011110101110
001111001110
110010010010
010010111111
111110011011
110110111011
111000111011
010100100000
100001010110
010010011111
010000101010
011011110101
001101101101
101110111110
001101101000
110100110010
011010010011
011011011110
011010101100
010101010010
101110000000
000001011011
110100101010
110010001111
111011111110
000101100010
110000010011
111110001010
000101100100
110001110100
011101111110
011010001001
110110101101
011111101101
111111101101
111001101000
011110100110
111011011001
010001110010
101001110111
010011110100
011011001010
110100011101
111001111001
010001100111
011100110000
101110001101
000100111111
001100100111
000010111001
110000100101
010101101101
100100110111
111110111100
100000101111
011110110111
100010001110
001000000110
111110110110
001100100010
000011111111
010000100010
000101111100
001000100111
011001000110
110001111110
001001011010
101111101111
001000010100
100100001100
001010101100
110100000101
001011101000
101011100101
011011111100
111000111100
110000001101
010011010101
001001110110
110110101011
101100010111
011000010101
010100011000
010001001100
010101110101
110100100110
111111101011
000110010111
100100110001
001111000011
010100000010
011101101010
010111110101
011101011001
101111010010
000000010110
100000100000
010001011010
000110011111
110000110111
111001110011
110100111110
000000001001
001000110001
000110000101
010110100000
001100010100
101010100101
111100001100
111010110110
000001100111
001011011110
000100100001
111001111000
110001011001
110101110101
101010110101
000011001010
001010110000
101111110101
101001111001
100010111101
001110101001
110011110111
101011001001
010001111100
000101100011
011110001100
111111110101
010010011001
111000110111
100110100000
010110110100
011111111110
001111101100
100010010010
101110010110
111010000101
011011000100
101000011001
001000001101
010110100111
000000011010
111101110000
000100000001
101000000010
011110101111
100010001111
101101110101
111110011101
001111111011
110000101100
000111100010
101111011111
011101101111
001011000101
100110101110
010100100001
010000011111
010100011001
011010100110
011010110111
101011101101
011101000001
010111011011
000010011100
011111011000
110011010100
011101101110
000100110100
010010111100
110000101000
010000110001
111011111001
110001101001
010101001010
111001010111
111001100000
010111000111
111011001101
001111011110
110100110111
110001010100
100001011010
101100001010
111111001101
111000010000
001001011101
111110100101
011110001110
111001101001
111011010001
001100001111
001110101111
010011111011
001001100001
010100111010
110000010101
011000001010
001111101111
111000000101
011011000011
111110010101
011000000111
010111000000
001100010000
110001100010
110011011100
101001101111
010010101111
011110101000
001000111100
100101110110
100100101011
100101000110
110100000100
001111110001
010001000100
110011011111
110110110001
010001111001
101000011101
111111101100
001111110100
111100011100
110111011011
000001001011
001011000100
000000010100
010001100010
010010001001
101011101110
000010000000
010101010000
101001000000
111010101011
000101111000
101100101011
111001100100
001011000011
010111101100
000100111110
000111111111
010001010000
110110000001
110101110100
001010100110
101110011011
011110001101
011100100001
111110010110
100100100001
110010011001
001010110100
011011110010
001110110001
010100110111
100000111001
100001011110
011011011100
101110000111
000001011010
000000111000
001001111111
110101101001
010011001100
101011100100
000011111101
011100100010
000010111011
111000010111
101101011101
100001010011
010001000000
111110001011
111101110010
100001101110
100000111000
010101010101
100110110001
010001101000
110001011011
011011010001
011101000000
101101011100
001101001001
101001100000
111011111000
110101111000
011001110101
011000111000
000011010100
011010010101
000010110110
110001101000
000001111011
000101110001
000110100011
011111000101
101010011001
110000011011
010010101101
110101000100
111001100011
010001100001
111001111111
000000101110
111010000111
100100000100
111110111001
101100011000
101111010110
010110111111
100001000110
110000110011
101000001111
000001001010
001100110100
100000110101
010111110010
010011011011
110101101101
010100111100
000010101011
111100001101
101001000011
001101110111
011001010001
111100111001
011100111010
111010010010
000101111011
000001110101
110110100001
111100110011
110110101010
110101010010
000000010000
100100010010
001000100110
010000110101
111011000000
101101100001
000110001101
110110110110
000010001101
010001101100
000000001000
000110101011
001000011101
001111110000
000100000100
010111110110
111111110110
101110100000
011010010111
111111101111
011010100001
110101111010
000010010111
011001011000
111010101010
100101100100
010110100101
011110111000
101001010110
101111001100
001111001000
011100101100
101100000010
110011100100
101001101000
110110000110
000111101011
111110001000
001011110000
001001100000
000111110111
000101111010
001010000110
011100000000
110100001000
111001001110
010111001100
101111111001
110010110111
010101010111
100111100110
110010100100
001011101100
110111001111
001001100111
010011101110
001101000100
000010111101
111101101110
111011011101
100101110100
101110110110
000110101101
110100100011
001001010010
111000101111
000010011110
001101110011
111010010001
010110001110
010000110100
010101111001
001010011011
010110100001
000001100110
100101111101
110000000101
010001010010
111101110100
001000011010
011011001001
010010010000
111111001100
001001011111
010001100100
100011111011
011000011100
100011011001
011111100111
110111101001
011011110111
010100000110
011011001100
000000101111
100000100100
101011001010
101101001111
000011111010
011101111010
000011110100
001011111110
101000010010
010000110010
001000000000
011001111100
111111000011
010101000000
110001111010
101001101101
001011000001
110010101010
110111110101
010100110110
101010111001
010111100110
101001000110
000011001000
110000001010
110001100110
111100011110
100100101111
001110001001
011011111010
000011000001
001111011011
000110000011
010010100111
110010001010
011100011001
101111101011
001100000100
101100101111
000101001100
010100010100
100100010001
100111111110
010111100111
001000011000
101010001001
110000110100
111110110001
100110100101
101000010100
100011000000
001110100100
110000110001`
