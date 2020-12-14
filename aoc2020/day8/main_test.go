package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	main "git.curve.tools/go/playground/adventofcode/aoc2020/day8"
)

func TestGetAccumulatorBeforeInfiniteLoop(t *testing.T) {
	input := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}

	want := 5

	gotb, goti := main.GetAccumulatorBeforeInfiniteLoop(input)
	assert.Equal(t, want, goti)
	assert.True(t, gotb)
}

func TestGetAccumulatorAfterFixing(t *testing.T) {
	input := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}

	want := 8

	got := main.GetAccumulatorAfterFixing(input)
	assert.Equal(t, want, got)
}
