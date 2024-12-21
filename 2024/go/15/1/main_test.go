package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Example(t *testing.T) {
	const expected = 10092
	assert.Equal(t, expected, Day1(Parse(exampleInput)))
}

func TestDay1(t *testing.T) {
	const expected = 1324
	assert.Equal(t, expected, Day1(Parse(input)))
}
