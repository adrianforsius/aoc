package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestDay1Example(t *testing.T) {
// 	const expected = 22
// 	assert.Equal(t, expected, Day1(Parse(exampleInput)))
// }

// not right: 762
// not right: 9226
// too low: 490
// to high: 27498
// not right: 1449
func TestDay1(t *testing.T) {
	const expected = 0
	assert.Equal(t, expected, Day1(Parse(input)))
}
