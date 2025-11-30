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

func TestDay1Example(t *testing.T) {
	const expected = 11
	assert.Equal(t, expected, Day1(Parse(Example())))
}

func TestDay1(t *testing.T) {
	const expected = 1879048
	assert.Equal(t, expected, Day1(Parse(Input())))
}
