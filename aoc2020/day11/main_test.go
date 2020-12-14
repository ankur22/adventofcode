package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	main "git.curve.tools/go/playground/adventofcode/aoc2020/day11"
)

func TestStableSeatOccupiedCount(t *testing.T) {
	floor := [][]byte{
		{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'L', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', '.', 'L', '.', '.', 'L', '.', '.'},
		{'L', 'L', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
		{'.', '.', 'L', '.', 'L', '.', '.', '.', '.', '.'},
		{'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L'},
		{'L', '.', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L'},
		{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
	}

	want := 37

	got := main.StableSeatOccupiedCount(floor)
	assert.Equal(t, want, got)
}

func TestLineOfSightStableSeatOccupiedCount(t *testing.T) {
	floor := [][]byte{
		{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'L', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', '.', 'L', '.', '.', 'L', '.', '.'},
		{'L', 'L', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
		{'.', '.', 'L', '.', 'L', '.', '.', '.', '.', '.'},
		{'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L'},
		{'L', '.', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L'},
		{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
	}

	want := 26

	got := main.LineOfSightStableSeatOccupiedCount(floor)
	assert.Equal(t, want, got)
}
