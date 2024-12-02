package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Example(t *testing.T) {
	left, right := Parse(exampleInput)

	const expected = 11
	assert.Equal(t, expected, Day1(left, right))
}

func TestDay1(t *testing.T) {
	left, right := Parse(input)

	const expected = 11
	assert.Equal(t, expected, Day1(left, right))
}
