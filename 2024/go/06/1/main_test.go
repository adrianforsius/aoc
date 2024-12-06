package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Example(t *testing.T) {
	const expected = 41
	assert.Equal(t, expected, Day1(Parse(exampleInput)))
}

func TestDay1Custom2Example(t *testing.T) {
	const expected = 31
	assert.Equal(t, expected, Day1(Parse(exampleCustomInput)))
}

func TestDay1(t *testing.T) {
	const expected = 0
	assert.Equal(t, expected, Day1(Parse(input)))
}
