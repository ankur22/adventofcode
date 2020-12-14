package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	main "git.curve.tools/go/playground/adventofcode/aoc2020/day9"
)

func TestFindNotSum(t *testing.T) {
	input := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}

	want := 127
	const preamble = 5

	got := main.FindNotSum(input, preamble)
	assert.Equal(t, want, got)
}

func TestFindContiguous(t *testing.T) {
	input := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}

	want := 62
	const preamble = 5

	got := main.FindContiguous(input, preamble)
	assert.Equal(t, want, got)
}
