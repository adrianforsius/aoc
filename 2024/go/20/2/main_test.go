package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Example(t *testing.T) {
	const expected = 22
	assert.Equal(t, expected, Day2(Parse(exampleInput), 50))
}

// func TestDay2(t *testing.T) {
// 	const expected = 0
// 	assert.Equal(t, expected, Day2(Parse(input), 100))
// }
