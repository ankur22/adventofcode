package main_test

import (
	"testing"

	main "github.com/ankur22/adventofcode/aoc2021/day6"
)

func TestCountLanternfishPart1(t *testing.T) {
	input := []string{"3", "4", "3", "1", "2"}
	want := 5934

	got := main.CountLanternfishPart1(input, 80)
	if got != want {
		t.Errorf("got %d not equal to want %d", got, want)
	}
}

func TestCountLanternfishPart2(t *testing.T) {
	input := []string{"3", "4", "3", "1", "2"}
	want := 26984457539

	got := main.CountLanternfishPart1(input, 256)
	if got != want {
		t.Errorf("got %d not equal to want %d", got, want)
	}
}
