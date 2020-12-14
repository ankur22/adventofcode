package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	main "git.curve.tools/go/playground/adventofcalendar20/day3"
)

func TestFindTreesPartOne(t *testing.T) {
	got := main.FindTreesPartOne(input, 3, 1)

	assert.Equal(t, 7, got)
}

func TestFindTreesPartTwo(t *testing.T) {
	got := main.FindTreesPartTwo(input)

	assert.Equal(t, 336, got)
}

var input = [][]byte{
	{'.', '.', '#', '#', '.', '.', '.', '.', '.', '.', '.'},
	{'#', '.', '.', '.', '#', '.', '.', '.', '#', '.', '.'},
	{'.', '#', '.', '.', '.', '.', '#', '.', '.', '#', '.'},
	{'.', '.', '#', '.', '#', '.', '.', '.', '#', '.', '#'},
	{'.', '#', '.', '.', '.', '#', '#', '.', '.', '#', '.'},
	{'.', '.', '#', '.', '#', '#', '.', '.', '.', '.', '.'},
	{'.', '#', '.', '#', '.', '#', '.', '.', '.', '.', '#'},
	{'.', '#', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
	{'#', '.', '#', '#', '.', '.', '.', '#', '.', '.', '.'},
	{'#', '.', '.', '.', '#', '#', '.', '.', '.', '.', '#'},
	{'.', '#', '.', '.', '#', '.', '.', '.', '#', '.', '#'},
}
