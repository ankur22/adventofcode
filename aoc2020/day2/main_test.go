package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	main "git.curve.tools/go/playground/adventofcalendar20/day2"
)

func TestNumOneValid(t *testing.T) {
	input := []main.Row{
		{
			StartRange: 1,
			EndRange:   3,
			Rule:       "a",
			Password:   "abcde",
		},
		{
			StartRange: 1,
			EndRange:   3,
			Rule:       "b",
			Password:   "cdefg",
		},
		{
			StartRange: 2,
			EndRange:   9,
			Rule:       "c",
			Password:   "ccccccccc",
		},
	}

	want := 2

	got := main.NumValidPartOne(input)

	assert.Equal(t, want, got)
}

func TestNumTwoValid(t *testing.T) {
	input := []main.Row{
		{
			StartRange: 1,
			EndRange:   3,
			Rule:       "a",
			Password:   "abcde",
		},
		{
			StartRange: 1,
			EndRange:   3,
			Rule:       "b",
			Password:   "cdefg",
		},
		{
			StartRange: 2,
			EndRange:   9,
			Rule:       "c",
			Password:   "ccccccccc",
		},
	}

	want := 1

	got := main.NumValidPartTwo(input)

	assert.Equal(t, want, got)
}
