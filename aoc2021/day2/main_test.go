package main_test

import (
	"testing"

	main "github.com/ankur22/adventofcode/aoc2021/day2"
)

func TestFindDepthChangesPart1(t *testing.T) {
	instructions := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	want := 150

	got := main.FindPositionPart1(instructions)

	if want != got {
		t.Errorf("got %d is not equal to want %d", got, want)
	}
}

func TestFindDepthChangesPart2(t *testing.T) {
	instructions := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	want := 900

	got := main.FindPositionPart2(instructions)

	if want != got {
		t.Errorf("got %d is not equal to want %d", got, want)
	}
}

func TestParseInstruction(t *testing.T) {
	instruction := "forward 5"
	wantDir := "forward"
	wantAmount := 5

	gotDir, gotAmount := main.ParseInstruction(instruction)

	if wantDir != gotDir {
		t.Errorf("got %s is not equal to want %s", gotDir, wantDir)
	}

	if wantAmount != gotAmount {
		t.Errorf("got %d is not equal to want %d", gotAmount, wantAmount)
	}
}
