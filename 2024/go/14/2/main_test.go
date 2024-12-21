package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestDay2Example(t *testing.T) {
// 	const expected = 2858
// 	assert.Equal(t, expected, Day2(Parse(exampleInput)))
// }

func TestDay2(t *testing.T) {
	const expected = 0
	assert.Equal(t, expected, Day2(Parse(input), 103, 101))
}
