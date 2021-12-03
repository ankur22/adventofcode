package main_test

import (
	"testing"

	main "github.com/ankur22/adventofcode/aoc2021/day3"
)

func TestGetPowerConsumptionPart1(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	want := 198

	got := main.GetPowerConsumptionPart1(input)

	if got != want {
		t.Errorf("got %d is not the same as want %d", got, want)
	}
}

func TestGetCO2O2Part2(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	want := 230

	got := main.GetCO2O2Part2(input)

	if got != want {
		t.Errorf("got %d is not the same as want %d", got, want)
	}
}
