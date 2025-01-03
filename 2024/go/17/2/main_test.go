package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Example(t *testing.T) {
	const expected int64 = 117440
	assert.Equal(t, expected, Day2(Parse(exampleInput)))
}

func TestDay2(t *testing.T) {
	const expected int64 = 0
	assert.Equal(t, expected, Day2(Parse(input)))
}
