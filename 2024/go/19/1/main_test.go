package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Example(t *testing.T) {
	const expected = 6
	assert.Equal(t, expected, Day1(Parse(exampleInput)))
}

// too low:
// 315
// 288
// func TestDay1(t *testing.T) {
// 	const expected = 0
// 	assert.Equal(t, expected, Day1(Parse(input)))
// }
