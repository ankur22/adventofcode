package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	main "git.curve.tools/go/playground/adventofcode/aoc2020/day6"
)

func TestSumYesAnswers(t *testing.T) {
	input := []string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
	}

	got := main.SumYesAnswers(input)

	assert.Equal(t, 11, got)
}

func TestSumAllSameYesAnswers(t *testing.T) {
	input := []string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
	}

	got := main.SumAllSameYesAnswers(input)

	assert.Equal(t, 6, got)
}
