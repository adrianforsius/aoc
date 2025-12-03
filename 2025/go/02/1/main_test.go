package main

import (
	"flag"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Input() []byte {
	fptr := flag.String("input", "input.txt", "file path to read from")
	flag.Parse()

	contents, err := os.ReadFile(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	return contents
}

func Example() []byte {
	fptr := flag.String("example", "example.txt", "file path to read from")
	flag.Parse()

	contents, err := os.ReadFile(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	return contents
}

func TestExample(t *testing.T) {
	const expected = 1_227_775_554
	assert.Equal(t, expected, Puzzle(Parse(Example())))
}

func TestInput(t *testing.T) {
	const expected = 38437576669 // too low
	assert.Equal(t, expected, Puzzle(Parse(Input())))
}
