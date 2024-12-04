package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Example(t *testing.T) {
	output := Parse(exampleInput)

	const expected = 48
	assert.Equal(t, expected, Day2(output))
}

func TestDay2(t *testing.T) {
	output := Parse(input)

	const expected = 111762583
	assert.Equal(t, expected, Day2(output))
}
