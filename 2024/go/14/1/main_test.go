package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Example(t *testing.T) {
	const expected = 12
	assert.Equal(t, expected, Day1(Parse(exampleInput), 7, 11))
}

func TestDay1(t *testing.T) {
	const expected = 0
	assert.Equal(t, expected, Day1(Parse(input), 103, 101))
}
