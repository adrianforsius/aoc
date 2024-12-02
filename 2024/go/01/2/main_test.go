package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Example(t *testing.T) {
	left, right := Parse(exampleInput)

	const output = 31
	assert.Equal(t, output, Day2(left, right))
}

func TestDay2(t *testing.T) {
	left, right := Parse(input)

	const output = 0
	assert.Equal(t, output, Day2(left, right))
}
