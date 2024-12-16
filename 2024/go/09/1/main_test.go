package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Example(t *testing.T) {
	const expected = 1928
	assert.Equal(t, expected, Day1(Parse(exampleInput)))
}

func TestDay1(t *testing.T) {
	const expected = 0
	assert.Equal(t, expected, Day1(Parse(input)))
}
