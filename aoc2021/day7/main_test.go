package main_test

import (
	"testing"

	main "github.com/ankur22/adventofcode/aoc2021/day7"
)

func TestFindHorizLeastCost(t *testing.T) {
	input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	want := 37

	got := main.FindHorizLeastCost(input)
	if got != want {
		t.Errorf("got %d is not the same and want %d", got, want)
	}
}

func TestFindHorizLeastCostPart2(t *testing.T) {
	input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	want := 168

	got := main.FindHorizLeastCostPart2(input)
	if got != want {
		t.Errorf("got %d is not the same and want %d", got, want)
	}
}
