package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Example(t *testing.T) {
	const expected = 55312
	assert.Equal(t, expected, Day2(Parse(exampleInput), 22))
}

func TestDay2(t *testing.T) {
	const expected = 0
	assert.Equal(t, expected, Day2(Parse(input), 75))
}
