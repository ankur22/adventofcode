package main_test

import (
	"testing"

	main "github.com/ankur22/adventofcode/aoc2021/day5"
)

func TestFindVentOverlapsPart1(t *testing.T) {
	input := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}
	want := 5

	got := main.FindVentOverlapsPart1(input)
	if got != want {
		t.Errorf("got %d is not the same as want %d", got, want)
	}
}

func TestFindVentOverlapsPart2(t *testing.T) {
	input := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}
	want := 12

	got := main.FindVentOverlapsPart2(input)
	if got != want {
		t.Errorf("got %d is not the same as want %d", got, want)
	}
}
