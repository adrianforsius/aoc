package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Example(t *testing.T) {
	const expected = 161
	assert.Equal(t, expected, Day1(Parse(exampleInput)))
}

func TestDay1(t *testing.T) {
	in := Parse(input)

	const expected = 0
	assert.Equal(t, expected, Day1(in))
}
