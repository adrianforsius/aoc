package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Example(t *testing.T) {
	reportCard := Parse(exampleInput)

	const expected = 4
	assert.Equal(t, expected, Day2(reportCard))
}

func TestDay2(t *testing.T) {
	reportCard := Parse(input)

	const expected = 2
	assert.Equal(t, expected, Day2(reportCard))
}
