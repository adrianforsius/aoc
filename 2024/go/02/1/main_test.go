package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Example(t *testing.T) {
	reportCard := Parse(exampleInput)

	const expected = 2
	assert.Equal(t, expected, Day1(reportCard))
}

func TestDay1(t *testing.T) {
	reportCard := Parse(input)

	const expected = 572
	assert.Equal(t, expected, Day1(reportCard))
}
