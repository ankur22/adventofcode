package main_test

import (
	"testing"

	main "github.com/ankur22/adventofcode/aoc2021/day1"
)

func TestFindDepthChangesPart1(t *testing.T) {
	depths := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	want := 7

	got := main.FindDepthChangesPart1(depths)

	if want != got {
		t.Errorf("got %d is not equal to want %d", got, want)
	}
}

func TestFindDepthChangesPart2(t *testing.T) {
	depths := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	want := 5

	got := main.FindDepthChangesPart2(depths)

	if want != got {
		t.Errorf("got %d is not equal to want %d", got, want)
	}
}
