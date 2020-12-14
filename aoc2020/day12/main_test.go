package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	main "git.curve.tools/go/playground/adventofcode/aoc2020/day12"
)

func TestFindManhattanDistance(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int64
	}{
		{
			name: "small",
			input: []string{
				"F10",
				"N3",
				"F7",
				"R90",
				"F11",
			},
			want: 25,
		},
		{
			name: "rotate only",
			input: []string{
				"R1080",
			},
			want: 0,
		},
		{
			name: "rotate left",
			input: []string{
				"L90",
				"F10",
				"L90",
				"F5",
			},
			want: 15,
		},
		{
			name: "rotate right",
			input: []string{
				"R180",
				"F10",
				"R180",
				"F5",
			},
			want: 5,
		},
		{
			name: "all dirs",
			input: []string{
				"N1",
				"E2",
				"S4",
				"W5",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := main.FindManhattanDistance(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestWaypointFindManhattanDistance(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int64
	}{
		{
			name: "small",
			input: []string{
				"F10",
				"N3",
				"F7",
				"R90",
				"F11",
			},
			want: 286,
		},
		// {
		// 	name: "rotate only",
		// 	input: []string{
		// 		"R1080",
		// 	},
		// 	want: 0,
		// },
		// {
		// 	name: "rotate left",
		// 	input: []string{
		// 		"L90",
		// 		"F10",
		// 		"L90",
		// 		"F5",
		// 	},
		// 	want: 15,
		// },
		// {
		// 	name: "rotate right",
		// 	input: []string{
		// 		"R180",
		// 		"F10",
		// 		"R180",
		// 		"F5",
		// 	},
		// 	want: 5,
		// },
		// {
		// 	name: "all dirs",
		// 	input: []string{
		// 		"N1",
		// 		"E2",
		// 		"S4",
		// 		"W5",
		// 	},
		// 	want: 6,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := main.WaypointFindManhattanDistance(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
